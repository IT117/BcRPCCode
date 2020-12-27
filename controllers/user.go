package controllers

type User struct {
	Count string `form:"count"`
	Arguments interface{} `form:"arguments"`
}
