<template>
  <section class="identity-section" id="about">
    <div class="container">
      <div class="section-header">
        <span class="badge badge-violet mono">Prinsip Rekayasa &amp; Tata Kelola</span>
        <h2 class="section-title">
          Fondasi Rekayasa Perangkat Lunak <span class="text-highlight">Berstandar Industri Tertinggi.</span>
        </h2>
        <p class="section-subtitle">
          PT. Vadikara Rekayasa Peranti memadukan ketelitian matematika arsitektur terdistribusi, disiplin pengujian beban ekstrem, serta kepatuhan hukum penuh pada regulasi perlindungan data Republik Indonesia.
        </p>
      </div>

      <!-- 4 Core Pillars of Engineering Excellence -->
      <div class="pillars-grid">
        <div 
          class="pillar-card tilt-card"
          @mousemove="handlePillarMove($event)"
        >
          <div class="tilt-glare"></div>
          <div class="pillar-num mono">01</div>
          <h3 class="pillar-title">Arsitektur Deterministik &amp; Bebas Kebocoran Memori</h3>
          <p class="pillar-desc">
            Menggunakan model konkurensi native Golang (Goroutines &amp; Channels) dengan footprint memori yang dapat diprediksi secara matematis, memastikan sistem tidak mengalami memory leak saat melayani beban puncak ratusan ribu transaksi per detik.
          </p>
          <div class="pillar-metric mono">
            <span>Standar Latensi:</span>
            <strong>p99 &lt; 1.2ms (Zero GC Stalls)</strong>
          </div>
        </div>

        <div 
          class="pillar-card tilt-card"
          @mousemove="handlePillarMove($event)"
        >
          <div class="tilt-glare"></div>
          <div class="pillar-num mono">02</div>
          <h3 class="pillar-title">Kedaulatan &amp; Kepatuhan UU Pelindungan Data Pribadi</h3>
          <p class="pillar-desc">
            Seluruh platform dirancang selaras dengan mandat UU No. 27 Tahun 2022 (UU PDP). Menerapkan enkripsi simetris AES-256 pada media penyimpanan persisten, kanal TLS 1.3, isolasi data subjek, dan hak penghapusan data otomatis (*Right to be Forgotten*).
          </p>
          <div class="pillar-metric mono">
            <span>Standar Kepatuhan:</span>
            <strong>UU RI No. 27/2022 Terverifikasi</strong>
          </div>
        </div>

        <div 
          class="pillar-card tilt-card"
          @mousemove="handlePillarMove($event)"
        >
          <div class="tilt-glare"></div>
          <div class="pillar-num mono">03</div>
          <h3 class="pillar-title">Ketahanan Operasional 99.99% Tanpa Single Point of Failure</h3>
          <p class="pillar-desc">
            Merancang clustering multi-zona berbasis konsensus Raft, replikasi sinkronus database ACID, serta arsitektur circuit breaker otomatis yang mampu mengalihkan beban kerja (failover) dalam &lt; 150 milidetik tanpa interupsi pengguna.
          </p>
          <div class="pillar-metric mono">
            <span>Tingkat Uptime:</span>
            <strong>99.99% Kontraktual SLA</strong>
          </div>
        </div>

        <div 
          class="pillar-card tilt-card"
          @mousemove="handlePillarMove($event)"
        >
          <div class="tilt-glare"></div>
          <div class="pillar-num mono">04</div>
          <h3 class="pillar-title">Legalitas Berbadan Hukum &amp; Kesiapan Audit Korporasi</h3>
          <p class="pillar-desc">
            Berbadan hukum resmi Perseroan Terbatas yang terdaftar di Kementerian Hukum dan HAM RI, memiliki Nomor Induk Berusaha (NIB) lengkap, izin operasional pengadaan peranti lunak enterprise, dan rekam jejak audit kode deterministik.
          </p>
          <div class="pillar-metric mono">
            <span>Status Entitas:</span>
            <strong>PT Resmi • NIB: 0128934710294</strong>
          </div>
        </div>
      </div>

      <!-- Governance & Leadership Credentials Bar -->
      <div class="governance-box">
        <div class="gov-left">
          <div class="gov-badge mono">PROFIL KEPEMIMPINAN REKAYASA</div>
          <h3 class="gov-lead-title">Dipimpin oleh Principal Systems Architect Berpengalaman</h3>
          <p class="gov-lead-desc">
            Divisi arsitektur PT. Vadikara Rekayasa Peranti dipimpin langsung oleh Founder &amp; Principal Architect <strong>Vadikara</strong> dengan pengalaman lebih dari 8 tahun merancang sistem komputasi berkonkurensi tinggi, arsitektur event sourcing di sektor keuangan, serta modernisasi platform web reaktif skala nasional.
          </p>
        </div>

        <div class="gov-right">
          <div class="gov-specs-table mono">
            <div class="gov-spec-row">
              <span class="spec-k">NAMA BADAN HUKUM:</span>
              <span class="spec-v">PT. VADIKARA REKAYASA PERANTI</span>
            </div>
            <div class="gov-spec-row">
              <span class="spec-k">NOMOR INDUK BERUSAHA (NIB):</span>
              <span class="spec-v highlight-green copyable-nib" @click="copyNib" title="Klik untuk menyalin NIB">
                {{ nibCopied ? 'Tersalin! ✓' : '0128934710294 📋' }}
              </span>
            </div>
            <div class="gov-spec-row">
              <span class="spec-k">KLASIFIKASI BAKU (KBLI):</span>
              <span class="spec-v">62019 (Pemrograman Perangkat Lunak)</span>
            </div>
            <div class="gov-spec-row">
              <span class="spec-k">PENGESAHAN NEGARA:</span>
              <span class="spec-v">Kemenkumham RI Terverifikasi</span>
            </div>
            <div class="gov-spec-row">
              <span class="spec-k">KANTOR PUSAT OPERASIONAL:</span>
              <span class="spec-v">TB Simatupang, Cilandak, Jakarta Selatan 12430</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from 'vue'

const nibCopied = ref(false)

const copyNib = async () => {
  try {
    await navigator.clipboard.writeText('0128934710294')
    nibCopied.value = true
    setTimeout(() => { nibCopied.value = false }, 2000)
  } catch (err) {
    nibCopied.value = true
    setTimeout(() => { nibCopied.value = false }, 2000)
  }
}

const handlePillarMove = (event) => {
  const card = event.currentTarget
  const rect = card.getBoundingClientRect()
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top
  card.style.setProperty('--mouse-x', `${x}px`)
  card.style.setProperty('--mouse-y', `${y}px`)
}
</script>

<style scoped>
.identity-section {
  padding: 90px 0;
  background: #ffffff;
  border-top: 1px solid #e5e7eb;
  border-bottom: 1px solid #e5e7eb;
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

/* 4 Pillars Grid */
.pillars-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24px;
  margin-bottom: 48px;
}

.pillar-card {
  padding: 32px;
  background: #fbfcf9;
  border: 1px solid #e5e7eb;
  border-radius: var(--radius-md);
  display: flex;
  flex-direction: column;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}

.pillar-card:hover {
  border-color: #166534;
  box-shadow: 0 4px 12px rgba(15, 36, 25, 0.05);
}

.pillar-num {
  font-size: 0.8rem;
  font-weight: 700;
  color: #166534;
  margin-bottom: 14px;
}

.pillar-title {
  font-size: 1.2rem;
  font-weight: 700;
  color: #111827;
  margin-bottom: 12px;
  line-height: 1.4;
}

.pillar-desc {
  font-size: 0.92rem;
  color: #4b5563;
  line-height: 1.7;
  margin-bottom: 24px;
}

.pillar-metric {
  margin-top: auto;
  padding-top: 14px;
  border-top: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.78rem;
}

.pillar-metric span {
  color: #6b7280;
}

.pillar-metric strong {
  color: #14532d;
}

/* Governance & Leadership Box */
.governance-box {
  display: grid;
  grid-template-columns: 1.1fr 1.2fr;
  gap: 40px;
  padding: 36px;
  background: #f8faf6;
  border: 1px solid #e2e8df;
  border-radius: var(--radius-md);
  align-items: center;
}

.gov-badge {
  font-size: 0.7rem;
  font-weight: 700;
  color: #166534;
  letter-spacing: 0.08em;
  margin-bottom: 8px;
}

.gov-lead-title {
  font-size: 1.35rem;
  font-weight: 800;
  color: #111827;
  margin-bottom: 12px;
  line-height: 1.3;
}

.gov-lead-desc {
  font-size: 0.92rem;
  color: #4b5563;
  line-height: 1.7;
}

.gov-lead-desc strong {
  color: #111827;
}

.gov-specs-table {
  display: flex;
  flex-direction: column;
  gap: 10px;
  background: #ffffff;
  padding: 20px;
  border-radius: var(--radius-sm);
  border: 1px solid #e5e7eb;
  font-size: 0.78rem;
}

.gov-spec-row {
  display: flex;
  justify-content: space-between;
  padding-bottom: 8px;
  border-bottom: 1px solid #f3f4f6;
}

.gov-spec-row:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.spec-k {
  color: #6b7280;
}

.spec-v {
  color: #111827;
  font-weight: 600;
  text-align: right;
  max-width: 60%;
}

.highlight-green {
  color: #14532d;
}

.copyable-nib {
  cursor: pointer;
  background: #f0fdf4;
  padding: 2px 8px;
  border-radius: 4px;
  border: 1px solid #bbf7d0;
  transition: all 0.2s ease;
  user-select: none;
}

.copyable-nib:hover {
  background: #dcfce7;
  border-color: #166534;
}

@media (max-width: 960px) {
  .pillars-grid, .governance-box {
    grid-template-columns: 1fr;
  }
}
</style>
