package dao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	return &UserDao{BaseDao{TableName: "admin_user", PrimaryKey: "id"}}
}
