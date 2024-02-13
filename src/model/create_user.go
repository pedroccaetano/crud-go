package model

import (
	"github.com/pedroccaetano/crud-go/src/configuration/logger"
	"github.com/pedroccaetano/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("jourmey", "createUser"))

	ud.EncryptPassword()

	return nil
}
