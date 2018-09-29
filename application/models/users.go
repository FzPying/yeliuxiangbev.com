package models

import "github.com/fzpying/beego/orm"

type Users struct {
	//用户序号
	Id   int
	Name string `orm:"size(100)"`
}


//新建用户
func AddUser(user *Users) (int64, error) {
	o := orm.NewOrm()             //数据库
	userId, err := o.Insert(user) //插入数据
	return userId, err
}

//查询账号
func GetUserById(userId int64) (*Users, error) {
	o := orm.NewOrm()                        //数据库
	user := new(Users)                       //TUser就是第9行struct的数据库，就是mysql的表
	qs := o.QueryTable("t_user")             //表名为t_user
	err := qs.Filter("id", userId).One(user) //One是指查询一条数据，One(user)是查询mysql表中一条数据
	return user, err
}

//手机号查询账号
func GetUserByTel(tel string) (*Users, error) {
	o := orm.NewOrm()
	user := new(Users)                     //TUser就是第9行struct的数据库，就是mysql的表
	qs := o.QueryTable("t_user")           //表名为t_user
	err := qs.Filter("tel", tel).One(user) //One是指查询一条数据，One(user)是查询mysql表中一条数据
	return user, err
}

//微信Id查询账号
func GetUserByVid(vid int64) (*Users, error) {
	o := orm.NewOrm()
	user := new(Users)                     //TUser就是第9行struct的数据库，就是mysql的表
	qs := o.QueryTable("t_user")           //表名为t_user
	err := qs.Filter("vid", vid).One(user) //One是指查询一条数据，One(user)是查询mysql表中一条数据
	return user, err
}
