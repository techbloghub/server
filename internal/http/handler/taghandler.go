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

		tags, err := fetchTags(c, searchParam, client)
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

func fetchTags(c *gin.Context, searchParam string, client *ent.Client) ([]*ent.Tag, error) {
	if searchParam != "" {
		return searchTagsByName(c, client, searchParam)
	}
	return client.Tag.Query().All(c)
}

func searchTagsByName(c *gin.Context, client *ent.Client, searchParam string) ([]*ent.Tag, error) {
	return client.Tag.Query().
		Where(tag.NameContainsFold(searchParam)).
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
