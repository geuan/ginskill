package validators

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	// 注册 自定义验证
	registerValidator("UserName", UserName("required,min=4").ToFunc())
}

type UserName string

func (this UserName) ToFunc() validator.Func {
	validatorError["UserName"]="用户名必须要在4-8位之间"
	return func(fl validator.FieldLevel) bool {
		v, ok := fl.Field().Interface().(string)
		if ok {
			return this.validate(v)
		}
		return false
	}
}

func (this UserName) validate(v string) bool {
	if err := myvalid.Var(v, string(this)); err != nil {
		return false
	}
	if len(v) > 8 {
		return false
	}
	return true
}

//var VUserName validator.Func = func(fl validator.FieldLevel) bool {
//	uname, ok := fl.Field().Interface().(string)
//	if ok && len(uname) >= 4 {
//		return true
//	}
//	return  false
//}
