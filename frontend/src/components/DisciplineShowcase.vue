<template>
  <section class="discipline-section" id="craft">
    <div class="studio-container">
      <div class="discipline-grid">
        <!-- Left: Dynamic Visual Showcase Tile (Scene.mp4 Frame 00:03) -->
        <div class="showcase-visual-col reveal-scale-init">
          <div class="showcase-tile-card">
            <div class="tile-image-wrapper">
              <img 
                :src="activeItem.resolvedImage" 
                :alt="activeItem.title" 
                class="tile-image"
              />
              <div class="tile-overlay-vignette"></div>
            </div>

            <!-- Dynamic Telemetry & Role Overlay -->
            <div class="tile-meta-info">
              <div class="meta-top">
                <span class="meta-tag mono">KEAHLIAN // 0{{ activeIndex + 1 }}</span>
                <span class="meta-pill mono">{{ activeItem.benchmark }}</span>
              </div>
              <h3 class="meta-title">{{ activeItem.title }}</h3>
              <p class="meta-desc">{{ activeItem.description }}</p>

              <div class="meta-stack-row">
                <span 
                  v-for="(t, i) in activeItem.stack" 
                  :key="i"
                  class="stack-chip mono"
                >
                  {{ t }}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Right: Huge Vertical Interactive Typography List (Scene.mp4 Frame 00:03) -->
        <div class="showcase-list-col">
          <div 
            class="typography-stack-list"
            @mouseleave="onMouseLeaveList"
          >
            <div 
              v-for="(item, idx) in disciplines" 
              :key="idx"
              class="type-item-row reveal-init"
              :class="[
                `delay-${Math.min(idx + 1, 6)}`,
                { 
                  'is-active': activeIndex === idx,
                  'is-dimmed': hoveredIndex !== null && hoveredIndex !== idx 
                }
              ]"
              @mouseenter="onHoverItem(idx)"
              @click="activeIndex = idx"
            >
              <span class="type-item-text">{{ item.title }}</span>
              <span class="type-item-arrow" v-if="activeIndex === idx">&rarr;</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useContent } from '../composables/useContent'

const { content, resolveImage } = useContent()

const activeIndex = ref(3)
const hoveredIndex = ref(null)

const disciplines = computed(() => {
  return content.disciplines || []
})

const activeItem = computed(() => {
  const list = disciplines.value
  if (!list.length) {
    return {
      title: 'Rekayasa Sistem',
      benchmark: '60 FPS',
      description: 'Sistem komputasi berdaya tahan tinggi.',
      stack: ['Golang', 'Vue 3'],
      resolvedImage: resolveImage('')
    }
  }
  const idx = Math.min(Math.max(0, activeIndex.value), list.length - 1)
  const item = list[idx] || list[0]
  return {
    ...item,
    resolvedImage: resolveImage(item.image)
  }
})

const onHoverItem = (idx) => {
  hoveredIndex.value = idx
  activeIndex.value = idx
}

const onMouseLeaveList = () => {
  hoveredIndex.value = null
}
</script>

<style scoped>
.discipline-section {
  padding: 80px 0 100px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  position: relative;
}

.discipline-grid {
  display: grid;
  grid-template-columns: 1fr 1.25fr;
  gap: 50px;
  align-items: center;
}

/* Left Visual Tile */
.showcase-tile-card {
  background: #121215;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-lg);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.6);
  transition: all 0.3s ease;
}

.tile-image-wrapper {
  position: relative;
  width: 100%;
  height: 380px;
  background: #000000;
  overflow: hidden;
}

.tile-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1), opacity 0.3s ease;
}

.showcase-tile-card:hover .tile-image {
  transform: scale(1.04);
}

.tile-overlay-vignette {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(180deg, transparent 40%, rgba(18, 18, 21, 0.95) 100%);
  pointer-events: none;
}

.tile-meta-info {
  padding: 24px 28px 28px;
  display: flex;
  flex-direction: column;
}

.meta-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.meta-tag {
  font-size: 0.72rem;
  color: #71717a;
  letter-spacing: 0.08em;
}

.meta-pill {
  font-size: 0.72rem;
  background: rgba(255, 255, 255, 0.08);
  color: #ffffff;
  padding: 2px 8px;
  border-radius: 4px;
}

.meta-title {
  font-size: 1.5rem;
  font-weight: 800;
  letter-spacing: -0.02em;
  color: #ffffff;
  margin-bottom: 8px;
}

.meta-desc {
  font-size: 0.88rem;
  color: #a1a1aa;
  line-height: 1.6;
  margin-bottom: 18px;
}

.meta-stack-row {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.stack-chip {
  font-size: 0.72rem;
  background: #1c1c20;
  border: 1px solid rgba(255, 255, 255, 0.08);
  color: #d4d4d8;
  padding: 3px 8px;
  border-radius: 4px;
}

/* Right Typography Stack List (Scene.mp4 00:03) */
.typography-stack-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.type-item-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
  padding: 8px 0;
  transition: all 0.25s cubic-bezier(0.16, 1, 0.3, 1);
}

.type-item-text {
  font-size: clamp(1.8rem, 3.2vw, 3rem);
  font-weight: 800;
  letter-spacing: -0.03em;
  line-height: 1.15;
  color: #ffffff;
  transition: color 0.2s ease, transform 0.2s ease;
}

.type-item-arrow {
  font-size: 2rem;
  color: #ffffff;
  margin-left: 16px;
  animation: float-horizontal 1.5s infinite ease-in-out;
}

@keyframes float-horizontal {
  0%, 100% { transform: translateX(0); }
  50% { transform: translateX(6px); }
}

/* Dimming states matching Scene.mp4 Frame 00:03 */
.type-item-row.is-dimmed .type-item-text {
  color: rgba(255, 255, 255, 0.18);
}

.type-item-row.is-active .type-item-text {
  color: #ffffff;
  transform: translateX(8px);
}

.type-item-row:hover .type-item-text {
  color: #ffffff !important;
  transform: translateX(8px);
}

@media (max-width: 960px) {
  .discipline-grid {
    grid-template-columns: 1fr;
    gap: 40px;
  }
  .tile-image-wrapper {
    height: 280px;
  }
}
</style>
