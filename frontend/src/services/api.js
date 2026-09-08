const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

function getHeaders(isAuth = false) {
  const headers = { 'Content-Type': 'application/json' }
  if (isAuth) {
    const token = localStorage.getItem('vk_admin_token')
    if (token) {
      headers['Authorization'] = `Bearer ${token}`
    }
  }
  return headers
}

async function request(endpoint, options = {}) {
  const url = `${BASE_URL}${endpoint}`
  const isAuth = options.auth || false
  const res = await fetch(url, {
    ...options,
    headers: {
      ...getHeaders(isAuth),
      ...options.headers,
    },
  })

  const data = await res.json().catch(() => ({}))
  if (!res.ok) {
    throw new Error(data.error || `HTTP error ${res.status}`)
  }
  return data
}

export const api = {
  // Public
  getHealth: () => request('/health'),
  getSiteContent: () => request('/content'),
  getProfile: () => request('/profile'),
  getProjects: () => request('/projects'),
  getProject: (slug) => request(`/projects/${slug}`),
  getServices: () => request('/services'),
  getArticles: () => request('/articles'),
  getArticle: (slug) => request(`/articles/${slug}`),
  sendMessage: (payload) => request('/messages', { method: 'POST', body: JSON.stringify(payload) }),
  
  // Complaints & Grievance (Layanan Pengaduan Resmi)
  sendComplaint: (payload) => request('/complaints', { method: 'POST', body: JSON.stringify(payload) }),
  getComplaintByTicket: (ticket) => request(`/complaints/${ticket}`),

  // Data Deletion Request (UU PDP)
  requestDataDeletion: (payload) => request('/privacy/data-deletion', { method: 'POST', body: JSON.stringify(payload) }),

  // Auth
  login: (email, password) => request('/auth/login', { method: 'POST', body: JSON.stringify({ email, password }) }),

  // Protected Admin CMS
  getMe: () => request('/admin/me', { auth: true }),
  deleteAccount: (password) => request('/admin/account', { method: 'DELETE', auth: true, body: JSON.stringify({ password }) }),
  getStats: () => request('/admin/stats', { auth: true }),
  updateProfile: (payload) => request('/admin/profile', { method: 'PUT', auth: true, body: JSON.stringify(payload) }),
  updateSiteContent: (payload) => request('/admin/content', { method: 'PUT', auth: true, body: JSON.stringify(payload) }),

  // Projects CRUD
  createProject: (payload) => request('/admin/projects', { method: 'POST', auth: true, body: JSON.stringify(payload) }),
  updateProject: (id, payload) => request(`/admin/projects/${id}`, { method: 'PUT', auth: true, body: JSON.stringify(payload) }),
  deleteProject: (id) => request(`/admin/projects/${id}`, { method: 'DELETE', auth: true }),

  // Services CRUD
  createService: (payload) => request('/admin/services', { method: 'POST', auth: true, body: JSON.stringify(payload) }),
  updateService: (id, payload) => request(`/admin/services/${id}`, { method: 'PUT', auth: true, body: JSON.stringify(payload) }),
  deleteService: (id) => request(`/admin/services/${id}`, { method: 'DELETE', auth: true }),

  // Articles CRUD
  createArticle: (payload) => request('/admin/articles', { method: 'POST', auth: true, body: JSON.stringify(payload) }),
  updateArticle: (id, payload) => request(`/admin/articles/${id}`, { method: 'PUT', auth: true, body: JSON.stringify(payload) }),
  deleteArticle: (id) => request(`/admin/articles/${id}`, { method: 'DELETE', auth: true }),

  // Inquiries / Messages
  getAdminMessages: () => request('/admin/messages', { auth: true }),
  updateMessageStatus: (id, status) => request(`/admin/messages/${id}`, { method: 'PUT', auth: true, body: JSON.stringify({ status }) }),
  deleteMessage: (id) => request(`/admin/messages/${id}`, { method: 'DELETE', auth: true }),

  // Complaints / Helpdesk
  getAdminComplaints: () => request('/admin/complaints', { auth: true }),
  updateComplaint: (id, status, resolution_notes) => request(`/admin/complaints/${id}`, { method: 'PUT', auth: true, body: JSON.stringify({ status, resolution_notes }) }),
  deleteComplaint: (id) => request(`/admin/complaints/${id}`, { method: 'DELETE', auth: true }),

  // Data Deletions (UU PDP)
  getAdminDataDeletions: () => request('/admin/data-deletions', { auth: true }),
  updateDataDeletionStatus: (id, status) => request(`/admin/data-deletions/${id}`, { method: 'PUT', auth: true, body: JSON.stringify({ status }) }),
}
