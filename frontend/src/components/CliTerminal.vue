<template>
  <div class="cli-overlay" v-if="isOpen" @click.self="$emit('close')">
    <div class="cli-window glass-panel">
      <!-- Terminal Header Bar -->
      <div class="cli-header">
        <div class="cli-controls">
          <span class="ctrl-dot dot-red" @click="$emit('close')"></span>
          <span class="ctrl-dot dot-yellow"></span>
          <span class="ctrl-dot dot-green"></span>
        </div>
        <div class="cli-title mono">vadikara@rekayasa-peranti: ~ (bash / Go runtime v1.26)</div>
        <div class="cli-actions">
          <button class="cli-btn-close mono" @click="$emit('close')">[esc to close]</button>
        </div>
      </div>

      <!-- Terminal Body -->
      <div class="cli-body mono" ref="bodyRef">
        <div class="cli-line intro-text">
          <span class="cyan-text">PT. VADIKARA REKAYASA PERANTI - CLI v2.6.4</span><br />
          Ketik <span class="highlight-text">'help'</span> untuk melihat daftar instruksi sistem.
        </div>

        <div v-for="(item, idx) in history" :key="idx" class="cli-history-item">
          <div class="cli-command-row">
            <span class="cli-prompt">vadikara@rekayasa:~$</span>
            <span class="cli-command">{{ item.cmd }}</span>
          </div>
          <div class="cli-output" v-html="item.output"></div>
        </div>

        <!-- Current Input Line -->
        <div class="cli-input-row">
          <span class="cli-prompt">vadikara@rekayasa:~$</span>
          <input 
            ref="inputRef"
            v-model="currentInput"
            @keydown.enter="handleCommand"
            @keydown.up.prevent="navigateHistory(-1)"
            @keydown.down.prevent="navigateHistory(1)"
            type="text" 
            class="cli-input mono" 
            autocomplete="off"
            spellcheck="false"
            autofocus
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick, watch } from 'vue'

const props = defineProps({
  isOpen: Boolean
})
const emit = defineEmits(['close'])

const inputRef = ref(null)
const bodyRef = ref(null)
const currentInput = ref('')
const commandHistory = ref([])
const historyIndex = ref(-1)

const history = ref([
  {
    cmd: 'sysinfo',
    output: `Entity: PT. Vadikara Rekayasa Peranti<br/>Principal: Dev Vadikara<br/>Status: Operational (99.99% SLA)<br/>NIB: 0128934710294`
  }
])

const handleCommand = () => {
  const cmd = currentInput.value.trim().toLowerCase()
  if (!cmd) return

  commandHistory.value.push(cmd)
  historyIndex.value = commandHistory.value.length

  let output = ''

  switch (cmd) {
    case 'help':
      output = `
Daftar Perintah Tersedia:<br/>
&nbsp;&nbsp;<span class="cyan-text">about</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;- Profil singkat Vadikara & PT. Rekayasa Peranti<br/>
&nbsp;&nbsp;<span class="cyan-text">services</span>&nbsp;&nbsp;- Layanan rekayasa perangkat lunak enterprise<br/>
&nbsp;&nbsp;<span class="cyan-text">projects</span>&nbsp;&nbsp;- Ringkasan proyek & performa sistem<br/>
&nbsp;&nbsp;<span class="cyan-text">skills</span>&nbsp;&nbsp;&nbsp;&nbsp;- Stack teknologi & arsitektur<br/>
&nbsp;&nbsp;<span class="cyan-text">contact</span>&nbsp;&nbsp;&nbsp;- Informasi kontak & hotline WhatsApp<br/>
&nbsp;&nbsp;<span class="cyan-text">complaint</span>&nbsp;- Layanan pengaduan resmi / tiket keluhan<br/>
&nbsp;&nbsp;<span class="cyan-text">privacy</span>&nbsp;&nbsp;&nbsp;- Hak privasi & penghapusan data (UU PDP)<br/>
&nbsp;&nbsp;<span class="cyan-text">clear</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;- Bersihkan layar terminal<br/>
&nbsp;&nbsp;<span class="cyan-text">exit</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;- Tutup jendela terminal`
      break
    case 'about':
      output = `
PT. VADIKARA REKAYASA PERANTI<br/>
Badan usaha spesialis rekayasa teknologi dan peranti lunak berdaya tahan tinggi.<br/>
Didirikan dan dipimpin oleh Vadikara (Principal Systems Architect) dengan track record membangun backend Golang 180k+ RPS, microservices gRPC terdistribusi, dan aplikasi web modern Vue 3.`
      break
    case 'services':
      output = `
1. High-Throughput Backend (Golang, gRPC, Kafka)<br/>
2. Modern Web Applications (Vue 3, Vite, HTML5 Canvas)<br/>
3. Enterprise Cloud Architecture & Kubernetes DevOps<br/>
4. AI Integration & Automated Reasoning Pipelines<br/>
5. Resilience, Security & Zero-Trust Audit`
      break
    case 'projects':
      output = `
• NovaMesh Engine: 185k req/sec, p99 &lt; 1.1ms<br/>
• OmniVortex ERP: IDR 65M/bulan GMV, 15k active users<br/>
• AegisGuard Mesh: Zero-Trust eBPF Packet Telemetry<br/>
• NeuralSense AI: 380ms Latency OCR Document Extraction<br/>
• FinPulse Core: 220k Order Matching/detik`
      break
    case 'skills':
      output = `
Core Languages: Golang, TypeScript, HTML5/CSS, SQL, Python<br/>
Infrastructure: Kubernetes, Docker, Terraform, Kafka, Redis, WireGuard<br/>
Frontend: Vue 3, Vite, Pinia, WebGL / Canvas API`
      break
    case 'contact':
      output = `
Email: contact@vadikara.com<br/>
WhatsApp: +62 812-8800-9921<br/>
Kantor: Gedung Rekayasa Digital Lt. 7, Jl. TB Simatupang No. 88, Jakarta Selatan`
      break
    case 'complaint':
      output = `
Untuk mengajukan pengaduan resmi atau klaim SLA:<br/>
Kunjungi menu <span class="highlight-text">/pengaduan</span> atau kirim email ke support@vadikara.com.<br/>
Nomor tiket resmi (format: VK-CL-2026-XXXX) akan diterbitkan otomatis.`
      break
    case 'privacy':
      output = `
PT. Vadikara Rekayasa Peranti tunduk pada UU No. 27/2022 (UU PDP).<br/>
Anda berhak mengajukan permohonan penghapusan data pribadi pada halaman <span class="highlight-text">/privacy-policy</span>.`
      break
    case 'clear':
      history.value = []
      currentInput.value = ''
      return
    case 'exit':
      emit('close')
      return
    case 'sudo rm -rf /':
      output = `<span class="red-text">[SECURITY ALERT] Unauthorized root purge attempt blocked by AegisGuard Zero-Trust Kernel Module. IP logged to security audit.</span>`
      break
    default:
      output = `<span class="red-text">Perintah tidak dikenal: '${cmd}'. Ketik 'help' untuk panduan.</span>`
  }

  history.value.push({
    cmd: currentInput.value,
    output
  })

  currentInput.value = ''

  nextTick(() => {
    if (bodyRef.value) {
      bodyRef.value.scrollTop = bodyRef.value.scrollHeight
    }
  })
}

const navigateHistory = (direction) => {
  if (commandHistory.value.length === 0) return
  historyIndex.value += direction
  if (historyIndex.value < 0) historyIndex.value = 0
  if (historyIndex.value >= commandHistory.value.length) {
    historyIndex.value = commandHistory.value.length
    currentInput.value = ''
    return
  }
  currentInput.value = commandHistory.value[historyIndex.value]
}

const handleKeyDown = (e) => {
  if (e.key === 'Escape' && props.isOpen) {
    emit('close')
  }
}

watch(() => props.isOpen, (newVal) => {
  if (newVal) {
    nextTick(() => {
      inputRef.value?.focus()
      if (bodyRef.value) {
        bodyRef.value.scrollTop = bodyRef.value.scrollHeight
      }
    })
  }
})

onMounted(() => {
  window.addEventListener('keydown', handleKeyDown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeyDown)
})
</script>

<style scoped>
.cli-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(4, 7, 12, 0.75);
  backdrop-filter: blur(10px);
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.cli-window {
  width: 100%;
  max-width: 820px;
  height: 520px;
  background: #0d1e15;
  border: 1px solid rgba(22, 163, 74, 0.35);
  border-radius: var(--radius-md);
  box-shadow: 0 25px 60px rgba(13, 30, 21, 0.35);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.cli-header {
  padding: 12px 18px;
  background: #14281c;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.cli-controls {
  display: flex;
  gap: 8px;
}

.ctrl-dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  cursor: pointer;
}

.dot-red { background: #ef4444; }
.dot-yellow { background: #f59e0b; }
.dot-green { background: #10b981; }

.cli-title {
  font-size: 0.8rem;
  color: #94a398;
}

.cli-btn-close {
  background: transparent;
  border: none;
  color: #768b7b;
  cursor: pointer;
  font-size: 0.75rem;
}

.cli-btn-close:hover {
  color: #fff;
}

.cli-body {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  font-size: 0.9rem;
  line-height: 1.6;
}

.intro-text {
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px dashed rgba(255, 255, 255, 0.1);
  color: #94a398;
}

.cli-history-item {
  margin-bottom: 14px;
}

.cli-command-row {
  display: flex;
  gap: 10px;
  margin-bottom: 4px;
}

.cli-prompt {
  color: #4ade80;
  font-weight: 600;
}

.cli-command {
  color: #fff;
}

.cli-output {
  color: #d1ded4;
  margin-left: 8px;
  padding-left: 8px;
  border-left: 2px solid rgba(74, 222, 128, 0.3);
}

.cli-input-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.cli-input {
  flex: 1;
  background: transparent;
  border: none;
  color: #4ade80;
  font-size: 0.9rem;
  outline: none;
  caret-color: #4ade80;
}

:deep(.cyan-text) { color: #4ade80; }
:deep(.highlight-text) { color: #fbbf24; font-weight: 600; }
:deep(.red-text) { color: #f43f5e; }
</style>
