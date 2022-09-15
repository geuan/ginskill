package UserModel

type UserModelAttrFunc func(u *UserModelImpl)

type UserModelAttrFuncs []UserModelAttrFunc

func WithUserID(id int) UserModelAttrFunc {
	return func(u *UserModelImpl) {
		u.UserId = id
	}
}

func WithUserName(username string) UserModelAttrFunc {
	return func(u *UserModelImpl) {
		u.UserName = username
	}
}

func (this UserModelAttrFuncs) Apply(u *UserModelImpl) {
	for _, f := range this {
		f(u)
	}
}
