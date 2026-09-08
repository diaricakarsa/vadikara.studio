import { ref, reactive, computed } from 'vue'
import { api } from '../services/api'

import fluidDropletsImg from '../assets/fluid_droplets.jpg'
import devWorkspaceImg from '../assets/dev_workspace.jpg'
import creativePosterImg from '../assets/creative_poster.jpg'
import labsGenerativeImg from '../assets/labs_generative.jpg'

// Map of preset assets for convenient selection and resolution
export const PRESET_IMAGES = [
  { label: 'Tetesan Cairan Organik (Hero / Droplets)', value: fluidDropletsImg, path: '/src/assets/fluid_droplets.jpg' },
  { label: 'Studio Kerja Dev (Workspace & Terminal)', value: devWorkspaceImg, path: '/src/assets/dev_workspace.jpg' },
  { label: 'Poster Kreatif Geometris (Creative Poster)', value: creativePosterImg, path: '/src/assets/creative_poster.jpg' },
  { label: 'Algoritma Generatif (Labs Generative)', value: labsGenerativeImg, path: '/src/assets/labs_generative.jpg' },
]

export function resolveImage(imgRef) {
  if (!imgRef) return fluidDropletsImg
  if (typeof imgRef === 'string') {
    if (imgRef.startsWith('data:image') || imgRef.startsWith('http://') || imgRef.startsWith('https://')) {
      return imgRef
    }
    if (imgRef.includes('fluid_droplets')) return fluidDropletsImg
    if (imgRef.includes('dev_workspace')) return devWorkspaceImg
    if (imgRef.includes('creative_poster')) return creativePosterImg
    if (imgRef.includes('labs_generative')) return labsGenerativeImg
  }
  return imgRef
}

// Global Singleton Reactive State with Natural Indonesian Defaults
const isLoaded = ref(false)
const isLoading = ref(false)
const lastError = ref(null)

const content = reactive({
  navbar: {
    brand_name: 'Vadikara',
    brand_sup: '®',
    action_label: 'GitHub',
    action_url: 'https://github.com/vadikara',
    links: [
      { label: 'Beranda', url: '#hero' },
      { label: 'Karya', url: '#work' },
      { label: 'Keahlian', url: '#craft' },
      { label: 'Tentang', url: '#manifesto' },
      { label: 'Labs', url: '#labs' },
      { label: 'Kontak', url: '#contact' }
    ]
  },
  hero: {
    title: 'Rekayasawan sistem kreatif &\nperanti lunak terdistribusi',
    sup: '®',
    reel_button_text: 'JELAJAHI REEL',
    reel_badge: 'VADIKARA // SHOWCASE REKAYASA SISTEM 2026',
    reel_url: '/Scene.mp4',
    reel_telemetry: 'Konkurensi Tinggi Golang • Web Kreatif 60 FPS • Sistem Terdistribusi Sub-milidetik',
    background_image: '/src/assets/fluid_droplets.jpg'
  },
  ticker: [
    'vadikara',
    'vadikara.dev',
    'rekayasa kreatif',
    'sistem terdistribusi',
    'antarmuka 60 fps',
    'inti golang',
    'arsitektur awan'
  ],
  disciplines: [
    {
      id: 1,
      title: 'Sistem Terdistribusi',
      benchmark: 'p99 < 1.2ms',
      description: 'Arsitektur klaster berdaya tahan tinggi dan toleran kesalahan dengan konsensus Raft, mitigasi split-brain, serta ketiadaan single point of failure.',
      stack: ['Golang 1.26', 'gRPC', 'Raft', 'Protobuf v3'],
      image: '/src/assets/dev_workspace.jpg'
    },
    {
      id: 2,
      title: 'Konkurensi Tinggi Go',
      benchmark: '350k+ RPS',
      description: 'Saluran eksekusi konkurensi berbasis Goroutine native, ring-buffer bebas kunci (lock-free), dan alokasi memori mendekati nol untuk beban ratusan ribu RPS.',
      stack: ['Go Runtime', 'Sync.Pool', 'Channels', 'pprof'],
      image: '/src/assets/dev_workspace.jpg'
    },
    {
      id: 3,
      title: 'Infrastruktur Cloud',
      benchmark: 'Zero-Downtime',
      description: 'Orkestrasi klaster Kubernetes deklaratif, kontainer distroless minimal risiko serangan, serta autoscaling horizontal otomatis.',
      stack: ['Kubernetes', 'Docker', 'Terraform', 'K3s'],
      image: '/src/assets/dev_workspace.jpg'
    },
    {
      id: 4,
      title: 'Web Kreatif & 60 FPS',
      benchmark: 'Sub-detik TTFB',
      description: 'Pengalaman web sensorik dengan animasi fisika cairan, tipografi kinetik, dan akselerasi GPU canvas yang deterministik dan konsisten 60 FPS.',
      stack: ['Vue 3', 'Modern CSS', 'Vite', 'Canvas / WebGL'],
      image: '/src/assets/creative_poster.jpg'
    },
    {
      id: 5,
      title: 'Jaring Peristiwa Kafka',
      benchmark: '180k pesan/detik',
      description: 'Tulang punggung berbasis event-driven terdistribusi menggunakan event sourcing, log partisi streaming, dan jaminan integritas transaksi mutlak.',
      stack: ['Apache Kafka', 'Event Sourcing', 'CQRS', 'Redis'],
      image: '/src/assets/labs_generative.jpg'
    },
    {
      id: 6,
      title: 'Keamanan Kernel & Zero-Trust',
      benchmark: 'Overhead 0%',
      description: 'Inspeksi paket mendalam dan penegakan kebijakan langsung pada kernel Linux dengan hook eBPF dan kriptografi mesh WireGuard.',
      stack: ['Linux eBPF', 'WireGuard', 'TLS 1.3', 'Zero-Trust'],
      image: '/src/assets/dev_workspace.jpg'
    },
    {
      id: 7,
      title: 'Internal Basis Data',
      benchmark: 'ACID Multi-Zona',
      description: 'Optimasi relasional PostgreSQL dengan pooling PgBouncer, replikasi streaming WAL, dan klaster cache in-memory Redis sub-milidetik.',
      stack: ['PostgreSQL 16', 'Redis Cluster', 'SQLite Pure-Go'],
      image: '/src/assets/fluid_droplets.jpg'
    },
    {
      id: 8,
      title: 'Pustaka Sumber Terbuka',
      benchmark: 'Lisensi MIT',
      description: 'Alat bantu pengembang, utilitas CLI, dan kerangka kerja uji performa ringan yang dirancang untuk komunitas rekayasa peranti lunak global.',
      stack: ['Go CLI', 'GitHub Actions', 'Engine Dokumentasi'],
      image: '/src/assets/creative_poster.jpg'
    }
  ],
  manifesto: {
    text_col1: 'Perangkat lunak hebat tidak memiliki tanggal kedaluwarsa. Kode bekerja dalam hitungan mikrodetik dan berskala tanpa kompromi. Kebebasan kreatif mandiri memungkinkan fokus pada proyek yang lebih sedikit namun mendalam—menyelami aspek intelektual, fungsional, dan artistik dari rekayasa kode.',
    text_col2: 'Menciptakan arsitektur yang berbobot, keandalan sistem tanpa akhir, serta interaksi web mikro yang mengangkat pengalaman manusia dalam berinteraksi dengan komputer secara intuitif.',
    button_text: 'Tentang Vadikara',
    button_url: '#contact'
  },
  labs: {
    tag: 'EKSPERIMEN // LABS 0.9',
    title: 'Vadikara Labs — arena bermain kode kreatif & sistem algoritmik.',
    background_image: '/src/assets/labs_generative.jpg',
    drawer_open_text: 'Jelajahi Labs',
    drawer_close_text: 'Tutup Eksperimen',
    experiments: [
      {
        id: 1,
        name: 'Visualisator Ring-Buffer Lock-Free',
        status: 'RUNTIME AKTIF',
        desc: 'Visualisasi interaktif perebutan thread dan keselarasan memori antrean cincin zero-alloc di bawah simulasi beban kerja sintetik.',
        tags: ['Go Assembly', 'Tata Letak Memori', 'Bebas Kunci']
      },
      {
        id: 2,
        name: 'Mesin Keadaan Konsensus Raft',
        status: 'SIMULASI KLASTER',
        desc: 'Simulasi pemilihan pemimpin klaster multi-node Raft dan replikasi log detak jantung dengan simulasi gangguan partisi jaringan acak.',
        tags: ['Konsensus Terdistribusi', 'Protokol Raft', 'Kekacauan Jaringan']
      },
      {
        id: 3,
        name: 'Kanvas Fourier Responsif Audio',
        status: 'WEBGL // SHADER',
        desc: 'Penganalisis spektrum domain frekuensi waktu nyata yang menggambar kelopak vektor generatif dan pita partikel cairan terakselerasi GPU.',
        tags: ['Web Audio API', 'GLSL Shaders', 'Canvas 60fps']
      }
    ]
  },
  contact: {
    tag: 'KOLABORASI // JALUR LANGSUNG',
    heading: 'Mari bangun sistem yang bermakna dan berdaya tahan tinggi.',
    subheading: 'Terbuka untuk konsultasi arsitektur sistem prinsipal, rekayasa backend Golang berkonkurensi tinggi, serta perancangan antarmuka web kreatif 60 FPS.',
    direct_email: 'vadikara.dev@gmail.com',
    github_url: 'https://github.com/vadikara',
    twitter_url: 'https://x.com/vadikara',
    linkedin_url: 'https://linkedin.com',
    topics: [
      'Sistem Terdistribusi',
      'Backend Golang 60k+ RPS',
      'Web Kreatif 60 FPS',
      'Arsitektur Cloud K8s'
    ]
  },
  footer: {
    logo_text: 'Vadikara',
    logo_sup: '®',
    tagline: 'Rekayasawan sistem kreatif & pengembang arsitektur terdistribusi. Merancang backend berkinerja tinggi serta antarmuka web 60 FPS yang berjiwa.',
    copyright_text: '© 2026 Vadikara. Seluruh hak cipta dilindungi.',
    built_text: 'Dirancang & Dibangun dengan Go 1.26 & Vue 3 • 60 FPS',
    index_links: [
      { label: 'Beranda', url: '#hero' },
      { label: 'Karya Terpilih', url: '#work' },
      { label: 'Keahlian & Disiplin', url: '#craft' },
      { label: 'Manifesto', url: '#manifesto' },
      { label: 'Vadikara Labs', url: '#labs' },
      { label: 'Kontak', url: '#contact' }
    ],
    social_links: [
      { label: 'GitHub ↗', url: 'https://github.com/vadikara' },
      { label: 'X // Twitter ↗', url: 'https://x.com/vadikara' },
      { label: 'LinkedIn ↗', url: 'https://linkedin.com' }
    ]
  },
  ambient: {
    track_title: '01 // RESONANSI CAIRAN',
    track_subtitle: 'Sintesis Harmonik 432 Hz',
    audio_url: '',
    enabled: true
  }
})

// Load content from API
export async function fetchContent() {
  isLoading.value = true
  try {
    const data = await api.getSiteContent()
    if (data) {
      if (data.navbar) Object.assign(content.navbar, data.navbar)
      if (data.hero) Object.assign(content.hero, data.hero)
      if (Array.isArray(data.ticker) && data.ticker.length > 0) content.ticker = data.ticker
      if (Array.isArray(data.disciplines) && data.disciplines.length > 0) content.disciplines = data.disciplines
      if (data.manifesto) Object.assign(content.manifesto, data.manifesto)
      if (data.labs) Object.assign(content.labs, data.labs)
      if (data.contact) Object.assign(content.contact, data.contact)
      if (data.footer) Object.assign(content.footer, data.footer)
      if (data.ambient) Object.assign(content.ambient, data.ambient)
    }
    isLoaded.value = true
    lastError.value = null
  } catch (err) {
    console.warn('[useContent] Using offline defaults, error:', err.message)
    lastError.value = err.message
  } finally {
    isLoading.value = false
  }
}

// Auto fetch once on initial import
fetchContent()

export function useContent() {
  return {
    content,
    isLoaded,
    isLoading,
    lastError,
    fetchContent,
    resolveImage,
    PRESET_IMAGES
  }
}
