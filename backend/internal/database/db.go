package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(1)

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	if err := seedInitialData(db); err != nil {
		log.Printf("Seed error or already seeded: %v", err)
	}

	DB = db
	return db, nil
}

func createTables(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT 'admin',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS company_profile (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			dev_name TEXT NOT NULL,
			dev_title TEXT NOT NULL,
			dev_bio TEXT NOT NULL,
			dev_skills TEXT NOT NULL,
			company_name TEXT NOT NULL,
			company_legal_name TEXT NOT NULL,
			company_nib TEXT NOT NULL,
			company_tagline TEXT NOT NULL,
			company_description TEXT NOT NULL,
			email TEXT NOT NULL,
			phone TEXT NOT NULL,
			whatsapp TEXT NOT NULL,
			address TEXT NOT NULL,
			github_url TEXT NOT NULL,
			linkedin_url TEXT NOT NULL,
			years_of_experience INTEGER DEFAULT 8,
			projects_completed INTEGER DEFAULT 40,
			satisfied_clients INTEGER DEFAULT 25,
			code_uptime_percent REAL DEFAULT 99.99,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			slug TEXT UNIQUE NOT NULL,
			description TEXT NOT NULL,
			full_overview TEXT NOT NULL,
			category TEXT NOT NULL,
			tech_stack TEXT NOT NULL,
			live_url TEXT DEFAULT '',
			github_url TEXT DEFAULT '',
			image_url TEXT DEFAULT '',
			featured BOOLEAN DEFAULT 0,
			architecture_summary TEXT DEFAULT '',
			metrics TEXT DEFAULT '',
			sort_order INTEGER DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS services (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			slug TEXT UNIQUE NOT NULL,
			icon TEXT NOT NULL,
			summary TEXT NOT NULL,
			capabilities TEXT NOT NULL,
			tech_stack TEXT NOT NULL,
			deliverables TEXT NOT NULL,
			sort_order INTEGER DEFAULT 0
		);`,
		`CREATE TABLE IF NOT EXISTS articles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			slug TEXT UNIQUE NOT NULL,
			excerpt TEXT NOT NULL,
			content TEXT NOT NULL,
			category TEXT NOT NULL,
			read_time TEXT NOT NULL,
			author TEXT NOT NULL,
			tags TEXT NOT NULL,
			published_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			organization TEXT DEFAULT '',
			budget_tier TEXT DEFAULT '',
			service_interest TEXT DEFAULT '',
			message TEXT NOT NULL,
			status TEXT DEFAULT 'new',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS complaints (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ticket_number TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			phone TEXT DEFAULT '',
			category TEXT NOT NULL,
			subject TEXT NOT NULL,
			description TEXT NOT NULL,
			status TEXT DEFAULT 'pending',
			resolution_notes TEXT DEFAULT '',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS data_deletion_requests (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			reason TEXT DEFAULT '',
			status TEXT DEFAULT 'pending',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS site_content (
			section_key TEXT PRIMARY KEY,
			content_json TEXT NOT NULL,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}
	return nil
}

func seedInitialData(db *sql.DB) error {
	// 1. Seed Admin User
	var userCount int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&userCount)
	if err == nil && userCount == 0 {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("vadikara2026!"), bcrypt.DefaultCost)
		_, err = db.Exec(`INSERT INTO users (username, email, password_hash, role) VALUES (?, ?, ?, ?)`,
			"vadikara", "admin@vadikara.com", string(hashed), "superadmin")
		if err != nil {
			log.Printf("Failed to seed admin: %v", err)
		} else {
			log.Printf("Admin user created: admin@vadikara.com / vadikara2026!")
		}
	}

	// 2. Seed Profile
	var profileCount int
	err = db.QueryRow("SELECT COUNT(*) FROM company_profile").Scan(&profileCount)
	if err == nil && profileCount == 0 {
		skillsJSON, _ := json.Marshal([]string{
			"Golang (Go)", "Distributed Systems", "Vue.js 3 / Vite", "HTML5 & Modern CSS", 
			"Cloud-Native / K8s", "PostgreSQL & Redis", "gRPC & Protocol Buffers", "Zero-Trust Architecture",
			"Docker & Microservices", "CI/CD & DevOps Automation", "High-Load Optimization", "System Modeling",
		})
		_, err = db.Exec(`INSERT INTO company_profile (
			dev_name, dev_title, dev_bio, dev_skills,
			company_name, company_legal_name, company_nib, company_tagline, company_description,
			email, phone, whatsapp, address, github_url, linkedin_url,
			years_of_experience, projects_completed, satisfied_clients, code_uptime_percent
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"Vadikara",
			"Principal Systems Architect & Lead Software Engineer",
			"Insinyur perangkat lunak dengan fokus pada performa ekstrem, konkurensi Golang, arsitektur sistem terdistribusi, dan aplikasi web modern beranimasi dinamis. Berpengalaman merancang fondasi teknologi yang andal, scalable, dan siap menghadapi jutaan transaksi.",
			string(skillsJSON),
			"PT. Vadikara Rekayasa Peranti",
			"PT. VADIKARA REKAYASA PERANTI",
			"0128934710294 (KBLI 62019 - Aktivitas Pemrograman Perangkat Lunak)",
			"High-Performance Software Engineering & Enterprise Digital Transformation",
			"PT. Vadikara Rekayasa Peranti adalah badan usaha rekayasa teknologi dan pengembangan peranti lunak berstandar industri tinggi. Kami mendampingi institusi korporat, startup skala berkembang, dan instansi strategis dalam membangun infrastruktur digital tahan uji, aman, dan efisien.",
			"contact@vadikara.com",
			"+62 812-8800-9921",
			"6281288009921",
			"Gedung Rekayasa Digital Lt. 7, Jl. TB Simatupang No. 88, Cilandak, Jakarta Selatan 12430",
			"https://github.com/vadikara",
			"https://linkedin.com/company/vadikara-rekayasa-peranti",
			8, 48, 34, 99.99,
		)
	}

	// 3. Seed Services
	var serviceCount int
	err = db.QueryRow("SELECT COUNT(*) FROM services").Scan(&serviceCount)
	if err == nil && serviceCount == 0 {
		services := []struct {
			title        string
			slug         string
			icon         string
			summary      string
			capabilities []string
			techStack    []string
			deliverables string
			sort         int
		}{
			{
				title:   "High-Throughput Backend & Microservices",
				slug:    "high-throughput-backend",
				icon:    "Server",
				summary: "Perancangan backend berkecepatan tinggi berbasis Golang dengan model konkurensi native (Goroutines), gRPC, dan arsitektur event-driven.",
				capabilities: []string{
					"Pemrosesan ratusan ribu Request Per Second (RPS)",
					"Arsitektur Microservices terisolasi dan fault-tolerant",
					"Komunikasi real-time via WebSockets dan gRPC",
					"Caching berlapis (In-memory, Redis, CDN)",
				},
				techStack:    []string{"Golang", "gRPC", "Kafka", "Redis", "PostgreSQL", "Protobuf"},
				deliverables: "Source code production-ready, unit & load test report, Dockerized stack, Swagger/OpenAPI docs.",
				sort:         1,
			},
			{
				title:   "Modern Interactive Web Applications (Vue 3 & HTML5)",
				slug:    "modern-web-applications",
				icon:    "Globe",
				summary: "Pengembangan front-end berkelas dunia dengan animasi dinamis penuh aksi, estetika premium, responsif 100%, dan performa instan.",
				capabilities: []string{
					"Aplikasi Single Page Application (SPA) & Progressive Web App (PWA)",
					"Animasi 60 FPS halus (Canvas, WebGL, CSS Physics)",
					"Design System kustom dan State Management reaktif",
					"Optimasi Core Web Vitals dan SEO teknis terstruktur",
				},
				techStack:    []string{"Vue 3", "Vite", "HTML5", "Modern CSS", "Pinia", "Canvas API"},
				deliverables: "Komponen modular, asset teroptimasi, dokumentasi UI/UX, bundle size minimal.",
				sort:         2,
			},
			{
				title:   "Enterprise Cloud Architecture & DevOps",
				slug:    "enterprise-cloud-devops",
				icon:    "Cpu",
				summary: "Infrastruktur cloud-native yang aman dan elastis dengan otomatisasi deployment CI/CD tanpa downtime (Zero Downtime).",
				capabilities: []string{
					"Orkestrasi Kubernetes (K8s) & Cluster autoscaling",
					"Infrastructure as Code (Terraform, Ansible)",
					"Pipeline CI/CD otomatis dengan pengujian terintegrasi",
					"Observability komprehensif (Prometheus, Grafana, OpenTelemetry)",
				},
				techStack:    []string{"Kubernetes", "Docker", "Terraform", "GitHub Actions", "Prometheus", "AWS/GCP"},
				deliverables: "Manifest IaC, pipeline deployment otomatis, dashboard monitoring 24/7.",
				sort:         3,
			},
			{
				title:   "AI Integration & Automated Reasoning Pipelines",
				slug:    "ai-integration-automation",
				icon:    "Sparkles",
				summary: "Integrasi kecerdasan buatan (LLM, computer vision, predictive telemetry) ke dalam alur kerja bisnis korporasi.",
				capabilities: []string{
					"Pipeline RAG (Retrieval-Augmented Generation) berbasis data internal",
					"Sistem inferensi dokumen otomatis & OCR cerdas",
					"Agentic workflows untuk otomatisasi tugas kompleks",
					"Optimasi latensi inferensi dan mitigasi biaya token API",
				},
				techStack:    []string{"Python/Go Bridge", "Vector DB (Qdrant/Milvus)", "Ollama", "OpenAI / Claude APIs"},
				deliverables: "API wrapper aman, database embedding, pipeline data ingest, audit trail.",
				sort:         4,
			},
			{
				title:   "Resilience, Security & Code Audit",
				slug:    "resilience-security-audit",
				icon:    "ShieldCheck",
				summary: "Audit mendalam terhadap kerentanan kode sumber, optimasi memori/CPU, serta penguatan perimeter keamanan sistem.",
				capabilities: []string{
					"Static & Dynamic Code Security Analysis (SAST/DAST)",
					"Stress testing & chaos engineering simulasi kegagalan",
					"Optimasi alokasi memori dan deteksi goroutine leaks",
					"Implementasi standar Zero-Trust & enkripsi end-to-end",
				},
				techStack:    []string{"Go pprof", "k6 Load Testing", "OWASP ZAP", "SonarQube", "WireGuard"},
				deliverables: "Laporan audit komprehensif, rekomendasi mitigasi langkah demi langkah, patch refactoring.",
				sort:         5,
			},
		}

		for _, s := range services {
			capsJSON, _ := json.Marshal(s.capabilities)
			techJSON, _ := json.Marshal(s.techStack)
			_, _ = db.Exec(`INSERT INTO services (title, slug, icon, summary, capabilities, tech_stack, deliverables, sort_order)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
				s.title, s.slug, s.icon, s.summary, string(capsJSON), string(techJSON), s.deliverables, s.sort)
		}
	}

	// 4. Seed Projects
	var projectCount int
	err = db.QueryRow("SELECT COUNT(*) FROM projects").Scan(&projectCount)
	if err == nil && projectCount == 0 {
		projects := []struct {
			title       string
			slug        string
			desc        string
			overview    string
			category    string
			tech        []string
			live        string
			github      string
			img         string
			featured    bool
			arch        string
			metrics     string
			sort        int
		}{
			{
				title:       "NovaMesh Distributed Engine",
				slug:        "novamesh-distributed-engine",
				desc:        "Backbone perpesanan event-driven berkecepatan ultra-tinggi yang menghubungkan puluhan microservice dengan latency sub-millisecond.",
				overview:    "Dirancang khusus untuk ekosistem transaksi finansial skala masif di mana setiap milidetik sangat berharga. Menggunakan Go runtime dengan optimasi zero-allocation dan custom buffer pool.",
				category:    "Microservices / Cloud",
				tech:        []string{"Golang", "gRPC", "Apache Kafka", "Redis Cluster", "Docker", "Prometheus"},
				live:        "https://novamesh.vadikara.dev",
				github:      "https://github.com/vadikara/novamesh-engine",
				img:         "https://images.unsplash.com/photo-1558494949-ef010cbdcc31?auto=format&fit=crop&w=1200&q=80",
				featured:    true,
				arch:        "Event Sourcing + CQRS pattern, Raft consensus for leader election, Lock-free ring buffer for queue processing.",
				metrics:     "185,000 req/sec • p99 latency < 1.1ms • 99.999% reliability",
				sort:        1,
			},
			{
				title:       "OmniVortex Enterprise Core ERP",
				slug:        "omnivortex-enterprise-core",
				desc:        "Platform ERP modern dan rantai pasok multi-gudang untuk korporasi ritel dengan konsolidasi akuntansi real-time.",
				overview:    "Sistem menyeluruh dari manajemen gudang otomatis, sinkronisasi inventori multi-cabang, hingga otomasi rekonsiliasi perbankan dengan antarmuka Vue 3 berkecepatan tinggi.",
				category:    "Enterprise Systems",
				tech:        []string{"Golang", "Vue 3", "PostgreSQL", "RabbitMQ", "HTML5/CSS", "MinIO"},
				live:        "https://omnivortex.vadikara.dev",
				github:      "https://github.com/vadikara/omnivortex-erp",
				img:         "https://images.unsplash.com/photo-1460925895917-afdab827c52f?auto=format&fit=crop&w=1200&q=80",
				featured:    true,
				arch:        "Modular Monolith transitioning to domain services, Row-level multi-tenant security, WebSocket live ledger stream.",
				metrics:     "IDR 65 Miliar GMV/Bulan • 15.000 Pengguna Aktif Bersamaan",
				sort:        2,
			},
			{
				title:       "AegisGuard Zero-Trust Mesh",
				slug:        "aegisguard-zero-trust",
				desc:        "Gateway keamanan jaringan berbasis kebijakan mikro-segmentasi dan pemantauan anomali lalu lintas data secara instan.",
				overview:    "Menghalangi pergerakan lateral ancaman siber dengan inspeksi paket mendalam dan karantina otomatis terhadap lalu lintas mencurigakan berbasis telemetry bPF.",
				category:    "Cybersecurity",
				tech:        []string{"Golang", "Linux eBPF", "Vue 3", "TimescaleDB", "WireGuard", "Grafana"},
				live:        "https://aegisguard.vadikara.dev",
				github:      "https://github.com/vadikara/aegisguard-mesh",
				img:         "https://images.unsplash.com/photo-1563986768609-322da13575f3?auto=format&fit=crop&w=1200&q=80",
				featured:    true,
				arch:        "Kernel-level packet filtering with eBPF hooks, distributed agent telemetry, decentralized token verification.",
				metrics:     "Zero security breach record • < 0.2% CPU overhead on host nodes",
				sort:        3,
			},
			{
				title:       "NeuralSense Document Reasoning AI",
				slug:        "neuralsense-ai-pipeline",
				desc:        "Pipeline computer vision dan pemrosesan bahasa alami untuk ekstraksi data dokumen legal dan faktur dalam < 400ms.",
				overview:    "Mengubah tumpukan ribuan dokumen fisik dan PDF menjadi data terstruktur berakurasi 99.4% yang siap diintegrasikan langsung ke sistem perbankan.",
				category:    "AI & Cloud",
				tech:        []string{"Python/Go Interop", "Vue 3", "TensorRT", "Redis", "AWS S3", "Tailored OCR"},
				live:        "https://neuralsense.vadikara.dev",
				github:      "https://github.com/vadikara/neuralsense-ai",
				img:         "https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?auto=format&fit=crop&w=1200&q=80",
				featured:    true,
				arch:        "Asynchronous worker pool in Go, GPU-accelerated inference workers, streaming WebSocket status updates.",
				metrics:     "99.4% Extraction Precision • 380ms Rata-rata Pemrosesan per Dokumen",
				sort:        4,
			},
			{
				title:       "FinPulse Real-Time Trading Engine",
				slug:        "finpulse-trading-engine",
				desc:        "Mesin pencocokan orderbook limit & market dengan latensi nano-second dan streaming ticker pasar waktu nyata.",
				overview:    "Dibuat dengan pendekatan memory-mapped files dan binary protocol untuk transaksi aset digital dengan jaminan integritas data 100%.",
				category:    "Enterprise Systems",
				tech:        []string{"Golang", "Vue 3 Reactive UI", "WebSockets", "SQLite In-Memory", "HTML5 Canvas"},
				live:        "https://finpulse.vadikara.dev",
				github:      "https://github.com/vadikara/finpulse-core",
				img:         "https://images.unsplash.com/photo-1642543492481-44e81e3914a7?auto=format&fit=crop&w=1200&q=80",
				featured:    false,
				arch:        "LMAX disruptor-inspired lock-free architecture in Go, append-only transaction ledger, snapshotting.",
				metrics:     "220,000 Matches/Detik • Zero Balance Discrepancy",
				sort:        5,
			},
		}

		for _, p := range projects {
			techJSON, _ := json.Marshal(p.tech)
			_, _ = db.Exec(`INSERT INTO projects (
				title, slug, description, full_overview, category, tech_stack, live_url, github_url, image_url, featured, architecture_summary, metrics, sort_order
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				p.title, p.slug, p.desc, p.overview, p.category, string(techJSON), p.live, p.github, p.img, p.featured, p.arch, p.metrics, p.sort)
		}
	}

	// 5. Seed Articles
	var articleCount int
	err = db.QueryRow("SELECT COUNT(*) FROM articles").Scan(&articleCount)
	if err == nil && articleCount == 0 {
		articles := []struct {
			title    string
			slug     string
			excerpt  string
			content  string
			category string
			readTime string
			author   string
			tags     []string
		}{
			{
				title:    "Membangun Microservice Berkecepatan 180K RPS dengan Golang dan gRPC",
				slug:     "membangun-microservice-180k-rps-golang",
				excerpt:  "Studi kasus arsitektur teknis bagaimana PT. Vadikara Rekayasa Peranti memangkas latensi p99 dan mengoptimalkan GC (Garbage Collector) Go.",
				content:  "Dalam rekayasa sistem terdistribusi berskala masif, performa tinggi bukan sekadar tentang seberapa cepat server memproses satu permintaan, melainkan bagaimana stabilitas p99 dan p99.9 terjaga saat beban lonjakan tiba-tiba terjadi. Pada studi kasus ini, kami mengulas teknik penggunaan sync.Pool untuk memory reuse, zero-copy buffer pooling, dan optimalisasi gRPC multiplexing.",
				category: "Engineering Architecture",
				readTime: "7 menit baca",
				author:   "Vadikara (Lead Architect)",
				tags:     []string{"Golang", "gRPC", "Performance", "Concurrency"},
			},
			{
				title:    "Sinergi Vue 3 dan Golang: Meracik Antarmuka Dinamis Tanpa Mengorbankan Performa",
				slug:     "sinergi-vue-3-dan-golang-antarmuka-dinamis",
				excerpt:  "Mengapa kombinasi Vue 3 Composition API dan Go backend adalah fondasi terbaik untuk dashboard korporat dan aplikasi modern.",
				content:  "Ketika antarmuka memerlukan respons instan, rendering 60 FPS, dan interaksi visual kaya data, arsitektur frontend dan backend harus selaras secara mendasar. Kami mengulas implementasi micro-frontends, caching responsif, dan state management reaktif.",
				category: "Frontend & Fullstack",
				readTime: "5 menit baca",
				author:   "Vadikara",
				tags:     []string{"Vue 3", "Vite", "Web Performance", "Clean Architecture"},
			},
			{
				title:    "Prinsip Rekayasa Perangkat Lunak Tahan Uji di PT. Vadikara Rekayasa Peranti",
				slug:     "prinsip-rekayasa-perangkat-lunak-tahan-uji",
				excerpt:  "Filosofi kami dalam memastikan reliabilitas 99.99%: Dari automated chaos engineering hingga kultur dokumentasi arsitektur.",
				content:  "Bagi kami di PT. Vadikara Rekayasa Peranti, peranti lunak yang baik bukan hanya yang berfungsi saat demo, melainkan yang terus berdiri kokoh saat database utama mengalami failover. Simak 5 pilar rekayasa yang kami terapkan pada seluruh proyek enterprise klien.",
				category: "Corporate Culture",
				readTime: "6 menit baca",
				author:   "Tim Rekayasa Vadikara",
				tags:     []string{"Culture", "DevOps", "Site Reliability", "Corporate"},
			},
		}

		for _, a := range articles {
			tagsJSON, _ := json.Marshal(a.tags)
			_, _ = db.Exec(`INSERT INTO articles (title, slug, excerpt, content, category, read_time, author, tags)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
				a.title, a.slug, a.excerpt, a.content, a.category, a.readTime, a.author, string(tagsJSON))
		}
	}

	// 6. Seed Sample Inquiry
	var messageCount int
	err = db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&messageCount)
	if err == nil && messageCount == 0 {
		_, _ = db.Exec(`INSERT INTO messages (name, email, organization, budget_tier, service_interest, message, status)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
			"Rian Pratama", "rian@megatech-enterprise.co.id", "PT Mega Tech Nusantara", "Enterprise (> IDR 150 Juta)",
			"High-Throughput Backend & Cloud Architecture",
			"Halo Tim PT. Vadikara Rekayasa Peranti, kami sedang membutuhkan perombakan sistem transaksi inti kami ke arsitektur Golang microservices. Mohon waktu untuk diskusi teknis.",
			"new",
		)
	}

	// 7. Seed Sample Complaint / Pengaduan
	var complaintCount int
	err = db.QueryRow("SELECT COUNT(*) FROM complaints").Scan(&complaintCount)
	if err == nil && complaintCount == 0 {
		_, _ = db.Exec(`INSERT INTO complaints (ticket_number, name, email, phone, category, subject, description, status, resolution_notes)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"VK-CL-2026-001",
			"Dimas Wicaksono",
			"dimas@klien-partner.id",
			"+62 811-2233-4455",
			"Kendala Teknis",
			"Permintaan Penyesuaian Timeout API Gateway pada Staging",
			"Mohon konfirmasi konfigurasi keep-alive timeout pada ingress gateway environment staging, kami mendapati intermittent timeout saat beban load testing 5.000 RPS.",
			"investigating",
			"Tim infrastruktur sedang mengaudit parameter keepalive dan buffer pool reverse proxy.",
		)
	}

	// 8. Seed Dynamic CMS Site Content
	{
		heroJSON := `{
			"title": "Rekayasawan sistem kreatif &\nperanti lunak terdistribusi",
			"reel_badge": "VADIKARA // SHOWCASE REKAYASA SISTEM 2026",
			"reel_url": "/Scene.mp4",
			"reel_telemetry": "Konkurensi Tinggi Golang • Web Kreatif 60 FPS • Sistem Terdistribusi Sub-milidetik",
			"background_image": "/src/assets/fluid_droplets.jpg"
		}`

		tickerJSON := `[
			"vadikara",
			"vadikara.dev",
			"rekayasa kreatif",
			"sistem terdistribusi",
			"antarmuka 60 fps",
			"inti golang",
			"arsitektur awan"
		]`

		disciplinesJSON := `[
			{
				"id": 1,
				"title": "Sistem Terdistribusi",
				"benchmark": "p99 < 1.2ms",
				"description": "Arsitektur klaster berdaya tahan tinggi dan toleran kesalahan dengan konsensus Raft, mitigasi split-brain, serta ketiadaan single point of failure.",
				"stack": ["Golang 1.26", "gRPC", "Raft", "Protobuf v3"],
				"image": "/src/assets/dev_workspace.jpg"
			},
			{
				"id": 2,
				"title": "Konkurensi Tinggi Go",
				"benchmark": "350k+ RPS",
				"description": "Saluran eksekusi konkurensi berbasis Goroutine native, ring-buffer bebas kunci (lock-free), dan alokasi memori mendekati nol untuk beban ratusan ribu RPS.",
				"stack": ["Go Runtime", "Sync.Pool", "Channels", "pprof"],
				"image": "/src/assets/dev_workspace.jpg"
			},
			{
				"id": 3,
				"title": "Infrastruktur Cloud",
				"benchmark": "Zero-Downtime",
				"description": "Orkestrasi klaster Kubernetes deklaratif, kontainer distroless minimal risiko serangan, serta autoscaling horizontal otomatis.",
				"stack": ["Kubernetes", "Docker", "Terraform", "K3s"],
				"image": "/src/assets/dev_workspace.jpg"
			},
			{
				"id": 4,
				"title": "Web Kreatif & 60 FPS",
				"benchmark": "Sub-detik TTFB",
				"description": "Pengalaman web sensorik dengan animasi fisika cairan, tipografi kinetik, dan akselerasi GPU canvas yang deterministik dan konsisten 60 FPS.",
				"stack": ["Vue 3", "Modern CSS", "Vite", "Canvas / WebGL"],
				"image": "/src/assets/creative_poster.jpg"
			},
			{
				"id": 5,
				"title": "Jaring Peristiwa Kafka",
				"benchmark": "180k pesan/detik",
				"description": "Tulang punggung berbasis event-driven terdistribusi menggunakan event sourcing, log partisi streaming, dan jaminan integritas transaksi mutlak.",
				"stack": ["Apache Kafka", "Event Sourcing", "CQRS", "Redis"],
				"image": "/src/assets/labs_generative.jpg"
			},
			{
				"id": 6,
				"title": "Keamanan Kernel & Zero-Trust",
				"benchmark": "Overhead 0%",
				"description": "Inspeksi paket mendalam dan penegakan kebijakan langsung pada kernel Linux dengan hook eBPF dan kriptografi mesh WireGuard.",
				"stack": ["Linux eBPF", "WireGuard", "TLS 1.3", "Zero-Trust"],
				"image": "/src/assets/dev_workspace.jpg"
			},
			{
				"id": 7,
				"title": "Internal Basis Data",
				"benchmark": "ACID Multi-Zona",
				"description": "Optimasi relasional PostgreSQL dengan pooling PgBouncer, replikasi streaming WAL, dan klaster cache in-memory Redis sub-milidetik.",
				"stack": ["PostgreSQL 16", "Redis Cluster", "SQLite Pure-Go"],
				"image": "/src/assets/fluid_droplets.jpg"
			},
			{
				"id": 8,
				"title": "Pustaka Sumber Terbuka",
				"benchmark": "Lisensi MIT",
				"description": "Alat bantu pengembang, utilitas CLI, dan kerangka kerja uji performa ringan yang dirancang untuk komunitas rekayasa peranti lunak global.",
				"stack": ["Go CLI", "GitHub Actions", "Engine Dokumentasi"],
				"image": "/src/assets/creative_poster.jpg"
			}
		]`

		manifestoJSON := `{
			"text_col1": "Perangkat lunak hebat tidak memiliki tanggal kedaluwarsa. Kode bekerja dalam hitungan mikrodetik dan berskala tanpa kompromi. Kebebasan kreatif mandiri memungkinkan fokus pada proyek yang lebih sedikit namun mendalam—menyelami aspek intelektual, fungsional, dan artistik dari rekayasa kode.",
			"text_col2": "Menciptakan arsitektur yang berbobot, keandalan sistem tanpa akhir, serta interaksi web mikro yang mengangkat pengalaman manusia dalam berinteraksi dengan komputer secara intuitif.",
			"button_text": "Tentang Vadikara",
			"button_url": "#contact"
		}`

		labsJSON := `{
			"tag": "EKSPERIMEN // LABS 0.9",
			"title": "Vadikara Labs — arena bermain kode kreatif & sistem algoritmik.",
			"background_image": "/src/assets/labs_generative.jpg",
			"experiments": [
				{
					"id": 1,
					"name": "Visualisator Ring-Buffer Lock-Free",
					"status": "RUNTIME AKTIF",
					"desc": "Visualisasi interaktif perebutan thread dan keselarasan memori antrean cincin zero-alloc di bawah simulasi beban kerja sintetik.",
					"tags": ["Go Assembly", "Tata Letak Memori", "Bebas Kunci"]
				},
				{
					"id": 2,
					"name": "Mesin Keadaan Konsensus Raft",
					"status": "SIMULASI KLASTER",
					"desc": "Simulasi pemilihan pemimpin klaster multi-node Raft dan replikasi log detak jantung dengan simulasi gangguan partisi jaringan acak.",
					"tags": ["Konsensus Terdistribusi", "Protokol Raft", "Kekacauan Jaringan"]
				},
				{
					"id": 3,
					"name": "Kanvas Fourier Responsif Audio",
					"status": "WEBGL // SHADER",
					"desc": "Penganalisis spektrum domain frekuensi waktu nyata yang menggambar kelopak vektor generatif dan pita partikel cairan terakselerasi GPU.",
					"tags": ["Web Audio API", "GLSL Shaders", "Canvas 60fps"]
				}
			]
		}`

		contactJSON := `{
			"tag": "KOLABORASI // JALUR LANGSUNG",
			"heading": "Mari bangun sistem yang bermakna dan berdaya tahan tinggi.",
			"subheading": "Terbuka untuk konsultasi arsitektur sistem prinsipal, rekayasa backend Golang berkonkurensi tinggi, serta perancangan antarmuka web kreatif 60 FPS.",
			"direct_email": "vadikara.dev@gmail.com",
			"github_url": "https://github.com/vadikara",
			"twitter_url": "https://x.com/vadikara",
			"linkedin_url": "https://linkedin.com",
			"topics": [
				"Sistem Terdistribusi",
				"Backend Golang 60k+ RPS",
				"Web Kreatif 60 FPS",
				"Arsitektur Cloud K8s"
			]
		}`

		navbarJSON := `{
			"brand_name": "Vadikara",
			"brand_sup": "®",
			"action_label": "GitHub",
			"action_url": "https://github.com/vadikara",
			"links": [
				{ "label": "Beranda", "url": "#hero" },
				{ "label": "Karya", "url": "#work" },
				{ "label": "Keahlian", "url": "#craft" },
				{ "label": "Tentang", "url": "#manifesto" },
				{ "label": "Labs", "url": "#labs" },
				{ "label": "Kontak", "url": "#contact" }
			]
		}`

		footerJSON := `{
			"logo_text": "Vadikara",
			"logo_sup": "®",
			"tagline": "Rekayasawan sistem kreatif & pengembang arsitektur terdistribusi. Merancang backend berkinerja tinggi serta antarmuka web 60 FPS yang berjiwa.",
			"copyright_text": "© 2026 Vadikara. Seluruh hak cipta dilindungi.",
			"built_text": "Dirancang & Dibangun dengan Go 1.26 & Vue 3 • 60 FPS",
			"index_links": [
				{ "label": "Beranda", "url": "#hero" },
				{ "label": "Karya Terpilih", "url": "#work" },
				{ "label": "Keahlian & Disiplin", "url": "#craft" },
				{ "label": "Manifesto", "url": "#manifesto" },
				{ "label": "Vadikara Labs", "url": "#labs" },
				{ "label": "Kontak", "url": "#contact" }
			],
			"social_links": [
				{ "label": "GitHub ↗", "url": "https://github.com/vadikara" },
				{ "label": "X // Twitter ↗", "url": "https://x.com/vadikara" },
				{ "label": "LinkedIn ↗", "url": "https://linkedin.com" }
			]
		}`

		ambientJSON := `{
			"track_title": "01 // RESONANSI CAIRAN",
			"track_subtitle": "Sintesis Harmonik 432 Hz",
			"audio_url": "",
			"enabled": true
		}`

		sections := map[string]string{
			"hero":        heroJSON,
			"ticker":      tickerJSON,
			"disciplines": disciplinesJSON,
			"manifesto":   manifestoJSON,
			"labs":        labsJSON,
			"contact":     contactJSON,
			"navbar":      navbarJSON,
			"footer":      footerJSON,
			"ambient":     ambientJSON,
		}

		for secKey, secVal := range sections {
			_, _ = db.Exec(`INSERT OR IGNORE INTO site_content (section_key, content_json) VALUES (?, ?)`, secKey, secVal)
		}
	}

	return nil
}
