package models

import (
	"math/rand"
	"snake/common"
)

type Innings struct {
	Id            int `json:"id"`
	UserId        int `json:"user_id"`
	CheckpointId  int `json:"checkpoint_id"`
	CurrentNumber int `json:"current_number"`
	Status        int `json:"status"`
}

func GetInningsByUserIdAndPointId(_user_id int, _checkpoint_id int) (innings Innings, err error) {
	Db := common.InitDB()
	Db.Find(&innings, "user_id=? and checkpoint_id=?", _user_id, _checkpoint_id)
	return
}

func UpdateInningsById(_id int, inData Innings) (err error) {
	Db := common.InitDB()
	Db.Table("innings").Where("id=?", _id).Update(&inData)
	return
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min+1) + min
}
