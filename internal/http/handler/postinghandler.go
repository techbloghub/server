package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/ent/company"
	"github.com/techbloghub/server/ent/posting"
	"github.com/techbloghub/server/internal/common"
)

type PostingSearchResponse struct {
	ID            int       `json:"posting_id"`
	Title         string    `json:"title"`
	Url           string    `json:"url"`
	Company       string    `json:"company"`
	Logo          string    `json:"logo"`
	Tags          []string  `json:"tags"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	PublishedTime time.Time `json:"published_time"`
}

type PostingSearchResponses struct {
	Count       int                     `json:"count"`
	Postings    []PostingSearchResponse `json:"postings"`
	HasNextPage bool                    `json:"has_next_page"`
}

// ✅ GetPostings을 더 깔끔하게 정리
func GetPostings(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		titleSearchParam := c.DefaultQuery("title", "")
		tagsSearchParam := c.DefaultQuery("tags", "")

		paging := common.GenerateTechPaging(c.Query("cursor"), c.Query("size"))

		handlePostingsQuery(client, titleSearchParam, tagsSearchParam, paging, c)
	}
}

func handlePostingsQuery(client *ent.Client, title, tags string, paging common.TechbloghubPaging, c *gin.Context) {
	searchParam := title
	if searchParam == "" {
		searchParam = tags
	}

	postings, err := fetchPostings(client, searchParam, paging, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalCount, err := countPostings(client, searchParam, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, PostingSearchResponses{
		Count:       totalCount,
		Postings:    convertToTitleSearchResponse(postings),
		HasNextPage: paging.HasNextPage(totalCount),
	})
}

func fetchPostings(client *ent.Client, param string, paging common.TechbloghubPaging, c *gin.Context) ([]*ent.Posting, error) {
	query := client.Posting.Query()

	if param != "" {
		query = applySearchFilter(query, param)
	}

	if paging.Cursor > common.CURSOR_DEFAULT {
		query = query.Where(posting.IDLT(paging.Cursor))
	}
	return query.Where(posting.HasCompany()).
		WithCompany(func(q *ent.CompanyQuery) {
			q.Select(
				company.FieldLogoURL,
				company.FieldName,
			)
		}).
		Order(
			ent.Desc(posting.FieldID),
			ent.Desc(posting.FieldID),
		).
		Limit(paging.Size).
		All(c.Request.Context())
}

func applySearchFilter(query *ent.PostingQuery, param string) *ent.PostingQuery {
	// before: [react,java]
	// after: ARRAY[react,java]
	if strings.HasPrefix(param, "[") && strings.HasSuffix(param, "]") {
		arrayQuery := "ARRAY" + param

		query = query.Where(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("%s @> %s", s.C("tags"), arrayQuery)))
		})
		// 단일 제목으로 검색하는 케이스는 제목 검색
	} else if param != "" {
		query = query.Where(posting.TitleContainsFold(param))
	}
	return query
}

func countPostings(client *ent.Client, param string, c *gin.Context) (int, error) {
	query := client.Posting.Query()
	if param != "" {
		query = applySearchFilter(query, param)
	}
	return query.Count(c)
}

// 대체 왜 못갖고오냐
func getCompanyInfo(posting *ent.Posting) (string, string) {
	if posting.Edges.Company != nil {
		return posting.Edges.Company.Name, posting.Edges.Company.LogoURL.String()
	}
	return "Unknown", ""
}

func convertToTitleSearchResponse(postings []*ent.Posting) []PostingSearchResponse {
	var responses []PostingSearchResponse

	for _, posting := range postings {
		companyName, logoURL := getCompanyInfo(posting)

		responses = append(responses, PostingSearchResponse{
			ID:            posting.ID,
			Title:         posting.Title,
			Url:           posting.URL.String(),
			Company:       companyName,
			Logo:          logoURL,
			Tags:          posting.Tags.ToStringSlice(),
			CreateTime:    posting.CreateTime,
			UpdateTime:    posting.UpdateTime,
			PublishedTime: posting.PublishedAt,
		})
	}

	return responses
}
