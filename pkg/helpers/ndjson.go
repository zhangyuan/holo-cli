package helpers

import (
	"reflect"
	"slices"

	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
)

type Column struct {
	Name      string
	DataTypes []reflect.Kind
}

func (c *Column) AddDataType(dataType reflect.Kind) {
	if !slices.Contains(c.DataTypes, dataType) {
		c.DataTypes = append(c.DataTypes, dataType)
	}
}

func (c *Column) HologresDataType() string {
	if slices.Contains(c.DataTypes, reflect.Slice) || slices.Contains(c.DataTypes, reflect.Map) {
		return "JSON"
	}
	return "TEXT"
}

func (c *Column) HologresName() string {
	return strcase.ToSnake(c.Name)
}

func GetColumnNames(ndjsonPath string) ([]string, error) {
	columns, err := GetColumns(ndjsonPath)
	if err != nil {
		return nil, err
	}

	return lo.Map(columns, func(item Column, index int) string {
		return item.HologresName()
	}), nil
}
