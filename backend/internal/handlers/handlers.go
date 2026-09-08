package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"

	"vadikara-backend/internal/database"
	"vadikara-backend/internal/middleware"
	"vadikara-backend/internal/models"
)

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// ---------------- AUTH HANDLERS ----------------

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user models.User
	var hashedPwd string
	row := database.DB.QueryRow(`SELECT id, username, email, password_hash, role, created_at FROM users WHERE email = ?`, req.Email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &hashedPwd, &user.Role, &user.CreatedAt)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "Email atau kata sandi tidak valid")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(req.Password)); err != nil {
		writeError(w, http.StatusUnauthorized, "Email atau kata sandi tidak valid")
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal membuat token otorisasi")
		return
	}

	writeJSON(w, http.StatusOK, models.LoginResponse{
		Token: token,
		User:  user,
	})
}

func Me(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	if !ok {
		writeError(w, http.StatusUnauthorized, "Sesi tidak valid")
		return
	}

	var user models.User
	row := database.DB.QueryRow(`SELECT id, username, email, role, created_at FROM users WHERE id = ?`, claims.UserID)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt); err != nil {
		writeError(w, http.StatusNotFound, "Pengguna tidak ditemukan")
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// Fitur Penghapusan Akun Admin
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	if !ok {
		writeError(w, http.StatusUnauthorized, "Sesi tidak valid")
		return
	}

	var req models.DeleteAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Payload tidak valid")
		return
	}

	// Verifikasi kata sandi sebelum hapus
	var hashedPwd string
	err := database.DB.QueryRow(`SELECT password_hash FROM users WHERE id = ?`, claims.UserID).Scan(&hashedPwd)
	if err != nil {
		writeError(w, http.StatusNotFound, "Akun tidak ditemukan")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(req.Password)); err != nil {
		writeError(w, http.StatusForbidden, "Kata sandi konfirmasi salah")
		return
	}

	// Cek apakah ini satu-satunya akun
	var count int
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&count)
	if count <= 1 {
		writeError(w, http.StatusBadRequest, "Tidak dapat menghapus satu-satunya akun administrator sistem. Buat akun admin lain terlebih dahulu.")
		return
	}

	_, err = database.DB.Exec(`DELETE FROM users WHERE id = ?`, claims.UserID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal menghapus akun: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Akun berhasil dihapus secara permanen dari sistem.",
	})
}

// ---------------- STATS HANDLER ----------------

func GetStats(w http.ResponseWriter, r *http.Request) {
	var stats models.DashboardStats
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM projects`).Scan(&stats.TotalProjects)
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM services`).Scan(&stats.TotalServices)
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM articles`).Scan(&stats.TotalArticles)
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM messages`).Scan(&stats.TotalInquiries)
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM messages WHERE status = 'new'`).Scan(&stats.NewInquiries)
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM complaints`).Scan(&stats.TotalComplaints)
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM complaints WHERE status = 'pending'`).Scan(&stats.PendingComplaints)

	writeJSON(w, http.StatusOK, stats)
}

// ---------------- PROFILE HANDLERS ----------------

func GetProfile(w http.ResponseWriter, r *http.Request) {
	var p models.CompanyProfile
	var skillsJSON string

	row := database.DB.QueryRow(`SELECT 
		id, dev_name, dev_title, dev_bio, dev_skills,
		company_name, company_legal_name, company_nib, company_tagline, company_description,
		email, phone, whatsapp, address, github_url, linkedin_url,
		years_of_experience, projects_completed, satisfied_clients, code_uptime_percent, updated_at
	FROM company_profile ORDER BY id LIMIT 1`)

	err := row.Scan(
		&p.ID, &p.DevName, &p.DevTitle, &p.DevBio, &skillsJSON,
		&p.CompanyName, &p.CompanyLegalName, &p.CompanyNIB, &p.CompanyTagline, &p.CompanyDescription,
		&p.Email, &p.Phone, &p.WhatsApp, &p.Address, &p.GithubURL, &p.LinkedinURL,
		&p.YearsOfExperience, &p.ProjectsCompleted, &p.SatisfiedClients, &p.CodeUptimePercent, &p.UpdatedAt,
	)

	if err != nil {
		writeError(w, http.StatusNotFound, "Profil tidak ditemukan")
		return
	}

	_ = json.Unmarshal([]byte(skillsJSON), &p.DevSkills)
	writeJSON(w, http.StatusOK, p)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var p models.CompanyProfile
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	skillsJSON, _ := json.Marshal(p.DevSkills)

	_, err := database.DB.Exec(`UPDATE company_profile SET
		dev_name = ?, dev_title = ?, dev_bio = ?, dev_skills = ?,
		company_name = ?, company_legal_name = ?, company_nib = ?, company_tagline = ?, company_description = ?,
		email = ?, phone = ?, whatsapp = ?, address = ?, github_url = ?, linkedin_url = ?,
		years_of_experience = ?, projects_completed = ?, satisfied_clients = ?, code_uptime_percent = ?,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = (SELECT id FROM company_profile ORDER BY id LIMIT 1)`,
		p.DevName, p.DevTitle, p.DevBio, string(skillsJSON),
		p.CompanyName, p.CompanyLegalName, p.CompanyNIB, p.CompanyTagline, p.CompanyDescription,
		p.Email, p.Phone, p.WhatsApp, p.Address, p.GithubURL, p.LinkedinURL,
		p.YearsOfExperience, p.ProjectsCompleted, p.SatisfiedClients, p.CodeUptimePercent,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui profil: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Profil berhasil diperbarui"})
}

// ---------------- PROJECTS HANDLERS ----------------

func GetProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT 
		id, title, slug, description, full_overview, category, tech_stack,
		live_url, github_url, image_url, featured, architecture_summary, metrics, sort_order, created_at
	FROM projects ORDER BY sort_order ASC, id DESC`)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil daftar proyek")
		return
	}
	defer rows.Close()

	projects := make([]models.Project, 0)
	for rows.Next() {
		var p models.Project
		var techJSON string
		if err := rows.Scan(
			&p.ID, &p.Title, &p.Slug, &p.Description, &p.FullOverview, &p.Category, &techJSON,
			&p.LiveURL, &p.GithubURL, &p.ImageURL, &p.Featured, &p.ArchitectureSummary, &p.Metrics, &p.SortOrder, &p.CreatedAt,
		); err == nil {
			_ = json.Unmarshal([]byte(techJSON), &p.TechStack)
			projects = append(projects, p)
		}
	}

	writeJSON(w, http.StatusOK, projects)
}

func GetProjectBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	var p models.Project
	var techJSON string

	row := database.DB.QueryRow(`SELECT 
		id, title, slug, description, full_overview, category, tech_stack,
		live_url, github_url, image_url, featured, architecture_summary, metrics, sort_order, created_at
	FROM projects WHERE slug = ?`, slug)

	err := row.Scan(
		&p.ID, &p.Title, &p.Slug, &p.Description, &p.FullOverview, &p.Category, &techJSON,
		&p.LiveURL, &p.GithubURL, &p.ImageURL, &p.Featured, &p.ArchitectureSummary, &p.Metrics, &p.SortOrder, &p.CreatedAt,
	)

	if err != nil {
		writeError(w, http.StatusNotFound, "Proyek tidak ditemukan")
		return
	}

	_ = json.Unmarshal([]byte(techJSON), &p.TechStack)
	writeJSON(w, http.StatusOK, p)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if p.Slug == "" {
		p.Slug = strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
	}

	techJSON, _ := json.Marshal(p.TechStack)
	res, err := database.DB.Exec(`INSERT INTO projects (
		title, slug, description, full_overview, category, tech_stack,
		live_url, github_url, image_url, featured, architecture_summary, metrics, sort_order
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.Title, p.Slug, p.Description, p.FullOverview, p.Category, string(techJSON),
		p.LiveURL, p.GithubURL, p.ImageURL, p.Featured, p.ArchitectureSummary, p.Metrics, p.SortOrder,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal membuat proyek: "+err.Error())
		return
	}

	p.ID, _ = res.LastInsertId()
	writeJSON(w, http.StatusCreated, p)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var p models.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	techJSON, _ := json.Marshal(p.TechStack)
	_, err = database.DB.Exec(`UPDATE projects SET
		title = ?, slug = ?, description = ?, full_overview = ?, category = ?, tech_stack = ?,
		live_url = ?, github_url = ?, image_url = ?, featured = ?, architecture_summary = ?, metrics = ?, sort_order = ?
		WHERE id = ?`,
		p.Title, p.Slug, p.Description, p.FullOverview, p.Category, string(techJSON),
		p.LiveURL, p.GithubURL, p.ImageURL, p.Featured, p.ArchitectureSummary, p.Metrics, p.SortOrder, id,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui proyek: "+err.Error())
		return
	}

	p.ID = id
	writeJSON(w, http.StatusOK, p)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	_, err := database.DB.Exec(`DELETE FROM projects WHERE id = ?`, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal menghapus proyek: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Proyek berhasil dihapus"})
}

// ---------------- SERVICES HANDLERS ----------------

func GetServices(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT 
		id, title, slug, icon, summary, capabilities, tech_stack, deliverables, sort_order
	FROM services ORDER BY sort_order ASC, id ASC`)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil layanan")
		return
	}
	defer rows.Close()

	services := make([]models.Service, 0)
	for rows.Next() {
		var s models.Service
		var capsJSON, techJSON string
		if err := rows.Scan(
			&s.ID, &s.Title, &s.Slug, &s.Icon, &s.Summary, &capsJSON, &techJSON, &s.Deliverables, &s.SortOrder,
		); err == nil {
			_ = json.Unmarshal([]byte(capsJSON), &s.Capabilities)
			_ = json.Unmarshal([]byte(techJSON), &s.TechStack)
			services = append(services, s)
		}
	}

	writeJSON(w, http.StatusOK, services)
}

func CreateService(w http.ResponseWriter, r *http.Request) {
	var s models.Service
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if s.Slug == "" {
		s.Slug = strings.ToLower(strings.ReplaceAll(s.Title, " ", "-"))
	}

	capsJSON, _ := json.Marshal(s.Capabilities)
	techJSON, _ := json.Marshal(s.TechStack)

	res, err := database.DB.Exec(`INSERT INTO services (
		title, slug, icon, summary, capabilities, tech_stack, deliverables, sort_order
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		s.Title, s.Slug, s.Icon, s.Summary, string(capsJSON), string(techJSON), s.Deliverables, s.SortOrder,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal menambah layanan: "+err.Error())
		return
	}

	s.ID, _ = res.LastInsertId()
	writeJSON(w, http.StatusCreated, s)
}

func UpdateService(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var s models.Service
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	capsJSON, _ := json.Marshal(s.Capabilities)
	techJSON, _ := json.Marshal(s.TechStack)

	_, err := database.DB.Exec(`UPDATE services SET
		title = ?, slug = ?, icon = ?, summary = ?, capabilities = ?, tech_stack = ?, deliverables = ?, sort_order = ?
		WHERE id = ?`,
		s.Title, s.Slug, s.Icon, s.Summary, string(capsJSON), string(techJSON), s.Deliverables, s.SortOrder, id,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui layanan: "+err.Error())
		return
	}

	s.ID = id
	writeJSON(w, http.StatusOK, s)
}

func DeleteService(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	_, err := database.DB.Exec(`DELETE FROM services WHERE id = ?`, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal menghapus layanan: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Layanan berhasil dihapus"})
}

// ---------------- ARTICLES HANDLERS ----------------

func GetArticles(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT 
		id, title, slug, excerpt, content, category, read_time, author, tags, published_at
	FROM articles ORDER BY id DESC`)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil artikel")
		return
	}
	defer rows.Close()

	articles := make([]models.Article, 0)
	for rows.Next() {
		var a models.Article
		var tagsJSON string
		if err := rows.Scan(
			&a.ID, &a.Title, &a.Slug, &a.Excerpt, &a.Content, &a.Category, &a.ReadTime, &a.Author, &tagsJSON, &a.PublishedAt,
		); err == nil {
			_ = json.Unmarshal([]byte(tagsJSON), &a.Tags)
			articles = append(articles, a)
		}
	}

	writeJSON(w, http.StatusOK, articles)
}

func GetArticleBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	var a models.Article
	var tagsJSON string

	row := database.DB.QueryRow(`SELECT 
		id, title, slug, excerpt, content, category, read_time, author, tags, published_at
	FROM articles WHERE slug = ?`, slug)

	err := row.Scan(
		&a.ID, &a.Title, &a.Slug, &a.Excerpt, &a.Content, &a.Category, &a.ReadTime, &a.Author, &tagsJSON, &a.PublishedAt,
	)

	if err != nil {
		writeError(w, http.StatusNotFound, "Artikel tidak ditemukan")
		return
	}

	_ = json.Unmarshal([]byte(tagsJSON), &a.Tags)
	writeJSON(w, http.StatusOK, a)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var a models.Article
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if a.Slug == "" {
		a.Slug = strings.ToLower(strings.ReplaceAll(a.Title, " ", "-"))
	}
	if a.Author == "" {
		a.Author = "Vadikara"
	}

	tagsJSON, _ := json.Marshal(a.Tags)
	res, err := database.DB.Exec(`INSERT INTO articles (
		title, slug, excerpt, content, category, read_time, author, tags
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		a.Title, a.Slug, a.Excerpt, a.Content, a.Category, a.ReadTime, a.Author, string(tagsJSON),
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal membuat artikel: "+err.Error())
		return
	}

	a.ID, _ = res.LastInsertId()
	writeJSON(w, http.StatusCreated, a)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var a models.Article
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	tagsJSON, _ := json.Marshal(a.Tags)
	_, err := database.DB.Exec(`UPDATE articles SET
		title = ?, slug = ?, excerpt = ?, content = ?, category = ?, read_time = ?, author = ?, tags = ?
		WHERE id = ?`,
		a.Title, a.Slug, a.Excerpt, a.Content, a.Category, a.ReadTime, a.Author, string(tagsJSON), id,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui artikel: "+err.Error())
		return
	}

	a.ID = id
	writeJSON(w, http.StatusOK, a)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	_, err := database.DB.Exec(`DELETE FROM articles WHERE id = ?`, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal menghapus artikel: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Artikel berhasil dihapus"})
}

// ---------------- MESSAGES (CONTACT PROPOSALS) HANDLERS ----------------

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m models.Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if m.Name == "" || m.Email == "" || m.Message == "" {
		writeError(w, http.StatusBadRequest, "Nama, email, dan pesan wajib diisi")
		return
	}

	m.Status = "new"
	res, err := database.DB.Exec(`INSERT INTO messages (
		name, email, organization, budget_tier, service_interest, message, status
	) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		m.Name, m.Email, m.Organization, m.BudgetTier, m.ServiceInterest, m.Message, m.Status,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengirim pesan: "+err.Error())
		return
	}

	m.ID, _ = res.LastInsertId()
	m.CreatedAt = time.Now()

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Pesan / Proposal penawaran Anda berhasil terkirim ke PT. Vadikara Rekayasa Peranti. Tim kami akan segera menindaklanjuti.",
		"data":    m,
	})
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT 
		id, name, email, organization, budget_tier, service_interest, message, status, created_at
	FROM messages ORDER BY id DESC`)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil pesan masuk")
		return
	}
	defer rows.Close()

	messages := make([]models.Message, 0)
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(
			&m.ID, &m.Name, &m.Email, &m.Organization, &m.BudgetTier, &m.ServiceInterest, &m.Message, &m.Status, &m.CreatedAt,
		); err == nil {
			messages = append(messages, m)
		}
	}

	writeJSON(w, http.StatusOK, messages)
}

func UpdateMessageStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var payload struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	_, err := database.DB.Exec(`UPDATE messages SET status = ? WHERE id = ?`, payload.Status, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui status: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Status pesan berhasil diperbarui"})
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	_, err := database.DB.Exec(`DELETE FROM messages WHERE id = ?`, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal menghapus pesan: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Pesan berhasil dihapus"})
}

// ---------------- COMPLAINTS (PENGAJUAN KELUHAN) HANDLERS ----------------

func CreateComplaint(w http.ResponseWriter, r *http.Request) {
	var c models.Complaint
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if c.Name == "" || c.Email == "" || c.Subject == "" || c.Description == "" {
		writeError(w, http.StatusBadRequest, "Nama, email, perihal, dan deskripsi keluhan wajib diisi")
		return
	}

	// Generate unique formal ticket number
	rNum := rand.Intn(9000) + 1000
	c.TicketNumber = fmt.Sprintf("VK-CL-%d-%04d", time.Now().Year(), rNum)
	c.Status = "pending"

	res, err := database.DB.Exec(`INSERT INTO complaints (
		ticket_number, name, email, phone, category, subject, description, status, resolution_notes
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		c.TicketNumber, c.Name, c.Email, c.Phone, c.Category, c.Subject, c.Description, c.Status, "",
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengajukan keluhan: "+err.Error())
		return
	}

	c.ID, _ = res.LastInsertId()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"message":       "Pengajuan keluhan resmi Anda telah berhasil dicatat.",
		"ticket_number": c.TicketNumber,
		"complaint":     c,
	})
}

func GetComplaintByTicket(w http.ResponseWriter, r *http.Request) {
	ticket := chi.URLParam(r, "ticket")
	var c models.Complaint

	row := database.DB.QueryRow(`SELECT 
		id, ticket_number, name, email, phone, category, subject, description, status, resolution_notes, created_at, updated_at
	FROM complaints WHERE ticket_number = ?`, ticket)

	err := row.Scan(
		&c.ID, &c.TicketNumber, &c.Name, &c.Email, &c.Phone, &c.Category, &c.Subject, &c.Description, &c.Status, &c.ResolutionNotes, &c.CreatedAt, &c.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, http.StatusNotFound, "Nomor tiket pengaduan tidak ditemukan")
			return
		}
		writeError(w, http.StatusInternalServerError, "Gagal mengambil data pengaduan: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, c)
}

func GetComplaints(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT 
		id, ticket_number, name, email, phone, category, subject, description, status, resolution_notes, created_at, updated_at
	FROM complaints ORDER BY id DESC`)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil daftar pengaduan")
		return
	}
	defer rows.Close()

	complaints := make([]models.Complaint, 0)
	for rows.Next() {
		var c models.Complaint
		if err := rows.Scan(
			&c.ID, &c.TicketNumber, &c.Name, &c.Email, &c.Phone, &c.Category, &c.Subject, &c.Description, &c.Status, &c.ResolutionNotes, &c.CreatedAt, &c.UpdatedAt,
		); err == nil {
			complaints = append(complaints, c)
		}
	}

	writeJSON(w, http.StatusOK, complaints)
}

func UpdateComplaint(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var payload struct {
		Status          string `json:"status"`
		ResolutionNotes string `json:"resolution_notes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	_, err := database.DB.Exec(`UPDATE complaints SET
		status = ?, resolution_notes = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?`, payload.Status, payload.ResolutionNotes, id)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui tiket pengaduan: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Status dan catatan penanganan pengaduan berhasil diperbarui"})
}

func DeleteComplaint(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	_, err := database.DB.Exec(`DELETE FROM complaints WHERE id = ?`, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal menghapus data pengaduan: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Data pengaduan berhasil dihapus"})
}

// ---------------- DATA DELETION REQUESTS (UU PDP) HANDLERS ----------------

func CreateDataDeletionRequest(w http.ResponseWriter, r *http.Request) {
	var req models.DataDeletionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Name == "" || req.Email == "" {
		writeError(w, http.StatusBadRequest, "Nama dan email wajib diisi untuk verifikasi penghapusan data")
		return
	}

	res, err := database.DB.Exec(`INSERT INTO data_deletion_requests (name, email, reason, status) VALUES (?, ?, ?, ?)`,
		req.Name, req.Email, req.Reason, "pending")
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memproses permohonan: "+err.Error())
		return
	}

	req.ID, _ = res.LastInsertId()
	req.CreatedAt = time.Now()
	req.Status = "pending"

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Permohonan penghapusan data pribadi Anda berdasarkan UU PDP telah kami terima dan akan diverifikasi dalam 3x24 jam kerja.",
		"request": req,
	})
}

func GetDataDeletionRequests(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT id, name, email, reason, status, created_at FROM data_deletion_requests ORDER BY id DESC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil daftar permohonan")
		return
	}
	defer rows.Close()

	list := make([]models.DataDeletionRequest, 0)
	for rows.Next() {
		var d models.DataDeletionRequest
		if err := rows.Scan(&d.ID, &d.Name, &d.Email, &d.Reason, &d.Status, &d.CreatedAt); err == nil {
			list = append(list, d)
		}
	}

	writeJSON(w, http.StatusOK, list)
}

func UpdateDataDeletionStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var payload struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	_, err := database.DB.Exec(`UPDATE data_deletion_requests SET status = ? WHERE id = ?`, payload.Status, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memperbarui status: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Status permohonan data deletion berhasil diperbarui"})
}

// ---------------- SITE CONTENT CMS HANDLERS ----------------

func GetSiteContent(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`SELECT section_key, content_json FROM site_content`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil data konten: "+err.Error())
		return
	}
	defer rows.Close()

	result := make(map[string]json.RawMessage)
	for rows.Next() {
		var key, val string
		if err := rows.Scan(&key, &val); err == nil {
			result[key] = json.RawMessage(val)
		}
	}

	writeJSON(w, http.StatusOK, result)
}

func UpdateSiteContent(w http.ResponseWriter, r *http.Request) {
	var payload map[string]json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "Payload tidak valid")
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal memulai transaksi: "+err.Error())
		return
	}
	defer tx.Rollback()

	for key, val := range payload {
		_, err := tx.Exec(`INSERT INTO site_content (section_key, content_json, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP)
			ON CONFLICT(section_key) DO UPDATE SET content_json = excluded.content_json, updated_at = CURRENT_TIMESTAMP`,
			key, string(val))
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Gagal menyimpan konten '"+key+"': "+err.Error())
			return
		}
	}

	if err := tx.Commit(); err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal menyimpan perubahan: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Konten website berhasil diperbarui",
	})
}

