<template>
  <div class="article-detail-page">
    <div class="container article-container" v-if="article">
      <router-link to="/articles" class="back-link mono">← Kembali ke Daftar Publikasi</router-link>

      <div class="article-header">
        <span class="badge badge-cyan mono">{{ article.category }}</span>
        <h1 class="detail-title">{{ article.title }}</h1>
        
        <div class="detail-meta mono">
          <span>Ditulis oleh <strong>{{ article.author }}</strong></span>
          <span>•</span>
          <span>{{ article.read_time }}</span>
        </div>
      </div>

      <div class="article-body glass-panel">
        <p class="lead-excerpt">{{ article.excerpt }}</p>
        <div class="content-text">
          <p>{{ article.content }}</p>
        </div>

        <div class="tags-row" v-if="article.tags">
          <span v-for="(t, i) in article.tags" :key="i" class="tag-chip mono">#{{ t }}</span>
        </div>
      </div>
    </div>

    <div class="container" v-else>
      <div class="empty-state glass-panel">
        <p>Memuat artikel...</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { api } from '../services/api'

const route = useRoute()
const article = ref(null)

onMounted(async () => {
  try {
    const slug = route.params.slug
    article.value = await api.getArticle(slug)
  } catch (err) {
    console.error(err)
  }
})
</script>

<style scoped>
.article-detail-page {
  padding: 130px 0 80px;
  position: relative;
  z-index: 1;
}

.article-container {
  max-width: 840px;
}

.back-link {
  color: var(--cyan-primary);
  text-decoration: none;
  font-size: 0.88rem;
  display: inline-block;
  margin-bottom: 24px;
}

.article-header {
  margin-bottom: 36px;
}

.detail-title {
  font-size: clamp(2rem, 3.8vw, 3rem);
  font-weight: 800;
  line-height: 1.2;
  margin: 16px 0;
  color: var(--text-main);
}

.detail-meta {
  display: flex;
  gap: 12px;
  font-size: 0.85rem;
  color: var(--text-dim);
}

.detail-meta strong {
  color: var(--text-main);
}

.article-body {
  padding: 48px;
  background: rgba(255, 255, 255, 0.94);
  border: 1px solid rgba(20, 60, 35, 0.08);
  box-shadow: var(--shadow-card);
}

.lead-excerpt {
  font-size: 1.15rem;
  color: var(--text-main);
  line-height: 1.8;
  font-weight: 500;
  margin-bottom: 30px;
  padding-bottom: 24px;
  border-bottom: 1px solid rgba(20, 60, 35, 0.08);
}

.content-text {
  color: var(--text-muted);
  font-size: 1.05rem;
  line-height: 1.9;
  margin-bottom: 40px;
}

.tags-row {
  display: flex;
  gap: 8px;
  padding-top: 20px;
  border-top: 1px solid rgba(20, 60, 35, 0.08);
}

.tag-chip {
  font-size: 0.78rem;
  color: #166534;
  background: #f4f7f2;
  border: 1px solid rgba(20, 60, 35, 0.1);
  padding: 4px 10px;
  border-radius: 4px;
  font-weight: 500;
}

@media (max-width: 600px) {
  .article-body {
    padding: 24px;
  }
}
</style>
