<template>
  <section class="labs-section" id="labs">
    <div class="studio-container">
      <div class="labs-card reveal-scale-init">
        <div class="labs-media-wrapper">
          <img 
            :src="labsBgImage" 
            alt="Vadikara Labs Algoritma Generatif" 
            class="labs-bg-image"
          />
          <div class="labs-overlay"></div>
        </div>

        <div class="labs-content reveal-init delay-1">
          <span class="labs-tag mono">{{ content.labs.tag || 'EKSPERIMEN // LABS 0.9' }}</span>
          <h2 class="labs-title">
            {{ content.labs.title || 'Vadikara Labs — arena bermain kode kreatif & sistem algoritmik.' }}
          </h2>

          <div class="labs-actions">
            <button class="manifesto-pill-btn" @click="toggleLabDrawer">
              <span>{{ isDrawerOpen ? 'Tutup Eksperimen' : 'Jelajahi Labs' }}</span>
              <span class="plus-circle">{{ isDrawerOpen ? '&times;' : '+' }}</span>
            </button>
          </div>
        </div>

        <!-- Expandable Interactive Labs Experiments Drawer -->
        <transition name="drawer-slide">
          <div v-if="isDrawerOpen" class="labs-drawer-panel">
            <div class="drawer-inner">
              <h3 class="drawer-title mono">// PROTOTIPE EKSPERIMENTAL AKTIF</h3>
              <div class="experiments-grid">
                <div 
                  v-for="(exp, i) in experiments" 
                  :key="exp.id || i"
                  class="experiment-item reveal-init"
                  :class="`delay-${i + 1}`"
                  @click="activeExperiment = exp"
                >
                  <div class="exp-top">
                    <span class="exp-id mono">0{{ i + 1 }}</span>
                    <span class="exp-status mono">{{ exp.status }}</span>
                  </div>
                  <h4 class="exp-name">{{ exp.name }}</h4>
                  <p class="exp-desc">{{ exp.desc }}</p>
                  <div class="exp-tags">
                    <span v-for="(t, idx) in exp.tags" :key="idx" class="exp-tag mono">{{ t }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </transition>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useContent } from '../composables/useContent'

const { content, resolveImage } = useContent()

const isDrawerOpen = ref(false)
const activeExperiment = ref(null)

const toggleLabDrawer = () => {
  isDrawerOpen.value = !isDrawerOpen.value
}

const labsBgImage = computed(() => resolveImage(content.labs.background_image))

const experiments = computed(() => {
  if (Array.isArray(content.labs.experiments) && content.labs.experiments.length > 0) {
    return content.labs.experiments
  }
  return [
    {
      id: 1,
      name: 'Visualisator Ring-Buffer Lock-Free',
      status: 'RUNTIME AKTIF',
      desc: 'Visualisasi interaktif perebutan thread dan keselarasan memori antrean cincin zero-alloc di bawah simulasi beban kerja sintetik.',
      tags: ['Go Assembly', 'Tata Letak Memori', 'Bebas Kunci']
    },
    {
      id: 2,
      name: 'Mesin Keadaan Konsensus Raft',
      status: 'SIMULASI KLASTER',
      desc: 'Simulasi pemilihan pemimpin klaster multi-node Raft dan replikasi log detak jantung dengan simulasi gangguan partisi jaringan acak.',
      tags: ['Konsensus Terdistribusi', 'Protokol Raft', 'Kekacauan Jaringan']
    },
    {
      id: 3,
      name: 'Kanvas Fourier Responsif Audio',
      status: 'WEBGL // SHADER',
      desc: 'Penganalisis spektrum domain frekuensi waktu nyata yang menggambar kelopak vektor generatif dan pita partikel cairan terakselerasi GPU.',
      tags: ['Web Audio API', 'GLSL Shaders', 'Canvas 60fps']
    }
  ]
})
</script>

<style scoped>
.labs-section {
  padding: 60px 0 80px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.labs-card {
  position: relative;
  background: #000000;
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: var(--radius-lg);
  overflow: hidden;
  min-height: 480px;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
}

.labs-media-wrapper {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
}

.labs-bg-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.75;
  transition: transform 0.6s ease;
}

.labs-card:hover .labs-bg-image {
  transform: scale(1.02);
}

.labs-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 70% 30%, transparent 20%, rgba(0, 0, 0, 0.85) 90%);
}

.labs-content {
  position: relative;
  z-index: 2;
  padding: 48px 48px;
  max-width: 820px;
}

.labs-tag {
  font-size: 0.72rem;
  color: #ef4444;
  letter-spacing: 0.08em;
  margin-bottom: 16px;
  display: block;
}

.labs-title {
  font-size: clamp(2rem, 3.8vw, 3.5rem);
  font-weight: 800;
  letter-spacing: -0.03em;
  line-height: 1.15;
  color: #ffffff;
  margin-bottom: 28px;
}

.manifesto-pill-btn {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  padding: 8px 10px 8px 22px;
  border-radius: var(--radius-full);
  background: rgba(255, 255, 255, 0.95);
  color: #000000;
  border: 1px solid #ffffff;
  font-size: 0.88rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.25s ease;
}

.manifesto-pill-btn:hover {
  background: #ffffff;
  transform: scale(1.04);
}

.plus-circle {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: #000000;
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.2rem;
  line-height: 1;
}

/* Expandable Drawer */
.labs-drawer-panel {
  position: relative;
  z-index: 3;
  background: #0d0d0f;
  border-top: 1px solid rgba(255, 255, 255, 0.15);
  padding: 36px 48px;
}

.drawer-title {
  font-size: 0.78rem;
  color: #71717a;
  letter-spacing: 0.08em;
  margin-bottom: 24px;
}

.experiments-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.experiment-item {
  background: #151518;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: var(--radius-sm);
  padding: 20px;
  cursor: pointer;
  transition: all 0.25s ease;
}

.experiment-item:hover {
  border-color: rgba(255, 255, 255, 0.35);
  transform: translateY(-3px);
}

.exp-top {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
}

.exp-id {
  font-size: 0.72rem;
  color: #71717a;
}

.exp-status {
  font-size: 0.65rem;
  color: #ef4444;
  background: rgba(239, 68, 68, 0.1);
  padding: 2px 6px;
  border-radius: 4px;
}

.exp-name {
  font-size: 1.05rem;
  font-weight: 700;
  color: #ffffff;
  margin-bottom: 6px;
}

.exp-desc {
  font-size: 0.82rem;
  color: #a1a1aa;
  line-height: 1.5;
  margin-bottom: 14px;
}

.exp-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.exp-tag {
  font-size: 0.68rem;
  background: #222226;
  color: #d4d4d8;
  padding: 2px 6px;
  border-radius: 3px;
}

.drawer-slide-enter-active, .drawer-slide-leave-active {
  transition: all 0.35s cubic-bezier(0.16, 1, 0.3, 1);
}

.drawer-slide-enter-from, .drawer-slide-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

@media (max-width: 900px) {
  .experiments-grid {
    grid-template-columns: 1fr;
  }
  .labs-content {
    padding: 30px 24px;
  }
  .labs-drawer-panel {
    padding: 24px;
  }
}
</style>
