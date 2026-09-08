package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"vadikara-backend/internal/database"
	"vadikara-backend/internal/handlers"
	"vadikara-backend/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "vadikara.db"
	}

	log.Printf("Initializing SQLite database at %s...", dbPath)
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	r := chi.NewRouter()

	// Global Middlewares
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(60 * time.Second))

	// CORS Setup
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// API Routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"status":    "healthy",
				"system":    "Vadikara Creative Systems Core API",
				"developer": "Vadikara",
				"uptime":    "99.99%",
				"timestamp": time.Now().Format(time.RFC3339),
			})
		})

		// Public Endpoints
		r.Get("/content", handlers.GetSiteContent)
		r.Get("/profile", handlers.GetProfile)
		r.Get("/projects", handlers.GetProjects)
		r.Get("/projects/{slug}", handlers.GetProjectBySlug)
		r.Get("/services", handlers.GetServices)
		r.Get("/articles", handlers.GetArticles)
		r.Get("/articles/{slug}", handlers.GetArticleBySlug)
		r.Post("/messages", handlers.CreateMessage)
		r.Post("/complaints", handlers.CreateComplaint)
		r.Get("/complaints/{ticket}", handlers.GetComplaintByTicket)
		r.Post("/privacy/data-deletion", handlers.CreateDataDeletionRequest)
		r.Post("/auth/login", handlers.Login)

		// Protected Admin / CMS Endpoints
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthRequired)

			r.Get("/admin/me", handlers.Me)
			r.Delete("/admin/account", handlers.DeleteAccount)
			r.Get("/admin/stats", handlers.GetStats)
			r.Put("/admin/profile", handlers.UpdateProfile)
			r.Put("/admin/content", handlers.UpdateSiteContent)

			// Projects
			r.Post("/admin/projects", handlers.CreateProject)
			r.Put("/admin/projects/{id}", handlers.UpdateProject)
			r.Delete("/admin/projects/{id}", handlers.DeleteProject)

			// Services
			r.Post("/admin/services", handlers.CreateService)
			r.Put("/admin/services/{id}", handlers.UpdateService)
			r.Delete("/admin/services/{id}", handlers.DeleteService)

			// Articles
			r.Post("/admin/articles", handlers.CreateArticle)
			r.Put("/admin/articles/{id}", handlers.UpdateArticle)
			r.Delete("/admin/articles/{id}", handlers.DeleteArticle)

			// Messages (Inquiries)
			r.Get("/admin/messages", handlers.GetMessages)
			r.Put("/admin/messages/{id}", handlers.UpdateMessageStatus)
			r.Delete("/admin/messages/{id}", handlers.DeleteMessage)

			// Complaints (Pengaduan & Helpdesk)
			r.Get("/admin/complaints", handlers.GetComplaints)
			r.Put("/admin/complaints/{id}", handlers.UpdateComplaint)
			r.Delete("/admin/complaints/{id}", handlers.DeleteComplaint)

			// Data Deletion Requests (UU PDP)
			r.Get("/admin/data-deletions", handlers.GetDataDeletionRequests)
			r.Put("/admin/data-deletions/{id}", handlers.UpdateDataDeletionStatus)
		})
	})

	// Static SPA File Server
	staticDir := filepath.Join("..", "frontend", "dist")
	if _, err := os.Stat(staticDir); err == nil {
		fileServer(r, "/", http.Dir(staticDir))
	} else {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head><title>Vadikara API Server</title></head>
			<body style="font-family: sans-serif; background: #0f172a; color: #f8fafc; padding: 40px; text-align: center;">
				<h1 style="color: #00f2fe;">PT. Vadikara Rekayasa Peranti</h1>
				<p>Backend API Engine is actively running on :`+port+`.</p>
				<p>Frontend is running on Vite Dev Server (port 5173).</p>
				<a href="http://localhost:5173" style="color: #38bdf8; text-decoration: underline;">Open Frontend Application</a>
			</body>
			</html>
			`)
		})
	}

	log.Printf("PT. Vadikara Rekayasa Peranti Backend listening on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))

		// Check if file exists, if not serve index.html for SPA routing
		f, err := root.Open(strings.TrimPrefix(r.URL.Path, pathPrefix))
		if err != nil {
			// fallback to index.html
			r.URL.Path = "/"
		} else {
			_ = f.Close()
		}
		fs.ServeHTTP(w, r)
	})
}
