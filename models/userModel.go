package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

const  TABLENAME_USER="user"
type User struct {
	Id int
	Username string
	Password string
	Power string
	Createtime string
	Logindateid string

}

func (u *User) TableName() string {
	return TableName(TABLENAME_USER)
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func FindByNameAndPwd(userNamem string,passWrod string)( *User,error) {
	u:=new(User)
	err := orm.NewOrm().QueryTable(TableName(TABLENAME_USER)).Filter("username",userNamem).Filter("password",passWrod).One(u)
	if err!=nil{
		panic(nil)
		return nil, err

	}
	return u,err
}
func UserGetById(id int) (*User, error) {
	u := new(User)

	err := orm.NewOrm().QueryTable(TableName(TABLENAME_USER)).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserAdds(u *User) (int64, error) {
	if u.Username==""{
		fmt.Println("用户名不能为空")
		return 0,fmt.Errorf("用户名不能为空")
	}
	if u.Password==""{
		fmt.Println("密码不能为空")
		return 0,fmt.Errorf("密码不能为空")
	}
	if u.Power==""{
		fmt.Println("权限为空")
		return 0,fmt.Errorf("权限为空")
	}
	if u.Createtime==""{
		fmt.Println("创建时间为空")
		return 0,fmt.Errorf("创建时间为空")
	}

	return orm.NewOrm().Insert(u)
}
