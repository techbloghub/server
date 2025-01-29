package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
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
		tagEntities, err := client.Tag.Query().All(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}

		tagCounts := len(tagEntities)
		tagResponses := make([]TagResponse, tagCounts)
		for i, tag := range tagEntities {
			tagResponses[i] = TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
		}

		c.JSON(200, TagResponses{
			Count: tagCounts,
			Tags:  tagResponses,
		})
	}
}
