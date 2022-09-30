package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BcChat struct {
	gorm.Model
	SenderMemberNumber string `gorm:"type:varchar(30);not null;index:ip_senderMemberNumber"`
	SendDate           time.Time
	NickName           string `gorm:"type:varchar(100);index:ip_nickName"`
	RoomName           string `gorm:"type:varchar(100);index:ip_roomName"`
	SendMessage        string `gorm:"type:text;"`
	SenderCode         string `gorm:"type:varchar(100)"`
}

func (bcChat *BcChat) IsEmpty() bool {
	if bcChat == nil {
		return true
	}
	if bcChat.ID == 0 {
		return true
	}
	return false
}
