package handler

import (
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
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

		if titleSearchParam == "" && tagsSearchParam == "" {
			posts, err := fetchPostings(client, "", paging, c)
			totalCount, err := countTotalPostings(client, "", c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, PostingSearchResponses{
				Count:       totalCount,
				Postings:    convertToTitleSearchResponse(posts),
				HasNextPage: paging.HasNextPage(totalCount),
			})
			return
		}

		if titleSearchParam != "" {
			counts, err := countTotalPostings(client, titleSearchParam, c)
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
				Count:       counts,
				Postings:    convertToTitleSearchResponse(postings),
				HasNextPage: paging.HasNextPage(counts),
			})
		} else {
			totalCount, err := countTotalPostings(client, tagsSearchParam, c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			postings, err := fetchPostings(client, tagsSearchParam, paging, c)
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
}

func fetchPostings(client *ent.Client, param string, paging common.TechbloghubPaging, c *gin.Context) ([]*ent.Posting, error) {
	query := client.Posting.Query().WithCompany()
	if strings.Contains(param, ",") {
		arr := strings.Split(param, ",")
		query = query.Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains("tags", arr))
		})
	} else if param != "" {
		query = query.Where(posting.TitleContainsFold(param))
	}
	if paging.Cursor > 0 {
		query = query.Where(posting.IDLT(paging.Cursor))
	}

	return query.Order(
		ent.Desc(posting.FieldPublishedAt),
		ent.Desc(posting.FieldID),
	).Limit(paging.Size).All(c)
}

func countTotalPostings(client *ent.Client, param string, c *gin.Context) (int, error) {
	query := client.Posting.Query()
	if param != "" {
		query = query.Where(posting.TitleContainsFold(param))
	} else if strings.Contains(param, ",") {
		arr := strings.Split(param, ",")
		query = query.Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains("tags", arr))
		})
	}
	return query.Count(c)
}

func convertToTitleSearchResponse(postings []*ent.Posting) []PostingSearchResponse {
	responses := make([]PostingSearchResponse, len(postings))
	for i, posting := range postings {
		responses[i] = PostingSearchResponse{
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
