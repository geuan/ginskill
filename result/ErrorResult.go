package result

import (
	"fmt"
	"ginskill/validators"
)

type ErrorResult struct {
	data interface{}
	err error
}

func (this *ErrorResult) Unwrap() interface{}  {
	if this.err != nil {
		validators.CheckErrors(this.err)  // 匹配到走这一步
		panic(this.err.Error())  // 没有匹配到走这一步
	}
	return  this.data
}

func Result(vs ...interface{}) *ErrorResult  {
	if len(vs) ==  1 {
		if vs[0] == nil {
			return &ErrorResult{data:nil,err:nil}
		}
		if e,ok := vs[0].(error);ok {
			return &ErrorResult{data:nil,err:e}
		}
	}
	if len(vs) ==  2 {
		if vs[1] == nil {
			return &ErrorResult{vs[0],nil}
		}
		if e,ok := vs[1].(error);ok {
			return &ErrorResult{vs[0],e}
		}
	}
	return &ErrorResult{nil,fmt.Errorf("error result format")}
}

