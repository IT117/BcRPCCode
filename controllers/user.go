package controllers

type User struct {
	Count string `form:"count"`
	Arguments interface{} `form:"arguments"`
	Arguments1 bool  `form:"arguments"`
}
