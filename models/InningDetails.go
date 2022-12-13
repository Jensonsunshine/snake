package models

import "snake/common"

type InningDetails struct {
	Id              int   `json:"id"`
	InningsId       int   `json:"innings_id"`        //局id
	UserId          int   `json:"user_id"`           //用户id
	CurrentNumber   int   `json:"current_number"`    //当前停留格数
	Number          int   `json:"number"`            //旋转出格数
	AfterMoveNumber int   `json:"after_move_number"` //应该移动到的格数
	Update          int64 `json:"update"`            //时间
}

func GetInningDetailsByUserIdAndInningsId(_user_id int, _innings_id int) (inningsDet InningDetails) {
	Db := common.InitDB()
	Db.Find(&inningsDet, "user_id=? and innings_id=?", _user_id, _innings_id)
	return
}

func GetInningDetailsListByUserIdAndPointId(_user_id int, _innings_id int) (inningsDet []InningDetails, err error) {
	Db := common.InitDB()
	Db.Find(&inningsDet, "user_id=? and innings_id=?", _user_id, _innings_id)
	return
}

func CreateInningDetails(InDet InningDetails) (err error) {
	Db := common.InitDB()
	Db.Create(&InDet)
	return
}
