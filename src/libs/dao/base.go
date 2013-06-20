package dao

import (
	"fmt"
	db "libs/db"
	"os"
	"path"
	"strings"
)

type BaseDao struct {
	TableName  string
	PrimaryKey string

	where string
	args  []interface{}

	limit   string
	fields  string
	orderby string
	groupby string
}

func NewBaseDao() *BaseDao {
	return &BaseDao{}
}

//params
func (dao *BaseDao) Wherep(id int) {
	dao.where = fmt.Sprintf(" WHERE %s = %d ", dao.PrimaryKey, id)
}

func (dao *BaseDao) Where(where string, args ...interface{}) *BaseDao {
	dao.where, dao.args = fmt.Sprintf(" WHERE %s ", where), args
	return dao
}

func (dao *BaseDao) Limit(start, offset int) *BaseDao {
	dao.limit = fmt.Sprintf(" LIMIT %d,%d ", start, offset)
	return dao
}

func (dao *BaseDao) OrderBy(orderby ...string) *BaseDao {
	dao.orderby = " ORDER BY " + strings.Join(orderby, ",") + " "
	return dao
}

func (dao *BaseDao) GroupBy(groupby string) {
	dao.groupby = groupby
}

//select query
func (dao *BaseDao) Gets() (results []map[string][]byte, err error) {
	query := fmt.Sprintf("SELECT * FROM %s%s%s%s", dao.GetTable(), dao.where, dao.orderby, dao.limit)
	results, err = dao.Db().FetchAll(query, dao.args...)
	return results, err
}

func (dao *BaseDao) Get() (result map[string][]byte, err error) {
	query := fmt.Sprintf("SELECT * FROM %s%s%s", dao.GetTable(), dao.where, dao.orderby)
	result, err = dao.Db().Fetch(query, dao.args...)
	return result, err
}

//count
func (dao *BaseDao) Count() (int64, error) {
	var total int64
	query := fmt.Sprintf("SELECT COUNT(*) AS total FROM %s %s", dao.GetTable(), dao.where)
	row, err := dao.Db().FetchRow(query, dao.args...)
	if err != nil {
		return total, err
	}
	err = row.Scan(&total)
	if err != nil {
		return total, err
	}
	return total, nil
}

func (dao *BaseDao) SetData(fields string, args ...interface{}) *BaseDao {
	dao.fields, dao.args = fields, args
	return dao
}

//update 
func (dao *BaseDao) Update() (int64, error) {
	query := fmt.Sprintf("UPDATE %s SET %s", dao.GetTable(), dao.fields)
	result, err := dao.Db().Execute(query, dao.args...)
	if err != nil {
		return 0, err
	}
	return result, err
}

//delete 
func (dao *BaseDao) Delete() (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", dao.GetTable(), dao.where)
	result, err := dao.Db().Execute(query, dao.args...)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (dao *BaseDao) GetTable() string {
	return dao.TableName
}

func (dao *BaseDao) Db() *db.Mysql {
	basepath, _ := os.Getwd()
	file := path.Join(basepath, "src/configs", "mysql.yaml")
	return db.NewMysql(file)
}
