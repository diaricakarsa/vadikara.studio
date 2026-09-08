<template>
  <canvas ref="canvasRef" class="particle-canvas" aria-hidden="true"></canvas>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const canvasRef = ref(null)
let ctx = null
let animationFrameId = null
let width = 0
let height = 0

const particles = []
const particleCount = 55
const connectionDistance = 140
const mouse = { x: null, y: null, radius: 160 }

class Particle {
  constructor(w, h) {
    this.x = Math.random() * w
    this.y = Math.random() * h
    this.vx = (Math.random() - 0.5) * 0.6
    this.vy = (Math.random() - 0.5) * 0.6
    this.size = Math.random() * 2.2 + 1.2
    
    // Natural tones: Spring Green, Woodland Sage, or Warm Sunlit Amber motes
    const rand = Math.random()
    if (rand > 0.6) {
      this.baseColor = 'rgba(22, 163, 74,' // Leaf green
    } else if (rand > 0.25) {
      this.baseColor = 'rgba(45, 106, 79,' // Forest sage
    } else {
      this.baseColor = 'rgba(217, 119, 6,'  // Sunlit warmth
    }
  }

  update(w, h) {
    this.x += this.vx
    this.y += this.vy

    if (this.x < 0 || this.x > w) this.vx *= -1
    if (this.y < 0 || this.y > h) this.vy *= -1

    // Gentle Mouse repulsion
    if (mouse.x !== null && mouse.y !== null) {
      const dx = mouse.x - this.x
      const dy = mouse.y - this.y
      const dist = Math.sqrt(dx * dx + dy * dy)
      if (dist < mouse.radius) {
        const force = (mouse.radius - dist) / mouse.radius
        this.x -= (dx / dist) * force * 2
        this.y -= (dy / dist) * force * 2
      }
    }
  }

  draw(ctx) {
    ctx.beginPath()
    ctx.arc(this.x, this.y, this.size, 0, Math.PI * 2)
    ctx.fillStyle = this.baseColor + ' 0.5)'
    ctx.shadowBlur = 6
    ctx.shadowColor = 'rgba(22, 163, 74, 0.25)'
    ctx.fill()
    ctx.shadowBlur = 0
  }
}

const handleResize = () => {
  if (!canvasRef.value) return
  const dpr = window.devicePixelRatio || 1
  width = window.innerWidth
  height = window.innerHeight
  canvasRef.value.width = width * dpr
  canvasRef.value.height = height * dpr
  canvasRef.value.style.width = `${width}px`
  canvasRef.value.style.height = `${height}px`
  ctx = canvasRef.value.getContext('2d')
  ctx.scale(dpr, dpr)
}

const handleMouseMove = (e) => {
  mouse.x = e.clientX
  mouse.y = e.clientY
}

const handleMouseLeave = () => {
  mouse.x = null
  mouse.y = null
}

const animate = () => {
  if (!ctx) return
  ctx.clearRect(0, 0, width, height)

  // Draw delicate organic connecting lines
  for (let i = 0; i < particles.length; i++) {
    particles[i].update(width, height)
    particles[i].draw(ctx)

    for (let j = i + 1; j < particles.length; j++) {
      const dx = particles[i].x - particles[j].x
      const dy = particles[i].y - particles[j].y
      const dist = Math.sqrt(dx * dx + dy * dy)

      if (dist < connectionDistance) {
        const opacity = (1 - dist / connectionDistance) * 0.22
        ctx.beginPath()
        ctx.moveTo(particles[i].x, particles[i].y)
        ctx.lineTo(particles[j].x, particles[j].y)
        ctx.strokeStyle = `rgba(34, 110, 70, ${opacity})`
        ctx.lineWidth = 0.65
        ctx.stroke()
      }
    }
  }

  animationFrameId = requestAnimationFrame(animate)
}

onMounted(() => {
  handleResize()
  for (let i = 0; i < particleCount; i++) {
    particles.push(new Particle(width, height))
  }
  window.addEventListener('resize', handleResize)
  window.addEventListener('mousemove', handleMouseMove)
  window.addEventListener('mouseleave', handleMouseLeave)
  animate()
})

onUnmounted(() => {
  if (animationFrameId) cancelAnimationFrame(animationFrameId)
  window.removeEventListener('resize', handleResize)
  window.removeEventListener('mousemove', handleMouseMove)
  window.removeEventListener('mouseleave', handleMouseLeave)
})
</script>

<style scoped>
.particle-canvas {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  pointer-events: none;
  z-index: 0;
  opacity: 0.75;
}
</style>
