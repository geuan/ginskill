package Getter

import (
	"fmt"
	"ginskill/data/mappers"
	"ginskill/models/UserModel"
	"ginskill/result"
)


var UserGetter IUserGetter

func init()  {
	UserGetter = NewUserGetterImpl()
	
}

type IUserGetter interface {
	GetUserList() []*UserModel.UserModelImpl
	GetUserByID(id int) *result.ErrorResult
}

type UserGetterImpl struct {
	userMapper *mappers.UserMapper
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{userMapper:&mappers.UserMapper{}}
}

func (this *UserGetterImpl) GetUserList() (users []*UserModel.UserModelImpl) {
	//dbs.Orm.Find(&users)
	sqlMapper := this.userMapper.GetUserList()
	//dbs.Orm.Raw(sqlMapper.Sql,sqlMapper.Args...).Find(&users)   // 注意这里要加三个点，解压缩切片
	sqlMapper.Query().Find(&users)
	return

}


func (this *UserGetterImpl)   GetUserByID(id int) *result.ErrorResult {
	user := UserModel.New()
	db := this.userMapper.GetUserDetail(id).Query().Find(&user)
	if db.Error != nil || db.RowsAffected == 0 {
		return result.Result(nil,fmt.Errorf("not found user,id=%d",id))
	}
	return result.Result(user,nil)
}




