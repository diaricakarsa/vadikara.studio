<template>
  <footer class="footer-wrapper reveal-init">
    <div class="studio-container">
      <div class="footer-inner">
        <div class="footer-top-row">
          <div class="footer-brand-col">
            <span class="footer-logo">{{ content.footer?.logo_text || 'Vadikara' }}<sup v-if="content.footer?.logo_sup">{{ content.footer.logo_sup }}</sup></span>
            <p class="footer-tagline">
              {{ content.footer?.tagline || 'Rekayasawan sistem kreatif & pengembang arsitektur terdistribusi. Merancang backend berkinerja tinggi serta antarmuka web 60 FPS yang berjiwa.' }}
            </p>
          </div>

          <div class="footer-nav-col">
            <span class="col-heading mono">INDEKS</span>
            <ul class="footer-links">
              <li v-for="(l, idx) in footerIndexLinks" :key="idx">
                <a :href="l.url">{{ l.label }}</a>
              </li>
            </ul>
          </div>

          <div class="footer-nav-col">
            <span class="col-heading mono">JEJARING</span>
            <ul class="footer-links">
              <li v-for="(s, idx) in footerSocialLinks" :key="idx">
                <a :href="s.url" target="_blank" rel="noopener">{{ s.label }}</a>
              </li>
            </ul>
          </div>
        </div>

        <div class="footer-bottom-row mono">
          <span class="copy-text">{{ content.footer?.copyright_text || '© 2026 Vadikara. Seluruh hak cipta dilindungi.' }}</span>
          <span class="built-text">{{ content.footer?.built_text || 'Dirancang & Dibangun dengan Go 1.26 & Vue 3 • 60 FPS' }}</span>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { computed } from 'vue'
import { useContent } from '../composables/useContent'

const { content } = useContent()

const footerIndexLinks = computed(() => {
  if (Array.isArray(content.footer?.index_links) && content.footer.index_links.length > 0) {
    return content.footer.index_links
  }
  return [
    { label: 'Beranda', url: '#hero' },
    { label: 'Karya Terpilih', url: '#work' },
    { label: 'Keahlian & Disiplin', url: '#craft' },
    { label: 'Manifesto', url: '#manifesto' },
    { label: 'Vadikara Labs', url: '#labs' },
    { label: 'Kontak', url: '#contact' }
  ]
})

const footerSocialLinks = computed(() => {
  if (Array.isArray(content.footer?.social_links) && content.footer.social_links.length > 0) {
    return content.footer.social_links
  }
  return [
    { label: 'GitHub ↗', url: 'https://github.com/vadikara' },
    { label: 'X // Twitter ↗', url: 'https://x.com/vadikara' },
    { label: 'LinkedIn ↗', url: 'https://linkedin.com' }
  ]
})
</script>

<style scoped>
.footer-wrapper {
  padding: 60px 0 40px;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
}

.footer-top-row {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr;
  gap: 40px;
  padding-bottom: 50px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
}

.footer-logo {
  font-size: 1.8rem;
  font-weight: 800;
  letter-spacing: -0.03em;
  color: #ffffff;
  display: block;
  margin-bottom: 12px;
}

.footer-logo sup {
  font-size: 0.55rem;
  margin-left: 2px;
}

.footer-tagline {
  font-size: 0.92rem;
  color: #71717a;
  max-width: 380px;
  line-height: 1.6;
}

.col-heading {
  font-size: 0.7rem;
  color: #52525b;
  letter-spacing: 0.08em;
  display: block;
  margin-bottom: 16px;
}

.footer-links {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.footer-links a {
  color: #a1a1aa;
  text-decoration: none;
  font-size: 0.88rem;
  transition: color 0.2s;
}

.footer-links a:hover {
  color: #ffffff;
}

.text-dim {
  color: #52525b !important;
}

.footer-bottom-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 30px;
  font-size: 0.75rem;
  color: #52525b;
}

@media (max-width: 768px) {
  .footer-top-row {
    grid-template-columns: 1fr;
    gap: 30px;
  }
  .footer-bottom-row {
    flex-direction: column;
    gap: 10px;
    align-items: flex-start;
  }
}
</style>
