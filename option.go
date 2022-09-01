package sqltestutil

import "fmt"

type Option func(*PostgresContainer)

func WithPassword(password string) Option {
	return func(container *PostgresContainer) {
		if len(password) == 0 {
			panic("sqltestutil: password option can not be empty")
		}
		container.password = password
	}
}
func WithUser(user string) Option {
	return func(container *PostgresContainer) {
		if len(user) == 0 {
			panic("sqltestutil: user option can not be empty")
		}
		container.user = user
	}
}
func WithPort(port uint16) Option {
	return func(container *PostgresContainer) {
		if port <= 1000 {
			panic("sqltestutil: port option can not be less than 1000")
		}
		container.port = fmt.Sprint(port)
	}
}

func WithVersion(version string) Option {
	return func(container *PostgresContainer) {
		if len(version) == 0 {
			panic("sqltestutil: version option can not be empty")
		}
		container.version = version
	}
}
func WithDBName(dbName string) Option {
	return func(container *PostgresContainer) {
		if len(dbName) == 0 {
			panic("sqltestutil: dbName option can not be empty")
		}
		container.dbName = dbName
	}
}
func WithContainerName(containerName string) Option {
	return func(container *PostgresContainer) {
		if len(containerName) == 0 {
			panic("sqltestutil: containerName option can not be empty")
		}
		container.containerName = containerName
	}
}
