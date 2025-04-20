package rss_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/internal/rss"
	"github.com/techbloghub/server/internal/testutils"
)

func TestSyncCompanyRSSFeed(t *testing.T) {
	testutils.TransactionalTest(t, func(t *testing.T, client *ent.Client) {
		ctx := context.Background()

		// Test with invalid RSS URL
		t.Run("Company with empty RSS URL", func(t *testing.T) {
			// Create test company with empty RSS URL
			emptyURL := &url.URL{}
			company, err := client.Company.Create().
				SetName("Empty RSS Company").
				SetBlogURL(&url.URL{Scheme: "https", Host: "example.com"}).
				SetLogoURL(&url.URL{Scheme: "https", Host: "example.com", Path: "/logo.png"}).
				SetRssURL(emptyURL).
				Save(ctx)
			assert.NoError(t, err)

			// Try to sync RSS feed (should fail due to invalid URL)
			count, err := rss.SyncCompanyRSSFeed(ctx, client, company.ID)
			assert.Error(t, err)
			assert.Equal(t, 0, count)
			assert.Contains(t, err.Error(), "failed to parse")
		})

		// Setup a mock RSS server
		server := setupMockRSSServer()
		defer server.Close()

		t.Run("Company with RSS URL", func(t *testing.T) {
			// Create test company with RSS URL
			serverURL, _ := url.Parse(server.URL + "/feed.xml")
			company, err := client.Company.Create().
				SetName("RSS Company").
				SetBlogURL(&url.URL{Scheme: "https", Host: "example.com"}).
				SetLogoURL(&url.URL{Scheme: "https", Host: "example.com", Path: "/logo.png"}).
				SetRssURL(serverURL).
				Save(ctx)
			assert.NoError(t, err)

			// Sync RSS feed
			count, err := rss.SyncCompanyRSSFeed(ctx, client, company.ID)
			assert.NoError(t, err)
			assert.Equal(t, 2, count) // We expect 2 postings from our mock RSS feed

			// Verify postings were created
			postings, err := client.Company.QueryPostings(company).All(ctx)
			assert.NoError(t, err)
			assert.Len(t, postings, 2)

			// Verify posting content
			assert.Equal(t, "Test Post 1", postings[0].Title)
			assert.Equal(t, "Test Post 2", postings[1].Title)
		})
	})
}

func TestSyncAllCompanyFeeds(t *testing.T) {
	testutils.TransactionalTest(t, func(t *testing.T, client *ent.Client) {
		ctx := context.Background()

		// Setup a mock RSS server
		server := setupMockRSSServer()
		defer server.Close()

		// Create test companies
		serverURL, _ := url.Parse(server.URL + "/feed.xml")

		// Company with RSS URL
		_, err := client.Company.Create().
			SetName("RSS Company 1").
			SetBlogURL(&url.URL{Scheme: "https", Host: "example1.com"}).
			SetLogoURL(&url.URL{Scheme: "https", Host: "example1.com", Path: "/logo1.png"}).
			SetRssURL(serverURL).
			Save(ctx)
		assert.NoError(t, err)

		// Another company with same RSS URL (to test deduplication)
		_, err = client.Company.Create().
			SetName("RSS Company 2").
			SetBlogURL(&url.URL{Scheme: "https", Host: "example2.com"}).
			SetLogoURL(&url.URL{Scheme: "https", Host: "example2.com", Path: "/logo2.png"}).
			SetRssURL(serverURL).
			Save(ctx)
		assert.NoError(t, err)

		// Company with empty RSS URL (won't be synced)
		_, err = client.Company.Create().
			SetName("No RSS Company").
			SetBlogURL(&url.URL{Scheme: "https", Host: "example3.com"}).
			SetLogoURL(&url.URL{Scheme: "https", Host: "example3.com", Path: "/logo3.png"}).
			SetRssURL(&url.URL{}).
			Save(ctx)
		assert.NoError(t, err)

		// Sync all RSS feeds
		count, err := rss.SyncAllCompanyFeeds(ctx, client)
		assert.NoError(t, err)
		assert.Equal(t, 2, count) // 2 postings from first company, 0 from second (deduplication)

		// Verify total postings
		postingsCount, err := client.Posting.Query().Count(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 2, postingsCount)
	})
}

// Setup a mock RSS server
func setupMockRSSServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a simple RSS feed with two items
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)

		// Current time for published dates
		now := time.Now().Format(time.RFC1123Z)

		// Simple RSS feed
		feed := `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>Test Blog</title>
    <link>https://example.com/blog</link>
    <description>A blog for testing RSS feeds</description>
    <item>
      <title>Test Post 1</title>
      <link>https://example.com/blog/post1</link>
      <pubDate>` + now + `</pubDate>
      <category>golang</category>
      <category>testing</category>
    </item>
    <item>
      <title>Test Post 2</title>
      <link>https://example.com/blog/post2</link>
      <pubDate>` + now + `</pubDate>
      <category>rss</category>
    </item>
  </channel>
</rss>`

		w.Write([]byte(feed))
	}))
}

