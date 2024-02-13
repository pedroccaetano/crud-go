package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroccaetano/crud-go/src/configuration/logger"
	"github.com/pedroccaetano/crud-go/src/configuration/validation"
	"github.com/pedroccaetano/crud-go/src/controller/model/request"
	"github.com/pedroccaetano/crud-go/src/model"
	"github.com/pedroccaetano/crud-go/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInteface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controoler",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "createUser"),
		)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user created successfully",
		zap.String("journey", "createUser"),
	)

	c.String(http.StatusCreated, "")
}
