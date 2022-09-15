package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetInfo(id int) (gin.H, error) {
	if id > 10 {
		return gin.H{"messgae": "test"}, nil
	} else {
		return nil, fmt.Errorf("test error")
	}

}
