package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
)

// Company 정보 Response 구조체
type CompanyResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
	BlogURL string `json:"blog_url"`
}

func ListCompanies(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		entCompanies, err := client.Company.Query().All(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}

		companies := make([]CompanyResponse, len(entCompanies))
		for i, company := range entCompanies {
			companies[i] = CompanyResponse{
				ID:      company.ID,
				Name:    company.Name,
				LogoURL: company.LogoURL.String(),
				BlogURL: company.BlogURL.String(),
			}
		}

		c.JSON(200, gin.H{
			"companies": companies,
		})
	}
}
