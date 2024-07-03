# holo-cli


## Usage

### Generate Create Table DDL from schema definition in YAML

```bash
./holo-cli schema2ddl -s examples/users.yaml -o users.sql
```

### COPY CSV file to Hologres

`COPY` to hologres with the default option `DELIMITER ',', FORMAT CSV, HEADER true`:

```bash
copyFromCsv -c tmp/users.csv -s public -t users --options
```

Or specify `COPY` command's options:

```bash
copyFromCsv -c tmp/users.csv -s public -t users --options "DELIMITER ',', FORMAT CSV, HEADER true, STREAM_MODE TRUE, ON_CONFLICT UPDATE"
```
