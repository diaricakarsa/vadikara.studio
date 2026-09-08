<template>
  <div class="services-page">
    <div class="container">
      <div class="page-header">
        <span class="badge badge-violet mono">Solusi Enterprise & Arsitektur</span>
        <h1 class="page-title">Solusi Rekayasa <span class="text-gradient-violet">Skala Industri.</span></h1>
        <p class="page-subtitle">
          Layanan end-to-end mulai dari perancangan microservices Go terdistribusi, pembuatan web modern reaktif, orkestrasi cloud Kubernetes, hingga audit ketahanan kode sumber.
        </p>
      </div>

      <ServicesSection />
    </div>
  </div>
</template>

<script setup>
import ServicesSection from '../components/ServicesSection.vue'
</script>

<style scoped>
.services-page {
  padding: 120px 0 60px;
  position: relative;
  z-index: 1;
}

.page-header {
  text-align: center;
  max-width: 800px;
  margin: 0 auto 30px;
}

.page-title {
  font-size: clamp(2.2rem, 4vw, 3.2rem);
  font-weight: 800;
  margin: 16px 0;
  color: var(--text-main);
}

.page-subtitle {
  font-size: 1.05rem;
  color: var(--text-muted);
}
</style>
