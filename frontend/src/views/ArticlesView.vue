<template>
  <div class="articles-page">
    <div class="container">
      <div class="page-header">
        <span class="badge badge-cyan mono">Publikasi Teknis & Catatan Arsitektur</span>
        <h1 class="page-title">Rekayasa & <span class="text-gradient-cyan">Tech Insights.</span></h1>
        <p class="page-subtitle">
          Dokumentasi pemikiran mendalam, studi kasus beban puncak, dan prinsip rekayasa yang kami terapkan di PT. Vadikara Rekayasa Peranti.
        </p>
      </div>

      <div class="articles-grid" v-if="articles.length > 0">
        <article 
          v-for="article in articles" 
          :key="article.id || article.slug" 
          class="article-card glass-panel"
        >
          <div class="article-meta">
            <span class="badge badge-cyan mono">{{ article.category }}</span>
            <span class="read-time mono">{{ article.read_time }}</span>
          </div>

          <h2 class="article-title">
            <router-link :to="`/articles/${article.slug}`">{{ article.title }}</router-link>
          </h2>

          <p class="article-excerpt">{{ article.excerpt }}</p>

          <div class="article-tags">
            <span v-for="(t, i) in article.tags" :key="i" class="tag-chip mono">#{{ t }}</span>
          </div>

          <div class="article-footer">
            <span class="author-info">Oleh <strong>{{ article.author }}</strong></span>
            <router-link :to="`/articles/${article.slug}`" class="read-more-link">
              Baca Selengkapnya →
            </router-link>
          </div>
        </article>
      </div>

      <div v-else class="empty-state glass-panel">
        <p>Memuat publikasi teknis...</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { api } from '../services/api'

const articles = ref([])

onMounted(async () => {
  try {
    const data = await api.getArticles()
    articles.value = data
  } catch (err) {
    console.error(err)
  }
})
</script>

<style scoped>
.articles-page {
  padding: 130px 0 80px;
  position: relative;
  z-index: 1;
}

.page-header {
  text-align: center;
  max-width: 800px;
  margin: 0 auto 50px;
}

.page-title {
  font-size: clamp(2.2rem, 4vw, 3.2rem);
  font-weight: 800;
  margin: 16px 0;
}

.page-subtitle {
  font-size: 1.05rem;
  color: var(--text-muted);
}

.articles-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 28px;
}

.article-card {
  padding: 32px;
  display: flex;
  flex-direction: column;
  background: rgba(255, 255, 255, 0.92);
  border: 1px solid rgba(20, 60, 35, 0.08);
  transition: transform 0.3s ease, border-color 0.3s ease, box-shadow 0.3s ease;
}

.article-card:hover {
  transform: translateY(-6px);
  border-color: rgba(22, 163, 74, 0.35);
  box-shadow: var(--shadow-card-hover);
}

.article-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 18px;
}

.read-time {
  font-size: 0.75rem;
  color: var(--text-dim);
}

.article-title {
  font-size: 1.35rem;
  font-weight: 700;
  line-height: 1.4;
  margin-bottom: 14px;
}

.article-title a {
  color: var(--text-main);
  text-decoration: none;
  transition: color 0.2s;
}

.article-title a:hover {
  color: #15803d;
}

.article-excerpt {
  color: var(--text-muted);
  font-size: 0.92rem;
  line-height: 1.6;
  margin-bottom: 20px;
}

.article-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 24px;
}

.tag-chip {
  font-size: 0.72rem;
  color: #166534;
  background: #f4f7f2;
  border: 1px solid rgba(20, 60, 35, 0.1);
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 500;
}

.article-footer {
  margin-top: auto;
  padding-top: 16px;
  border-top: 1px solid rgba(20, 60, 35, 0.08);
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.85rem;
}

.author-info {
  color: var(--text-dim);
}

.read-more-link {
  color: #15803d;
  text-decoration: none;
  font-weight: 600;
}

@media (max-width: 1024px) {
  .articles-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .articles-grid {
    grid-template-columns: 1fr;
  }
}
</style>
