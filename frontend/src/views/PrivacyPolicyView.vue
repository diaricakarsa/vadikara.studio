<template>
  <div class="privacy-view">
    <div class="container page-container">
      <div class="page-header">
        <span class="badge badge-cyan mono">Kepatuhan UU No. 27 Tahun 2022 (UU PDP)</span>
        <h1 class="page-title">Kebijakan Privasi & <span class="text-gradient-cyan">Pelindungan Data.</span></h1>
        <p class="page-subtitle">
          PT. Vadikara Rekayasa Peranti memegang teguh prinsip integritas, kerahasiaan, dan kedaulatan data pribadi mitra, calon klien, dan pengunjung situs sesuai peraturan perundang-undangan Republik Indonesia.
        </p>
      </div>

      <div class="policy-layout">
        <!-- Main Document Body -->
        <div class="policy-document glass-panel">
          <section class="doc-section">
            <h2>1. Landasan Hukum & Definisi</h2>
            <p>
              Kebijakan Privasi ini disusun berlandaskan Undang-Undang Republik Indonesia Nomor 27 Tahun 2022 tentang Pelindungan Data Pribadi (UU PDP). 
              PT. Vadikara Rekayasa Peranti (selanjutnya disebut "Perusahaan", "Kami", atau "Pengendali Data Pribadi") berkedudukan di Jakarta Selatan, DKI Jakarta, dengan Nomor Induk Berusaha (NIB) 0128934710294.
            </p>
          </section>

          <section class="doc-section">
            <h2>2. Data Pribadi yang Kami Kumpulkan</h2>
            <p>Kami hanya mengumpulkan data yang relevan dan terbatas untuk keperluan komunikasi teknis dan administrasi kemitraan:</p>
            <ul>
              <li><strong>Data Kontak:</strong> Nama lengkap, alamat surat elektronik (email), nomor telepon/WhatsApp, dan nama institusi/perusahaan yang Anda berikan melalui formulir kontak atau proposal.</li>
              <li><strong>Data Pengaduan & Layanan:</strong> Rincian keluhan, nomor tiket bantuan teknis, dan catatan resolusi.</li>
              <li><strong>Data Teknis Otomatis:</strong> Alamat Protokol Internet (IP address), log peramban, dan telemetri koneksi untuk keperluan pencegahan serangan siber (DDoS mitigation & Zero-Trust).</li>
            </ul>
          </section>

          <section class="doc-section">
            <h2>3. Tujuan Pemrosesan & Dasar Kebijakan</h2>
            <p>Data pribadi Anda diproses secara sah atas dasar:</p>
            <ul>
              <li>Persetujuan eksplisit subjek data saat mengirimkan formulir permohonan proposal teknis;</li>
              <li>Pemenuhan kewajiban kontraktual perjanjian kerja sama rekayasa perangkat lunak;</li>
              <li>Kepentingan yang sah dalam menjaga ketahanan infrastruktur komputasi kami dari ancaman intrusi.</li>
            </ul>
          </section>

          <section class="doc-section">
            <h2>4. Standar Keamanan & Enkripsi</h2>
            <p>
              Kami menerapkan standar enkripsi TLS 1.3 pada transmisi data dan enkripsi AES-256 pada media penyimpanan persisten. Seluruh data disimpan dalam pusat data berstandar Tier-3 dengan isolasi jaringan Zero-Trust. Kami tidak pernah memperjualbelikan, menyewakan, atau menyebarkan data pribadi Anda kepada pihak ketiga mana pun tanpa persetujuan tertulis.
            </p>
          </section>

          <section class="doc-section" id="rights">
            <h2>5. Hak Subjek Data Pribadi (Pasal 5 s/d Pasal 13 UU PDP)</h2>
            <p>Sebagai Subjek Data Pribadi, Anda memiliki hak-hak yang dilindungi hukum, antara lain:</p>
            <ul>
              <li><strong>Hak Mendapatkan Informasi:</strong> Mengetahui kejelasan identitas pengendali dan tujuan penggunaan data.</li>
              <li><strong>Hak Akses & Salinan Data:</strong> Memperoleh akses terhadap rekaman data pribadi Anda.</li>
              <li><strong>Hak Pembaruan / Ralat:</strong> Mengubah data yang keliru atau tidak akurat.</li>
              <li><strong>Hak Penghapusan Data (Right to be Forgotten):</strong> Menghapus dan memusnahkan data pribadi Anda dari sistem kami bila tidak lagi diperlukan untuk pelaksanaan kontrak.</li>
            </ul>
          </section>

          <!-- FORMULIR HAK PENGHAPUSAN DATA PRIBADI -->
          <section class="doc-section deletion-highlight" id="deletion">
            <div class="deletion-header">
              <span class="badge badge-amber mono">UU PDP PASAL 8 & 9</span>
              <h3>Formulir Permohonan Penghapusan Data Pribadi</h3>
              <p>
                Jika Anda sebelumnya pernah mengirimkan pesan, proposal, atau pengaduan dan ingin seluruh data pribadi Anda dihapus secara permanen dari server PT. Vadikara Rekayasa Peranti, silakan ajukan permohonan resmi di bawah ini:
              </p>
            </div>

            <div v-if="deletionSuccess" class="success-box">
              <span>✅</span>
              <div>
                <strong>Permohonan Penghapusan Berhasil Diterima</strong>
                <p>{{ deletionSuccess }}</p>
              </div>
            </div>

            <form v-else @submit.prevent="submitDeletionRequest" class="deletion-form">
              <div v-if="deletionError" class="error-box">
                {{ deletionError }}
              </div>

              <div class="form-row">
                <div class="form-group">
                  <label class="form-label">Nama Terdaftar *</label>
                  <input v-model="delForm.name" type="text" class="form-input" placeholder="Nama Anda saat mengirim kontak" required />
                </div>
                <div class="form-group">
                  <label class="form-label">Email Terdaftar *</label>
                  <input v-model="delForm.email" type="email" class="form-input" placeholder="email@yang-dihapus.com" required />
                </div>
              </div>

              <div class="form-group">
                <label class="form-label">Alasan Permohonan Penghapusan</label>
                <textarea v-model="delForm.reason" class="form-textarea" placeholder="contoh: Proyek telah selesai / Menggunakan hak penghapusan data berdasarkan UU PDP Pasal 8..." rows="3"></textarea>
              </div>

              <button type="submit" class="btn btn-danger mono" :disabled="delLoading">
                {{ delLoading ? 'Memproses Permohonan...' : 'Ajukan Penghapusan Data Pribadi 🗑️' }}
              </button>
            </form>
          </section>

          <section class="doc-section">
            <h2>6. Kontak Pejabat Pelindungan Data (DPO)</h2>
            <p>
              Untuk pertanyaan, penegakan hak privasi, atau permintaan data audit lebih lanjut, hubungi Petugas Pelindungan Data Pribadi kami:
            </p>
            <div class="dpo-card mono">
              <strong>Data Protection Officer (DPO)</strong><br />
              PT. Vadikara Rekayasa Peranti<br />
              Email: <span class="cyan-text">privacy@vadikara.com</span> / <span class="cyan-text">dpo@vadikara.com</span><br />
              Alamat: Gedung Rekayasa Digital Lt. 7, Jl. TB Simatupang No. 88, Jakarta Selatan 12430
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { api } from '../services/api'

const delLoading = ref(false)
const deletionSuccess = ref('')
const deletionError = ref('')

const delForm = reactive({
  name: '',
  email: '',
  reason: ''
})

const submitDeletionRequest = async () => {
  delLoading.value = true
  deletionError.value = ''
  deletionSuccess.value = ''

  try {
    const res = await api.requestDataDeletion(delForm)
    deletionSuccess.value = res.message || 'Permohonan Anda berhasil diajukan dan akan diproses dalam 3x24 jam kerja.'
    delForm.name = ''
    delForm.email = ''
    delForm.reason = ''
  } catch (err) {
    deletionError.value = err.message || 'Gagal memproses permohonan. Pastikan isian sudah lengkap.'
  } finally {
    delLoading.value = false
  }
}
</script>

<style scoped>
.privacy-view {
  padding: 130px 0 80px;
  position: relative;
  z-index: 1;
}

.page-header {
  text-align: center;
  max-width: 820px;
  margin: 0 auto 50px;
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

.policy-document {
  max-width: 880px;
  margin: 0 auto;
  padding: 48px;
  background: rgba(255, 255, 255, 0.94);
  border: 1px solid rgba(20, 60, 35, 0.08);
  box-shadow: var(--shadow-card);
}

.doc-section {
  margin-bottom: 40px;
}

.doc-section h2 {
  font-size: 1.4rem;
  font-weight: 700;
  margin-bottom: 14px;
  color: var(--text-main);
  border-bottom: 1px solid rgba(20, 60, 35, 0.08);
  padding-bottom: 8px;
}

.doc-section p, .doc-section li {
  color: var(--text-muted);
  line-height: 1.8;
  font-size: 0.98rem;
}

.doc-section ul {
  margin-left: 20px;
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.deletion-highlight {
  background: #fffbeb;
  border: 1px solid rgba(217, 119, 6, 0.25);
  padding: 30px;
  border-radius: var(--radius-md);
}

.deletion-header h3 {
  font-size: 1.3rem;
  margin: 12px 0 8px;
  color: #b45309;
}

.deletion-form {
  margin-top: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.dpo-card {
  background: #f8faf6;
  padding: 20px;
  border-radius: var(--radius-sm);
  border: 1px solid rgba(20, 60, 35, 0.08);
  font-size: 0.88rem;
  line-height: 1.8;
  margin-top: 14px;
  color: var(--text-main);
}

.cyan-text { color: var(--cyan-primary); }

.success-box {
  display: flex;
  gap: 14px;
  background: #ecfdf5;
  border: 1px solid rgba(22, 163, 74, 0.25);
  color: #15803d;
  padding: 20px;
  border-radius: var(--radius-sm);
  margin-top: 20px;
}

.error-box {
  background: #fff1f2;
  border: 1px solid rgba(225, 29, 72, 0.25);
  color: #e11d48;
  padding: 12px 16px;
  border-radius: var(--radius-sm);
  margin-bottom: 16px;
  font-size: 0.9rem;
}

@media (max-width: 700px) {
  .policy-document {
    padding: 28px;
  }
  .form-row {
    grid-template-columns: 1fr;
  }
}
</style>
