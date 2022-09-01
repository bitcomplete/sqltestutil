package sqltestutil_test

import (
	"context"
	"github.com/mohammad-ahmadi-de/sqltestutil"
	"testing"
)

func TestStartPostgresContainer(t *testing.T) {
	c, err := sqltestutil.StartPostgresContainer(context.Background(), sqltestutil.WithPort(5321))
	if err != nil {
		t.Fatal(err)
	}
	c.Shutdown(context.Background())
}
