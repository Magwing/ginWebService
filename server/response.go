package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// StatusOK 200
	StatusOK = http.StatusOK
	// StatusBadRequest 400
	StatusBadRequest = http.StatusBadRequest
	// StatusUnauthorized 401
	StatusUnauthorized = http.StatusUnauthorized
	// StatusForbidden 403
	StatusForbidden = http.StatusForbidden
	// StatusMethodNotAllowed 405
	StatusMethodNotAllowed = http.StatusMethodNotAllowed
	// StatusRequestTimeout 408
	StatusRequestTimeout = http.StatusRequestTimeout
	// StatusInternalServerError 500
	StatusInternalServerError = http.StatusInternalServerError
)

type MetaData struct {
	Status  int
	Message string
}

type ReturnStr struct {
	Meta   MetaData
	Result interface{}
}

func WriteResponse(c *gin.Context, status int, res interface{}, err error) {

	message := "OK"
	if err != nil {
		message = err.Error()
	}
	c.JSON(StatusOK, ReturnStr{
		Meta: MetaData{
			Status:  status,
			Message: message,
		},
		Result: res,
	})
}
