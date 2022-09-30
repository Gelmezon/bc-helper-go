package service

import (
	"awesomeProject/model"
	"awesomeProject/serializer"
)

type BcChatService struct {
}

func (service *BcChatService) Create(chat *model.BcChat) error {
	client := model.DbClient
	err := client.Create(&chat).Error
	if err != nil {
		return err
	}
	return nil
}
func (service *BcChatService) Modify(chat *model.BcChat) error {
	client := model.DbClient
	err := client.Save(&chat).Error
	if err != nil {
		return err
	}
	return nil
}
func (service *BcChatService) List(chat *model.BcChat) (chats []*model.BcChat, count int) {
	model.DbClient.Where(chat).Find(&chats).Count(&count)
	return
}
func (service *BcChatService) Page(page *serializer.PageParam, chat *model.BcChat) (chats []*model.BcChat, count int) {
	model.DbClient.Where(chat).Offset(page.Begin()).Limit(page.End()).Find(&chats).Count(&count)
	return
}
func (service *BcChatService) OneById(id int) (chat *model.BcChat) {
	chat = &model.BcChat{}
	model.DbClient.Where("id=?", id).First(chat)
	return chat
}
func (service *BcChatService) Remove(id int) error {
	err := model.DbClient.Where("id=?", id).Delete(model.BcChat{}).Error
	if err != nil {
		return err
	}
	return nil
}
