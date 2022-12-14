package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

type JsonResult struct {
	Message string   `json:"message"`
	Code 	string 	`json:"code"`
	Result 	interface{}  `json:"result"`
}

func NewJsonResult(message string, code string, result interface{}) *JsonResult {
	return &JsonResult{Message: message, Code: code, Result: result}
}


var  ResultPool *sync.Pool

func init()  {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJsonResult("", "",nil)
		},
	}
}



// 借鉴学习，思维方式
type ResultFunc  func(message string,code string,result interface{}) func(output Output)
type Output func(c *gin.Context,v interface{})

func R(c *gin.Context)  ResultFunc {
	return func(message string, code string, result interface{})  func(output Output) {  // 装饰器
		r := ResultPool.Get().(*JsonResult)
		defer ResultPool.Put(r)
		r.Message= message
		r.Code = code
		r.Result = result
		//c.JSON(200,r)
		return func(output Output) {
			output(c,r)
		}
	}
}

func OK(c *gin.Context,v interface{})  {
	c.JSON(200,v)
}

func Error(c *gin.Context, v interface{})  {
	c.JSON(400,v)
}

func OK2String(c *gin.Context,v interface{})  {
	c.String(200,fmt.Sprintf("%v",v))

}