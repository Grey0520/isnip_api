package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	UserID       uint64 `json:"user_id,string" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
	UserName     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	AccessToken  string
	RefreshToken string
}

// UnmarshalJSON 为User类型实现自定义的UnmarshalJSON方法
func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		u.UserName = required.UserName
		u.Password = required.Password
	}
	return
}

type RegisterForm struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UnmarshalJSON 为RegisterForm类型实现自定义的UnmarshalJSON方法
func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else if required.Password != required.ConfirmPassword {
		err = errors.New("两次密码不一致")
	} else {
		r.UserName = required.UserName
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}
	return
}
