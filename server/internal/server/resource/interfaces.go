package resource

import "github.com/gin-gonic/gin"

type PgService struct {
	service *Service
}

type Service interface {
	Login(c *gin.Context)
	GetShortStat(c *gin.Context)
	GetWeekStat(c *gin.Context)
	AddResource(c *gin.Context)
	CheckResource(c *gin.Context)
	DeleteResource(c *gin.Context)
	UpdateResource(c *gin.Context)
	GetGeneralStat(c *gin.Context)
	GetCertificates(c *gin.Context)
	UserIdentity(c *gin.Context)
	GetUserInfo(c *gin.Context)
	AddOwner(c *gin.Context)
	AddUsData(c *gin.Context)
	ChangeUser(c *gin.Context)
}
