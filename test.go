package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

const logErrMsg string="GetLogger failed"

type User struct {
	Name string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type LoggerError struct {
}

func(err LoggerError)Error()string{
	return logErrMsg
}

func GetLogger(e *echo.Echo) (*log.Logger,error){
	if l, ok := e.Logger.(*log.Logger); ok {
		fmt.Println("GetLogger OK")
		return l, nil
	}
	err:=LoggerError{}
	return nil,err
}

func main(){
	e:=echo.New()
	logger,err:=GetLogger(e)
	if err!=nil{
		fmt.Println(err)
	}
	if err!=nil{
		logger.Fatal(err)
	}
	e.POST("/users",func(c echo.Context)(err error){
		u:=new(User)
		if err=c.Bind(u);err!=nil{
			return err
		}
		return c.JSON(http.StatusOK,u)
	})
	e.Start(":3000")
}

