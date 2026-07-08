package model

import "time"

// URL represents a shortened URL mapping.
type URL struct {
	ID          int64      `json:"id"`
	ShortCode   string     `json:"short_code"`
	OriginalURL string     `json:"original_url"`
	CreatedAt   time.Time  `json:"created_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	IsCustom    bool       `json:"is_custom"`
	UserID      *string    `json:"user_id,omitempty"`
	ClickCount  int        `json:"click_count"`
}

// Click represents a single click event for analytics.
type Click struct {
	ID         int64     `json:"id"`
	ShortCode  string    `json:"short_code"`
	ClickedAt  time.Time `json:"clicked_at"`
	UserAgent  string    `json:"user_agent,omitempty"`
	IPAddress  string    `json:"ip_address,omitempty"`
	Referer    string    `json:"referer,omitempty"`
	Country    string    `json:"country,omitempty"`
	DeviceType string    `json:"device_type,omitempty"`
}

// ShortenRequest is the request body for creating a short URL.
type ShortenRequest struct {
	URL            string `json:"url"`
	CustomAlias    string `json:"custom_alias,omitempty"`
	ExpiresInHours int    `json:"expires_in_hours,omitempty"`
}

// ShortenResponse is the response after creating a short URL.
type ShortenResponse struct {
	ShortURL    string     `json:"short_url"`
	ShortCode   string     `json:"short_code"`
	OriginalURL string     `json:"original_url"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

// StatsResponse holds analytics data for a short URL.
type StatsResponse struct {
	ShortCode   string         `json:"short_code"`
	OriginalURL string         `json:"original_url"`
	TotalClicks int            `json:"total_clicks"`
	ClicksByDay []ClickByDay   `json:"clicks_by_day"`
	TopReferers []string       `json:"top_referers"`
	Devices     map[string]int `json:"devices"`
}

// ClickByDay represents aggregated clicks per day.
type ClickByDay struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

// ErrorResponse is a standard JSON error body.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
