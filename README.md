# answer-service
answer-map REST web service

## Setting up a local database
- [Install postgresql.](https://www.postgresql.org/download/)
- Add postgres binaries to path environment variable.
- Run `pg_ctl -D'.\data' initdb`.
- Run `pg_ctl -D'.\data' -l logfile start`.
- Run `psql -d postgres -U lewis -a -f .\sql\create_schema.sql`.

## Setting up sqlboiler
- Run `go install github.com/volatiletech/sqlboiler/v4@latest`.
- Run `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest`.
- Check `sqlboiler.toml` matches with your database. The database with the schema needs to be running. [See here for help.](https://github.com/volatiletech/sqlboiler#configuration)
- Check that the destination directory is not present or empty. This will avoid conflicts when overwriting existing files.
- Run `sqlboiler psql`.
- Run `go mod tidy` to import the dependencies of the generated code.
