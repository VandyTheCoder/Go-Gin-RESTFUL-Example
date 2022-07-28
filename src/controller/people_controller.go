package controller

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"me.lucifer/backend/src/model/people"
	"me.lucifer/backend/src/template/request/person_create"
	"me.lucifer/backend/src/template/request/person_update"
)

func PeopleGet(repo people.Getter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		results := repo.GetAll()
		ctx.JSON(http.StatusOK, results)
	}
}

func PeopleCreate(repo people.Adder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := person_create.Request{}
		ctx.Bind(&body)
		
		person := people.Person{
			LastName: 	body.LastName,
			FirstName: 	body.FirstName,
			Gender: 		body.Gender,
		}
		repo.Add(person)

		ctx.Status(http.StatusAccepted)
	}
}

func PeopleShow(repo *people.Repo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil { panic(err) }
		result := repo.Find(id)
		ctx.JSON(http.StatusOK, result)
	}
}

func PeopleUpdate(repo *people.Repo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil { panic(err) }
		index := id - 1

		body := person_update.Request{}
		ctx.Bind(&body)
		
		repo.People[index].LastName = body.LastName
		repo.People[index].FirstName = body.FirstName
		repo.People[index].Gender = body.Gender
		
		ctx.Status(http.StatusAccepted)
	}
}