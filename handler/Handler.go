package handler

import (
	"fmt"
	"net/http"
	"rest-api-gorm/repository"
	"rest-api-gorm/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerService struct {
	service service.Service
}

func NewHandler(service service.Service) *handlerService {
	return &handlerService{service}
}

func (h handlerService) CreateApi(c *gin.Context) {
	var persons service.PersonRequest
	err := c.ShouldBindJSON(&persons)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			messageError := fmt.Sprintf("data Error Field : %s, dan di tag %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, messageError)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}
	person, err := h.service.Create(persons)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": person,
	})

}

func (h handlerService) FindAll(c *gin.Context) {

	persons, err := h.service.FindAll()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "data Tidak ditemukan",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": persons,
	})

}

func (h handlerService) FindById(c *gin.Context) {

	paramCon := c.Param("id")
	params, err := strconv.Atoi(paramCon)

	if err != nil {
		panic(err)
	}

	person, err := h.service.FindById(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"data": "id tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": person,
	})
}

func (h handlerService) Update(c *gin.Context) {

	var updatePerson service.PersonUpdate

	id := c.Param("id")
	getId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	data, err := h.service.FindById(getId)

	if err != nil {
		panic(err)
	}

	updatePerson.Nama = data.Nama
	updatePerson.Alamat = data.Alamat

	err = c.ShouldBindJSON(&updatePerson)
	if err != nil {
		errorMessage := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages := fmt.Sprintf("error field :%s and Tag :%s", e.Field(), e.ActualTag())
			errorMessage = append(errorMessage, errorMessages)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	data, err = h.service.Update(getId, updatePerson)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func (h handlerService) Delate(c *gin.Context) {
	getId := c.Param("id")
	id, err := strconv.Atoi(getId)
	if err != nil {
		panic(err)
	}

	err = h.service.Delete(id, repository.Person{})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "data berhasil dihapus",
	})
}
