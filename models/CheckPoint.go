package models

import (
	"snake/common"
)

/*
*关卡
 */

type Checkpoint struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"total"`
}

func GetPointById(_id int) (cpoints Checkpoint, err error) {
	Db := common.InitDB()
	Db.Find(&cpoints, _id)
	return
}
