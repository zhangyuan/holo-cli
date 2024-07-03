# holo-cli


## Usage

### Generate Create Table DDL from schema definition in YAML

```bash
./holo-cli schema2ddl -s examples/users.yaml -o tmp/users.sql
```

### COPY CSV file to Hologres

`COPY` to Hologres with the default option `DELIMITER ',', FORMAT CSV, HEADER true`:

```bash
./holo-cli copyFromCsv -c tmp/users.csv -s public -t users --options
```

Or specify `COPY` command's options:

```bash
./holo-cli copyFromCsv -c tmp/users.csv -s public -t users --options "DELIMITER ',', FORMAT CSV, HEADER true, STREAM_MODE TRUE, ON_CONFLICT UPDATE"
```

### COPY SQL query result to local file

`COPY` from hologres to local file with the default option `DELIMITER ',', FORMAT CSV, HEADER true`:

```bash
./holo-cli copyToLocal -q "select * from users" -o tmp/file.csv
```

Or specify `COPY` command's options:

```bash
./holo-cli copyToLocal -q "select * from users" -o tmp/file.text --options "FORMAT TEXT"
```
