package rss

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/ent/posting"
	"github.com/techbloghub/server/internal/schemasupport"
)

// FetchRSSItems fetches RSS items from a given URL
func FetchRSSItems(rssURL *url.URL) ([]*gofeed.Item, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(rssURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	return feed.Items, nil
}

// ProcessRSSItem converts an RSS item to a posting create input
func ProcessRSSItem(ctx context.Context, client *ent.Client, item *gofeed.Item, companyID int) (*ent.PostingCreate, error) {
	// Skip items without links
	if item.Link == "" {
		return nil, fmt.Errorf("RSS item has no link")
	}

	// Parse URL
	postURL, err := url.Parse(item.Link)
	if err != nil {
		return nil, fmt.Errorf("failed to parse item URL: %w", err)
	}

	// Get published date or use current time if not available
	publishedAt := time.Now()
	if item.PublishedParsed != nil {
		publishedAt = *item.PublishedParsed
	}

	// Create the posting
	postingCreate := client.Posting.Create().
		SetTitle(item.Title).
		SetURL(postURL).
		SetCompanyID(companyID).
		SetPublishedAt(publishedAt)

	// Add tags if available
	if item.Categories != nil && len(item.Categories) > 0 {
		postingCreate.SetTags(schemasupport.NewPostingTags(item.Categories))
	}

	return postingCreate, nil
}

// SyncCompanyRSSFeed syncs postings from a company's RSS feed
func SyncCompanyRSSFeed(ctx context.Context, client *ent.Client, companyID int) (int, error) {
	// Get the company with its RSS URL
	company, err := client.Company.Get(ctx, companyID)
	if err != nil {
		return 0, fmt.Errorf("failed to find company: %w", err)
	}

	// Check if the company has an RSS URL
	if company.RssURL == nil {
		return 0, fmt.Errorf("company does not have an RSS URL")
	}

	// Fetch postings from RSS feed
	items, err := FetchRSSItems(company.RssURL)
	if err != nil {
		return 0, err
	}

	// Process items and save to database
	count := 0
	for _, item := range items {
		postingCreate, err := ProcessRSSItem(ctx, client, item, companyID)
		if err != nil {
			log.Printf("Failed to process RSS item: %v", err)
			continue
		}

		// Check if posting already exists by URL
		urlValue, _ := postingCreate.Mutation().URL()
		exists, err := client.Posting.Query().
			Where(posting.URLEQ(urlValue)).
			Exist(ctx)
		if err != nil {
			log.Printf("Failed to check if posting exists: %v", err)
			continue
		}

		// Only save if posting doesn't exist
		if !exists {
			_, err = postingCreate.Save(ctx)
			if err != nil {
				log.Printf("Failed to save posting: %v", err)
				continue
			}
			count++
		}
	}

	return count, nil
}

// SyncAllCompanyFeeds syncs RSS feeds for all companies with RSS URLs
func SyncAllCompanyFeeds(ctx context.Context, client *ent.Client) (int, error) {
	// Get all companies that have RSS URLs
	companies, err := client.Company.Query().All(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to query companies: %w", err)
	}

	totalCount := 0
	for _, company := range companies {
		if company.RssURL != nil {
			count, err := SyncCompanyRSSFeed(ctx, client, company.ID)
			if err != nil {
				log.Printf("Failed to sync RSS feed for company %s: %v", company.Name, err)
				continue
			}
			totalCount += count
			log.Printf("Synced %d new postings for %s", count, company.Name)
		}
	}

	return totalCount, nil
}
