package service

import (
	"context"
	"gin_consult/dao"
	"gin_consult/pkg/e"
	"gin_consult/serializer"
	logging "github.com/sirupsen/logrus"
)

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (service *ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}