package loader

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
)

type Loader struct {
	dsn string
}

func NewLoader(dsn string) *Loader {
	return &Loader{
		dsn: dsn,
	}
}

func (loader *Loader) CopyFromCsv(csvPath string, schemaName, tableName string, options string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, loader.dsn)
	if err != nil {
		return 0, err
	}
	defer conn.Close(context.Background())

	columnNames, err := getCsvColumnNames(csvPath)
	if err != nil {
		return 0, err
	}

	csvFile, err := os.Open(csvPath)
	if err != nil {
		return 0, err
	}
	defer csvFile.Close()

	fullTableName := fmt.Sprintf("\"%s\".\"%s\"", schemaName, tableName)

	quotedColumnNames := lo.Map(columnNames, func(columnName string, index int) string {
		return fmt.Sprintf("\"%s\"", columnName)
	})

	if options == "" {
		options = `DELIMITER ',', FORMAT CSV, HEADER true`
	}

	sql := fmt.Sprintf(`COPY %s(%s) FROM STDIN WITH (%s)`,
		fullTableName, strings.Join(quotedColumnNames, ","), options,
	)

	commandTag, err := conn.PgConn().CopyFrom(context.Background(), csvFile, sql)
	if err != nil {
		return 0, err
	}
	return commandTag.RowsAffected(), nil
}

func getCsvColumnNames(path string) ([]string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	columnNames, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	return columnNames, nil
}
