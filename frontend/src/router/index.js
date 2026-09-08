import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import ProjectsView from '../views/ProjectsView.vue'
import ServicesView from '../views/ServicesView.vue'
import ArticlesView from '../views/ArticlesView.vue'
import ArticleDetailView from '../views/ArticleDetailView.vue'
import ComplaintView from '../views/ComplaintView.vue'
import PrivacyPolicyView from '../views/PrivacyPolicyView.vue'

import AdminLogin from '../views/admin/AdminLogin.vue'
import AdminDashboard from '../views/admin/AdminDashboard.vue'

const routes = [
  { path: '/', name: 'Home', component: HomeView },
  { path: '/projects', name: 'Projects', component: ProjectsView },
  { path: '/services', name: 'Services', component: ServicesView },
  { path: '/articles', name: 'Articles', component: ArticlesView },
  { path: '/articles/:slug', name: 'ArticleDetail', component: ArticleDetailView },
  { path: '/pengaduan', name: 'Complaint', component: ComplaintView },
  { path: '/privacy-policy', name: 'PrivacyPolicy', component: PrivacyPolicyView },
  
  // Admin / CMS
  { path: '/admin/login', name: 'AdminLogin', component: AdminLogin },
  { 
    path: '/admin', 
    name: 'AdminDashboard', 
    component: AdminDashboard,
    meta: { requiresAuth: true }
  },
  // Fallback
  { path: '/:pathMatch(.*)*', redirect: '/' }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else if (to.hash) {
      return { el: to.hash, behavior: 'smooth' }
    } else {
      return { top: 0, behavior: 'smooth' }
    }
  }
})

// Navigation Guard for Admin
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('vk_admin_token')
  if (to.meta.requiresAuth && !token) {
    next({ name: 'AdminLogin' })
  } else if (to.name === 'AdminLogin' && token) {
    next({ name: 'AdminDashboard' })
  } else {
    next()
  }
})

export default router
