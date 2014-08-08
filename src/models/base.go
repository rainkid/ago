package models

import (
	"fmt"
	db "libs/db"
	utils "libs/utils"
	"os"
	"path"
	"strings"
)

type Model struct {
	TableName  string
	PrimaryKey string
	Data       map[string]string

	where string
	args  []interface{}

	limit   string
	orderby string
	groupby string
}

func NewModel(tableName, primaryKey string) *Model {
	return &Model{TableName: tableName, PrimaryKey: primaryKey}
}

//params
func (model *Model) Wherep(id interface{}) *Model {
	model.where = fmt.Sprintf(" WHERE %s = %v ", model.PrimaryKey, id)
	return model
}

func (model *Model) Where(where string, args ...interface{}) *Model {
	model.where, model.args = fmt.Sprintf(" WHERE %s ", where), args
	return model
}

func (model *Model) Limit(start, offset int) *Model {
	model.limit = fmt.Sprintf(" LIMIT %d,%d ", start, offset)
	return model
}

func (model *Model) OrderBy(orderby ...string) *Model {
	model.orderby = " ORDER BY " + strings.Join(orderby, ",") + " "
	return model
}

func (model *Model) GroupBy(groupby string) {
	model.groupby = groupby
}

//select query
func (model *Model) Gets() (results []map[string]interface{}, err error) {
	query := fmt.Sprintf("SELECT * FROM %s%s%s%s", model.GetTable(), model.where, model.orderby, model.limit)
	results, err = model.Db().FetchAll(query, model.args...)
	return results, err
}

func (model *Model) Get() (result map[string]interface{}, err error) {
	query := fmt.Sprintf("SELECT * FROM %s%s%s", model.GetTable(), model.where, model.orderby)
	result, err = model.Db().Fetch(query, model.args...)
	return result, err
}

//count
func (model *Model) Count() (int64, error) {
	var total int64
	query := fmt.Sprintf("SELECT COUNT(*) AS total FROM %s %s", model.GetTable(), model.where)
	row, err := model.Db().FetchRow(query, model.args...)
	if err != nil {
		return total, err
	}
	err = row.Scan(&total)
	if err != nil {
		return total, err
	}
	return total, nil
}

func (model *Model) CleanData() *Model {
	for key := range model.Data {
		delete(model.Data, key)
	}
	return model
}

//insert data
func (model *Model) Insert(mData map[string]string) (int64, error) {
	data, args := model.CookMap(mData, " =?, ", ", ")
	query := fmt.Sprintf("INSERT INTO %s SET %s", model.GetTable(), data)
	result, err := model.Db().Execute(query, args...)
	if err != nil {
		return 0, err
	}

	return result, err
}

//update
func (model *Model) Update(mData map[string]string) (int64, error) {
	data, dargs := model.CookMap(mData, " =?, ", ", ")
	args := utils.MapMerge(dargs, model.args)
	query := fmt.Sprintf("UPDATE %s SET %s%s", model.GetTable(), data, model.where)
	result, err := model.Db().Execute(query, args...)
	if err != nil {
		return 0, err
	}
	return result, err
}

//delete
func (model *Model) Delete() (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s %s", model.GetTable(), model.where)
	result, err := model.Db().Execute(query, model.args...)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (model *Model) CookMap(data map[string]string, sep string, cutset string) (string, []interface{}) {
	var fields []string
	var values []interface{}
	for field, value := range data {
		fields = append(fields, field)
		values = append(values, value)
	}
	return strings.Trim(strings.Join(fields, sep)+sep, cutset), values
}

func (model *Model) GetTable() string {
	return model.TableName
}

func (model *Model) Db() *db.Mysql {
	basepath, _ := os.Getwd()
	file := path.Join(basepath, "src/configs", "mysql.ini")
	return db.NewMysql(file)
}
