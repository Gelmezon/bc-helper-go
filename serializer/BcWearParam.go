package serializer

import (
	"awesomeProject/model"
	"awesomeProject/support"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type BcWearParam struct {
	ID                   uint             `form:"id" json:"id"`
	UpdateMemberNumber   string           `form:"updateMemberNumber" json:"updateMemberNumber"`
	UpdateMemberNickName string           `form:"updateMemberNickName" json:"updateMemberNickName"`
	WearName             string           `form:"wearName" json:"wearName"`
	WearData             *json.RawMessage `form:"wearData" json:"wearData"`
	AutoSave             bool             `form:"autoSave" json:"autoSave"`
	Md5                  string           `form:"md5" json:"md5"`
	RandomCode           string           `form:"randomCode" json:"randomCode"`
	PageParam
}

func (query *BcWearParam) BuildDBData(point *model.BcWear) (err error) {
	var wearDataString string
	var md5 string
	autoSave := false
	if query.WearData != nil {
		marshal, err := json.Marshal(query.WearData)
		if err != nil {
			return err
		}
		wearDataString = string(marshal)
		md5 = support.GeneratorMD5(wearDataString)
	}
	if len(query.WearName) != 0 && len(query.UpdateMemberNumber) != 0 {
		autoSave = true
	}
	*point = model.BcWear{
		Model: gorm.Model{
			ID: query.ID,
		},
		UpdateMemberNumber:   query.UpdateMemberNumber,
		UpdateMemberNickName: query.UpdateMemberNickName,
		WearData:             wearDataString,
		WearName:             query.WearName,
		AutoSave:             autoSave,
		Md5:                  md5,
		RandomCode:           query.RandomCode,
	}
	return
}

func (query *BcWearParam) SerializeParam() bool {
	return true
}

func SerializeBcWear(bcWear *model.BcWear) (query *BcWearParam, err error) {
	var jsonData *json.RawMessage
	if len(bcWear.WearData) != 0 {
		byteList := []byte(bcWear.WearData)
		err = json.Unmarshal(byteList, &jsonData)
		if err != nil {
			return
		}
	}
	query = &BcWearParam{
		ID:                   bcWear.ID,
		UpdateMemberNumber:   bcWear.UpdateMemberNumber,
		UpdateMemberNickName: bcWear.UpdateMemberNickName,
		WearData:             jsonData,
		WearName:             bcWear.WearName,
		AutoSave:             bcWear.AutoSave,
		Md5:                  bcWear.Md5,
		RandomCode:           bcWear.RandomCode,
	}
	return
}
