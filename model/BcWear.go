package model

import (
	"github.com/jinzhu/gorm"
)

type BcWear struct {
	gorm.Model
	UpdateMemberNumber   string `gorm:"type:varchar(30);not null;index:ip_updateMemberNumber"`
	UpdateMemberNickName string `gorm:"type:varchar(100);"`
	WearData             string `gorm:"type:text;"`
	WearName             string `gorm:"type:varchar(100);not null;index:ip_wearName"`
	AutoSave             bool   `gorm:"type:boolean;"`
	Md5                  string `gorm:"type:varchar(100);index:md5_idx"`
	RandomCode           string `gorm:"type:varchar(100)"`
}

func (bcWear *BcWear) IsEmpty() bool {
	if bcWear == nil {
		return true
	}
	if bcWear.ID == 0 {
		return true
	}
	return false
}
