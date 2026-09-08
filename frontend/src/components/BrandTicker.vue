<template>
  <div class="brand-ticker-section reveal-init">
    <div class="ticker-inner">
      <div class="ticker-track anim-marquee">
        <!-- First loop -->
        <template v-for="(item, idx) in tickerItems" :key="'a-' + idx">
          <span class="ticker-item font-hero" :class="{ 'text-dim': idx % 2 === 1 }">{{ item }}</span>
          <span class="ticker-dot">•</span>
        </template>
        <!-- Duplicate loop for seamless infinite scroll -->
        <template v-for="(item, idx) in tickerItems" :key="'b-' + idx">
          <span class="ticker-item font-hero" :class="{ 'text-dim': idx % 2 === 1 }">{{ item }}</span>
          <span class="ticker-dot">•</span>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useContent } from '../composables/useContent'

const { content } = useContent()

const tickerItems = computed(() => {
  if (Array.isArray(content.ticker) && content.ticker.length > 0) {
    return content.ticker
  }
  return [
    'vadikara',
    'vadikara.dev',
    'rekayasa kreatif',
    'sistem terdistribusi',
    'antarmuka 60 fps',
    'inti golang',
    'arsitektur awan'
  ]
})
</script>

<style scoped>
.brand-ticker-section {
  padding: 60px 0 40px;
  overflow: hidden;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  position: relative;
}

.ticker-inner {
  display: flex;
  overflow: hidden;
  user-select: none;
}

.ticker-track {
  display: flex;
  align-items: center;
  gap: 32px;
  white-space: nowrap;
}

.ticker-item {
  font-size: clamp(3rem, 7vw, 6.2rem);
  font-weight: 800;
  letter-spacing: -0.04em;
  color: #ffffff;
  line-height: 1;
  transition: color 0.3s ease;
}

.ticker-item.text-dim {
  color: #3f3f46;
}

.ticker-item:hover {
  color: #ffffff !important;
}

.ticker-dot {
  font-size: 2.2rem;
  color: #27272a;
}
</style>
