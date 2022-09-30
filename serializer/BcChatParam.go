package serializer

import (
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
	"time"
)

type BcChatParam struct {
	ID                 uint      `json:"id"`
	SenderMemberNumber string    `json:"sender_member_number"`
	SendDate           time.Time `json:"send_date"`
	NickName           string    `json:"nick_name"`
	RoomName           string    `json:"room_name"`
	SendMessage        string    `json:"send_message"`
	SenderCode         string    `json:"sender_code"`
	PageParam
}

func (query *BcChatParam) SerializeParam() bool {
	return true
}

func (query *BcChatParam) BuildDBData(point *model.BcChat) (err error) {
	point = &model.BcChat{
		Model: gorm.Model{
			ID: query.ID,
		},
		SenderMemberNumber: query.SenderMemberNumber,
		SendDate:           query.SendDate,
		NickName:           query.NickName,
		RoomName:           query.RoomName,
		SendMessage:        query.SendMessage,
		SenderCode:         query.SenderCode,
	}
	return
}

func SerializeBcChat(bcChat *model.BcChat) (query *BcChatParam, err error) {
	query = &BcChatParam{
		ID:                 bcChat.ID,
		SenderMemberNumber: bcChat.SenderMemberNumber,
		SendDate:           bcChat.SendDate,
		NickName:           bcChat.NickName,
		RoomName:           bcChat.RoomName,
		SendMessage:        bcChat.SendMessage,
		SenderCode:         bcChat.SenderCode,
	}
	return query, nil
}
