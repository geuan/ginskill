package mappers

import (
	"ginskill/dbs"
	"gorm.io/gorm"
)

type SqlMapper struct {
	Sql string     //  SQL语句
	Args []interface{}   // 参数集合
}

func NewSqlMapper(sql string,args []interface{}) *SqlMapper  {
	return &SqlMapper{Sql:sql,Args:args}
}

// 查询
func (this *SqlMapper) Query() *gorm.DB  {
	return dbs.Orm.Raw(this.Sql,this.Args...)
}

// update delete insert
func (this *SqlMapper) Exec() *gorm.DB  {
	return dbs.Orm.Exec(this.Sql,this.Args)
}

func Mapper(sql string,args []interface{},err error) *SqlMapper  {
	if err != nil {
		panic(err.Error())
	}
	return NewSqlMapper(sql,args)
}
