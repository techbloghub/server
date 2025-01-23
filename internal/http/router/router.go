package router

import (
	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/internal/http/handler"
)

func InitRouter(r *gin.Engine, client *ent.Client) {
	// PingPong 테스트
	r.GET("/ping", handler.PingPong)

	// 회사 리스트 조회
	// curl -X GET http://localhost:8080/companies
	r.GET("/companies", handler.ListCompanies(client))
}
