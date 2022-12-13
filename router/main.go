package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"snake/models"
	"strconv"
	"time"
)

type JsonData struct {
	Status          int `json:"status"`            //当前局状态
	AfterMoveNumber int `json:"after_move_number"` //旋转骰子后点数
	Number          int `json:"number"`            //骰子点数
}

// 路由入口
func pathRouter() *gin.Engine {
	r := gin.Default()
	return r
}

// 路由分组
func GroupRouter() {
	r := pathRouter()

	//关卡
	CheckPoint := r.Group("/check_point")
	{
		CheckPoint.GET("/index", func(context *gin.Context) {
			Point, _ := models.GetPointById(1)

			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "OK",
				"data": Point,
			})
		})
	}

	//游戏局
	Inning := r.Group("/inning")
	{
		//开始游戏
		Inning.GET("/get_dice_num", func(context *gin.Context) {
			innings_id := context.DefaultQuery("innings_id", "")
			user_id := context.DefaultQuery("user_id", "")
			inningsId, _ := strconv.Atoi(innings_id)
			userId, _ := strconv.Atoi(user_id)

			//当前局
			Innings, err := models.GetInningsByUserIdAndPointId(userId, inningsId)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  "数据错误",
					"data": "",
				})
				panic("数据错误")
			}

			if Innings.Status == 1 {
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  "当前局已经成功",
					"data": "",
				})
				panic("当前局已经成功")
			}

			//骰子旋转出的随机数
			randInt := models.RandInt(1, 6)
			//获取关卡数据
			Point, err := models.GetPointById(1)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  "数据错误",
					"data": "",
				})
				panic("数据错误")
			}

			//判断当前点数
			AfterMoveNumber := Innings.CurrentNumber + randInt

			inningStatus := 0
			if AfterMoveNumber == Point.Total {
				inningStatus = 1 //成功
			} else if AfterMoveNumber > Point.Total {
				AfterMoveNumber = AfterMoveNumber - Point.Total
			}

			//插入旋转记录
			updateInt := time.Now().Unix()
			fmt.Println(updateInt)
			InningDetail := models.InningDetails{
				InningsId:       inningsId,
				UserId:          userId,
				CurrentNumber:   Innings.CurrentNumber,
				Number:          randInt,
				AfterMoveNumber: AfterMoveNumber,
				Update:          updateInt,
			}

			InningsData := models.Innings{
				CurrentNumber: AfterMoveNumber,
				Status:        inningStatus,
			}
			models.CreateInningDetails(InningDetail)          //添加旋转记录
			models.UpdateInningsById(Innings.Id, InningsData) //更改本局信息

			jsonData := JsonData{
				Status:          inningStatus,
				AfterMoveNumber: AfterMoveNumber,
				Number:          randInt,
			}
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "OK",
				"data": jsonData,
			})
		})

		//回放
		Inning.GET("/get_inning_list", func(context *gin.Context) {
			innings_id := context.DefaultQuery("innings_id", "")
			user_id := context.DefaultQuery("user_id", "")
			inningsId, _ := strconv.Atoi(innings_id)
			userId, _ := strconv.Atoi(user_id)

			Innings, err := models.GetInningDetailsListByUserIdAndPointId(userId, inningsId)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  "数据错误",
					"data": "",
				})
			}

			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "OK",
				"data": Innings,
			})
		})
	}

	r.Run(":8080")
}
