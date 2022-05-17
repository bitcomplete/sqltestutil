package sqltestutil

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

// Suite is a testify suite [1] that provides a database connection for running
// tests against a SQL database. For each test that is run, a new transaction is
// started, and then rolled back at the end of the test so that each test can
// operate on a clean slate. Here's an example of how to use it:
//
//     type ExampleTestSuite struct {
//         sqltestutil.Suite
//     }
//
//     func (s *ExampleTestSuite) TestExample() {
//         _, err := s.Tx().Exec("INSERT INTO foo (bar) VALUES (?)", "baz")
//         s.Assert().NoError(err)
//     }
//
//     func TestExampleTestSuite(t *testing.T) {
//         suite.Run(t, &ExampleTestSuite{
//		       Suite: sqltestutil.Suite{
//                 Context: context.Background(),
//                 DriverName: "pgx",
//			       DataSourceName: "postgres://localhost:5432/example",
//             },
//         })
//     }
//
// [1]: https://pkg.go.dev/github.com/stretchr/testify@v1.7.0/suite#Suite
type Suite struct {
	suite.Suite

	// Context is a required field for constructing a Suite, and is used for
	// database operations within a suite. It's public because it's convenient to
	// have access to it in tests.
	context.Context

	// DriverName is a required field for constructing a Suite, and is used to
	// connect to the underlying SQL database.
	DriverName string

	// DataSourceName is a required field for constructing a Suite, and is used to
	// connect to the underlying SQL database.
	DataSourceName string

	db *sqlx.DB
	tx *sqlx.Tx
}

// DB returns the underlying SQL connection.
func (s *Suite) DB() *sqlx.DB {
	return s.db
}

// Tx returns the transaction for the current test.
func (s *Suite) Tx() *sqlx.Tx {
	return s.tx
}

func (s *Suite) SetupTest() {
	s.Require().Nil(s.tx)
	var err error
	s.tx, err = s.db.BeginTxx(s.Context, nil)
	s.Require().NoError(err)
}

func (s *Suite) TearDownTest() {
	if s.tx != nil {
		err := s.tx.Rollback()
		s.Require().NoError(err)
		s.tx = nil
	}
}

func (s *Suite) SetupSuite() {
	db, err := sqlx.Open(s.DriverName, s.DataSourceName)
	s.Require().NoError(err)
	s.db = db
}

func (s *Suite) TeardownSuite() {
	if err := s.db.Close(); err != nil {
		fmt.Println("error in database close:", err)
	}
}
