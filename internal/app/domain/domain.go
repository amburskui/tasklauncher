package domain

import (
	"time"
)

// Task interface
type Task interface {
	GetCommand() []string
}

// BrowserTask struct
type BrowserTask struct {
	ID          string              `json:"id"`
	Title       string              `json:"title"`
	Application BrowserApplications `json:"browser_application"`
	Arguments   Arguments           `json:"arguments"`
	CreatedAt   *time.Time          `json:"created_at"`
	CreatedBy   string              `json:"created_by"`
	UpdatedAt   *time.Time          `json:"updated_at"`
	UpdatedBy   string              `json:"updated_by"`
	PublishedAt string              `json:"published_at"`
	PublishedBy string              `json:"published_by"`
}

// Application interface
type Application interface {
	GetPath() []string
}

// BrowserApplications is a slice of "BrowserApplication"s
type BrowserApplications []*BrowserApplication

// Arguments is a slice of "string"s
type Arguments []string

// BrowserApplication struct
type BrowserApplication struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}
