<template>
  <div class="app-root">
    <!-- Public App Shell with Studio Size Dark Master Canvas (Scene.mp4 Benchmark) -->
    <div class="studio-canvas" v-if="!isAdminRoute">
      <!-- Minimalist Header -->
      <Navbar />

      <!-- Main Content Outlet -->
      <main class="main-content">
        <router-view />
      </main>

      <!-- Minimalist Studio Footer -->
      <Footer />

      <!-- Floating Interactive Ambient Sound Player (Scene.mp4 Frame 00:07) -->
      <AmbientPlayer />
    </div>

    <!-- Admin / Backoffice focused wrapper -->
    <div v-else class="admin-wrapper">
      <router-view />
    </div>
  </div>
</template>

<script setup>
import { computed, watchEffect } from 'vue'
import { useRoute } from 'vue-router'
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'
import AmbientPlayer from './components/AmbientPlayer.vue'

const route = useRoute()

const isAdminRoute = computed(() => {
  return route.path.startsWith('/admin')
})

watchEffect(() => {
  if (typeof document !== 'undefined') {
    if (isAdminRoute.value) {
      document.body.classList.add('is-admin-mode')
    } else {
      document.body.classList.remove('is-admin-mode')
    }
  }
})
</script>

<style scoped>
.app-root {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.main-content {
  flex: 1;
  position: relative;
  z-index: 1;
}

.admin-wrapper {
  width: 100%;
  min-height: 100vh;
}
</style>
