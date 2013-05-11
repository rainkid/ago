package models

import (
	"database/sql"
	"fmt"
	mdb "libs/db"
	"os"
	"path"
	"strings"
)

type DataMap map[string]interface{}

type ModelInterface interface {
	IsValid() bool
	Valid() (int64, string)
}

type Model struct {
	ModelInterface
	TableName string
	Data      DataMap
	Filters   DataMap
}

func (m *Model) InitData() {
	if m.Data == nil {
		m.Data = make(DataMap)
	}
}

func (m *Model) InitFilter() {
	if m.Filters == nil {
		m.Filters = make(DataMap)
	}
}

func (m *Model) Set(key string, value interface{}) *Model {
	m.InitData()
	m.Data[key] = value
	return m
}

func (m *Model) Sets(values DataMap) *Model {
	m.InitData()
	if len(values) > 0 {
		for key, value := range values {
			m.Data[key] = value
		}
	}
	return m
}

func (m *Model) Where(key string, value interface{}) *Model {
	m.InitFilter()
	m.Filters[key] = value
	return m
}

func (m *Model) Add() (int64, error) {
	db := m.DBHandler()

	data := m.SqlWhere(m.Data, ",")
	sql := fmt.Sprintf("INSERT INTO %s SET %s", m.TableName, data)

	res, err := db.Execute(sql)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (m *Model) Get() (*sql.Row, error) {
	db := m.DBHandler()

	where := m.SqlWhere(m.Filters, " AND")
	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT 0,1 ", m.TableName, where)
	row, err := db.Fetch(sql)
	return row, err
}

func (m *Model) Gets() (*sql.Rows, error) {
	db := m.DBHandler()

	where := m.SqlWhere(m.Filters, " AND")
	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s", m.TableName, where)

	rows, err := db.FetchAll(sql)
	return rows, err
}

func (m *Model) Delete() (int64, error) {
	db := m.DBHandler()

	where := m.SqlWhere(m.Filters, " AND")
	sql := fmt.Sprintf("DELETE FROM %s WHERE %s", m.TableName, where)

	res, err := db.Execute(sql)
	if err != nil {
		return 0, nil
	}
	affect, err := res.RowsAffected()
	return affect, err

}

func (m *Model) SqlWhere(dm DataMap, split string) string {
	strs := []string{}
	for k, v := range dm {
		strs = append(strs, fmt.Sprintf("%s %s='%s'", split, k, v))
	}
	return strings.Trim(strings.Trim(strings.Trim(strings.Join(strs, ""), " "), ","), split)
}

func (m *Model) Valid() (int64, string) {
	return 0, ""
}

func (m *Model) GetData() DataMap {
	m.InitData()
	return m.Data
}

func (m *Model) DBHandler() *mdb.Mysql {
	basepath, _ := os.Getwd()
	file := path.Join(basepath, "src/configs", "db.yaml")
	return &mdb.Mysql{ConfigFile: file}
}

func (m *Model) IsValid() bool {
	if code, _ := m.Valid(); code != 0 {
		return false
	}
	return true
}
