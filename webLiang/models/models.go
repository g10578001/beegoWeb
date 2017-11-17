package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Username string `orm:"size(256);unique"`
	Email    string `orm:"size(256);unique"`
	Password string `orm:"size(256);unique"`
}

func (usr *User) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}

func init() {
	orm.RegisterModel(new(User))
}
