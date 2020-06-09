// Package assist
package iorm

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const (
	MySQL    = "mysql"
	Postgres = "postgres"
)

var (
	errInsertInvalidType = errors.New("yiigo: invalid data type of InsertSQL() / PGInsertSQL(), expects: struct, *struct, []struct, []*struct, map[string]interface{}, []map[string]interface{}")
	errUpdateInvalidType = errors.New("yiigo: invalid data type of UpdateSQL() / PGUpdateSQL(), expects: struct, *struct, map[string]interface{}")
)

// InsertSQL returns mysql insert sql and binds.
// param data expects: `struct`, `*struct`, `[]struct`, `[]*struct`, `map[string]interface{}`, `[]map[string]interface{}`.
func InsertSQL(table string, data interface{}) (sql string, binds []interface{}) {
	v := reflect.Indirect(reflect.ValueOf(data))

	switch v.Kind() {
	case reflect.Map:
		if x, ok := data.(map[string]interface{}); ok {
			sql, binds = singleInsertWithMap(MySQL, table, x)
		}
	case reflect.Struct:
		sql, binds = singleInsertWithStruct(MySQL, table, v)
	case reflect.Slice:
		count := v.Len()
		if count == 0 {
			return sql, binds
		}
		e := v.Type().Elem()

		switch e.Kind() {
		case reflect.Map:
			x, ok := data.([]map[string]interface{})
			if !ok {
				panic(errInsertInvalidType)
			}
			sql, binds = batchInsertWithMap(MySQL, table, x, count)
		case reflect.Struct:
			sql, binds = batchInsertWithStruct(MySQL, table, v, count)
		case reflect.Ptr:
			if e.Elem().Kind() != reflect.Struct {
				panic(errInsertInvalidType)
			}
			sql, binds = batchInsertWithStruct(MySQL, table, v, count)
		default:
			panic(errInsertInvalidType)
		}
	default:
		panic(errInsertInvalidType)
	}
	return sql, binds
}

// UpdateSQL returns mysql update sql and binds.
// param query expects eg: "UPDATE `table` SET ? WHERE `id` = ?".
// param data expects: `struct`, `*struct`, `map[string]interface{}`.
func UpdateSQL(query string, data interface{}, args ...interface{}) (sql string, binds []interface{}) {
	binds = make([]interface{}, 0)

	v := reflect.Indirect(reflect.ValueOf(data))
	switch v.Kind() {
	case reflect.Map:
		x, ok := data.(map[string]interface{})
		if !ok {
			panic(errUpdateInvalidType)
		}

		sql, binds = updateWithMap(MySQL, query, x, args...)
	case reflect.Struct:
		sql, binds = updateWithStruct(MySQL, query, v, args...)
	default:
		panic(errUpdateInvalidType)
	}

	return sql, binds
}

// PGInsertSQL returns postgres insert sql and binds.
// param data expects: `struct`, `*struct`, `[]struct`, `[]*struct`, `map[string]interface{}`, `[]map[string]interface{}`.
func PGInsertSQL(table string, data interface{}) (sql string, binds []interface{}) {
	binds = make([]interface{}, 0)

	v := reflect.Indirect(reflect.ValueOf(data))
	switch v.Kind() {
	case reflect.Map:
		if x, ok := data.(map[string]interface{}); ok {
			sql, binds = singleInsertWithMap(Postgres, table, x)
		}
	case reflect.Struct:
		sql, binds = singleInsertWithStruct(Postgres, table, v)
	case reflect.Slice:
		count := v.Len()
		if count == 0 {
			return sql, binds
		}

		e := v.Type().Elem()
		switch e.Kind() {
		case reflect.Map:
			x, ok := data.([]map[string]interface{})
			if !ok {
				panic(errInsertInvalidType)
			}
			sql, binds = batchInsertWithMap(Postgres, table, x, count)
		case reflect.Struct:
			sql, binds = batchInsertWithStruct(Postgres, table, v, count)
		case reflect.Ptr:
			if e.Elem().Kind() != reflect.Struct {
				panic(errInsertInvalidType)
			}

			sql, binds = batchInsertWithStruct(Postgres, table, v, count)
		default:
			panic(errInsertInvalidType)
		}
	default:
		panic(errInsertInvalidType)
	}

	return sql, binds
}

// PGUpdateSQL returns postgres update sql and binds.
// param query expects eg: "UPDATE `table` SET $1 WHERE `id` = $2".
// param data expects: `struct`, `*struct`, `map[string]interface{}`.
func PGUpdateSQL(query string, data interface{}, args ...interface{}) (sql string, binds []interface{}) {
	binds = make([]interface{}, 0)

	v := reflect.Indirect(reflect.ValueOf(data))
	switch v.Kind() {
	case reflect.Map:
		x, ok := data.(map[string]interface{})
		if !ok {
			panic(errUpdateInvalidType)
		}
		sql, binds = updateWithMap(Postgres, query, x, args...)
	case reflect.Struct:
		sql, binds = updateWithStruct(Postgres, query, v, args...)
	}

	return sql, binds
}

func singleInsertWithMap(driver string, table string, data map[string]interface{}) (sql string, binds []interface{}) {
	fieldNum := len(data)
	columns := make([]string, 0, fieldNum)
	placeholders := make([]string, 0, fieldNum)
	binds = make([]interface{}, 0, fieldNum)

	switch driver {
	case MySQL:
		for k, v := range data {
			columns = append(columns, fmt.Sprintf("`%s`", k))
			placeholders = append(placeholders, "?")
			binds = append(binds, v)
		}
		sql = fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	case Postgres:
		bindIndex := 0
		for k, v := range data {
			bindIndex++
			columns = append(columns, fmt.Sprintf(`"%s"`, k))
			placeholders = append(placeholders, fmt.Sprintf("$%d", bindIndex))
			binds = append(binds, v)
		}
		sql = fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s) RETURNING "id"`, table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	}

	return sql, binds
}

func singleInsertWithStruct(driver string, table string, v reflect.Value) (sql string, binds []interface{}) {
	fieldNum := v.NumField()

	columns := make([]string, 0, fieldNum)
	placeholders := make([]string, 0, fieldNum)
	binds = make([]interface{}, 0, fieldNum)

	t := v.Type()

	switch driver {
	case MySQL:
		for i := 0; i < fieldNum; i++ {
			column := t.Field(i).Tag.Get("db")
			if column == "-" {
				continue
			}
			if column == "" {
				column = t.Field(i).Name
			}
			columns = append(columns, fmt.Sprintf("`%s`", column))
			placeholders = append(placeholders, "?")
			binds = append(binds, v.Field(i).Interface())
		}

		sql = fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	case Postgres:
		bindIndex := 0
		for i := 0; i < fieldNum; i++ {
			column := t.Field(i).Tag.Get("db")
			if column == "-" {
				continue
			}
			bindIndex++
			if column == "" {
				column = t.Field(i).Name
			}

			columns = append(columns, fmt.Sprintf(`"%s"`, column))
			placeholders = append(placeholders, fmt.Sprintf("$%d", bindIndex))
			binds = append(binds, v.Field(i).Interface())
		}
		sql = fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s) RETURNING "id"`, table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	}

	return sql, binds
}

func batchInsertWithMap(driver string, table string, data []map[string]interface{}, count int) (sql string, binds []interface{}) {
	fieldNum := len(data[0])

	fields := make([]string, 0, fieldNum)
	columns := make([]string, 0, fieldNum)
	placeholders := make([]string, 0, fieldNum)
	binds = make([]interface{}, 0, fieldNum*count)

	switch driver {
	case MySQL:
		for k := range data[0] {
			fields = append(fields, k)
			columns = append(columns, fmt.Sprintf("`%s`", k))
		}

		for _, x := range data {
			phrs := make([]string, 0, fieldNum)
			for _, v := range fields {
				phrs = append(phrs, "?")
				binds = append(binds, x[v])
			}
			placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(phrs, ", ")))
		}

		sql = fmt.Sprintf("INSERT INTO `%s` (%s) VALUES %s", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	case Postgres:
		for k := range data[0] {
			fields = append(fields, k)
			columns = append(columns, fmt.Sprintf(`"%s"`, k))
		}
		bindIndex := 0

		for _, x := range data {
			phrs := make([]string, 0, fieldNum)
			for _, v := range fields {
				bindIndex++
				phrs = append(phrs, fmt.Sprintf("$%d", bindIndex))
				binds = append(binds, x[v])
			}
			placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(phrs, ", ")))
		}
		sql = fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES %s`, table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	}

	return sql, binds
}

func batchInsertWithStruct(driver string, table string, v reflect.Value, count int) (sql string, binds []interface{}) {
	first := reflect.Indirect(v.Index(0))

	fieldNum := first.NumField()
	columns := make([]string, 0, fieldNum)
	placeholders := make([]string, 0, fieldNum)
	binds = make([]interface{}, 0, fieldNum*count)

	t := first.Type()

	switch driver {
	case MySQL:
		for i := 0; i < count; i++ {
			phrs := make([]string, 0, fieldNum)
			for j := 0; j < fieldNum; j++ {
				column := t.Field(j).Tag.Get("db")
				if column == "-" {
					continue
				}
				if i == 0 {
					if column == "" {
						column = t.Field(j).Name
					}
					columns = append(columns, fmt.Sprintf("`%s`", column))
				}
				phrs = append(phrs, "?")
				binds = append(binds, reflect.Indirect(v.Index(i)).Field(j).Interface())
			}
			placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(phrs, ", ")))
		}

		sql = fmt.Sprintf("INSERT INTO `%s` (%s) VALUES %s", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	case Postgres:
		bindIndex := 0

		for i := 0; i < count; i++ {
			phrs := make([]string, 0, fieldNum)
			for j := 0; j < fieldNum; j++ {
				column := t.Field(j).Tag.Get("db")
				if column == "-" {
					continue
				}
				bindIndex++
				if i == 0 {
					if column == "" {
						column = t.Field(j).Name
					}
					columns = append(columns, fmt.Sprintf(`"%s"`, column))
				}
				phrs = append(phrs, fmt.Sprintf("$%d", bindIndex))
				binds = append(binds, reflect.Indirect(v.Index(i)).Field(j).Interface())
			}
			placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(phrs, ", ")))
		}

		sql = fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES %s`, table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	}

	return sql, binds
}

func updateWithMap(driver string, query string, data map[string]interface{}, args ...interface{}) (sql string, binds []interface{}) {
	dataLen := len(data)
	argsLen := len(args)
	binds = make([]interface{}, 0, dataLen+argsLen)
	sets := make([]string, 0, dataLen)

	switch driver {
	case MySQL:
		for k, v := range data {
			sets = append(sets, fmt.Sprintf("`%s` = ?", k))
			binds = append(binds, v)
		}
		sql = strings.Replace(query, "?", strings.Join(sets, ", "), 1)
		binds = append(binds, args...)
	case Postgres:
		bindIndex := 0
		for k, v := range data {
			bindIndex++
			sets = append(sets, fmt.Sprintf(`"%s" = $%d`, k, bindIndex))
			binds = append(binds, v)
		}

		oldnew := make([]string, 0, argsLen*2)
		for i := 1; i <= argsLen; i++ {
			oldnew = append(oldnew, fmt.Sprintf("$%d", i+1), fmt.Sprintf("$%d", dataLen+i))
		}

		r := strings.NewReplacer(oldnew...)
		query = r.Replace(query)
		sql = strings.Replace(query, "$1", strings.Join(sets, ", "), 1)
		binds = append(binds, args...)
	}

	return sql, binds
}

func updateWithStruct(driver string, query string, v reflect.Value, args ...interface{}) (sql string, binds []interface{}) {
	fieldNum := v.NumField()
	argsLen := len(args)
	sets := make([]string, 0, fieldNum)
	binds = make([]interface{}, 0, fieldNum+argsLen)
	t := v.Type()

	switch driver {
	case MySQL:
		for i := 0; i < fieldNum; i++ {
			column := t.Field(i).Tag.Get("db")
			if column == "-" {
				continue
			}
			if column == "" {
				column = t.Field(i).Name
			}

			sets = append(sets, fmt.Sprintf("`%s` = ?", column))
			binds = append(binds, v.Field(i).Interface())
		}

		sql = strings.Replace(query, "?", strings.Join(sets, ", "), 1)
		binds = append(binds, args...)
	case Postgres:
		bindIndex := 0

		for i := 0; i < fieldNum; i++ {
			column := t.Field(i).Tag.Get("db")
			if column == "-" {
				continue
			}
			bindIndex++
			if column == "" {
				column = t.Field(i).Name
			}
			sets = append(sets, fmt.Sprintf(`"%s" = $%d`, column, bindIndex))
			binds = append(binds, v.Field(i).Interface())
		}

		oldnew := make([]string, 0, argsLen*2)
		for i := 1; i <= argsLen; i++ {
			oldnew = append(oldnew, fmt.Sprintf("$%d", i+1), fmt.Sprintf("$%d", bindIndex+i))
		}

		r := strings.NewReplacer(oldnew...)
		query = r.Replace(query)
		sql = strings.Replace(query, "$1", strings.Join(sets, ", "), 1)
		binds = append(binds, args...)
	}
	return sql, binds
}
