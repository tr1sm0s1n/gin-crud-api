package main

import "time"

type (
	Certificate struct {
		ID     int       `json:"id" binding:"required"`
		Name   string    `json:"name" binding:"required"`
		Course string    `json:"course" binding:"required"`
		Grade  string    `json:"grade" binding:"required"`
		Date   time.Time `json:"date" binding:"required"`

		Status     string  `json:"status" binding:"required"`
		Score      float64 `json:"score" binding:"required"`
		Percentage float64 `json:"percentage" binding:"required"`
		Rank       int     `json:"rank" binding:"required"`
		IsVerified bool    `json:"is_verified"`
		IsExpired  bool    `json:"is_expired"`

		IssueDate  time.Time `json:"issue_date" binding:"required"`
		ExpiryDate time.Time `json:"expiry_date" binding:"required"`
		CreatedAt  time.Time `json:"created_at" binding:"required"`
		UpdatedAt  time.Time `json:"updated_at" binding:"required"`

		Student Student `json:"student" binding:"required"`
		Issuer  Issuer  `json:"issuer" binding:"required"`

		Tags       []string `json:"tags" binding:"required"`
		Categories []string `json:"categories" binding:"required"`
		Skills     []string `json:"skills" binding:"required"`

		Modules []Module `json:"modules" binding:"required"`

		Achievements []Achievement `json:"achievements" binding:"required"`

		Documents []Document `json:"documents" binding:"required"`

		Metadata map[string]string `json:"metadata"`

		Statistics struct {
			TotalStudents  int     `json:"total_students"`
			AverageScore   float64 `json:"average_score"`
			HighestScore   float64 `json:"highest_score"`
			CompletionRate float64 `json:"completion_rate"`
			GlobalRanking  int     `json:"global_ranking"`
			CountryRanking int     `json:"country_ranking"`
		} `json:"statistics"`

		Settings struct {
			ShowScore     bool `json:"show_score"`
			ShowRank      bool `json:"show_rank"`
			Downloadable  bool `json:"downloadable"`
			PublicProfile bool `json:"public_profile"`
		} `json:"settings"`

		Links struct {
			CertificateURL string `json:"certificate_url"`
			VerifyURL      string `json:"verify_url"`
			DownloadURL    string `json:"download_url"`
		} `json:"links"`

		CustomFields map[string]any `json:"custom_fields"`

		AuditTrail []struct {
			Action    string    `json:"action"`
			User      string    `json:"user"`
			IPAddress string    `json:"ip_address"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"audit_trail"`
	}

	Student struct {
		ID        int      `json:"id" binding:"required"`
		Name      string   `json:"name" binding:"required"`
		Email     string   `json:"email" binding:"required"`
		Phone     string   `json:"phone" binding:"required"`
		Country   string   `json:"country" binding:"required"`
		City      string   `json:"city" binding:"required"`
		Interests []string `json:"interests" binding:"required"`
	}

	Issuer struct {
		ID      int    `json:"id" binding:"required"`
		Name    string `json:"name" binding:"required"`
		Website string `json:"website"`
		Logo    string `json:"logo"`
		Email   string `json:"email" binding:"required"`
	}

	Module struct {
		ID          int       `json:"id" binding:"required"`
		Name        string    `json:"name" binding:"required"`
		Score       float64   `json:"score" binding:"required"`
		MaxScore    float64   `json:"max_score" binding:"required"`
		Passed      bool      `json:"passed" binding:"required"`
		CompletedAt time.Time `json:"completed_at" binding:"required"`
	}

	Achievement struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		BadgeURL    string `json:"badge_url"`
	}

	Document struct {
		ID   int    `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
		Type string `json:"type" binding:"required"`
		Size int64  `json:"size" binding:"required"`
		URL  string `json:"url" binding:"required"`
	}
)
