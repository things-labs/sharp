// +build sqlite3

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSqlite3(source string) gorm.Dialector {
	return sqlite.Open(source)
}
