<template>
  <div class="floating-player-dock" v-if="isVisible && content.ambient?.enabled !== false">
    <!-- Hidden HTML5 Audio Element for custom audio files or streams -->
    <audio 
      v-if="hasCustomAudio" 
      ref="htmlAudioRef" 
      :src="content.ambient.audio_url" 
      loop 
      preload="auto"
      @play="isPlaying = true"
      @pause="isPlaying = false"
      @ended="isPlaying = false"
    ></audio>

    <button class="player-toggle-btn" @click="togglePlay" :aria-label="isPlaying ? 'Jeda Audio' : 'Putar Audio'">
      <span v-if="!isPlaying" class="play-icon">&#9658;</span>
      <span v-else class="pause-icon">&#10074;&#10074;</span>
    </button>

    <div class="player-track-info">
      <div class="track-thumb">
        <img src="../assets/fluid_droplets.jpg" alt="Art Tetesan" class="thumb-img" />
      </div>
      <div class="track-text">
        <span class="track-title mono">{{ content.ambient?.track_title || '01 // RESONANSI CAIRAN' }}</span>
        <span class="track-artist mono">{{ content.ambient?.track_subtitle || 'Sintesis Harmonik 432 Hz' }}</span>
      </div>
    </div>

    <!-- Live Audio Spectrum Visualizer Bars -->
    <div class="player-spectrum" :class="{ 'spectrum-active': isPlaying }">
      <span class="spec-bar b1"></span>
      <span class="spec-bar b2"></span>
      <span class="spec-bar b3"></span>
      <span class="spec-bar b4"></span>
      <span class="spec-bar b5"></span>
    </div>

    <!-- Volume / Sound generator toggle -->
    <button class="player-vol-btn" @click="toggleMute" :title="isMuted ? 'Bunyikan Suara' : 'Bisukan Suara'">
      <span v-if="!isMuted">🔊</span>
      <span v-else>🔇</span>
    </button>

    <!-- Close button (Scene.mp4 00:07) -->
    <button class="player-close-btn" @click="isVisible = false" title="Tutup Pemutar">
      &times;
    </button>
  </div>
</template>

<script setup>
import { ref, computed, watch, onUnmounted } from 'vue'
import { useContent } from '../composables/useContent'

const { content } = useContent()
const isVisible = ref(true)
const isPlaying = ref(false)
const isMuted = ref(false)
const htmlAudioRef = ref(null)

const hasCustomAudio = computed(() => {
  return Boolean(content.ambient?.audio_url && content.ambient.audio_url.trim().length > 0)
})

// Web Audio API ambient drone synthesizer (Pure browser native fallback)
let audioCtx = null
let oscillator = null
let gainNode = null

const initSynthAudio = () => {
  if (audioCtx) return
  try {
    const AudioContext = window.AudioContext || window.webkitAudioContext
    audioCtx = new AudioContext()
    
    // Ambient warm analog 432Hz harmonic drone
    oscillator = audioCtx.createOscillator()
    gainNode = audioCtx.createGain()
    
    oscillator.type = 'sine'
    oscillator.frequency.setValueAtTime(108, audioCtx.currentTime) // 432 Hz / 4 harmonic sub-octave
    gainNode.gain.setValueAtTime(0.04, audioCtx.currentTime)
    
    oscillator.connect(gainNode)
    gainNode.connect(audioCtx.destination)
    oscillator.start()
  } catch (e) {
    // Graceful fallback
  }
}

const togglePlay = () => {
  if (hasCustomAudio.value && htmlAudioRef.value) {
    // Custom audio file mode
    if (isPlaying.value) {
      htmlAudioRef.value.pause()
      isPlaying.value = false
    } else {
      htmlAudioRef.value.muted = isMuted.value
      htmlAudioRef.value.play().then(() => {
        isPlaying.value = true
      }).catch(err => {
        console.warn('Playback error:', err)
      })
    }
  } else {
    // Generative synthesizer fallback mode
    if (!audioCtx) initSynthAudio()
    
    if (audioCtx && audioCtx.state === 'suspended') {
      audioCtx.resume()
    }

    isPlaying.value = !isPlaying.value

    if (gainNode && audioCtx) {
      if (isPlaying.value && !isMuted.value) {
        gainNode.gain.setTargetAtTime(0.04, audioCtx.currentTime, 0.2)
      } else {
        gainNode.gain.setTargetAtTime(0, audioCtx.currentTime, 0.2)
      }
    }
  }
}

const toggleMute = () => {
  isMuted.value = !isMuted.value
  
  if (hasCustomAudio.value && htmlAudioRef.value) {
    htmlAudioRef.value.muted = isMuted.value
  } else if (gainNode && audioCtx) {
    gainNode.gain.setTargetAtTime(isMuted.value ? 0 : 0.04, audioCtx.currentTime, 0.1)
  }
}

// Watch for audio URL changes from CMS
watch(() => content.ambient?.audio_url, (newUrl) => {
  if (isPlaying.value) {
    if (htmlAudioRef.value) {
      htmlAudioRef.value.pause()
    }
    if (gainNode && audioCtx) {
      gainNode.gain.setTargetAtTime(0, audioCtx.currentTime, 0.1)
    }
    isPlaying.value = false
  }
})

onUnmounted(() => {
  if (htmlAudioRef.value) {
    htmlAudioRef.value.pause()
  }
  if (oscillator) {
    try { oscillator.stop() } catch (e) {}
  }
})
</script>

<style scoped>
.floating-player-dock {
  position: fixed;
  bottom: 24px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 1000;
  background: rgba(14, 14, 16, 0.92);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.18);
  border-radius: var(--radius-full);
  padding: 8px 16px 8px 10px;
  display: flex;
  align-items: center;
  gap: 14px;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.7);
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  user-select: none;
}

.floating-player-dock:hover {
  border-color: rgba(255, 255, 255, 0.35);
  box-shadow: 0 25px 60px rgba(0, 0, 0, 0.9);
}

.player-toggle-btn {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: #ffffff;
  border: none;
  color: #000000;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.2s;
  flex-shrink: 0;
}

.player-toggle-btn:hover {
  transform: scale(1.1);
}

.play-icon {
  font-size: 0.85rem;
  margin-left: 2px;
}

.pause-icon {
  font-size: 0.75rem;
}

.player-track-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.track-thumb {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  overflow: hidden;
  background: #000;
  flex-shrink: 0;
}

.thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.track-text {
  display: flex;
  flex-direction: column;
}

.track-title {
  font-size: 0.72rem;
  font-weight: 700;
  color: #ffffff;
  line-height: 1.2;
}

.track-artist {
  font-size: 0.6rem;
  color: #71717a;
}

/* Spectrum Bars */
.player-spectrum {
  display: flex;
  align-items: flex-end;
  gap: 3px;
  height: 16px;
  padding: 0 4px;
}

.spec-bar {
  width: 2.5px;
  background: #52525b;
  border-radius: 2px;
  height: 4px;
  transition: height 0.2s ease, background-color 0.2s ease;
}

.spectrum-active .spec-bar {
  background: #ffffff;
  animation: bar-bounce 0.8s infinite ease-in-out alternate;
}

.spectrum-active .b1 { animation-delay: 0.1s; }
.spectrum-active .b2 { animation-delay: 0.3s; }
.spectrum-active .b3 { animation-delay: 0.5s; }
.spectrum-active .b4 { animation-delay: 0.2s; }
.spectrum-active .b5 { animation-delay: 0.4s; }

@keyframes bar-bounce {
  0% { height: 4px; }
  100% { height: 16px; }
}

.player-vol-btn, .player-close-btn {
  background: transparent;
  border: none;
  color: #a1a1aa;
  cursor: pointer;
  font-size: 0.9rem;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 4px;
  transition: color 0.2s;
}

.player-vol-btn:hover, .player-close-btn:hover {
  color: #ffffff;
}

.player-close-btn {
  font-size: 1.2rem;
  line-height: 1;
  margin-left: 2px;
}

@media (max-width: 640px) {
  .floating-player-dock {
    bottom: 16px;
    padding: 6px 12px;
    gap: 10px;
  }
  .track-artist {
    display: none;
  }
}
</style>
