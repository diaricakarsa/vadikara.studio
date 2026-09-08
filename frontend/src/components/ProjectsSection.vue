<template>
  <section class="projects-section" id="work">
    <div class="studio-container">
      <!-- Top Row: Section Title & Navigation Arrows (Scene.mp4 00:04 & 00:07) -->
      <div class="work-header-row reveal-init">
        <h2 class="work-title">Karya terpilih</h2>
        <div class="work-controls">
          <button class="view-all-pill mono" @click="scrollToManifesto">
            <span>Jelajahi Semua</span>
          </button>
          <div class="arrow-group">
            <button class="circle-btn" @click="scrollCarousel(-1)" aria-label="Proyek Sebelumnya">
              &larr;
            </button>
            <button class="circle-btn" @click="scrollCarousel(1)" aria-label="Proyek Selanjutnya">
              &rarr;
            </button>
          </div>
        </div>
      </div>

      <!-- Horizontal Project Carousel (Scene.mp4 00:04) -->
      <div class="carousel-viewport" ref="carouselRef">
        <div class="carousel-track">
          <div 
            v-for="(project, idx) in displayProjects" 
            :key="project.id"
            class="project-slide-card reveal-scale-init"
            :class="`delay-${Math.min(idx + 1, 4)}`"
            @click="openProjectModal(project)"
          >
            <div class="slide-image-wrapper">
              <img :src="project.image" :alt="project.title" class="slide-image" />
              <div class="slide-overlay-badge">
                <span class="slide-cat mono">{{ project.tag }}</span>
              </div>
            </div>

            <div class="slide-meta">
              <div class="slide-title-row">
                <h3 class="slide-title">{{ project.title }}</h3>
                <span class="slide-arrow">&nearr;</span>
              </div>
              <p class="slide-desc">{{ project.summary }}</p>
              <div class="slide-metric mono">
                <span>{{ project.metric }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Developer Statement & Manifesto (Scene.mp4 00:04 & 00:06) -->
      <div class="manifesto-block" id="manifesto">
        <div class="manifesto-grid reveal-init">
          <div class="manifesto-col">
            <p class="manifesto-text">
              {{ content.manifesto.text_col1 || 'Perangkat lunak hebat tidak memiliki tanggal kedaluwarsa. Kode bekerja dalam hitungan mikrodetik dan berskala tanpa kompromi. Kebebasan kreatif mandiri memungkinkan fokus pada proyek yang lebih sedikit namun mendalam—menyelami aspek intelektual, fungsional, dan artistik dari rekayasa kode.' }}
            </p>
          </div>
          <div class="manifesto-col">
            <p class="manifesto-text text-muted">
              {{ content.manifesto.text_col2 || 'Menciptakan arsitektur yang berbobot, keandalan sistem tanpa akhir, serta interaksi web mikro yang mengangkat pengalaman manusia dalam berinteraksi dengan komputer secara intuitif.' }}
            </p>
          </div>
        </div>

        <div class="manifesto-action-row reveal-init delay-2">
          <a :href="content.manifesto.button_url || '#contact'" class="manifesto-pill-btn">
            <span>{{ content.manifesto.button_text || 'Tentang Vadikara' }}</span>
            <span class="plus-circle">+</span>
          </a>
        </div>
      </div>
    </div>

    <!-- Project Detail Modal -->
    <transition name="fade">
      <div v-if="selectedProject" class="project-modal-backdrop" @click.self="selectedProject = null">
        <div class="project-modal-card">
          <div class="modal-header">
            <div>
              <span class="badge-tag mono">{{ selectedProject.tag }}</span>
              <h3 class="modal-project-title">{{ selectedProject.title }}</h3>
            </div>
            <button class="modal-close" @click="selectedProject = null" aria-label="Tutup">&times;</button>
          </div>

          <div class="modal-body">
            <div class="modal-img-wrapper">
              <img :src="selectedProject.image" :alt="selectedProject.title" class="modal-img" />
            </div>

            <div class="modal-spec-box mono">
              <strong>BENCHMARK PERFORMA TERUJI:</strong>
              <p class="green-text">{{ selectedProject.metric }}</p>
            </div>

            <p class="modal-desc">{{ selectedProject.fullDesc }}</p>

            <div class="modal-stack-tags">
              <span v-for="(t, i) in selectedProject.tech" :key="i" class="stack-chip mono">{{ t }}</span>
            </div>
          </div>

          <div class="modal-footer">
            <a v-if="selectedProject.github" :href="selectedProject.github" target="_blank" rel="noopener" class="pill-btn">
              Lihat Repositori Sumber &nearr;
            </a>
            <a href="#contact" class="circle-btn-solid circle-btn" @click="selectedProject = null">
              &rarr;
            </a>
          </div>
        </div>
      </div>
    </transition>
  </section>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useContent } from '../composables/useContent'
import { api } from '../services/api'
import creativePosterImg from '../assets/creative_poster.jpg'
import devWorkspaceImg from '../assets/dev_workspace.jpg'
import fluidDropletsImg from '../assets/fluid_droplets.jpg'
import labsGenerativeImg from '../assets/labs_generative.jpg'

const { content, resolveImage } = useContent()

const carouselRef = ref(null)
const selectedProject = ref(null)
const backendProjects = ref([])

const defaultFeaturedProjects = [
  {
    id: 1,
    title: 'NovaMesh Inti Terdistribusi',
    tag: 'SISTEM // GOLANG',
    summary: 'Tulang punggung event-driven berkecepatan ultra-tinggi yang menghubungkan layanan transaksi dengan latensi sub-milidetik.',
    metric: '185,000 req/s • p99 < 1.1ms',
    image: creativePosterImg,
    fullDesc: 'Runtime Go terkustomisasi dengan alokasi buffer nol, konsensus Raft untuk pemilihan leader terdistribusi, serta antrean ring-buffer lock-free.',
    tech: ['Go 1.26', 'gRPC', 'Kafka', 'Redis Cluster', 'Prometheus'],
    github: 'https://github.com/vadikara'
  },
  {
    id: 2,
    title: 'AegisGuard Jaring Zero-Trust',
    tag: 'KEAMANAN // eBPF',
    summary: 'Gateway mikro-segmentasi tingkat kernel yang mendeteksi anomali jaringan dan lonjakan paket data secara seketika.',
    metric: 'Kernel Linux eBPF • < 0.2% CPU Overhead',
    image: devWorkspaceImg,
    fullDesc: 'Mencegah pergerakan ancaman lateral melalui inspeksi paket mendalam langsung pada perimeter kernel Linux tanpa mengorbankan performa jaringan host.',
    tech: ['Linux eBPF', 'Golang', 'WireGuard', 'TimescaleDB'],
    github: 'https://github.com/vadikara'
  },
  {
    id: 3,
    title: 'Kinetic Fluid Antarmuka Web',
    tag: 'WEB KREATIF // CANVAS',
    summary: 'Sistem interaksi cairan terakselerasi GPU 60 FPS yang dirancang untuk pengalaman web sensorik mewah.',
    metric: 'Fluiditas 60 FPS • Nol Frame Drop',
    image: fluidDropletsImg,
    fullDesc: 'Simulasi fisika cairan organik, refraksi cahaya kaustik, dan dinamika kursor mouse interaktif menggunakan komputasi shader modern.',
    tech: ['Vue 3', 'Canvas 2D/WebGL', 'Vite', 'Modern CSS'],
    github: 'https://github.com/vadikara'
  },
  {
    id: 4,
    title: 'Synapse Pialang Peristiwa',
    tag: 'EVENT SOURCING // CLOUD',
    summary: 'Broker peristiwa terpisah yang memproses log buku besar streaming dengan pengurutan mutlak dan ketahanan multi-zona.',
    metric: 'Nol Data Loss • Multi-Zona Raft',
    image: labsGenerativeImg,
    fullDesc: 'Platform distribusi log peristiwa berkecepatan tinggi yang dibangun untuk mendukung audit trail finansial dan mesin keadaan terdistribusi real-time.',
    tech: ['Kafka 3.7', 'Golang', 'Docker', 'K3s'],
    github: 'https://github.com/vadikara'
  }
]

const displayProjects = computed(() => {
  if (backendProjects.value.length > 0) {
    return backendProjects.value
  }
  return defaultFeaturedProjects
})

onMounted(async () => {
  try {
    const list = await api.getProjects()
    if (Array.isArray(list) && list.length > 0) {
      backendProjects.value = list.map(p => ({
        id: p.id,
        title: p.title,
        tag: p.category ? p.category.toUpperCase() : 'SISTEM',
        summary: p.description,
        metric: p.metrics || 'Latensi sub-milidetik',
        image: resolveImage(p.image_url),
        fullDesc: p.full_overview || p.architecture_summary || p.description,
        tech: p.tech_stack || [],
        github: p.github_url,
        live: p.live_url
      }))
    }
  } catch (err) {
    console.warn('[ProjectsSection] Fallback to default featured projects')
  }
})

const scrollCarousel = (direction) => {
  if (!carouselRef.value) return
  const scrollAmount = 460 * direction
  carouselRef.value.scrollBy({ left: scrollAmount, behavior: 'smooth' })
}

const openProjectModal = (project) => {
  selectedProject.value = project
}

const scrollToManifesto = () => {
  const el = document.getElementById('manifesto')
  if (el) el.scrollIntoView({ behavior: 'smooth' })
}
</script>

<style scoped>
.projects-section {
  padding: 80px 0 100px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  position: relative;
}

.work-header-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 40px;
}

.work-title {
  font-size: clamp(2.2rem, 4vw, 3.6rem);
  font-weight: 800;
  letter-spacing: -0.03em;
  color: #ffffff;
  margin: 0;
}

.work-controls {
  display: flex;
  align-items: center;
  gap: 14px;
}

.view-all-pill {
  padding: 8px 18px;
  border-radius: var(--radius-full);
  background: rgba(255, 255, 255, 0.08);
  border: 1px solid rgba(255, 255, 255, 0.15);
  color: #ffffff;
  font-size: 0.78rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.view-all-pill:hover {
  background: #ffffff;
  color: #000000;
}

.arrow-group {
  display: flex;
  gap: 8px;
}

/* Horizontal Carousel */
.carousel-viewport {
  width: 100%;
  overflow-x: auto;
  scroll-behavior: smooth;
  scrollbar-width: none;
  -ms-overflow-style: none;
  padding-bottom: 20px;
}

.carousel-viewport::-webkit-scrollbar {
  display: none;
}

.carousel-track {
  display: flex;
  gap: 28px;
  width: max-content;
}

.project-slide-card {
  width: 440px;
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-lg);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

.project-slide-card:hover {
  border-color: rgba(255, 255, 255, 0.4);
  transform: translateY(-6px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
}

.slide-image-wrapper {
  position: relative;
  width: 100%;
  height: 290px;
  background: #000000;
  overflow: hidden;
}

.slide-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}

.project-slide-card:hover .slide-image {
  transform: scale(1.06);
}

.slide-overlay-badge {
  position: absolute;
  top: 16px;
  left: 16px;
  background: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255, 255, 255, 0.12);
  padding: 4px 10px;
  border-radius: 4px;
}

.slide-cat {
  font-size: 0.68rem;
  color: #ffffff;
  font-weight: 700;
  letter-spacing: 0.06em;
}

.slide-meta {
  padding: 24px;
  display: flex;
  flex-direction: column;
  flex: 1;
}

.slide-title-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.slide-title {
  font-size: 1.3rem;
  font-weight: 800;
  color: #ffffff;
  letter-spacing: -0.02em;
}

.slide-arrow {
  font-size: 1.2rem;
  color: #71717a;
  transition: transform 0.2s ease, color 0.2s ease;
}

.project-slide-card:hover .slide-arrow {
  transform: translate(3px, -3px);
  color: #ffffff;
}

.slide-desc {
  font-size: 0.88rem;
  color: #a1a1aa;
  line-height: 1.55;
  margin-bottom: 16px;
}

.slide-metric {
  margin-top: auto;
  font-size: 0.76rem;
  color: #22c55e;
  background: rgba(34, 197, 94, 0.08);
  border: 1px solid rgba(34, 197, 94, 0.2);
  padding: 4px 10px;
  border-radius: 4px;
  width: fit-content;
}

/* Developer Manifesto (Scene.mp4 00:04 & 00:06) */
.manifesto-block {
  margin-top: 80px;
  padding-top: 60px;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
}

.manifesto-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 40px;
  margin-bottom: 36px;
}

.manifesto-text {
  font-size: clamp(1.2rem, 2vw, 1.65rem);
  font-weight: 700;
  line-height: 1.45;
  letter-spacing: -0.02em;
  color: #ffffff;
}

.manifesto-text.text-muted {
  color: #a1a1aa;
}

.manifesto-action-row {
  display: flex;
}

.manifesto-pill-btn {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  padding: 8px 10px 8px 22px;
  border-radius: var(--radius-full);
  background: rgba(255, 255, 255, 0.08);
  border: 1px solid rgba(255, 255, 255, 0.15);
  color: #ffffff;
  font-size: 0.88rem;
  font-weight: 700;
  text-decoration: none;
  transition: all 0.25s ease;
}

.manifesto-pill-btn:hover {
  background: #ffffff;
  color: #000000;
  border-color: #ffffff;
}

.plus-circle {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.15);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.1rem;
  transition: all 0.2s ease;
}

.manifesto-pill-btn:hover .plus-circle {
  background: #000000;
  color: #ffffff;
}

/* Modal */
.project-modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.85);
  backdrop-filter: blur(16px);
  z-index: 2500;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}

.project-modal-card {
  width: 100%;
  max-width: 680px;
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.15);
  border-radius: var(--radius-md);
  padding: 30px;
  box-shadow: 0 30px 90px rgba(0, 0, 0, 0.8);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 20px;
}

.badge-tag {
  font-size: 0.68rem;
  color: #71717a;
  letter-spacing: 0.08em;
}

.modal-project-title {
  font-size: 1.6rem;
  font-weight: 800;
  color: #ffffff;
  margin-top: 4px;
}

.modal-close {
  background: transparent;
  border: none;
  font-size: 2rem;
  color: #a1a1aa;
  cursor: pointer;
  line-height: 1;
}

.modal-img-wrapper {
  width: 100%;
  height: 240px;
  border-radius: var(--radius-sm);
  overflow: hidden;
  margin-bottom: 20px;
}

.modal-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.modal-spec-box {
  background: #18181b;
  border: 1px solid rgba(255, 255, 255, 0.08);
  padding: 12px 16px;
  border-radius: 6px;
  font-size: 0.82rem;
  margin-bottom: 16px;
}

.green-text {
  color: #22c55e;
  margin-top: 4px;
}

.modal-desc {
  font-size: 0.92rem;
  color: #a1a1aa;
  line-height: 1.6;
  margin-bottom: 20px;
}

.modal-stack-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 24px;
}

.stack-chip {
  font-size: 0.75rem;
  background: #27272a;
  color: #e4e4e7;
  padding: 4px 10px;
  border-radius: 4px;
}

.modal-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
}

@media (max-width: 768px) {
  .manifesto-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }
  .project-slide-card {
    width: 320px;
  }
  .slide-image-wrapper {
    height: 220px;
  }
}
</style>
