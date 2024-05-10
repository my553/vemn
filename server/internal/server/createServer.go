package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"vemn/internal/server/resource"
)

type Service struct {
	service *Create
}

type Create interface {
	CreateServer()
}

func (service *Service) CreateServer() {
	method := resource.PgService{}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Authorization"},
	}))

	api := router.Group("/api", method.UserIdentity)
	{
		api.GET("/chart-statistic", method.GetShortStat)
		api.GET("/week-statistic", method.GetWeekStat)
		api.GET("/general-statistic", method.GetGeneralStat)
		api.GET("/certificates", method.GetCertificates)
		api.GET("/statistic", method.GetStatistic)
		api.GET("/user-info", method.GetUserInfo)

		api.POST("/add-resource", method.AddResource)
		api.POST("/check-resource", method.CheckResource)
		api.POST("/delete-resource", method.DeleteResource)
		api.POST("/update-resource", method.UpdateResource)

		api.POST("/change-user", method.ChangeUser)
		api.POST("/add-owner", method.AddOwner)
		api.POST("/add-usdata", method.AddUsData)
	}
	router.POST("/login", method.Login)

	//err := router.RunTLS(":8080", "certificate/*.cibint.mosreg.ru.crt", "certificate/*.cibint.mosreg.ru.key")
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
