<template>
  <div class="home-view">
    <!-- Hero Section: Headline & Organic Fluid Droplet Stage (Scene.mp4 00:00 & 00:08) -->
    <HeroSection />

    <!-- Typography Ticker: vadikara -> vadikara.dev (Scene.mp4 00:01 & 00:02) -->
    <BrandTicker />

    <!-- Craft & Disciplines: Vertical Typography List with Left Showcase Tile (Scene.mp4 00:03) -->
    <DisciplineShowcase />

    <!-- Featured Work: Horizontal Carousel & Developer Manifesto (Scene.mp4 00:04 & 00:06) -->
    <ProjectsSection />

    <!-- Vadikara Labs: Experimental Code & Generative Art (Scene.mp4 00:05 & 00:06) -->
    <LabsSection />

    <!-- Direct Developer Contact & Collaboration -->
    <ContactSection />
  </div>
</template>

<script setup>
import HeroSection from '../components/HeroSection.vue'
import BrandTicker from '../components/BrandTicker.vue'
import DisciplineShowcase from '../components/DisciplineShowcase.vue'
import ProjectsSection from '../components/ProjectsSection.vue'
import LabsSection from '../components/LabsSection.vue'
import ContactSection from '../components/ContactSection.vue'
</script>

<style scoped>
.home-view {
  display: flex;
  flex-direction: column;
  position: relative;
}
</style>
