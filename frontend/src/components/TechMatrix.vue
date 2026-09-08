<template>
  <section class="tech-matrix-section" id="tech-matrix">
    <div class="container">
      <div class="section-header">
        <span class="badge badge-violet mono">Stack &amp; Standar Instrumen Rekayasa</span>
        <h2 class="section-title">
          Instrumen Produksi <span class="text-highlight">Teruji &amp; Deterministik.</span>
        </h2>
        <p class="section-subtitle">
          Kami meniadakan spekulasi teknis. Seluruh pustaka, runtime, dan infrastruktur dipilih secara ketat untuk menjamin throughput tinggi, keamanan memori, dan kemudahan pemeliharaan jangka panjang.
        </p>

        <!-- Interactive Search and Category Filter -->
        <div class="matrix-filter-controls">
          <div class="filter-category-row">
            <button 
              v-for="cat in filterCategories" 
              :key="cat"
              class="cat-filter-btn mono"
              :class="{ 'active': activeCategory === cat }"
              @click="activeCategory = cat"
            >
              {{ cat }}
            </button>
          </div>

          <div class="matrix-search-box">
            <span class="search-icon">🔎</span>
            <input 
              v-model="searchQuery" 
              type="text" 
              class="matrix-search-input mono" 
              placeholder="Cari pustaka, runtime, protokol (contoh: Go, Kafka, Redis, eBPF, K8s)..." 
            />
            <button v-if="searchQuery" class="clear-search-btn" @click="searchQuery = ''">&times;</button>
          </div>
        </div>
      </div>

      <!-- Filtered Matrix Groups -->
      <div class="matrix-grid">
        <div 
          class="matrix-card" 
          v-for="(group, idx) in filteredGroups" 
          :key="idx"
        >
          <div class="matrix-card-header">
            <span class="group-icon">{{ group.icon }}</span>
            <div>
              <h3 class="group-title">{{ group.title }}</h3>
              <span class="group-sub mono">{{ group.subtitle }}</span>
            </div>
          </div>

          <div class="tech-spec-table">
            <div 
              v-for="(skill, sIdx) in group.skills" 
              :key="sIdx"
              class="spec-row"
              @mouseenter="hoveredSkill = skill.name"
              @mouseleave="hoveredSkill = null"
              :class="{ 'row-highlight': hoveredSkill === skill.name }"
            >
              <div class="spec-left">
                <span class="spec-name">{{ skill.name }}</span>
                <span class="spec-version mono">{{ skill.version }}</span>
              </div>
              <div class="spec-right">
                <span class="spec-role mono">{{ skill.role }}</span>
                <span class="spec-ping mono" title="Production Grade Verified">✓</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-if="filteredGroups.length === 0" class="empty-results mono">
        <p>Tidak ditemukan instrumen dengan kriteria pencarian "{{ searchQuery }}".</p>
        <button class="btn btn-secondary btn-sm" @click="resetFilters">Reset Filter Pencarian</button>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed } from 'vue'

const activeCategory = ref('Semua Bidang')
const searchQuery = ref('')
const hoveredSkill = ref(null)

const filterCategories = ['Semua Bidang', 'Backend & Concurrency', 'Cloud & Zero-Trust', 'Enterprise Web & UI', 'Data Fabric & Cache']

const groups = [
  {
    category: 'Backend & Concurrency',
    title: "Backend & Concurrency Runtime",
    subtitle: "High-Throughput / Sub-millisecond Execution",
    icon: "⚡",
    skills: [
      { name: "Golang Core", version: "v1.26 Production", role: "Native Goroutines & Concurrency" },
      { name: "gRPC & Protocol Buffers", version: "Protobuf v3", role: "High-Performance RPC Transport" },
      { name: "Apache Kafka", version: "Event Mesh", role: "Distributed Log Partitioning" },
      { name: "REST & HTTP/2 Engine", version: "Chi / Gin Clean Arc", role: "Enterprise External Ingress" }
    ]
  },
  {
    category: 'Cloud & Zero-Trust',
    title: "Cloud Infrastructure & Zero-Trust Mesh",
    subtitle: "Multi-Zone Resilient Orchestration",
    icon: "☁️",
    skills: [
      { name: "Kubernetes (K8s) & K3s", version: "Cluster Ops", role: "Container Autoscaling & Rollouts" },
      { name: "Docker Runtime", version: "Multi-Stage Distroless", role: "Minimal Attack Surface Containers" },
      { name: "Linux eBPF & WireGuard", version: "Kernel Security", role: "Zero-Trust Encrypted Mesh" },
      { name: "Prometheus & Grafana", version: "Telemetry Engine", role: "p99 Latency & Health Observability" }
    ]
  },
  {
    category: 'Enterprise Web & UI',
    title: "Enterprise Web Applications & UI",
    subtitle: "Deterministic State & 60 FPS Fluidity",
    icon: "🌐",
    skills: [
      { name: "Vue 3 Composition API", version: "Production Lead", role: "Reactive Enterprise Client" },
      { name: "Modern CSS & HTML5", version: "Sub-second TTFB", role: "Responsive Architecture & WCAG" },
      { name: "Vite Build Toolchain", version: "Rollup Engine", role: "Optimized Tree-shaken Bundles" },
      { name: "Pinia State Store", version: "Deterministic Store", role: "Audit-Ready State Synchronization" }
    ]
  },
  {
    category: 'Data Fabric & Cache',
    title: "Data Persistence & In-Memory Fabric",
    subtitle: "ACID Compliance & Instant Cache",
    icon: "🗄️",
    skills: [
      { name: "PostgreSQL Primary", version: "PgBouncer Pool", role: "ACID Transactions & Replicas" },
      { name: "Redis In-Memory Cluster", version: "Redis v7", role: "Sub-ms Query Caching & Pub/Sub" },
      { name: "SQLite In-Memory", version: "Pure-Go Driver", role: "Ultra-Fast Embedded Querying" },
      { name: "Object Storage S3 Vault", version: "MinIO / AWS S3", role: "Encrypted Blob Document Storage" }
    ]
  }
]

const filteredGroups = computed(() => {
  const query = searchQuery.value.trim().toLowerCase()

  return groups
    .filter(group => {
      if (activeCategory.value === 'Semua Bidang') return true
      return group.category === activeCategory.value
    })
    .map(group => {
      if (!query) return group

      const matchedSkills = group.skills.filter(s => 
        s.name.toLowerCase().includes(query) ||
        s.role.toLowerCase().includes(query) ||
        s.version.toLowerCase().includes(query)
      )

      return {
        ...group,
        skills: matchedSkills
      }
    })
    .filter(group => group.skills.length > 0)
})

const resetFilters = () => {
  activeCategory.value = 'Semua Bidang'
  searchQuery.value = ''
}
</script>

<style scoped>
.tech-matrix-section {
  padding: 90px 0;
  background: #ffffff;
  border-top: 1px solid #e5e7eb;
}

.section-header {
  text-align: center;
  max-width: 820px;
  margin: 0 auto 50px;
}

.section-title {
  font-size: clamp(2rem, 3.4vw, 2.6rem);
  font-weight: 800;
  line-height: 1.25;
  letter-spacing: -0.02em;
  color: #111827;
  margin: 16px 0;
}

.text-highlight {
  color: #14532d;
}

.section-subtitle {
  font-size: 1.05rem;
  color: #4b5563;
  line-height: 1.7;
  margin-bottom: 30px;
}

/* Interactive Filter Controls */
.matrix-filter-controls {
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-items: center;
}

.filter-category-row {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 8px;
}

.cat-filter-btn {
  padding: 6px 14px;
  border-radius: var(--radius-xs);
  background: #f4f5f1;
  border: 1px solid #e5e7eb;
  color: #4b5563;
  font-size: 0.78rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.cat-filter-btn:hover {
  background: #e2e8df;
  color: #111827;
}

.cat-filter-btn.active {
  background: #14532d;
  border-color: #14532d;
  color: #ffffff;
  box-shadow: 0 2px 8px rgba(20, 83, 45, 0.2);
}

.matrix-search-box {
  position: relative;
  width: 100%;
  max-width: 520px;
  display: flex;
  align-items: center;
}

.search-icon {
  position: absolute;
  left: 14px;
  font-size: 0.9rem;
  color: #9ca3af;
  pointer-events: none;
}

.matrix-search-input {
  width: 100%;
  padding: 10px 36px 10px 38px;
  background: #fbfcf9;
  border: 1px solid #d1d5db;
  border-radius: var(--radius-full);
  font-size: 0.84rem;
  color: #111827;
  outline: none;
  transition: all 0.2s ease;
}

.matrix-search-input:focus {
  background: #ffffff;
  border-color: #166534;
  box-shadow: 0 0 0 3px rgba(22, 101, 52, 0.12);
}

.clear-search-btn {
  position: absolute;
  right: 12px;
  background: transparent;
  border: none;
  font-size: 1.2rem;
  color: #9ca3af;
  cursor: pointer;
}

.clear-search-btn:hover {
  color: #111827;
}

.matrix-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24px;
}

.matrix-card {
  padding: 30px;
  background: #fbfcf9;
  border: 1px solid #e5e7eb;
  border-radius: var(--radius-md);
  transition: all 0.2s ease;
}

.matrix-card:hover {
  border-color: #14532d;
  box-shadow: 0 8px 20px rgba(15, 36, 25, 0.06);
}

.matrix-card-header {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 20px;
  padding-bottom: 14px;
  border-bottom: 1px solid #e5e7eb;
}

.group-icon {
  font-size: 1.5rem;
}

.group-title {
  font-size: 1.15rem;
  font-weight: 700;
  color: #111827;
  margin-bottom: 2px;
}

.group-sub {
  font-size: 0.72rem;
  color: #6b7280;
}

.tech-spec-table {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.spec-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 12px;
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: var(--radius-sm);
  font-size: 0.85rem;
  transition: all 0.2s ease;
}

.spec-row.row-highlight {
  border-color: #166534;
  background: #f0fdf4;
  transform: translateX(3px);
}

.spec-left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.spec-name {
  font-weight: 700;
  color: #111827;
}

.spec-version {
  font-size: 0.7rem;
  color: #166534;
  background: #f0fdf4;
  padding: 2px 6px;
  border-radius: 4px;
  border: 1px solid #bbf7d0;
}

.spec-right {
  display: flex;
  align-items: center;
  gap: 10px;
}

.spec-role {
  font-size: 0.75rem;
  color: #4b5563;
}

.spec-ping {
  color: #15803d;
  font-weight: 800;
  font-size: 0.8rem;
}

.empty-results {
  text-align: center;
  padding: 40px 20px;
  color: #6b7280;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
}

@media (max-width: 900px) {
  .matrix-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 600px) {
  .spec-row {
    flex-direction: column;
    align-items: flex-start;
    gap: 6px;
  }
}
</style>
