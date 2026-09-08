<template>
  <header class="navbar-wrapper reveal-init">
    <div class="navbar-inner">
      <router-link to="/" class="brand-logo">
        <span class="brand-name">{{ content.navbar?.brand_name || 'Vadikara' }}<sup v-if="content.navbar?.brand_sup">{{ content.navbar.brand_sup }}</sup></span>
      </router-link>

      <nav class="nav-links" :class="{ 'nav-open': isMobileMenuOpen }">
        <a 
          v-for="(link, idx) in navLinks" 
          :key="idx"
          :href="link.url" 
          class="nav-link" 
          @click="closeMenu"
        >
          {{ link.label }}
        </a>
      </nav>

      <div class="nav-actions">
        <a 
          v-if="content.navbar?.action_url" 
          :href="content.navbar.action_url" 
          target="_blank" 
          rel="noopener" 
          class="pill-btn nav-pill-btn"
        >
          <span>{{ content.navbar?.action_label || 'GitHub' }}</span>
          <span class="arrow-up-right">&nearr;</span>
        </a>
        <button class="mobile-toggle" @click="toggleMenu" aria-label="Buka Menu Navigasi">
          <span :class="{ 'open': isMobileMenuOpen }"></span>
        </button>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useContent } from '../composables/useContent'

const { content } = useContent()
const isMobileMenuOpen = ref(false)

const navLinks = computed(() => {
  if (Array.isArray(content.navbar?.links) && content.navbar.links.length > 0) {
    return content.navbar.links
  }
  return [
    { label: 'Beranda', url: '#hero' },
    { label: 'Karya', url: '#work' },
    { label: 'Keahlian', url: '#craft' },
    { label: 'Tentang', url: '#manifesto' },
    { label: 'Labs', url: '#labs' },
    { label: 'Kontak', url: '#contact' }
  ]
})

const toggleMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

const closeMenu = () => {
  isMobileMenuOpen.value = false
}
</script>

<style scoped>
.navbar-wrapper {
  width: 100%;
  padding: 28px 40px 10px;
  background: transparent;
  z-index: 100;
  position: relative;
}

.navbar-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.brand-logo {
  text-decoration: none;
  color: #ffffff;
}

.brand-name {
  font-size: 1.55rem;
  font-weight: 800;
  letter-spacing: -0.03em;
  color: #ffffff;
}

.brand-name sup {
  font-size: 0.65rem;
  margin-left: 2px;
  font-weight: 600;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 28px;
}

.nav-link {
  color: #a1a1aa;
  text-decoration: none;
  font-size: 0.9rem;
  font-weight: 500;
  transition: color 0.2s ease;
  position: relative;
}

.nav-link:hover {
  color: #ffffff;
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.nav-pill-btn {
  padding: 8px 18px;
  font-size: 0.78rem;
}

.arrow-up-right {
  font-size: 0.9rem;
  transition: transform 0.2s ease;
}

.nav-pill-btn:hover .arrow-up-right {
  transform: translate(2px, -2px);
}

.mobile-toggle {
  display: none;
  background: transparent;
  border: none;
  cursor: pointer;
  padding: 6px;
}

.mobile-toggle span {
  display: block;
  width: 22px;
  height: 2px;
  background: #ffffff;
  position: relative;
  transition: all 0.3s;
}

.mobile-toggle span::before,
.mobile-toggle span::after {
  content: '';
  position: absolute;
  width: 22px;
  height: 2px;
  background: #ffffff;
  transition: all 0.3s;
}

.mobile-toggle span::before { top: -6px; }
.mobile-toggle span::after { bottom: -6px; }

.mobile-toggle span.open { background: transparent; }
.mobile-toggle span.open::before { transform: rotate(45deg); top: 0; }
.mobile-toggle span.open::after { transform: rotate(-45deg); bottom: 0; }

@media (max-width: 860px) {
  .navbar-wrapper {
    padding: 20px 24px 10px;
  }
  .mobile-toggle {
    display: block;
  }
  .nav-links {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: #09090b;
    flex-direction: column;
    justify-content: center;
    padding: 40px;
    gap: 26px;
    z-index: 999;
    transform: translateY(-100%);
    transition: transform 0.35s cubic-bezier(0.16, 1, 0.3, 1);
  }
  .nav-links.nav-open {
    transform: translateY(0);
  }
  .nav-link {
    font-size: 1.5rem;
    font-weight: 700;
  }
}
</style>
