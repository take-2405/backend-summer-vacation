package model

import (
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/db"
)

type User struct {
	Name	string `json:"name"`
}

type Task1 struct {
	Timestamp string `json:"timestamp"`
	Detail  Task1detail `json:"detail"`
}

type Task1detail struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

type Task2 struct {
	Year int`json:"year"`
	Month int `json:"month""`
	Day int `json:"day"`
}
type Task2Response struct {
	Week string `json:"week"`
}

type SignUp struct{
	Id string `json:id`
	Password string `json:password`
}

func InsertSignUpData(id,hashPass,token string)error{
	stmt, err := db.Conn.Prepare("INSERT INTO signUp VALUES (?,?,?)")
	if err != nil {
		return err
	}
	//fmt.Println(fmt.Sprintf("%s", hashPass))
	_, err = stmt.Exec(id,hashPass,token)
	return err
}