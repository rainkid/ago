package service

import (
	dao "libs/dao"
)

type BaseService struct {
}

func (base *BaseService) GetList(page, perpage int, where string, args ...interface{}) (int64, []map[string][]byte) {
	if page < 1 {
		page = 1
	}
	start := (page - 1) * perpage
	result, err := base.GetDao().Where(where, args...).Limit(start, perpage).Gets()
	if err != nil {
		return 0, nil
	}
	total, err := base.Count(where, args...)
	if err != nil {
		return 0, nil
	}
	return total, result
}

func (base *BaseService) Get(where string, args ...interface{}) (map[string][]byte, error) {
	return base.GetDao().Where(where, args...).Get()
}

func (base *BaseService) Count(where string, args ...interface{}) (int64, error) {
	return base.GetDao().Where(where, args...).Count()
}

func (base *BaseService) GetDao() *BaseDao {
	return dao.NewBaseDao()
}

func (base *BaseService) CookMap(data map[string]interface{}, sep string) (string, []interface{}) {
	var fields []string
	var values []interface{}
	for field, value := range data {
		fields = append(fields, field)
		values = append(values, value)
	}
	return strings.Join(fields, sep) + sep, values
}
