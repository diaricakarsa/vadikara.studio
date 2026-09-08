<template>
  <section class="hero-section" id="hero">
    <div class="hero-header-content reveal-init">
      <h1 class="hero-main-title" v-html="formattedHeroTitle"></h1>
    </div>

    <!-- Organic Fluid Droplets Interactive Stage (Scene.mp4 00:00 & 00:08) -->
    <div 
      class="fluid-droplet-stage reveal-scale-init delay-1"
      @mousemove="handleStageMouseMove($event)"
      @mouseleave="handleStageMouseLeave"
    >
      <div class="fluid-media-container" :style="stageTransform">
        <img 
          :src="heroBgImage" 
          alt="Vadikara Showcase Stage" 
          class="fluid-image"
        />
        <div class="fluid-overlay-glare"></div>

        <!-- Floating Interactive Action Badge (PLAY REEL ↗ from Scene.mp4 00:00) -->
        <div class="floating-reel-anchor reveal-init delay-2">
          <button class="play-reel-btn" @click="toggleReelModal" aria-label="Putar Showcase Reel">
            <span class="reel-text mono">JELAJAHI REEL</span>
            <span class="reel-icon">&nearr;</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Interactive Video/Reel Modal -->
    <transition name="fade">
      <div v-if="isReelModalOpen" class="reel-modal-backdrop" @click.self="toggleReelModal">
        <div class="reel-modal-card">
          <div class="reel-modal-header">
            <span class="mono reel-label">{{ content.hero.reel_badge || 'VADIKARA // SHOWCASE REKAYASA SISTEM 2026' }}</span>
            <button class="reel-close" @click="toggleReelModal" aria-label="Tutup">&times;</button>
          </div>
          <div class="reel-video-wrapper">
            <iframe 
              v-if="videoMeta.isYoutube || videoMeta.isVimeo"
              :src="videoMeta.embedUrl" 
              title="Vadikara Showcase Reel"
              class="modal-video-frame"
              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
              allowfullscreen
            ></iframe>
            <video 
              v-else
              :src="videoMeta.directUrl || '/Scene.mp4'" 
              autoplay 
              loop 
              controls 
              class="modal-video"
            ></video>
          </div>
          <div class="reel-footer mono">
            <span>{{ content.hero.reel_telemetry || 'Konkurensi Tinggi Golang • Web Kreatif 60 FPS • Sistem Terdistribusi Sub-milidetik' }}</span>
          </div>
        </div>
      </div>
    </transition>
  </section>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useContent } from '../composables/useContent'
import { parseVideoUrl } from '../utils/media'

const { content, resolveImage } = useContent()

const heroBgImage = computed(() => resolveImage(content.hero.background_image))
const videoMeta = computed(() => parseVideoUrl(content.hero.reel_url))

const formattedHeroTitle = computed(() => {
  const t = content.hero.title || 'Rekayasawan sistem kreatif &\nperanti lunak terdistribusi'
  const html = t.replace(/\n/g, '<br />')
  return html.includes('sup') ? html : `${html}<sup>&reg;</sup>`
})

const isReelModalOpen = ref(false)
const mouseX = ref(0)
const mouseY = ref(0)
const isHovering = ref(false)

const handleStageMouseMove = (event) => {
  const stage = event.currentTarget
  const rect = stage.getBoundingClientRect()
  const x = (event.clientX - rect.left - rect.width / 2) / (rect.width / 2)
  const y = (event.clientY - rect.top - rect.height / 2) / (rect.height / 2)
  mouseX.value = x
  mouseY.value = y
  isHovering.value = true
}

const handleStageMouseLeave = () => {
  mouseX.value = 0
  mouseY.value = 0
  isHovering.value = false
}

const stageTransform = computed(() => {
  if (!isHovering.value) return 'transform: perspective(1000px) rotateX(0deg) rotateY(0deg) scale(1)'
  const rotX = -mouseY.value * 4
  const rotY = mouseX.value * 4
  return `transform: perspective(1000px) rotateX(${rotX.toFixed(2)}deg) rotateY(${rotY.toFixed(2)}deg) scale(1.01)`
})

const toggleReelModal = () => {
  isReelModalOpen.value = !isReelModalOpen.value
}
</script>

<style scoped>
.hero-section {
  padding: 30px 40px 0;
  position: relative;
  display: flex;
  flex-direction: column;
}

.hero-header-content {
  margin-bottom: 36px;
}

.hero-main-title {
  font-size: clamp(2.6rem, 5.4vw, 5.2rem);
  font-weight: 800;
  letter-spacing: -0.04em;
  line-height: 1.05;
  color: #ffffff;
  margin: 0;
}

.hero-main-title sup {
  font-size: 0.35em;
  font-weight: 600;
  margin-left: 4px;
}

/* Organic Fluid Droplets Stage */
.fluid-droplet-stage {
  width: 100%;
  border-radius: var(--radius-lg);
  overflow: hidden;
  position: relative;
  background: #000000;
  cursor: pointer;
}

.fluid-media-container {
  position: relative;
  width: 100%;
  height: clamp(380px, 50vw, 560px);
  border-radius: var(--radius-lg);
  overflow: hidden;
  transition: transform 0.25s cubic-bezier(0.16, 1, 0.3, 1);
}

.fluid-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center 60%;
  display: block;
}

.fluid-overlay-glare {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 50% 30%, rgba(255, 255, 255, 0.12), transparent 70%);
  pointer-events: none;
}

/* Floating Action Badge (Scene.mp4 Frame 00:00) */
.floating-reel-anchor {
  position: absolute;
  bottom: 40px;
  right: 48px;
  z-index: 10;
}

.play-reel-btn {
  width: 130px;
  height: 130px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: #000000;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 16px 36px rgba(0, 0, 0, 0.3);
  transition: all 0.35s cubic-bezier(0.16, 1, 0.3, 1);
  gap: 4px;
}

.play-reel-btn:hover {
  transform: scale(1.1) rotate(5deg);
  background: #ffffff;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.45);
}

.reel-text {
  font-size: 0.68rem;
  font-weight: 700;
  letter-spacing: 0.08em;
}

.reel-icon {
  font-size: 1.4rem;
  font-weight: 700;
  transition: transform 0.2s ease;
}

.play-reel-btn:hover .reel-icon {
  transform: translate(3px, -3px);
}

/* Modal */
.reel-modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.85);
  backdrop-filter: blur(16px);
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}

.reel-modal-card {
  width: 100%;
  max-width: 900px;
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.15);
  border-radius: var(--radius-md);
  overflow: hidden;
  box-shadow: 0 30px 90px rgba(0, 0, 0, 0.8);
}

.reel-modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 22px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.reel-label {
  font-size: 0.76rem;
  color: #a1a1aa;
  letter-spacing: 0.06em;
}

.reel-close {
  background: transparent;
  border: none;
  font-size: 1.8rem;
  color: #ffffff;
  cursor: pointer;
  line-height: 1;
}

.reel-video-wrapper {
  width: 100%;
  aspect-ratio: 16 / 9;
  background: #000000;
}

.modal-video,
.modal-video-frame {
  width: 100%;
  height: 100%;
  border: none;
  object-fit: contain;
  display: block;
}

.reel-footer {
  padding: 14px 22px;
  font-size: 0.72rem;
  color: #71717a;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.25s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

@media (max-width: 768px) {
  .hero-section {
    padding: 20px 20px 0;
  }
  .floating-reel-anchor {
    bottom: 20px;
    right: 20px;
  }
  .play-reel-btn {
    width: 90px;
    height: 90px;
  }
  .reel-text {
    font-size: 0.58rem;
  }
}
</style>
