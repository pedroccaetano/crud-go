package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pedroccaetano/crud-go/src/configuration/validation"
	"github.com/pedroccaetano/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
