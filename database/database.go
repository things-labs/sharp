package database

import (
	"fmt"

	"github.com/thinkgos/go-core-package/lib/univ"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config 数据库配置
type Config struct {
	Dialect  string            `yaml:"dialect" json:"dialect"` // mysql sqlite3 postgres
	Username string            `yaml:"username" json:"username"`
	Password string            `yaml:"password" json:"password"`
	Protocol string            `yaml:"protocol" json:"protocol"`
	Host     string            `yaml:"host" json:"host"`
	Port     string            `yaml:"port" json:"port"`
	DbName   string            `yaml:"dbName" json:"dbName"`
	Extend   map[string]string `yaml:"extend" json:"extend"`
	LogMode  bool              `yaml:"logMode" json:"logMode"`
}

func New(c Config, config *gorm.Config, dialectorNews ...func(c Config) gorm.Dialector) (*gorm.DB, error) {
	var dialect gorm.Dialector

	switch c.Dialect {
	case "mysql":
		values := make(univ.Values)
		values.Add("charset", "utf8mb4")
		values.Add("parseTime", "True")
		values.Add("loc", "Local")
		for k, v := range c.Extend {
			values.Add(k, v)
		}
		dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s",
			c.Username, c.Password, c.Protocol, c.Host, c.Port, c.DbName, values.Encode("=", "&")) // DSN data source name
		dialect = mysql.New(mysql.Config{
			DSN: dsn,
			// DefaultStringSize:         256,   // string 类型字段的默认长度
			// DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			// DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			// DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			// SkipInitializeWithVersion: false, // 根据版本自动配置
		})
	case "postgres":
		values := make(univ.Values)
		values.Add("user", c.Username)
		values.Add("password", c.Password)
		values.Add("host", c.Host)
		values.Add("port", c.Port)
		values.Add("dbname", c.DbName)
		for k, v := range c.Extend {
			values.Add(k, v)
		}
		dsn := values.Encode("=", " ")
		dialect = postgres.New(postgres.Config{
			DSN: dsn,
		})
	case "sqlite3":
		dialect = newSqlite3(c.DbName)
	case "extend":
		if len(dialectorNews) == 0 {
			panic("select option dialector should give a dialector new function")
		}
		dialectorNew := dialectorNews[0]
		dialect = dialectorNew(c)
	default:
		panic("please select database driver one of [mysql|postgres|sqlite3|extend], if use sqlite3, build tags with sqlite3!")
	}
	return gorm.Open(dialect, config)
}
