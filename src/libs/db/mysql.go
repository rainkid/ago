package db

import (
	"database/sql"
	"dogo"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

type Mysql struct {
	ConfigFile string
}

func (m *Mysql) FetchAll(s string) (*sql.Rows, error) {
	dsn := m.Dsn()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func (m *Mysql) Fetch(s string) (*sql.Row, error) {
	dsn := m.Dsn()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("select * from admin_user limit 0,1")
	return row, nil
}

func (m *Mysql) Execute(s string) (sql.Result, error) {
	dsn := m.Dsn()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(s)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec()
	if err != nil {
		return nil, err
	}
	return res, err
}

func (m *Mysql) Dsn() string {
	config, _ := dogo.NewConfig(m.ConfigFile)
	env, _ := config.String("base", "env")

	username, _ := config.String(env, "username")
	password, _ := config.String(env, "password")
	host, _ := config.String(env, "host")
	port, _ := config.Int(env, "port")
	dbname, _ := config.String(env, "dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, dbname)
	return dsn
}
