# sqltestutil

[![Documentation](https://godoc.org/github.com/bitcomplete/sqltestutil?status.svg)](http://godoc.org/github.com/bitcomplete/sqltestutil)

Utilities for testing Golang code that runs SQL.

## Usage

### PostgresContainer

PostgresContainer is a Docker container running Postgres that can be used to
cheaply start a throwaway Postgres instance for testing.

### RunMigration

RunMigration reads all of the files matching *.up.sql in a directory and
executes them in lexicographical order against the provided DB.

### LoadScenario

LoadScenario reads a YAML "scenario" file and uses it to populate the given DB.

### Suite

Suite is a [testify
suite](https://pkg.go.dev/github.com/stretchr/testify@v1.7.0/suite#Suite) that
provides a database connection for running tests against a SQL database.
