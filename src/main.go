package main

import (
	"github.com/gin-gonic/gin"
	"me.lucifer/backend/src/controller"
	"me.lucifer/backend/src/model/people"
)

func main() {
	prefix := "/api/v1"

	people_repo := people.Init()

	router := gin.Default()
	router.GET("/", controller.RootController())
	router.GET(prefix			+"/people",				controller.PeopleGet(people_repo))
	router.GET(prefix   	+"/people/:id", 	controller.PeopleShow(people_repo))
	router.POST(prefix		+"/people", 			controller.PeopleCreate(people_repo))
	router.PUT(prefix			+"/people/:id", 	controller.PeopleUpdate(people_repo))

	router.Run()
}