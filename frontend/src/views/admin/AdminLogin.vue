<template>
  <div class="login-view">
    <div class="ambient-glow"></div>

    <div class="login-card glass-panel">
      <!-- Top Brand Mark -->
      <div class="login-header">
        <div class="logo-box">
          <span class="logo-char mono">V</span>
          <span class="logo-pulse"></span>
        </div>
        <h1 class="login-title">Vadikara CMS</h1>
        <p class="login-subtitle mono">PUSAT KONTROL KONTEN PORTFOLIO</p>
      </div>

      <!-- Error Message -->
      <transition name="fade">
        <div v-if="errorMsg" class="error-box mono">
          <span>⚠️</span>
          <span>{{ errorMsg }}</span>
        </div>
      </transition>

      <!-- Form -->
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label class="form-label mono">EMAIL ADMINISTRATOR</label>
          <div class="input-wrap">
            <span class="input-icon">✉️</span>
            <input 
              v-model="email" 
              type="email" 
              class="form-input mono" 
              placeholder="admin@vadikara.com" 
              autocomplete="email"
              required 
            />
          </div>
        </div>

        <div class="form-group">
          <label class="form-label mono">KATA SANDI</label>
          <div class="input-wrap">
            <span class="input-icon">🔒</span>
            <input 
              v-model="password" 
              :type="showPassword ? 'text' : 'password'" 
              class="form-input mono" 
              placeholder="••••••••••••" 
              autocomplete="current-password"
              required 
            />
            <button 
              type="button" 
              class="eye-btn" 
              @click="showPassword = !showPassword"
              :aria-label="showPassword ? 'Sembunyikan sandi' : 'Tampilkan sandi'"
            >
              {{ showPassword ? '👁️' : '👁️‍🗨️' }}
            </button>
          </div>
        </div>

        <!-- Quick Fill Helper Card -->
        <div class="credential-helper" @click="fillDefaultCredentials" title="Klik untuk mengisi otomatis">
          <div class="helper-top">
            <span class="helper-label mono">KREDENSIAL DEFAULT SISTEM</span>
            <span class="helper-action mono">KLIK UNTUK ISI ⚡</span>
          </div>
          <div class="helper-values mono">
            <code>admin@vadikara.com</code> &bull; <code>vadikara2026!</code>
          </div>
        </div>

        <button type="submit" class="btn btn-primary btn-submit mono" :disabled="loading">
          <span v-if="loading">Memverifikasi Otorisasi...</span>
          <span v-else>Masuk ke Portal CMS &rarr;</span>
        </button>
      </form>

      <div class="login-footer">
        <router-link to="/" class="back-link mono">
          &larr; Kembali ke Website Publik
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { api } from '../../services/api'

const router = useRouter()
const email = ref('admin@vadikara.com')
const password = ref('vadikara2026!')
const showPassword = ref(false)
const loading = ref(false)
const errorMsg = ref('')

const fillDefaultCredentials = () => {
  email.value = 'admin@vadikara.com'
  password.value = 'vadikara2026!'
}

const handleLogin = async () => {
  loading.value = true
  errorMsg.value = ''

  try {
    const res = await api.login(email.value, password.value)
    localStorage.setItem('vk_admin_token', res.token)
    localStorage.setItem('vk_admin_user', JSON.stringify(res.user))
    router.push('/admin')
  } catch (err) {
    errorMsg.value = err.message || 'Email atau kata sandi tidak valid.'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-view {
  min-height: 100vh;
  width: 100%;
  background: #080808;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  position: relative;
  overflow: hidden;
}

.ambient-glow {
  position: absolute;
  top: 15%;
  left: 50%;
  transform: translateX(-50%);
  width: 600px;
  height: 400px;
  background: radial-gradient(circle, rgba(0, 242, 254, 0.12) 0%, rgba(34, 197, 94, 0.04) 50%, transparent 70%);
  pointer-events: none;
  filter: blur(40px);
}

.login-card {
  width: 100%;
  max-width: 460px;
  padding: 44px 38px;
  background: rgba(17, 17, 19, 0.88);
  backdrop-filter: blur(24px);
  -webkit-backdrop-filter: blur(24px);
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: var(--radius-md);
  box-shadow: 0 30px 80px -10px rgba(0, 0, 0, 0.8),
              0 0 0 1px rgba(255, 255, 255, 0.05);
  position: relative;
  z-index: 10;
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.logo-box {
  width: 58px;
  height: 58px;
  margin: 0 auto 18px;
  background: linear-gradient(135deg, rgba(0, 242, 254, 0.15), rgba(34, 197, 94, 0.1));
  border: 1px solid rgba(0, 242, 254, 0.35);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  box-shadow: 0 8px 24px rgba(0, 242, 254, 0.15);
}

.logo-char {
  font-size: 2rem;
  font-weight: 800;
  color: #00f2fe;
  letter-spacing: -0.02em;
}

.logo-pulse {
  position: absolute;
  width: 100%;
  height: 100%;
  border-radius: 16px;
  border: 1px solid rgba(0, 242, 254, 0.4);
  animation: pulseAura 3s infinite ease-out;
}

@keyframes pulseAura {
  0% { transform: scale(1); opacity: 0.8; }
  100% { transform: scale(1.3); opacity: 0; }
}

.login-title {
  font-size: 1.75rem;
  font-weight: 800;
  color: #ffffff;
  margin-bottom: 6px;
  letter-spacing: -0.03em;
}

.login-subtitle {
  font-size: 0.74rem;
  color: #a1a1aa;
  letter-spacing: 0.12em;
  font-weight: 600;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-size: 0.72rem;
  font-weight: 700;
  color: #a1a1aa;
  letter-spacing: 0.08em;
}

.input-wrap {
  position: relative;
  display: flex;
  align-items: center;
}

.input-icon {
  position: absolute;
  left: 14px;
  font-size: 0.95rem;
  color: #71717a;
  pointer-events: none;
}

.form-input {
  width: 100%;
  padding: 13px 16px 13px 44px;
  background: #09090b;
  border: 1px solid rgba(255, 255, 255, 0.14);
  border-radius: var(--radius-sm);
  color: #ffffff;
  font-size: 0.92rem;
  transition: all 0.2s ease;
}

.form-input:focus {
  outline: none;
  border-color: #00f2fe;
  background: #0d0d10;
  box-shadow: 0 0 0 3px rgba(0, 242, 254, 0.15);
}

.eye-btn {
  position: absolute;
  right: 12px;
  background: transparent;
  border: none;
  color: #a1a1aa;
  cursor: pointer;
  padding: 4px;
  font-size: 1rem;
}

.credential-helper {
  background: rgba(0, 242, 254, 0.04);
  border: 1px dashed rgba(0, 242, 254, 0.25);
  border-radius: var(--radius-sm);
  padding: 12px 14px;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.credential-helper:hover {
  background: rgba(0, 242, 254, 0.08);
  border-color: rgba(0, 242, 254, 0.5);
  transform: translateY(-1px);
}

.helper-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.helper-label {
  font-size: 0.65rem;
  color: #71717a;
  font-weight: 700;
  letter-spacing: 0.06em;
}

.helper-action {
  font-size: 0.65rem;
  color: #00f2fe;
  font-weight: 700;
}

.helper-values {
  font-size: 0.74rem;
  color: #e4e4e7;
}

.helper-values code {
  color: #00f2fe;
  background: rgba(0, 242, 254, 0.1);
  padding: 2px 6px;
  border-radius: 4px;
}

.btn-submit {
  width: 100%;
  padding: 14px;
  background: #ffffff;
  color: #000000;
  font-weight: 700;
  font-size: 0.88rem;
  letter-spacing: 0.04em;
  border-radius: var(--radius-sm);
  border: none;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.16, 1, 0.3, 1);
  margin-top: 4px;
}

.btn-submit:hover:not(:disabled) {
  background: #00f2fe;
  box-shadow: 0 10px 25px rgba(0, 242, 254, 0.35);
  transform: translateY(-2px);
}

.btn-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.error-box {
  background: rgba(239, 68, 68, 0.12);
  border: 1px solid rgba(239, 68, 68, 0.35);
  color: #fca5a5;
  padding: 10px 14px;
  border-radius: var(--radius-sm);
  font-size: 0.78rem;
  margin-bottom: 16px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.login-footer {
  text-align: center;
  margin-top: 28px;
  padding-top: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
}

.back-link {
  color: #71717a;
  text-decoration: none;
  font-size: 0.78rem;
  transition: color 0.2s ease;
}

.back-link:hover {
  color: #ffffff;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>
