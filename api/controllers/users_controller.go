package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xmaten/expenses-tracker-api/api/models"
	"github.com/xmaten/expenses-tracker-api/api/utils/formaterror"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateUser(c *gin.Context) {
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error": errList,
		})

		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error": errList,
		})
		return
	}
	user.Prepare()
	errorMessages := user.Validate("")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error": errList,
		})
		return
	}

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error": errList,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"response": userCreated,
	})
}

func (server *Server) GetUser(c *gin.Context) {
	errList = map[string]string{}

	userID := c.Param("id")

	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": errList,
		})
		return
	}
	user := models.User{}

	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		errList["No_user"] = "No user found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error": errList,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": userGotten,
	})
}



















