package service

import (
	"awesomeProject/model"
	"awesomeProject/serializer"
)

type BcWearService struct {
}

func (service *BcWearService) Create(wear *model.BcWear) error {
	client := model.DbClient
	err := client.Create(&wear).Error
	if err != nil {
		return err
	}
	return nil
}
func (service *BcWearService) Modify(wear *model.BcWear) error {
	client := model.DbClient
	err := client.Save(&wear).Error
	if err != nil {
		return err
	}
	return nil
}
func (service *BcWearService) List(wear *model.BcWear) (wears []*model.BcWear, count int) {
	model.DbClient.Where(wear).Find(&wears).Count(&count)
	return
}
func (service *BcWearService) Page(page *serializer.PageParam, wear *model.BcWear) (wears []*model.BcWear, count int) {
	model.DbClient.Where(wear).Offset(page.Begin()).Limit(page.End()).Find(&wears).Count(&count)
	return
}
func (service *BcWearService) OneById(id int) (wear *model.BcWear) {
	wear = &model.BcWear{}
	model.DbClient.Where("id=?", id).First(wear)
	return wear
}
func (service *BcWearService) Remove(id int) error {
	err := model.DbClient.Where("id=?", id).Delete(model.BcWear{}).Error
	if err != nil {
		return err
	}
	return nil
}
