package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/ent/posting"
	"github.com/techbloghub/server/internal/common"
)

type TitleSearchResponse struct {
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
	Count       int                   `json:"count"`
	Postings    []TitleSearchResponse `json:"postings"`
	HasNextPage bool                  `json:"has_next_page"`
}

func GetPostings(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		titleSearchParam := c.DefaultQuery("title", "")
		tagsSearchParam := c.DefaultQuery("tags", "")
		paging := common.GenerateTechPaging(c.Query("cursor"), c.Query("size"))

		totalCount, err := countTotalPostings(client, titleSearchParam, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		postings, err := fetchPostings(client, titleSearchParam, paging, c)
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
}

func fetchPostings(client *ent.Client, title string, paging common.TechbloghubPaging, c *gin.Context) ([]*ent.Posting, error) {
	query := client.Posting.Query().WithCompany()
	if title != "" {
		query = query.Where(posting.TitleContainsFold(title))
	}
	if paging.Cursor > 0 {
		query = query.Where(posting.IDLT(paging.Cursor))
	}

	return query.Order(
		ent.Desc(posting.FieldPublishedAt),
		ent.Desc(posting.FieldID),
	).Limit(paging.Size).All(c)
}

func countTotalPostings(client *ent.Client, title string, c *gin.Context) (int, error) {
	query := client.Posting.Query()
	if title != "" {
		query = query.Where(posting.TitleContainsFold(title))
	}
	return query.Count(c)
}

func convertToTitleSearchResponse(postings []*ent.Posting) []TitleSearchResponse {
	responses := make([]TitleSearchResponse, len(postings))
	for i, posting := range postings {
		responses[i] = TitleSearchResponse{
			ID:            posting.ID,
			Title:         posting.Title,
			Url:           posting.URL.String(),
			Company:       posting.Edges.Company.Name,
			Logo:          posting.Edges.Company.LogoURL.String(),
			Tags:          posting.Tags.ToStringSlice(),
			CreateTime:    posting.CreateTime,
			UpdateTime:    posting.UpdateTime,
			PublishedTime: posting.PublishedAt,
		}
	}
	return responses
}
