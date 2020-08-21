package model

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
	Year string`json:"year"`
	Month string `json:"month""`
	Day string `json:"day"`
}
type Task2Response struct {
	Week string `json:"week"`
}