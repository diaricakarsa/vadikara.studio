package models

import "time"

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type CompanyProfile struct {
	ID                 int64     `json:"id"`
	DevName            string    `json:"dev_name"`
	DevTitle           string    `json:"dev_title"`
	DevBio             string    `json:"dev_bio"`
	DevSkills          []string  `json:"dev_skills"`
	CompanyName        string    `json:"company_name"`
	CompanyLegalName   string    `json:"company_legal_name"`
	CompanyNIB         string    `json:"company_nib"`
	CompanyTagline     string    `json:"company_tagline"`
	CompanyDescription string    `json:"company_description"`
	Email              string    `json:"email"`
	Phone              string    `json:"phone"`
	WhatsApp           string    `json:"whatsapp"`
	Address            string    `json:"address"`
	GithubURL          string    `json:"github_url"`
	LinkedinURL        string    `json:"linkedin_url"`
	YearsOfExperience  int       `json:"years_of_experience"`
	ProjectsCompleted  int       `json:"projects_completed"`
	SatisfiedClients   int       `json:"satisfied_clients"`
	CodeUptimePercent  float64   `json:"code_uptime_percent"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type Project struct {
	ID                  int64     `json:"id"`
	Title               string    `json:"title"`
	Slug                string    `json:"slug"`
	Description         string    `json:"description"`
	FullOverview        string    `json:"full_overview"`
	Category            string    `json:"category"`
	TechStack           []string  `json:"tech_stack"`
	LiveURL             string    `json:"live_url"`
	GithubURL           string    `json:"github_url"`
	ImageURL            string    `json:"image_url"`
	Featured            bool      `json:"featured"`
	ArchitectureSummary string    `json:"architecture_summary"`
	Metrics             string    `json:"metrics"`
	SortOrder           int       `json:"sort_order"`
	CreatedAt           time.Time `json:"created_at"`
}

type Service struct {
	ID           int64    `json:"id"`
	Title        string   `json:"title"`
	Slug         string   `json:"slug"`
	Icon         string   `json:"icon"`
	Summary      string   `json:"summary"`
	Capabilities []string `json:"capabilities"`
	TechStack    []string `json:"tech_stack"`
	Deliverables string   `json:"deliverables"`
	SortOrder    int      `json:"sort_order"`
}

type Article struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Excerpt     string    `json:"excerpt"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	ReadTime    string    `json:"read_time"`
	Author      string    `json:"author"`
	Tags        []string  `json:"tags"`
	PublishedAt time.Time `json:"published_at"`
}

type Message struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Organization    string    `json:"organization"`
	BudgetTier      string    `json:"budget_tier"`
	ServiceInterest string    `json:"service_interest"`
	Message         string    `json:"message"`
	Status          string    `json:"status"` // "new", "reviewed", "contacted"
	CreatedAt       time.Time `json:"created_at"`
}

// Complaint / Layanan Pengaduan Resmi
type Complaint struct {
	ID              int64     `json:"id"`
	TicketNumber    string    `json:"ticket_number"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Category        string    `json:"category"` // "SLA", "Kualitas", "Keamanan", "Bug", "Lainnya"
	Subject         string    `json:"subject"`
	Description     string    `json:"description"`
	Status          string    `json:"status"` // "pending", "investigating", "resolved", "rejected"
	ResolutionNotes string    `json:"resolution_notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Data Deletion Request (UU PDP / Hak Penghapusan Data Pribadi)
type DataDeletionRequest struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status"` // "pending", "processed"
	CreatedAt   time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteAccountRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type DashboardStats struct {
	TotalProjects       int `json:"total_projects"`
	TotalServices       int `json:"total_services"`
	TotalArticles       int `json:"total_articles"`
	TotalInquiries      int `json:"total_inquiries"`
	NewInquiries        int `json:"new_inquiries"`
	TotalComplaints     int `json:"total_complaints"`
	PendingComplaints   int `json:"pending_complaints"`
}

// ---------------- DYNAMIC CMS SITE CONTENT ----------------

type HeroContent struct {
	Title           string `json:"title"`
	ReelBadge       string `json:"reel_badge"`
	ReelURL         string `json:"reel_url"`
	ReelTelemetry   string `json:"reel_telemetry"`
	BackgroundImage string `json:"background_image"`
}

type DisciplineItem struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Benchmark   string   `json:"benchmark"`
	Description string   `json:"description"`
	Stack       []string `json:"stack"`
	Image       string   `json:"image"`
}

type ManifestoContent struct {
	TextCol1   string `json:"text_col1"`
	TextCol2   string `json:"text_col2"`
	ButtonText string `json:"button_text"`
	ButtonURL  string `json:"button_url"`
}

type LabExperimentItem struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Desc   string   `json:"desc"`
	Tags   []string `json:"tags"`
}

type LabsContent struct {
	Tag             string              `json:"tag"`
	Title           string              `json:"title"`
	BackgroundImage string              `json:"background_image"`
	Experiments     []LabExperimentItem `json:"experiments"`
}

type ContactContent struct {
	Tag         string   `json:"tag"`
	Heading     string   `json:"heading"`
	Subheading  string   `json:"subheading"`
	DirectEmail string   `json:"direct_email"`
	GithubURL   string   `json:"github_url"`
	TwitterURL  string   `json:"twitter_url"`
	LinkedinURL string   `json:"linkedin_url"`
	Topics      []string `json:"topics"`
}

type SiteContent struct {
	Hero        HeroContent        `json:"hero"`
	Ticker      []string           `json:"ticker"`
	Disciplines []DisciplineItem   `json:"disciplines"`
	Manifesto   ManifestoContent   `json:"manifesto"`
	Labs        LabsContent        `json:"labs"`
	Contact     ContactContent     `json:"contact"`
}

