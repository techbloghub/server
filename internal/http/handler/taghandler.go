package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/ent/tag"
)

type TagResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TagResponses struct {
	Count int           `json:"count"`
	Tags  []TagResponse `json:"tags"`
}

func GetTags(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchParam := c.DefaultQuery("search", "")

		var tags []*ent.Tag
		var err error

		if searchParam != "" {
			tags, err = searchTagsByName(c, client, searchParam)
		} else {
			tags, err = client.Tag.Query().All(c)
		}

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		tagResponses := makeTagResponses(tags)
		c.JSON(200, TagResponses{
			Count: len(tagResponses),
			Tags:  tagResponses,
		})
	}
}

func searchTagsByName(c *gin.Context, client *ent.Client, searchParam string) ([]*ent.Tag, error) {
	return client.Tag.Query().
		Where(tag.NameContains(searchParam)).
		All(c)
}

func makeTagResponses(tags []*ent.Tag) []TagResponse {
	tagResponses := make([]TagResponse, len(tags))
	for i, tag := range tags {
		tagResponses[i] = TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}
	return tagResponses
}
