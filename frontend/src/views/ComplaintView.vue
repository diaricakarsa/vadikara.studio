<template>
  <div class="complaint-view">
    <div class="container page-container">
      <div class="page-header">
        <span class="badge badge-amber mono">Layanan Pengaduan Resmi & Whistleblowing</span>
        <h1 class="page-title">Pengajuan Keluhan & <span class="text-gradient-gold">Bantuan SLA.</span></h1>
        <p class="page-subtitle">
          PT. Vadikara Rekayasa Peranti berkomitmen menjaga standar transparansi, integritas layanan, dan pemenuhan Service Level Agreement (SLA). Setiap pengaduan akan mendapatkan nomor tiket resmi dan ditindaklanjuti oleh komite manajemen teknis.
        </p>
      </div>

      <div class="tabs-container">
        <button 
          class="tab-btn" 
          :class="{ 'active': activeTab === 'submit' }"
          @click="activeTab = 'submit'"
        >
          📝 Ajukan Keluhan Baru
        </button>
        <button 
          class="tab-btn" 
          :class="{ 'active': activeTab === 'track' }"
          @click="activeTab = 'track'"
        >
          🔍 Lacak Status Tiket
        </button>
      </div>

      <!-- TAB 1: FORM PENGAJUAN KELUHAN -->
      <div v-if="activeTab === 'submit'" class="form-container glass-panel">
        <div v-if="submittedTicket" class="success-ticket-box">
          <div class="ticket-badge mono">TIKET RESMI DITERBITKAN</div>
          <h2 class="ticket-number mono">{{ submittedTicket.ticket_number }}</h2>
          <p class="ticket-desc">
            Pengaduan Anda telah tercatat dalam sistem helpdesk resmi PT. Vadikara Rekayasa Peranti. 
            Simpan nomor tiket di atas untuk memantau proses investigasi teknis dan resolusi penanganan.
          </p>
          <div class="ticket-summary mono">
            <div><strong>Pelapor:</strong> {{ submittedTicket.name }} ({{ submittedTicket.email }})</div>
            <div><strong>Kategori:</strong> {{ submittedTicket.category }}</div>
            <div><strong>Subjek:</strong> {{ submittedTicket.subject }}</div>
            <div><strong>Status:</strong> <span class="badge badge-amber">PENDING INVESTIGATION</span></div>
          </div>
          <button class="btn btn-primary" @click="resetForm">Ajukan Keluhan Lainnya</button>
        </div>

        <form v-else @submit.prevent="submitComplaint">
          <div v-if="errorMsg" class="error-box">
            <span>⚠️</span> {{ errorMsg }}
          </div>

          <div class="form-grid">
            <div class="form-group">
              <label class="form-label">Nama Lengkap Pelapor *</label>
              <input v-model="form.name" type="text" class="form-input" placeholder="Nama Anda" required />
            </div>

            <div class="form-group">
              <label class="form-label">Email Korespondensi *</label>
              <input v-model="form.email" type="email" class="form-input" placeholder="email@perusahaan.com" required />
            </div>
          </div>

          <div class="form-grid">
            <div class="form-group">
              <label class="form-label">Nomor Telepon / WhatsApp</label>
              <input v-model="form.phone" type="text" class="form-input" placeholder="+62 81..." />
            </div>

            <div class="form-group">
              <label class="form-label">Kategori Pengaduan *</label>
              <select v-model="form.category" class="form-select" required>
                <option value="">Pilih Kategori...</option>
                <option value="Kendala Teknis & Bug">Kendala Teknis & Bug Sistem</option>
                <option value="Pelanggaran SLA (Downtime)">Pelanggaran SLA & Waktu Operasional</option>
                <option value="Keamanan & Privasi Data">Keamanan Sistem & Privasi Data</option>
                <option value="Kualitas Layanan & Komunikasi">Kualitas Layanan & Tim Rekayasa</option>
                <option value="Lainnya">Lainnya / Whistleblowing</option>
              </select>
            </div>
          </div>

          <div class="form-group">
            <label class="form-label">Perihal / Pokok Masalah *</label>
            <input v-model="form.subject" type="text" class="form-input" placeholder="Ringkasan singkat kendala yang dialami" required />
          </div>

          <div class="form-group">
            <label class="form-label">Kronologi & Deskripsi Lengkap *</label>
            <textarea 
              v-model="form.description" 
              class="form-textarea" 
              placeholder="Jelaskan detail kendala, waktu kejadian, dampak yang dirasakan, serta harapan penyelesaian..."
              rows="6"
              required
            ></textarea>
          </div>

          <div class="form-actions">
            <button type="submit" class="btn btn-primary btn-lg" :disabled="loading">
              <span v-if="loading">Menerbitkan Tiket...</span>
              <span v-else>Kirimkan Pengaduan Resmi 🛡️</span>
            </button>
          </div>
        </form>
      </div>

      <!-- TAB 2: LACAK STATUS TIKET -->
      <div v-if="activeTab === 'track'" class="track-container glass-panel">
        <h3 class="track-title">Pencarian Tiket Pengaduan</h3>
        <p class="track-subtitle">Masukkan nomor tiket pengaduan yang Anda terima (contoh: <code>VK-CL-2026-001</code>).</p>

        <form @submit.prevent="lookupTicket" class="track-search-form">
          <input 
            v-model="searchTicketNumber" 
            type="text" 
            class="form-input mono" 
            placeholder="VK-CL-2026-XXXX" 
            required 
          />
          <button type="submit" class="btn btn-primary" :disabled="trackingLoading">
            {{ trackingLoading ? 'Mencari...' : 'Cek Tiket' }}
          </button>
        </form>

        <div v-if="trackError" class="error-box">
          {{ trackError }}
        </div>

        <div v-if="ticketResult" class="ticket-result-card glass-panel">
          <div class="res-header">
            <div>
              <span class="mono cyan-text">NOMOR TIKET</span>
              <h3 class="mono">{{ ticketResult.ticket_number }}</h3>
            </div>
            <span class="badge" :class="getStatusBadgeClass(ticketResult.status)">
              {{ ticketResult.status.toUpperCase() }}
            </span>
          </div>

          <div class="res-body">
            <div class="res-row">
              <span class="res-key">Perihal:</span>
              <span class="res-val">{{ ticketResult.subject }}</span>
            </div>
            <div class="res-row">
              <span class="res-key">Kategori:</span>
              <span class="res-val">{{ ticketResult.category }}</span>
            </div>
            <div class="res-row">
              <span class="res-key">Tanggal Diajukan:</span>
              <span class="res-val mono">{{ formatDate(ticketResult.created_at) }}</span>
            </div>
            <div class="res-row">
              <span class="res-key">Deskripsi Laporan:</span>
              <span class="res-val">{{ ticketResult.description }}</span>
            </div>
            <div class="res-row notes-row" v-if="ticketResult.resolution_notes">
              <span class="res-key green-text">Catatan Tim Penanganan:</span>
              <div class="res-val notes-box mono">{{ ticketResult.resolution_notes }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { api } from '../services/api'

const activeTab = ref('submit')
const loading = ref(false)
const trackingLoading = ref(false)
const errorMsg = ref('')
const trackError = ref('')
const submittedTicket = ref(null)
const searchTicketNumber = ref('')
const ticketResult = ref(null)

const form = reactive({
  name: '',
  email: '',
  phone: '',
  category: '',
  subject: '',
  description: ''
})

const submitComplaint = async () => {
  loading.value = true
  errorMsg.value = ''

  try {
    const res = await api.sendComplaint(form)
    submittedTicket.value = res.complaint
  } catch (err) {
    errorMsg.value = err.message || 'Gagal mengirimkan pengaduan. Silakan periksa kembali isian Anda.'
  } finally {
    loading.value = false
  }
}

const lookupTicket = async () => {
  if (!searchTicketNumber.value.trim()) return
  trackingLoading.value = true
  trackError.value = ''
  ticketResult.value = null

  try {
    const res = await api.getComplaintByTicket(searchTicketNumber.value.trim())
    ticketResult.value = res
  } catch (err) {
    trackError.value = err.message || 'Nomor tiket tidak ditemukan dalam basis data sistem kami.'
  } finally {
    trackingLoading.value = false
  }
}

const resetForm = () => {
  submittedTicket.value = null
  form.name = ''
  form.email = ''
  form.phone = ''
  form.category = ''
  form.subject = ''
  form.description = ''
}

const getStatusBadgeClass = (status) => {
  switch (status) {
    case 'resolved': return 'badge-live'
    case 'investigating': return 'badge-cyan'
    case 'rejected': return 'badge-amber'
    default: return 'badge-violet'
  }
}

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>

<style scoped>
.complaint-view {
  padding: 130px 0 80px;
  position: relative;
  z-index: 1;
}

.page-header {
  text-align: center;
  max-width: 800px;
  margin: 0 auto 40px;
}

.page-title {
  font-size: clamp(2.2rem, 4vw, 3.2rem);
  font-weight: 800;
  margin: 16px 0;
}

.page-subtitle {
  font-size: 1.05rem;
  color: var(--text-muted);
  line-height: 1.7;
}

.tabs-container {
  display: flex;
  justify-content: center;
  gap: 12px;
  margin-bottom: 36px;
}

.tab-btn {
  padding: 12px 24px;
  border-radius: var(--radius-full);
  background: #ffffff;
  border: 1px solid rgba(20, 60, 35, 0.12);
  color: var(--text-muted);
  font-size: 0.95rem;
  font-weight: 600;
  cursor: pointer;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.02);
  transition: all 0.25s;
}

.tab-btn.active {
  background: #d97706;
  border-color: #d97706;
  color: #ffffff;
  box-shadow: 0 4px 14px rgba(217, 119, 6, 0.3);
}

.form-container, .track-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 40px;
  background: rgba(255, 255, 255, 0.92);
  border: 1px solid rgba(20, 60, 35, 0.08);
  box-shadow: var(--shadow-card);
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.form-actions {
  margin-top: 10px;
}

.success-ticket-box {
  text-align: center;
  padding: 20px;
}

.ticket-badge {
  font-size: 0.75rem;
  color: #d97706;
  font-weight: 700;
  letter-spacing: 0.1em;
  margin-bottom: 8px;
}

.ticket-number {
  font-size: 2.2rem;
  color: var(--cyan-primary);
  margin-bottom: 16px;
  letter-spacing: 0.05em;
}

.ticket-desc {
  color: var(--text-muted);
  font-size: 0.95rem;
  line-height: 1.6;
  max-width: 600px;
  margin: 0 auto 24px;
}

.ticket-summary {
  background: #f8faf6;
  padding: 20px;
  border-radius: var(--radius-sm);
  text-align: left;
  display: flex;
  flex-direction: column;
  gap: 8px;
  font-size: 0.88rem;
  margin-bottom: 30px;
  border: 1px solid rgba(20, 60, 35, 0.08);
  color: var(--text-main);
}

.track-search-form {
  display: flex;
  gap: 12px;
  margin-top: 20px;
  margin-bottom: 30px;
}

.ticket-result-card {
  padding: 24px;
  background: #ffffff;
  border: 1px solid rgba(20, 60, 35, 0.1);
}

.res-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 16px;
  border-bottom: 1px solid rgba(20, 60, 35, 0.08);
  margin-bottom: 16px;
}

.res-row {
  display: flex;
  margin-bottom: 12px;
  font-size: 0.9rem;
}

.res-key {
  width: 160px;
  color: var(--text-dim);
  font-weight: 600;
  flex-shrink: 0;
}

.res-val {
  color: var(--text-main);
  font-weight: 500;
  flex: 1;
}

.notes-box {
  background: #ecfdf5;
  border: 1px solid rgba(22, 163, 74, 0.25);
  padding: 12px;
  border-radius: var(--radius-sm);
  color: #15803d;
}

.green-text { color: #15803d; }
.cyan-text { color: var(--cyan-primary); }

.error-box {
  background: #fff1f2;
  border: 1px solid rgba(225, 29, 72, 0.25);
  color: #e11d48;
  padding: 12px 16px;
  border-radius: var(--radius-sm);
  margin-bottom: 20px;
}

@media (max-width: 650px) {
  .form-grid {
    grid-template-columns: 1fr;
  }
  .res-row {
    flex-direction: column;
    gap: 4px;
  }
}
</style>
