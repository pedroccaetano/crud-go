package service

import (
	"github.com/pedroccaetano/crud-go/src/configuration/rest_err"
	"github.com/pedroccaetano/crud-go/src/model"
)

func NewUserDomainService() UserDomaiService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomaiService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
