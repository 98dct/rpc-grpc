package main

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var user = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "木系",
		Phone: "15109201334",
	},
	"2": {
		Id:    "2",
		Name:  "小沐",
		Phone: "18922345678",
	},
}
