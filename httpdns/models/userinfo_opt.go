package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id          int
	Username    string
	Password    string
	Lasttime    time.Time
	Phonenumber string
}

func User_search(key string) (User, error) {
	var user User
	o := orm.NewOrm()

	err := o.Raw("SELECT id, username, password, lasttime, phonenumber FROM userinfo WHERE username = ?", key).QueryRow(&user)
	if err == nil {
		return user, nil
	} else {
		return user, err
	}

}

func User_update(user User) error {
	o := orm.NewOrm()

	res, err := o.Raw("UPDATE userinfo SET lasttime = ? where username = ? ", user.Lasttime, user.Username).Exec()

	if err == nil {
		_, _ = res.RowsAffected()
		return nil
	} else {
		return errors.New("Update error")
	}

}

func User_insert(user User) error {
	o := orm.NewOrm()

	_, err := o.Raw("INSERT INTO userinfo (username, password, lasttime, phonenumber) VALUE(?,?,?,?);", user.Username, user.Password, user.Lasttime, user.Phonenumber).Exec()
	if err != nil {
		fmt.Println(err)
		return errors.New("Insert error!")
	} else {
		return nil
	}
}
