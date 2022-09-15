package handlers

import (
	"ginskill/data/Getter"
	"ginskill/data/setter"
	"ginskill/models/UserModel"
	"ginskill/result"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context)  {
	//user := UserModel.New()
	//result.Result(c.ShouldBind(user)).Unwrap()
	//if user.UserId > 10 {
	//	R(c)("userlist","100001","userlist")(OK)
	//} else {
	//	R(c)("userlist","100001","userlist")(Error)
	//}

	R(c)("userlist","100001",Getter.UserGetter.GetUserList())(OK)
}

func UserDetail(c *gin.Context)  {
	id := &struct {Id int `uri:"id" binding:"required,gt=1"`}{}   // 代码参考
	result.Result(c.ShouldBindUri(id)).Unwrap()
	R(c)("userdetail","100002",Getter.UserGetter.GetUserByID(id.Id).Unwrap())(OK)

}

func UserSave(c *gin.Context)  {
	u := UserModel.New()
	result.Result(c.ShouldBindJSON(u)).Unwrap()
	R(c)("saveuser","100003",setter.UserSetter.SaveUser(u).Unwrap())(OK)

}