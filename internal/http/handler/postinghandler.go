package handler

import (
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
	CreateTime    time.Time `json:"createTime"`
	UpdateTime    time.Time `json:"updateTime"`
	PublishedTime time.Time `json:"publishedTime"`
}

type PostingSearchResponses struct {
	Count int `json:"id"`
	//Tags   []TagSearchResponse   `json:"tags"`
	Postings    []TitleSearchResponse `json:"titles"`
	HasNextPage bool                  `json:"hasNextPage"`
}

func GetPostings(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		titleSearchParam := c.DefaultQuery("title", "")
		paging := common.GenerateTechPaging(
			c.Query("cursor"),
			c.Query("size"),
		)

		query := client.Posting.
			Query().
			WithCompany()
		if titleSearchParam != "" {
			query = query.Where(posting.TitleContainsFold(titleSearchParam))
		}
		if paging.Cursor > 0 {
			query = query.Where(posting.IDLT(paging.Cursor))
		}
		postings, err := query.
			Order(
				ent.Desc(posting.FieldPublishedAt),
				ent.Desc(posting.FieldID),
			).
			Limit(paging.Size).
			All(c)
		postingsByTitle := make([]TitleSearchResponse, len(postings))
		if err != nil {
			for i, posting := range postings {
				postingsByTitle[i] = TitleSearchResponse{
					ID:            posting.ID,
					Title:         posting.Title,
					Url:           posting.URL.String(),
					Company:       posting.Edges.Company.Name,
					Logo:          posting.Edges.Company.LogoURL.String(),
					Tags:          posting.Tags.ToStringArray(),
					CreateTime:    posting.CreateTime,
					UpdateTime:    posting.UpdateTime,
					PublishedTime: posting.PublishedAt,
				}
			}
		}

		totalCount, err := client.Posting.Query().
			Where(posting.TitleContainsFold(titleSearchParam)).
			Count(c)
		c.JSON(200, PostingSearchResponses{
			Count:       totalCount,
			Postings:    postingsByTitle,
			HasNextPage: paging.HasNextPage(totalCount),
		})
	}
}

//tagSearchParam := c.DefaultQuery("tag", "")

//postingsByTag := []TagSearchResponse{}
//if tagSearchParam != "" {
//	result, err := client.Posting.
//		Query().
//		Where(schema.TagsContains(tagSearchParam)).
//		All(c)
//
//	if err != nil {
//		postingsByTag = make([]TagSearchResponse, len(result))
//		for i, posting := range result {
//			postingsByTag[i] = TagSearchResponse{
//				ID:            posting.ID,
//				Title:         posting.Title,
//				Url:           posting.URL.String(),
//				CreateTime:    posting.CreateTime,
//				UpdateTime:    posting.UpdateTime,
//				PublishedTime: posting.PublishedAt,
//			}
//		}
//	}
//}
//type TagSearchResponse struct {
//	ID            int       `json:"posting_id"`
//	Title         string    `json:"title"`
//	Url           string    `json:"url"`
//	Company       string    `json:"company"`
//	CreateTime    time.Time `json:"createTime"`
//	UpdateTime    time.Time `json:"updateTime"`
//	PublishedTime time.Time `json:"publishedTime"`
//}
