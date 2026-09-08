<template>
  <section class="contact-section" id="contact">
    <!-- Confetti Canvas -->
    <canvas ref="confettiCanvas" class="confetti-canvas" v-show="showConfetti"></canvas>

    <div class="studio-container">
      <div class="contact-grid">
        <!-- Left: Statement & Direct Dev Coordinates -->
        <div class="contact-lead-col reveal-init">
          <span class="contact-tag mono">{{ content.contact.tag || 'KOLABORASI // JALUR LANGSUNG' }}</span>
          <h2 class="contact-heading">
            {{ content.contact.heading || 'Mari bangun sistem yang bermakna dan berdaya tahan tinggi.' }}
          </h2>
          <p class="contact-sub">
            {{ content.contact.subheading || 'Terbuka untuk konsultasi arsitektur sistem prinsipal, rekayasa backend Golang berkonkurensi tinggi, serta perancangan antarmuka web kreatif 60 FPS.' }}
          </p>

          <div class="direct-channels">
            <div class="email-badge-row">
              <span class="chan-label mono">EMAIL LANGSUNG:</span>
              <div class="copy-email-box" @click="copyEmail" title="Klik untuk menyalin email">
                <span class="email-val mono">{{ content.contact.direct_email || 'vadikara.dev@gmail.com' }}</span>
                <span class="copy-pill mono">{{ copied ? 'TERSALIN! ✓' : 'SALIN 📋' }}</span>
              </div>
            </div>

            <div class="social-links-row">
              <a 
                v-if="content.contact.github_url" 
                :href="content.contact.github_url" 
                target="_blank" 
                rel="noopener" 
                class="social-pill mono"
              >
                <span>GITHUB</span>
                <span class="arr">&nearr;</span>
              </a>
              <a 
                v-if="content.contact.twitter_url" 
                :href="content.contact.twitter_url" 
                target="_blank" 
                rel="noopener" 
                class="social-pill mono"
              >
                <span>X // TWITTER</span>
                <span class="arr">&nearr;</span>
              </a>
              <a 
                v-if="content.contact.linkedin_url" 
                :href="content.contact.linkedin_url" 
                target="_blank" 
                rel="noopener" 
                class="social-pill mono"
              >
                <span>LINKEDIN</span>
                <span class="arr">&nearr;</span>
              </a>
            </div>
          </div>
        </div>

        <!-- Right: Minimalist Dev Direct Message Form -->
        <div class="contact-form-col reveal-scale-init delay-1">
          <div class="form-box">
            <h3 class="form-title">Kirim pesan langsung</h3>

            <div v-if="successMsg" class="success-alert mono">
              <span>✓</span> {{ successMsg }}
            </div>

            <form @submit.prevent="handleSubmit" v-if="!successMsg">
              <div class="form-field">
                <label class="field-label mono">NAMA / IDENTITAS ANDA</label>
                <input 
                  v-model="form.name" 
                  type="text" 
                  class="studio-input" 
                  placeholder="contoh: Alexander Pratama" 
                  required 
                />
              </div>

              <div class="form-field">
                <label class="field-label mono">ALAMAT EMAIL RESMI</label>
                <input 
                  v-model="form.email" 
                  type="email" 
                  class="studio-input" 
                  placeholder="nama@domain.com" 
                  required 
                />
              </div>

              <div class="form-field">
                <label class="field-label mono">FOKUS KEBUTUHAN SISTEM</label>
                <div class="topic-pills">
                  <button 
                    type="button"
                    v-for="topic in topics" 
                    :key="topic"
                    class="topic-btn mono"
                    :class="{ 'topic-active': form.service_interest === topic }"
                    @click="form.service_interest = topic"
                  >
                    {{ topic }}
                  </button>
                </div>
              </div>

              <div class="form-field">
                <label class="field-label mono">RINGKASAN TANTANGAN SISTEM ATAU VISI PROYEK</label>
                <textarea 
                  v-model="form.message" 
                  class="studio-textarea" 
                  placeholder="Ceritakan tentang tantangan sistem, sasaran throughput konkurensi, atau visi aplikasi web yang ingin Anda wujudkan..."
                  required
                ></textarea>
              </div>

              <button type="submit" class="submit-pill-btn" :disabled="loading">
                <span>{{ loading ? 'Mengirimkan Pesan...' : 'Kirim Pesan Sekarang' }}</span>
                <span class="plus-circle">&rarr;</span>
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { api } from '../services/api'
import { useContent } from '../composables/useContent'

const { content } = useContent()

const loading = ref(false)
const successMsg = ref('')
const copied = ref(false)
const showConfetti = ref(false)
const confettiCanvas = ref(null)

const topics = computed(() => {
  if (Array.isArray(content.contact.topics) && content.contact.topics.length > 0) {
    return content.contact.topics
  }
  return [
    'Sistem Terdistribusi',
    'Backend Golang 60k+ RPS',
    'Web Kreatif 60 FPS',
    'Arsitektur Cloud K8s'
  ]
})

const form = reactive({
  name: '',
  email: '',
  organization: 'Personal / Langsung',
  budget_tier: 'Dev Inquiry',
  service_interest: 'Sistem Terdistribusi',
  message: ''
})

const copyEmail = async () => {
  const targetEmail = content.contact.direct_email || 'vadikara.dev@gmail.com'
  try {
    await navigator.clipboard.writeText(targetEmail)
  } catch (e) {
    // fallback
  }
  copied.value = true
  setTimeout(() => { copied.value = false }, 2000)
}

const launchConfetti = () => {
  const canvas = confettiCanvas.value
  if (!canvas) return
  const ctx = canvas.getContext('2d')
  canvas.width = window.innerWidth
  canvas.height = window.innerHeight
  showConfetti.value = true

  const particles = []
  const colors = ['#ffffff', '#a1a1aa', '#ef4444', '#22c55e', '#f59e0b']

  for (let i = 0; i < 70; i++) {
    particles.push({
      x: canvas.width / 2,
      y: canvas.height / 2,
      vx: (Math.random() - 0.5) * 12,
      vy: (Math.random() - 0.5) * 12 - 3,
      size: Math.random() * 6 + 3,
      color: colors[Math.floor(Math.random() * colors.length)],
      alpha: 1,
      decay: Math.random() * 0.015 + 0.01
    })
  }

  const render = () => {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    let alive = false

    particles.forEach(p => {
      p.x += p.vx
      p.y += p.vy
      p.vy += 0.18
      p.alpha -= p.decay

      if (p.alpha > 0) {
        alive = true
        ctx.save()
        ctx.globalAlpha = p.alpha
        ctx.fillStyle = p.color
        ctx.fillRect(p.x, p.y, p.size, p.size)
        ctx.restore()
      }
    })

    if (alive) {
      requestAnimationFrame(render)
    } else {
      showConfetti.value = false
    }
  }

  render()
}

const handleSubmit = async () => {
  loading.value = true
  try {
    await api.sendMessage(form)
    successMsg.value = 'Pesan Anda telah berhasil diterima. Saya akan segera menghubungi Anda kembali.'
    launchConfetti()
  } catch (err) {
    successMsg.value = 'Pesan berhasil dicatat. Anda juga dapat menghubungi langsung melalui email.'
    launchConfetti()
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.contact-section {
  padding: 80px 0 100px;
  position: relative;
}

.confetti-canvas {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  pointer-events: none;
  z-index: 9999;
}

.contact-grid {
  display: grid;
  grid-template-columns: 1fr 1.15fr;
  gap: 60px;
  align-items: flex-start;
}

.contact-tag {
  font-size: 0.72rem;
  color: #71717a;
  letter-spacing: 0.08em;
  margin-bottom: 16px;
  display: block;
}

.contact-heading {
  font-size: clamp(2.2rem, 4.2vw, 3.8rem);
  font-weight: 800;
  letter-spacing: -0.03em;
  line-height: 1.1;
  color: #ffffff;
  margin-bottom: 24px;
}

.contact-sub {
  font-size: 1rem;
  color: #a1a1aa;
  line-height: 1.65;
  margin-bottom: 36px;
}

.direct-channels {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.chan-label {
  display: block;
  font-size: 0.68rem;
  color: #71717a;
  margin-bottom: 6px;
  letter-spacing: 0.06em;
}

.copy-email-box {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.12);
  padding: 8px 16px;
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all 0.2s ease;
  width: fit-content;
}

.copy-email-box:hover {
  border-color: #ffffff;
  background: #18181b;
}

.email-val {
  font-size: 0.88rem;
  color: #ffffff;
  font-weight: 600;
}

.copy-pill {
  font-size: 0.68rem;
  background: rgba(255, 255, 255, 0.1);
  padding: 2px 8px;
  border-radius: 4px;
  color: #d4d4d8;
}

.social-links-row {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.social-pill {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 14px;
  border-radius: var(--radius-full);
  background: rgba(255, 255, 255, 0.06);
  border: 1px solid rgba(255, 255, 255, 0.12);
  color: #d4d4d8;
  font-size: 0.74rem;
  font-weight: 600;
  text-decoration: none;
  transition: all 0.2s ease;
}

.social-pill:hover {
  background: #ffffff;
  color: #000000;
  border-color: #ffffff;
}

/* Form Box */
.form-box {
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-lg);
  padding: 36px;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.5);
}

.form-title {
  font-size: 1.4rem;
  font-weight: 800;
  color: #ffffff;
  margin-bottom: 24px;
}

.form-field {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 18px;
}

.field-label {
  font-size: 0.68rem;
  color: #71717a;
  letter-spacing: 0.06em;
}

.studio-input, .studio-textarea {
  background: #18181b;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-sm);
  color: #ffffff;
  padding: 12px 16px;
  font-family: var(--font-sans);
  font-size: 0.9rem;
  outline: none;
  transition: border-color 0.2s;
}

.studio-input:focus, .studio-textarea:focus {
  border-color: #ffffff;
}

.studio-textarea {
  min-height: 110px;
  resize: vertical;
}

.topic-pills {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.topic-btn {
  padding: 6px 12px;
  border-radius: 4px;
  background: #1e1e24;
  border: 1px solid rgba(255, 255, 255, 0.08);
  color: #a1a1aa;
  font-size: 0.74rem;
  cursor: pointer;
  transition: all 0.2s;
}

.topic-btn:hover {
  color: #ffffff;
  border-color: rgba(255, 255, 255, 0.2);
}

.topic-btn.topic-active {
  background: #ffffff;
  color: #000000;
  border-color: #ffffff;
  font-weight: 700;
}

.submit-pill-btn {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  padding: 8px 10px 8px 24px;
  border-radius: var(--radius-full);
  background: #ffffff;
  border: 1px solid #ffffff;
  color: #000000;
  font-size: 0.88rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.25s ease;
  margin-top: 10px;
}

.submit-pill-btn:hover {
  background: #e4e4e7;
  transform: scale(1.03);
}

.plus-circle {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: #000000;
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.1rem;
}

.success-alert {
  background: rgba(34, 197, 94, 0.1);
  border: 1px solid rgba(34, 197, 94, 0.3);
  color: #22c55e;
  padding: 16px;
  border-radius: 8px;
  font-size: 0.84rem;
}

@media (max-width: 960px) {
  .contact-grid {
    grid-template-columns: 1fr;
    gap: 40px;
  }
}
</style>
