package sqltestutil

import (
	"context"
	"database/sql/driver"
	"io/ioutil"
	"path/filepath"
	"sort"
)

// RunMigrations reads all of the files matching *.up.sql in migrationDir and
// executes them in lexicographical order against the provided db. A typical
// convention is to use a numeric prefix for each new migration, e.g.:
//
//   001_create_users.up.sql
//   002_create_posts.up.sql
//   003_create_comments.up.sql
//
// Note that this function does not check whether the migration has already been
// run. Its primary purpose is to initialize a test database.
func RunMigrations(ctx context.Context, db driver.ExecerContext, migrationDir string, files ...string) error {
	filenames, err := filepath.Glob(filepath.Join(migrationDir, "*.up.sql"))
	if err != nil {
		return err
	}
	var filter map[string]struct{} = nil
	if len(files) > 0 {
		filter = make(map[string]struct{})
		for i := range files {
			filter[files[i]] = struct{}{}
		}
	}
	sort.Strings(filenames)
	for _, filename := range filenames {
		if len(files) > 0 {
			if _, exist := filter[filepath.Base(filename)]; !exist {
				continue
			}
		}
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		_, err = db.ExecContext(ctx, string(data), nil)
		if err != nil {
			return err
		}
	}
	return nil
}
