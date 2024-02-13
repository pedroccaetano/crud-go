package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroccaetano/crud-go/src/configuration/logger"
	"github.com/pedroccaetano/crud-go/src/configuration/validation"
	"github.com/pedroccaetano/crud-go/src/controller/model/request"
	"github.com/pedroccaetano/crud-go/src/model"
	"github.com/pedroccaetano/crud-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInteface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertDomainToResponse(domain)

	logger.Info("user created successfully",
		zap.String("journey", "createUser"),
		zap.Any("response", response),
	)

	c.JSON(http.StatusCreated, response)
}
