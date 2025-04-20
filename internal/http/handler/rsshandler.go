package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/internal/rss"
)

type RssSyncResponse struct {
	Status    string `json:"status"`
	Count     int    `json:"count"`
	CompanyID int    `json:"company_id,omitempty"`
	Message   string `json:"message,omitempty"`
}

// SyncCompanyRSSFeed syncs postings from a specific company's RSS feed
func SyncCompanyRSSFeed(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyID, err := strconv.Atoi(c.Param("company_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, RssSyncResponse{
				Status:  "error",
				Message: "Invalid company ID",
			})
			return
		}

		count, err := rss.SyncCompanyRSSFeed(c.Request.Context(), client, companyID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, RssSyncResponse{
				Status:    "error",
				CompanyID: companyID,
				Message:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, RssSyncResponse{
			Status:    "success",
			Count:     count,
			CompanyID: companyID,
			Message:   "RSS feed sync completed successfully",
		})
	}
}

// SyncAllRSSFeeds syncs postings from all companies with RSS feeds
func SyncAllRSSFeeds(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := rss.SyncAllCompanyFeeds(c.Request.Context(), client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, RssSyncResponse{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, RssSyncResponse{
			Status:  "success",
			Count:   count,
			Message: "All RSS feeds synced successfully",
		})
	}
}

