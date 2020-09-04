package controller

import (
	"log"
	"net/http"
	"time"
	"crypto/sha256"
	//"golang.org/x/crypto/bcrypt"
	// import gin library
	"github.com/gin-gonic/gin"

	// import sample API packages
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/model"
)

var(
	user model.User
	task1 model.Task1
	task2 model.Task2
	task2res model.Task2Response
	signUp model.SignUp
)

type Controller struct {
}

func (ctrl *Controller)HelloWorld(context *gin.Context) {
	context.JSON(200, gin.H{"message": "hello world"})
}

func (ctrl *Controller)SayHello(context *gin.Context) {
	err := context.BindJSON(&user)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON")
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	context.JSON(200, gin.H{"message": "hello " + user.Name})
}

// 課題1
// 説明：
// 現在の日付と時間を返す.
// JSONの生成は gin.H を用いても良い
// 
// リクエスト => なし
// レスポンス => 
// {
//   "timestamp": string,
//   "detail": {
//     "date": string, //例： 2020-09-02
//     "time": string, //例: 00:00:00
//   }
// }
func (ctrl *Controller)Task1(context *gin.Context) {
	time := time.Now()
	const layout1 = "2006-01-02"
	const layout2 = "15:04:05"
	task1.Timestamp=time.String()
	task1.Detail.Time=time.Format(layout1)
	task1.Detail.Date=time.Format(layout2)
	context.JSON(200,model.Task1{task1.Timestamp,task1.Detail} )
	return
}

// 課題2
// 説明：
// ツェラーの公式でリクエストで投げた日付の曜日を返す
// JSONの生成は encoding/json を使用すること
// 
// リクエスト => 
// {
//   "year": Int,
//   "month": Int,
//   "day": Int,
// }
// レスポンス => 
// {
//   "week": string //例： Monday
// }
func (ctrl *Controller)Task2(context *gin.Context) {
	err := context.ShouldBindJSON(&task2)
	if err != nil {
		log.Println("[ERROR] Faild Bind JSON")
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	task2res.Week=DistinguishDaysOfWeek(calcDaysOfWeek(task2.Year,task2.Month,task2.Day))
	context.JSON(200, model.Task2Response{task2res.Week})
}

func calcDaysOfWeek(year,month,day int)int{
	var numDaysOfWeek int
	numDaysOfWeek=(year+(year/4)-(year/100)+(year/400)+(13*month+8)/5+day)%7
	return numDaysOfWeek
}

func DistinguishDaysOfWeek(numDaysOfWeek int)string{
	switch numDaysOfWeek{
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "TuesDay"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	}
	return ""
}

// 課題3
// 説明：
// ユーザーIDとパスワードをデータベースに登録して, 発行したトークンを返す
// パスワードはハッシュ化したものをデータベースに登録する
// JSONの生成は encoding/json を使用すること
// 
// リクエスト => 
// {
//   "id": string,
//   "password": string,
// }
// レスポンス => 
// {
//   "token": string
// }
func (ctrl *Controller)SignUp(context *gin.Context) {
	if err := context.BindJSON(&signUp) ; err != nil {
		log.Println("[ERROR] Faild Bind JSON")
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	pass := []byte(signUp.Passeord)
	hashed:= sha256.Sum256(pass)
	log.Println(string(hashed[0]))
	if err :=model.InsertSignUpData(signUp.Id,"");err!=nil{
		log.Println(err)
		context.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	context.JSON(200, "")
}

// 課題4
// 説明：
// ユーザーIDとパスワードをデータベースに登録されたものかを照合する
// 照合が終わったら結果を返す
// JSONの生成は encoding/json を使用すること
// 
// リクエスト => 
// {
//   "id": string,
//   "password": string
// }
// レスポンス => 
// {
//   "certification": boolean
// }
func (ctrl *Controller)SignIn(context *gin.Context) {
}