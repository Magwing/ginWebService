package server

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var handler *gin.Engine
var once sync.Once

func GetHandler() *gin.Engine {
	once.Do(func() {
		handler = gin.Default()
		RegisteMiddleware(handler)
	})
	return handler
}

func GET(path string, actFunc func(c *gin.Context) (int, interface{}, error)) {
	GetHandler().GET(path, wrapHandler(actFunc))
}

func PUT(path string, actFunc func(c *gin.Context) (int, interface{}, error)) {
	GetHandler().PUT(path, wrapHandler(actFunc))
}

func POST(path string, actFunc func(c *gin.Context) (int, interface{}, error)) {
	GetHandler().POST(path, wrapHandler(actFunc))
}

func DELETE(path string, actFunc func(c *gin.Context) (int, interface{}, error)) {
	GetHandler().DELETE(path, wrapHandler(actFunc))
}

func OPTIONS(path string, actFunc func(c *gin.Context) (int, interface{}, error)) {
	GetHandler().OPTIONS(path, wrapHandler(actFunc))
}

func wrapHandler(actFunc func(c *gin.Context) (status int, res interface{}, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		sucChan := make(chan interface{}, 1)
		var (
			status int
			res    interface{}
			err    error
		)
		go func(c *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					status, res, err = StatusInternalServerError, nil, errors.New("unknown error")
					log.Println("panic error ")
					sucChan <- true
				}
			}()
			status, res, err = actFunc(c)
			sucChan <- true
		}(c)

		select {
		case <-time.After(time.Second):
			status, res, err = StatusRequestTimeout, nil, errors.New("time out")
			log.Println("time out")
		case <-sucChan:
			log.Println("http request success")
		}

		WriteResponse(c, status, res, err)
	}
}
