package main

import "time"

var certificates = []Certificate{
	{
		ID:         1,
		Name:       "John Doe",
		Course:     "Advanced React",
		Grade:      "A",
		Date:       time.Date(2025, 5, 10, 0, 0, 0, 0, time.UTC),
		Status:     "Completed",
		Score:      95.5,
		Percentage: 95.5,
		Rank:       3,
		IsVerified: true,
		IsExpired:  false,
		IssueDate:  time.Date(2025, 5, 10, 0, 0, 0, 0, time.UTC),
		ExpiryDate: time.Date(2028, 5, 10, 0, 0, 0, 0, time.UTC),
		CreatedAt:  time.Date(2025, 5, 10, 10, 30, 0, 0, time.UTC),
		UpdatedAt:  time.Date(2025, 5, 10, 14, 20, 0, 0, time.UTC),

		Student: Student{
			ID:        101,
			Name:      "John Doe",
			Email:     "john@example.com",
			Phone:     "+1-555-1234",
			Country:   "USA",
			City:      "New York",
			Interests: []string{"React", "TypeScript", "UI Design"},
		},

		Issuer: Issuer{
			ID:      1,
			Name:    "Tech Academy",
			Website: "https://academy.example.com",
			Logo:    "https://picsum.photos/100",
			Email:   "support@academy.example.com",
		},

		Tags:       []string{"frontend", "react", "typescript"},
		Categories: []string{"Web Development", "Frontend"},
		Skills:     []string{"React", "Redux", "Hooks"},

		Modules: []Module{
			{
				ID:          1,
				Name:        "React Fundamentals",
				Score:       98,
				MaxScore:    100,
				Passed:      true,
				CompletedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:          2,
				Name:        "State Management",
				Score:       92,
				MaxScore:    100,
				Passed:      true,
				CompletedAt: time.Date(2025, 4, 15, 0, 0, 0, 0, time.UTC),
			},
		},

		Achievements: []Achievement{
			{
				ID:          1,
				Title:       "Top Performer",
				Description: "Scored above 95%",
				BadgeURL:    "https://picsum.photos/50",
			},
			{
				ID:          2,
				Title:       "Perfect Attendance",
				Description: "Attended all sessions",
				BadgeURL:    "https://picsum.photos/51",
			},
		},

		Documents: []Document{
			{
				ID:   1,
				Name: "Certificate.pdf",
				Type: "pdf",
				Size: 245678,
				URL:  "https://example.com/files/certificate.pdf",
			},
		},

		Metadata: map[string]string{
			"language": "English",
			"track":    "Frontend",
			"region":   "North America",
		},

		CustomFields: map[string]any{
			"github":          "johndoe",
			"yearsExperience": 4,
			"preferredIDE":    "VS Code",
			"available":       true,
		},

		AuditTrail: []struct {
			Action    string    `json:"action"`
			User      string    `json:"user"`
			IPAddress string    `json:"ip_address"`
			Timestamp time.Time `json:"timestamp"`
		}{
			{
				Action:    "created",
				User:      "admin",
				Timestamp: time.Date(2025, 5, 10, 10, 30, 0, 0, time.UTC),
				IPAddress: "192.168.1.10",
			},
			{
				Action:    "verified",
				User:      "reviewer",
				Timestamp: time.Date(2025, 5, 10, 14, 20, 0, 0, time.UTC),
				IPAddress: "192.168.1.20",
			},
		},
	},

	{
		ID:         2,
		Name:       "Jane Smith",
		Course:     "Vue.js Masterclass",
		Grade:      "B+",
		Date:       time.Date(2025, 3, 21, 0, 0, 0, 0, time.UTC),
		Status:     "Completed",
		Score:      87.2,
		Percentage: 87.2,
		Rank:       18,
		IsVerified: true,
		IsExpired:  false,
		IssueDate:  time.Date(2025, 3, 21, 0, 0, 0, 0, time.UTC),
		ExpiryDate: time.Date(2028, 3, 21, 0, 0, 0, 0, time.UTC),
		CreatedAt:  time.Date(2025, 3, 21, 8, 15, 0, 0, time.UTC),
		UpdatedAt:  time.Date(2025, 3, 21, 9, 30, 0, 0, time.UTC),

		Student: Student{
			ID:        102,
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			Phone:     "+44-777-123456",
			Country:   "UK",
			City:      "London",
			Interests: []string{"Vue", "Nuxt", "CSS"},
		},

		Issuer: Issuer{
			ID:      2,
			Name:    "Frontend Institute",
			Website: "https://frontend.example.com",
			Logo:    "https://picsum.photos/101",
			Email:   "hello@frontend.example.com",
		},

		Tags:       []string{"vue", "frontend"},
		Categories: []string{"Web Development"},
		Skills:     []string{"Vue", "Pinia", "Composition API"},

		Modules: []Module{
			{
				ID:          10,
				Name:        "Vue Basics",
				Score:       89,
				MaxScore:    100,
				Passed:      true,
				CompletedAt: time.Date(2025, 2, 25, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:          11,
				Name:        "Nuxt Development",
				Score:       85,
				MaxScore:    100,
				Passed:      true,
				CompletedAt: time.Date(2025, 3, 3, 0, 0, 0, 0, time.UTC),
			},
		},

		Achievements: []Achievement{},
		Documents: []Document{
			{
				ID:   2,
				Name: "vue-certificate.pdf",
				Type: "pdf",
				Size: 345890,
				URL:  "https://example.com/files/vue.pdf",
			},
		},

		Metadata: map[string]string{
			"language": "English",
			"track":    "Vue",
		},

		CustomFields: map[string]any{
			"portfolio": "https://janesmith.dev",
			"mentor":    false,
		},

		AuditTrail: []struct {
			Action    string    `json:"action"`
			User      string    `json:"user"`
			IPAddress string    `json:"ip_address"`
			Timestamp time.Time `json:"timestamp"`
		}{
			{
				Action:    "created",
				User:      "system",
				Timestamp: time.Date(2025, 3, 15, 6, 36, 0, 0, time.UTC),
				IPAddress: "10.0.0.1",
			},
		},
	},
}

func getIndex(id int) (int, bool) {
	for i, cert := range certificates {
		if cert.ID == id {
			return i, true
		}
	}
	return -1, false
}
