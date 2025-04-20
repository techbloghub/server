package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/techbloghub/server/config"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/internal/rss"
)

// Scheduler manages all scheduled tasks
type Scheduler struct {
	client *ent.Client
	cron   *cron.Cron
	config *config.SchedulerConfig
}

// NewScheduler creates a new scheduler
func NewScheduler(client *ent.Client, config *config.SchedulerConfig) *Scheduler {
	// Create a new cron scheduler with seconds field enabled
	c := cron.New(cron.WithSeconds())
	return &Scheduler{
		client: client,
		cron:   c,
		config: config,
	}
}

// Start starts the scheduler
func (s *Scheduler) Start() {
	s.setupJobs()
	s.cron.Start()
	log.Println("Scheduler started")
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
	ctx := s.cron.Stop()
	<-ctx.Done()
	log.Println("Scheduler stopped")
}

// setupJobs sets up all scheduled jobs
func (s *Scheduler) setupJobs() {
	// Use the cron expression from config or fall back to default (hourly)
	cronExpr := s.config.RssSyncIntervalCron
	if cronExpr == "" {
		cronExpr = "0 0 * * * *" // Default: every hour
	}

	// Add RSS sync job with configured cron expression
	_, err := s.cron.AddFunc(cronExpr, func() {
		s.syncRSSFeeds()
	})
	if err != nil {
		log.Printf("Failed to add RSS sync job: %v", err)
	} else {
		log.Printf("RSS sync job scheduled with cron expression: %s", cronExpr)
	}
}

// syncRSSFeeds syncs all company RSS feeds
func (s *Scheduler) syncRSSFeeds() {
	log.Println("Starting RSS feed sync...")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	count, err := rss.SyncAllCompanyFeeds(ctx, s.client)
	if err != nil {
		log.Printf("Error syncing RSS feeds: %v", err)
		return
	}

	log.Printf("RSS feed sync completed. Added %d new postings", count)
}

