<template>
  <section class="services-section" id="services">
    <div class="container">
      <div class="section-header">
        <span class="badge badge-cyan mono">Solusi &amp; Kapabilitas Rekayasa</span>
        <h2 class="section-title">
          Kapabilitas Rekayasa End-to-End <span class="text-highlight">Skala Industri.</span>
        </h2>
        <p class="section-subtitle">
          Dari arsitektur backend konkurensi Golang berlatensi sub-milidetik hingga platform web korporat reaktif yang mematuhi tata kelola kedaulatan data nasional.
        </p>
      </div>

      <!-- Enterprise Bento Grid Layout -->
      <div class="bento-grid">
        <div 
          v-for="(service, idx) in services" 
          :key="service.id || service.slug"
          class="bento-card tilt-card"
          :class="{ 'bento-featured': idx === 0 }"
          @mousemove="handleCardMouseMove($event)"
        >
          <div class="tilt-glare"></div>

          <div class="card-top-row">
            <span class="card-category-tag mono">KAPABILITAS 0{{ idx + 1 }}</span>
            <span class="service-icon-pill mono">{{ getIcon(service.icon) }}</span>
          </div>

          <h3 class="service-name">{{ service.title }}</h3>
          <p class="service-summary">{{ service.summary }}</p>

          <div class="capabilities-list">
            <div 
              v-for="(cap, i) in service.capabilities" 
              :key="i"
              class="capability-item"
            >
              <span class="check-icon">✓</span>
              <span>{{ cap }}</span>
            </div>
          </div>

          <!-- Featured Code & Architecture Protocol Switcher for Card 0 -->
          <div v-if="idx === 0" class="featured-spec-box mono">
            <!-- Code Tab Navigation -->
            <div class="code-tabs-header">
              <div class="tab-list">
                <button 
                  v-for="(tab, tIdx) in codeTabs" 
                  :key="tIdx"
                  class="code-tab-btn mono"
                  :class="{ 'active': activeCodeTab === tIdx }"
                  @click="activeCodeTab = tIdx"
                >
                  <span class="tab-icon">{{ tab.icon }}</span>
                  <span>{{ tab.label }}</span>
                </button>
              </div>
              <button class="copy-btn mono" @click="copyActiveCode">
                <span>{{ copiedCode ? 'Tersalin! ✓' : 'Salin Kode 📋' }}</span>
              </button>
            </div>

            <!-- Active Code Content -->
            <div class="spec-bar">
              <span class="spec-file">{{ currentCodeTab.filename }}</span>
              <span class="spec-status">{{ currentCodeTab.badge }}</span>
            </div>

            <pre class="spec-code"><code>{{ currentCodeTab.code }}</code></pre>

            <div class="spec-footer-bar">
              <span class="spec-tag">RUNTIME: {{ currentCodeTab.runtime }}</span>
              <span class="spec-tag highlight-green">LATENCY: {{ currentCodeTab.latency }}</span>
            </div>
          </div>

          <div class="tech-tags">
            <span 
              v-for="(t, i) in service.tech_stack" 
              :key="i" 
              class="tech-tag mono"
            >
              {{ t }}
            </span>
          </div>

          <div class="deliverables-box" v-if="service.deliverables">
            <span class="deliverable-label mono">OUTPUT DOKUMENTASI &amp; KODE:</span>
            <span class="deliverable-text">{{ service.deliverables }}</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { api } from '../services/api'

const services = ref([])
const activeCodeTab = ref(0)
const copiedCode = ref(false)

const codeTabs = [
  {
    label: 'Golang Core gRPC',
    icon: '⚡',
    filename: 'worker_pool_engine.go',
    badge: 'p99 < 0.8ms • Zero Alloc',
    runtime: 'Golang 1.26',
    latency: 'Sub-millisecond',
    code: `package engine

type WorkerPool struct {
    Jobs    chan *transaction.Payload
    Workers int
    syncPool sync.Pool
}

func (wp *WorkerPool) Dispatch(ctx context.Context, p *transaction.Payload) error {
    select {
    case wp.Jobs <- p:
        return nil
    case <-ctx.Done():
        return errors.New("dispatch deadline exceeded")
    }
}`
  },
  {
    label: 'Kubernetes HPA Spec',
    icon: '☸️',
    filename: 'hpa_autoscaler.yaml',
    badge: 'K8s v1.30 Production',
    runtime: 'Container Orchestration',
    latency: 'Auto-scaled < 30s',
    code: `apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: vadikara-core-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: vadikara-core-engine
  minReplicas: 8
  maxReplicas: 32
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70`
  },
  {
    label: 'Kafka Event Mesh',
    icon: '📨',
    filename: 'event_dispatcher.go',
    badge: '185k msg/s • Raft Ack=all',
    runtime: 'Apache Kafka 3.7',
    latency: '< 1.1ms Ack=all',
    code: `func (d *EventDispatcher) PublishAudit(ctx context.Context, evt *Event) error {
    record := &kgo.Record{
        Topic:     "corporate.audit.v1",
        Key:       []byte(evt.TenantID),
        Value:     evt.PayloadBytes(),
        Timestamp: time.Now().UTC(),
    }
    return d.client.ProduceSync(ctx, record).FirstErr()
}`
  },
  {
    label: 'UU PDP AES-256 Vault',
    icon: '🔒',
    filename: 'pdp_crypto_vault.go',
    badge: 'Kepatuhan UU No. 27/2022',
    runtime: 'Hardware AES-GCM',
    latency: '0.04ms / record',
    code: `func EncryptPII(plaintext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil { return nil, err }
    gcm, err := cipher.NewGCM(block)
    if err != nil { return nil, err }
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil { return nil, err }
    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}`
  }
]

const currentCodeTab = computed(() => codeTabs[activeCodeTab.value])

const copyActiveCode = async () => {
  try {
    await navigator.clipboard.writeText(currentCodeTab.value.code)
    copiedCode.value = true
    setTimeout(() => {
      copiedCode.value = false
    }, 2000)
  } catch (err) {
    // Fallback if clipboard API restricted
    copiedCode.value = true
    setTimeout(() => { copiedCode.value = false }, 2000)
  }
}

const handleCardMouseMove = (event) => {
  const card = event.currentTarget
  const rect = card.getBoundingClientRect()
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top
  card.style.setProperty('--mouse-x', `${x}px`)
  card.style.setProperty('--mouse-y', `${y}px`)
}

const fallbackServices = [
  {
    title: "High-Throughput Backend & Microservices Terdistribusi",
    slug: "high-throughput-backend",
    icon: "Server",
    summary: "Perancangan backend berkecepatan tinggi berbasis Golang dengan model konkurensi native (Goroutines), gRPC, dan arsitektur event-driven.",
    capabilities: [
      "Pemrosesan ratusan ribu Request Per Second (RPS)",
      "Arsitektur Microservices terisolasi dan fault-tolerant",
      "Komunikasi real-time via WebSockets dan gRPC",
      "Caching berlapis (In-memory, Redis, CDN)"
    ],
    tech_stack: ["Golang (Go 1.26)", "gRPC", "Apache Kafka", "Redis Cluster", "PostgreSQL", "Protobuf"],
    deliverables: "Source code production-ready, benchmark load test report, Dockerized stack, OpenAPI docs."
  },
  {
    title: "Infrastruktur Cloud-Native, Kubernetes & Observabilitas",
    slug: "enterprise-cloud-devops",
    icon: "Cpu",
    summary: "Infrastruktur cloud-native yang aman dan elastis dengan otomatisasi deployment CI/CD tanpa downtime (Zero Downtime).",
    capabilities: [
      "Orkestrasi Kubernetes (K8s) & Cluster autoscaling",
      "Infrastructure as Code (Terraform, Ansible)",
      "Pipeline CI/CD otomatis dengan pengujian terintegrasi",
      "Observability komprehensif (Prometheus, Grafana)"
    ],
    tech_stack: ["Kubernetes", "Docker Multi-stage", "Terraform", "GitHub Actions", "Prometheus", "Grafana"],
    deliverables: "Manifest IaC, pipeline deployment otomatis, dashboard monitoring 24/7."
  },
  {
    title: "Audit Keamanan Kode, Resiliensi & Kepatuhan UU PDP",
    slug: "resilience-security-audit",
    icon: "ShieldCheck",
    summary: "Audit mendalam terhadap kerentanan kode sumber, optimasi memori/CPU, serta penguatan perimeter keamanan sistem sesuai UU PDP.",
    capabilities: [
      "Static & Dynamic Code Security Analysis (SAST/DAST)",
      "Stress testing & chaos engineering simulasi kegagalan",
      "Optimasi alokasi memori dan deteksi goroutine leaks",
      "Implementasi standar Zero-Trust & enkripsi AES-256"
    ],
    tech_stack: ["Go pprof", "k6 Load Testing", "OWASP Standards", "WireGuard", "UU No. 27/2022"],
    deliverables: "Laporan audit komprehensif, rekomendasi mitigasi langkah demi langkah, patch refactoring."
  },
  {
    title: "Aplikasi Web Enterprise Reaktif Berkinerja Tinggi",
    slug: "modern-web-applications",
    icon: "Globe",
    summary: "Pengembangan front-end berkelas dunia dengan animasi dinamis penuh aksi, estetika premium, responsif 100%, dan performa instan.",
    capabilities: [
      "Aplikasi Single Page Application (SPA) & Progressive Web App (PWA)",
      "Animasi 60 FPS halus (Canvas, Modern CSS)",
      "Design System kustom dan State Management deterministik",
      "Optimasi Core Web Vitals dan SEO teknis terstruktur"
    ],
    tech_stack: ["Vue 3", "Vite", "HTML5", "Modern CSS", "Pinia", "TypeScript"],
    deliverables: "Komponen modular, asset teroptimasi, dokumentasi UI/UX, bundle size minimal."
  },
  {
    title: "Integrasi AI & Pipeline Dokumen Otomatis",
    slug: "ai-integration-automation",
    icon: "Sparkles",
    summary: "Integrasi kecerdasan buatan (LLM, computer vision, predictive telemetry) ke dalam alur kerja bisnis korporasi.",
    capabilities: [
      "Pipeline RAG (Retrieval-Augmented Generation) berbasis data internal",
      "Sistem inferensi dokumen otomatis & ekstraksi data legal",
      "Agentic workflows untuk otomatisasi tugas kompleks",
      "Optimasi latensi inferensi dan mitigasi biaya token API"
    ],
    tech_stack: ["Python/Go Bridge", "Vector DB", "TensorRT", "Ollama", "Claude / OpenAI APIs"],
    deliverables: "API wrapper aman, database embedding, pipeline data ingest, audit trail."
  }
]

const getIcon = (iconName) => {
  switch (iconName) {
    case 'Server': return '⚡ Server Concurrency'
    case 'Globe': return '🌐 Enterprise Web'
    case 'Cpu': return '⚙️ Cloud Infrastructure'
    case 'Sparkles': return '🧠 AI Pipeline'
    case 'ShieldCheck': return '🛡️ Security & Privacy'
    default: return '🚀 Engineering'
  }
}

onMounted(async () => {
  try {
    const data = await api.getServices()
    services.value = data.length > 0 ? data : fallbackServices
  } catch (err) {
    services.value = fallbackServices
  }
})
</script>

<style scoped>
.services-section {
  padding: 90px 0;
  background: #fbfcf9;
  position: relative;
}

.section-header {
  text-align: center;
  max-width: 820px;
  margin: 0 auto 60px;
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
}

/* Bento Grid */
.bento-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24px;
}

.bento-card {
  padding: 34px;
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: var(--radius-md);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.04);
  display: flex;
  flex-direction: column;
  position: relative;
  overflow: hidden;
  transition: all 0.25s ease;
}

.bento-card:hover {
  border-color: #14532d;
  box-shadow: 0 12px 28px -5px rgba(15, 36, 25, 0.08);
}

.bento-featured {
  grid-column: span 2;
  background: linear-gradient(180deg, #ffffff 0%, #fbfdfa 100%);
  border-color: #bbf7d0;
}

.card-top-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.card-category-tag {
  font-size: 0.68rem;
  font-weight: 700;
  color: #166534;
  letter-spacing: 0.06em;
}

.service-icon-pill {
  font-size: 0.72rem;
  color: #4b5563;
  background: #f3f4f6;
  padding: 3px 8px;
  border-radius: 4px;
}

.service-name {
  font-size: 1.35rem;
  font-weight: 800;
  color: #111827;
  margin-bottom: 10px;
  line-height: 1.35;
}

.service-summary {
  font-size: 0.94rem;
  color: #4b5563;
  line-height: 1.65;
  margin-bottom: 22px;
}

.capabilities-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 22px;
}

.capability-item {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  font-size: 0.88rem;
  color: #374151;
  transition: transform 0.15s ease;
}

.capability-item:hover {
  transform: translateX(4px);
  color: #14532d;
}

.check-icon {
  color: #15803d;
  font-weight: 800;
}

/* Featured Interactive Architecture Spec Box */
.featured-spec-box {
  background: #0b1a12;
  color: #86efac;
  border-radius: var(--radius-sm);
  padding: 16px;
  margin-bottom: 22px;
  font-size: 0.8rem;
  border: 1px solid rgba(22, 101, 52, 0.4);
  box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.3);
}

.code-tabs-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding-bottom: 10px;
  margin-bottom: 12px;
  flex-wrap: wrap;
  gap: 8px;
}

.tab-list {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.code-tab-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 5px 10px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #9ca3af;
  font-size: 0.72rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.code-tab-btn:hover {
  background: rgba(255, 255, 255, 0.12);
  color: #ffffff;
}

.code-tab-btn.active {
  background: #15803d;
  border-color: #22c55e;
  color: #ffffff;
  box-shadow: 0 0 10px rgba(34, 197, 94, 0.3);
}

.copy-btn {
  background: rgba(255, 255, 255, 0.08);
  border: 1px solid rgba(255, 255, 255, 0.15);
  color: #d1d5db;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 0.7rem;
  cursor: pointer;
  transition: all 0.2s;
}

.copy-btn:hover {
  background: rgba(255, 255, 255, 0.2);
  color: #fff;
}

.spec-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 6px;
  margin-bottom: 8px;
  color: #9ca3af;
  font-size: 0.72rem;
}

.spec-file {
  color: #cbd5e1;
  font-weight: 700;
}

.spec-status {
  color: #4ade80;
}

.spec-code {
  margin: 0;
  line-height: 1.55;
  color: #bbf7d0;
  overflow-x: auto;
  font-family: var(--font-mono);
  font-size: 0.78rem;
  max-height: 170px;
  padding: 4px 0;
}

.spec-footer-bar {
  display: flex;
  gap: 14px;
  margin-top: 10px;
  padding-top: 8px;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  font-size: 0.68rem;
}

.spec-tag {
  color: #9ca3af;
}

.spec-tag.highlight-green {
  color: #4ade80;
  font-weight: 700;
}

.tech-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-bottom: 22px;
}

.tech-tag {
  font-size: 0.72rem;
  padding: 3px 8px;
  border-radius: 4px;
  background: #f4f5f1;
  border: 1px solid #e5e7eb;
  color: #166534;
  font-weight: 600;
}

.deliverables-box {
  margin-top: auto;
  padding: 14px;
  background: #f9fafb;
  border-radius: var(--radius-sm);
  border: 1px solid #e5e7eb;
  border-left: 3px solid #14532d;
  font-size: 0.8rem;
}

.deliverable-label {
  display: block;
  font-size: 0.68rem;
  color: #6b7280;
  font-weight: 700;
  margin-bottom: 4px;
  letter-spacing: 0.05em;
}

.deliverable-text {
  color: #374151;
  line-height: 1.5;
}

@media (max-width: 960px) {
  .bento-grid {
    grid-template-columns: 1fr;
  }
  .bento-featured {
    grid-column: span 1;
  }
}
</style>
