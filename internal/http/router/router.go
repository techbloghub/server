package router

import (
	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/internal/http/handler"
)

func InitRouter(r *gin.Engine, client *ent.Client) {
	// PingPong 테스트
	r.GET("/ping", handler.PingPong)
	r.HEAD("/ping", handler.DoNothing)

	// 회사 리스트 조회
	// curl -X GET http://localhost:8080/companies
	r.GET("/companies", handler.ListCompanies(client))

	// 태그 전체 목록 조회
	r.GET("/tags", handler.GetTags(client))

	// 포스팅(게시글 조회)
	r.GET("/postings", handler.GetPostings(client))
	
	// RSS feed sync endpoints
	// These would typically be protected by authentication/authorization in production
	rssGroup := r.Group("/rss")
	{
		// Sync all company RSS feeds
		// curl -X POST http://localhost:8080/rss/sync
		rssGroup.POST("/sync", handler.SyncAllRSSFeeds(client))
		
		// Sync a specific company's RSS feed
		// curl -X POST http://localhost:8080/rss/sync/company/1
		rssGroup.POST("/sync/company/:company_id", handler.SyncCompanyRSSFeed(client))
	}
}
