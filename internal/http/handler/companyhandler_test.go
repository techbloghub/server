package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/internal/http/handler"
	"github.com/techbloghub/server/internal/testutils"
)

func TestListCompanies(t *testing.T) {
	cl, tx := testutils.SetupDB(t)
	defer testutils.TearDown(cl, tx)

	txClient := tx.Client()

	// seedn data 추가
	seedCompanies(t, txClient)

	// Create a new Gin router for test environment
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/companies", handler.ListCompanies(txClient))

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/companies", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var actualResponse handler.CompanyListResponse
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expectedResponse := []map[string]interface{}{
		{
			"name":     "Company A",
			"logo_url": "http://example.com/logoA.png",
			"blog_url": "http://example.com/blogA",
		},
		{
			"name":     "Company B",
			"logo_url": "http://example.com/logoB.png",
			"blog_url": "http://example.com/blogB",
		},
	}

	for i, company := range actualResponse.Companies {
		assert.Equal(t, expectedResponse[i]["name"], company.Name)
		assert.Equal(t, expectedResponse[i]["logo_url"], company.LogoURL)
		assert.Equal(t, expectedResponse[i]["blog_url"], company.BlogURL)
	}
}

func seedCompanies(t *testing.T, client *ent.Client) {
	_, err := client.Company.CreateBulk(
		client.Company.Create().
			SetName("Company A").
			SetLogoURL(&url.URL{Scheme: "http", Host: "example.com", Path: "/logoA.png"}).
			SetBlogURL(&url.URL{Scheme: "http", Host: "example.com", Path: "/blogA"}).
			SetRssURL(&url.URL{Scheme: "http", Host: "example.com", Path: "/rssA"}),
		client.Company.Create().
			SetName("Company B").
			SetLogoURL(&url.URL{Scheme: "http", Host: "example.com", Path: "/logoB.png"}).
			SetBlogURL(&url.URL{Scheme: "http", Host: "example.com", Path: "/blogB"}).
			SetRssURL(&url.URL{Scheme: "http", Host: "example.com", Path: "/rssB"}),
	).Save(context.Background())
	if err != nil {
		t.Fatalf("failed to seed companies: %v", err)
	}
}
