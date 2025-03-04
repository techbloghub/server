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

func GetPostings(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		titleSearchParam := c.DefaultQuery("title", "")
		tagsSearchParam := c.DefaultQuery("tags", "")
		paging := common.GenerateTechPaging(c.Query("cursor"), c.Query("size"))

		postings, err := fetchPostings(client, titleSearchParam, tagsSearchParam, paging, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		totalCount, err := countPostings(client, titleSearchParam, tagsSearchParam, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, PostingSearchResponses{
			Count:       totalCount,
			Postings:    toPostingSearchResponses(postings),
			HasNextPage: paging.HasNextPage(totalCount),
		})
	}
}

func applySearchFilters(query *ent.PostingQuery, title, tags string) *ent.PostingQuery {
	if title != "" {
		query = query.Where(posting.TitleContainsFold(title))
	}
	// 태그 조회 케이스 -> postgresql ARRAY 검색 형식에 맞춰야함
	// tags: [react,java]
	// query : where tags @> ARRAY['react','java']
	if tags != "" {
		trimmed := strings.Trim(tags, "[]")
		arr := strings.Split(trimmed, ",")
		arrayQuery := "ARRAY" + fmt.Sprintf("['%s']", strings.Join(arr, ","))

		query = query.Where(func(s *sql.Selector) {
			s.Where(sql.ExprP(fmt.Sprintf("%s @> %s", s.C("tags"), arrayQuery)))
		})
	}

	return query
}

func fetchPostings(client *ent.Client, title, tags string, paging common.TechbloghubPaging, c *gin.Context) ([]*ent.Posting, error) {
	query := client.Posting.Query()
	query = applySearchFilters(query, title, tags)

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
			ent.Desc(posting.FieldPublishedAt),
			ent.Desc(posting.FieldID),
		).
		Limit(paging.Size).
		All(c.Request.Context())
}

func countPostings(client *ent.Client, title, tags string, c *gin.Context) (int, error) {
	query := client.Posting.Query()
	query = applySearchFilters(query, title, tags)
	return query.Count(c)
}

// 대체 왜 못갖고오냐!!!!
func getCompanyInfo(posting *ent.Posting) (string, string) {
	if posting.Edges.Company != nil {
		return posting.Edges.Company.Name, posting.Edges.Company.LogoURL.String()
	}
	return "Unknown", ""
}

func toPostingSearchResponses(postings []*ent.Posting) []PostingSearchResponse {
	responses := make([]PostingSearchResponse, len(postings))

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
