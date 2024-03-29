// Copyright (c) 2023 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package models

import (
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/exhash"
	"github.com/ergoapi/util/exid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type User struct {
	Model
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
	Banned   bool   `gorm:"column:banned" json:"banned"`
	Token    string `gorm:"column:token" json:"token"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "system_user"
}

func init() {
	Migrate(User{})
}

func (u *User) Save() error {
	var uu User
	err := GDB.Model(User{}).Where("username = ?", u.Username).Last(&uu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if uu.ID > 0 {
		uu.Role = u.Role
		uu.Banned = u.Banned
		uu.Email = u.Email
		return GDB.Save(&uu).Error
	}
	return GDB.Create(&u).Error
}

func (u *User) New() *User {
	tx := GDB.Save(&u)
	if tx.Error != nil {
		logrus.Debugf("添加人员save: %v, err: %v", u.Username, tx.Error.Error())
		return nil
	}
	var nw User
	if err := tx.Row().Scan(&nw); err != nil {
		logrus.Debugf("添加人员scan: %v, err: %v", u.Username, err.Error())
		return nil
	}
	return &nw
}

func (u *User) Exist() bool {
	tx := GDB.Model(User{}).Where("username = ?", u.Username).Find(&User{})
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return false
	}
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// InitAdmin init
func InitAdmin() {
	val, err := ConfigsGet("initadmin")
	if err != nil {
		logrus.Fatalf("cannot query initadmin err: %v", err)
	}
	if val != "" {
		logrus.Infof(color.SGreen("exist initadmin %v success...", val))
		return
	}
	user := viper.GetString("server.admin.user")
	adminuser := User{
		Username: user,
		Password: exhash.MD5(viper.GetString("server.admin.pass")),
		Email:    viper.GetString("server.admin.mail"),
		Banned:   false,
		Token:    exid.GenUID(user),
	}
	if err := adminuser.Save(); err != nil {
		logrus.Fatalf("init admin in mysql err: %v", err)
	}
	err = ConfigsSet("initadmin", "done")
	if err != nil {
		logrus.Fatalf("init initadmin in mysql err: %v", err)
	}
	logrus.Infof(color.SGreen("init  admin %v success...", user))
}
