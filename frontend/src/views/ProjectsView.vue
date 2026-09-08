<template>
  <div class="projects-page">
    <div class="container">
      <div class="page-header">
        <span class="badge badge-cyan mono">Koleksi Portofolio Lengkap</span>
        <h1 class="page-title">Sistem & Arsitektur <span class="text-gradient-cyan">Rekayasa Vadikara.</span></h1>
        <p class="page-subtitle">
          Eksplorasi mendalam mengenai perancangan arsitektur, optimasi latensi gRPC, pipeline AI, dan platform enterprise yang telah kami luncurkan.
        </p>
      </div>

      <ProjectsSection />
    </div>
  </div>
</template>

<script setup>
import ProjectsSection from '../components/ProjectsSection.vue'
</script>

<style scoped>
.projects-page {
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
