<template>
  <div class="admin-wrapper">
    <!-- Top Navigation Bar -->
    <header class="admin-topbar glass-panel">
      <div class="topbar-brand">
        <router-link to="/" class="brand-link">
          <div class="brand-dot"></div>
          <span class="mono brand-title">VADIKARA // PUSAT KONTROL CMS</span>
        </router-link>
        <span class="badge badge-cyan mono">Studio Portfolio Manager</span>
      </div>

      <div class="topbar-actions">
        <button class="btn btn-secondary btn-sm mono" @click="loadAllData" title="Sinkronkan data dari basis data">
          🔄 Sinkronkan Data
        </button>
        <a href="/" target="_blank" class="btn btn-primary btn-sm mono" title="Buka website publik di tab baru">
          🌐 Lihat Situs Langsung &nearr;
        </a>
        <button class="btn btn-secondary btn-sm mono" @click="handleLogout">
          🚪 Keluar
        </button>
      </div>
    </header>

    <div class="admin-main">
      <!-- Admin Sidebar Navigation -->
      <aside class="admin-sidebar glass-panel">
        <div class="sidebar-user">
          <div class="user-avatar mono">V</div>
          <div class="user-meta">
            <strong class="user-name">Vadikara</strong>
            <span class="user-role mono">SUPERADMIN</span>
          </div>
        </div>

        <nav class="sidebar-nav">
          <button 
            v-for="t in tabs" 
            :key="t.id"
            class="nav-tab-btn mono"
            :class="{ 'active': activeTab === t.id }"
            @click="activeTab = t.id"
          >
            <span class="tab-icon">{{ t.icon }}</span>
            <span class="tab-label">{{ t.label }}</span>
            <span v-if="t.badge" class="tab-badge">{{ t.badge }}</span>
          </button>
        </nav>

        <div class="sidebar-footer-info mono">
          <div class="engine-badge">
            <span class="dot green-pulse"></span>
            <span>API &amp; SQLite Terkoneksi</span>
          </div>
        </div>
      </aside>

      <!-- Admin Content Canvas -->
      <main class="admin-content">
        <!-- GLOBAL TOAST ALERTS -->
        <transition name="fade">
          <div v-if="successAlert" class="alert alert-success mono">
            <span>✅</span> {{ successAlert }}
          </div>
        </transition>
        <transition name="fade">
          <div v-if="errorAlert" class="alert alert-error mono">
            <span>⚠️</span> {{ errorAlert }}
          </div>
        </transition>

        <!-- 1. TAB: RINGKASAN (OVERVIEW) -->
        <section v-if="activeTab === 'overview'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Ringkasan Sistem &amp; Konten</h2>
              <p class="pane-desc">Status seluruh modul konten website, karya aktif, dan kotak masuk pesan.</p>
            </div>
            <button class="btn btn-primary btn-sm mono" @click="activeTab = 'hero'">
              ✏️ Mulai Kelola Konten
            </button>
          </div>

          <div class="stats-cards-grid">
            <div class="stat-card glass-panel" @click="activeTab = 'projects'">
              <span class="stat-label mono">KARYA TERPILIH</span>
              <div class="stat-val text-gradient-cyan mono">{{ projects.length }}</div>
              <span class="stat-sub">Proyek sistem produksi</span>
            </div>

            <div class="stat-card glass-panel" @click="activeTab = 'disciplines'">
              <span class="stat-label mono">KEAHLIAN &amp; DISIPLIN</span>
              <div class="stat-val text-gradient-violet mono">{{ siteContent.disciplines?.length || 0 }}</div>
              <span class="stat-sub">Showcase kinetik aktif</span>
            </div>

            <div class="stat-card glass-panel" @click="activeTab = 'labs'">
              <span class="stat-label mono">EKSPERIMEN LABS</span>
              <div class="stat-val text-gradient-gold mono">{{ siteContent.labs?.experiments?.length || 0 }}</div>
              <span class="stat-sub">Prototipe aktif di drawer</span>
            </div>

            <div class="stat-card glass-panel" @click="activeTab = 'messages'">
              <span class="stat-label mono">PESAN MASUK</span>
              <div class="stat-val text-gradient-cyan mono">{{ messages.length }}</div>
              <span class="stat-sub">{{ newMessagesCount }} pesan baru</span>
            </div>
          </div>

          <!-- Quick Access Cards -->
          <div class="overview-details-grid">
            <div class="glass-panel overview-box">
              <h3 class="box-title">Daftar Modul Konten Terhubung</h3>
              <div class="module-status-list">
                <div class="module-status-row" @click="activeTab = 'navbar'">
                  <div>
                    <strong>Navigasi &amp; Header</strong>
                    <p class="table-sub mono">Logo, tautan menu ({{ siteContent.navbar?.links?.length || 0 }} item), tombol GitHub</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>

                <div class="module-status-row" @click="activeTab = 'hero'">
                  <div>
                    <strong>Hero Section &amp; Reel Showcase</strong>
                    <p class="table-sub mono">Headline puitis, video reel modal, gambar panggung cairan</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>

                <div class="module-status-row" @click="activeTab = 'ticker'">
                  <div>
                    <strong>Ticker Marquee Teks Berjalan</strong>
                    <p class="table-sub mono">{{ siteContent.ticker?.length || 0 }} Kata Kunci Berputar</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>

                <div class="module-status-row" @click="activeTab = 'disciplines'">
                  <div>
                    <strong>Showcase Keahlian (Scene.mp4 00:03)</strong>
                    <p class="table-sub mono">{{ siteContent.disciplines?.length || 0 }} Disiplin Terdaftar</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>

                <div class="module-status-row" @click="activeTab = 'projects'">
                  <div>
                    <strong>Karya Terpilih (Projects Carousel)</strong>
                    <p class="table-sub mono">{{ projects.length }} Proyek dengan Metrik Benchmark</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>

                <div class="module-status-row" @click="activeTab = 'manifesto'">
                  <div>
                    <strong>Manifesto &amp; Filosofi Pengembang</strong>
                    <p class="table-sub mono">2 Kolom Pernyataan Kode &amp; Tombol</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>

                <div class="module-status-row" @click="activeTab = 'labs'">
                  <div>
                    <strong>Vadikara Labs &amp; Prototipe</strong>
                    <p class="table-sub mono">{{ siteContent.labs?.experiments?.length || 0 }} Eksperimen Aktif</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>

                <div class="module-status-row" @click="activeTab = 'contact'">
                  <div>
                    <strong>Kontak Langsung &amp; Sosial Media</strong>
                    <p class="table-sub mono">{{ siteContent.contact?.direct_email }}</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>

                <div class="module-status-row" @click="activeTab = 'footer'">
                  <div>
                    <strong>Footer &amp; Ambient Sound</strong>
                    <p class="table-sub mono">Tagline, tautan indeks, audio synthesizer</p>
                  </div>
                  <span class="badge badge-live mono">DINAMIS ✓</span>
                </div>
              </div>
            </div>

            <!-- Recent Inquiries Box -->
            <div class="glass-panel overview-box">
              <div class="box-header-row">
                <h3 class="box-title">Pesan Masuk Terbaru</h3>
                <button class="btn btn-secondary btn-xs mono" @click="activeTab = 'messages'">Buka Inbox</button>
              </div>

              <div v-if="messages.length === 0" class="empty-hint mono">
                Belum ada pesan masuk dari pengunjung website.
              </div>
              <div v-else class="recent-list">
                <div v-for="m in messages.slice(0, 4)" :key="m.id" class="recent-item">
                  <div class="recent-header">
                    <div>
                      <strong>{{ m.name }}</strong>
                      <span class="table-sub mono"> ({{ m.email }})</span>
                    </div>
                    <span class="badge" :class="m.status === 'new' ? 'badge-amber' : 'badge-cyan'">
                      {{ m.status === 'new' ? 'BARU' : 'DITINJAU' }}
                    </span>
                  </div>
                  <div class="recent-topic mono text-dim">{{ m.service_interest || 'Dev Inquiry' }}</div>
                  <p class="recent-text">{{ m.message }}</p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- 2. TAB: NAVIGASI & HEADER -->
        <section v-if="activeTab === 'navbar'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Navigasi &amp; Header Website</h2>
              <p class="pane-desc">Kelola nama brand, tombol aksi GitHub, dan tautan menu navigasi utama.</p>
            </div>
            <button class="btn btn-primary btn-sm mono" @click="saveNavbar" :disabled="isSaving">
              💾 {{ isSaving ? 'Menyimpan...' : 'Simpan Navigasi' }}
            </button>
          </div>

          <div class="glass-panel content-form-card">
            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">TEKS LOGO BRAND</label>
                <input v-model="siteContent.navbar.brand_name" type="text" class="form-input mono" placeholder="Vadikara" />
              </div>
              <div class="form-group">
                <label class="form-label mono">SIMBOL SUPER-SCRIPT (SUP)</label>
                <input v-model="siteContent.navbar.brand_sup" type="text" class="form-input mono" placeholder="®" />
              </div>
            </div>

            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">LABEL TOMBOL AKSI KANAN</label>
                <input v-model="siteContent.navbar.action_label" type="text" class="form-input mono" placeholder="GitHub" />
              </div>
              <div class="form-group">
                <label class="form-label mono">URL TOMBOL AKSI KANAN</label>
                <input v-model="siteContent.navbar.action_url" type="url" class="form-input mono" placeholder="https://github.com/vadikara" />
              </div>
            </div>

            <h3 class="subheading mono" style="margin-top: 24px;">// TAUTAN MENU NAVIGASI</h3>
            <div class="links-manage-list">
              <div v-for="(l, idx) in siteContent.navbar.links" :key="idx" class="link-manage-row">
                <span class="item-number mono">#0{{ idx + 1 }}</span>
                <input v-model="l.label" type="text" class="form-input" placeholder="Label Menu (contoh: Karya)" style="flex: 1;" />
                <input v-model="l.url" type="text" class="form-input mono" placeholder="URL Target (#work atau https://...)" style="flex: 1.5;" />
                <button class="btn btn-danger btn-xs mono" @click="removeNavbarLink(idx)">✕</button>
              </div>
            </div>

            <button class="btn btn-secondary btn-sm mono" @click="addNavbarLink" style="align-self: flex-start;">
              ➕ Tambah Tautan Menu
            </button>
          </div>
        </section>

        <!-- 3. TAB: HERO & REEL -->
        <section v-if="activeTab === 'hero'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Hero Section &amp; Video Reel</h2>
              <p class="pane-desc">Atur headline puitis, teks lencana, tautan video reel, dan gambar latar panggung cairan.</p>
            </div>
            <button class="btn btn-primary btn-sm mono" @click="saveHero" :disabled="isSaving">
              💾 {{ isSaving ? 'Menyimpan...' : 'Simpan Hero' }}
            </button>
          </div>

          <div class="glass-panel content-form-card">
            <div class="form-group">
              <label class="form-label mono">HEADLINE UTAMA (Mendukung baris baru \n)</label>
              <textarea 
                v-model="siteContent.hero.title" 
                class="form-textarea headline-input" 
                rows="3"
                placeholder="Rekayasawan sistem kreatif &\nperanti lunak terdistribusi"
              ></textarea>
              <span class="field-hint mono">Gunakan tombol Enter untuk memecah baris teks seperti pada Scene.mp4.</span>
            </div>

            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">LENCANA MODAL REEL (REEL BADGE)</label>
                <input 
                  v-model="siteContent.hero.reel_badge" 
                  type="text" 
                  class="form-input mono" 
                  placeholder="VADIKARA // SHOWCASE REKAYASA SISTEM 2026"
                />
              </div>

              <div class="form-group">
                <div class="label-with-badge">
                  <label class="form-label mono">URL VIDEO REEL (YouTube / Vimeo / MP4)</label>
                  <span v-if="heroVideoMeta.isYoutube" class="badge badge-amber mono">🔴 YOUTUBE (ID: {{ heroVideoMeta.id }})</span>
                  <span v-else-if="heroVideoMeta.isVimeo" class="badge badge-cyan mono">🔵 VIMEO (ID: {{ heroVideoMeta.id }})</span>
                  <span v-else class="badge badge-live mono">🟢 VIDEO LANGSUNG</span>
                </div>
                <input 
                  v-model="siteContent.hero.reel_url" 
                  type="text" 
                  class="form-input mono" 
                  placeholder="https://youtu.be/... atau /Scene.mp4 atau berkas video"
                />
                <div class="reel-quick-actions">
                  <button type="button" class="btn btn-secondary btn-xs mono" @click="siteContent.hero.reel_url = '/Scene.mp4'">
                    🎬 Gunakan Bawaan (/Scene.mp4)
                  </button>
                  <label class="btn btn-secondary btn-xs mono file-upload-label">
                    📁 Unggah Video MP4/WebM...
                    <input type="file" accept="video/mp4,video/webm" @change="handleVideoUpload" style="display: none;" />
                  </label>
                </div>
              </div>
            </div>

            <!-- Video Live Preview in CMS -->
            <div class="form-group video-preview-group">
              <div class="video-preview-header">
                <label class="form-label mono">PRATINJAU LANGSUNG VIDEO REEL (LIVE PREVIEW CMS)</label>
                <span class="mono text-dim" style="font-size: 0.72rem;">{{ heroVideoMeta.label }}</span>
              </div>
              <div class="video-preview-player-box">
                <!-- If YouTube or Vimeo -->
                <iframe 
                  v-if="heroVideoMeta.isYoutube || heroVideoMeta.isVimeo"
                  :src="heroVideoMeta.previewEmbedUrl" 
                  class="cms-video-frame"
                  title="Pratinjau Video Reel"
                  allow="accelerometer; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                  allowfullscreen
                ></iframe>
                <!-- If Direct Video -->
                <video 
                  v-else
                  :src="heroVideoMeta.directUrl || '/Scene.mp4'" 
                  controls 
                  class="cms-video-player"
                ></video>
              </div>
              <span class="field-hint mono">Pratinjau di atas memverifikasi bahwa URL video (YouTube/MP4) dapat diputar sebelum Anda menyimpannya ke website.</span>
            </div>

            <div class="form-group">
              <label class="form-label mono">TEKS FOOTER TELEMETRI REEL</label>
              <input 
                v-model="siteContent.hero.reel_telemetry" 
                type="text" 
                class="form-input mono" 
                placeholder="Konkurensi Tinggi Golang • Web Kreatif 60 FPS • Sistem Terdistribusi Sub-milidetik"
              />
            </div>

            <!-- Image Management Widget -->
            <div class="form-group">
              <label class="form-label mono">GAMBAR LATAR PANGGUNG CAIRAN ORGANIK</label>
              <div class="image-manager-box">
                <!-- Preview Thumb -->
                <div class="current-image-preview">
                  <img :src="resolveImage(siteContent.hero.background_image)" alt="Preview" class="preview-img" />
                  <span class="preview-badge mono">GAMBAR AKTIF</span>
                </div>

                <div class="image-options-col">
                  <!-- Preset Selection -->
                  <span class="mono text-dim" style="font-size: 0.72rem;">1. PILIH DARI PRESET STUDIO:</span>
                  <div class="preset-picker-grid">
                    <div 
                      v-for="preset in PRESET_IMAGES" 
                      :key="preset.path"
                      class="preset-thumb-card"
                      :class="{ 'is-selected': siteContent.hero.background_image === preset.path }"
                      @click="siteContent.hero.background_image = preset.path"
                    >
                      <img :src="preset.value" :alt="preset.label" class="preset-img" />
                      <span class="preset-label mono">{{ preset.label.split('(')[0] }}</span>
                    </div>
                  </div>

                  <!-- Upload Option -->
                  <div class="upload-option-row">
                    <span class="mono text-dim" style="font-size: 0.72rem;">2. ATAU UNGGAH DARI PERANGKAT:</span>
                    <label class="upload-btn mono">
                      📁 Pilih File Gambar...
                      <input type="file" accept="image/*" @change="e => handleFileUpload(e, siteContent.hero, 'background_image')" style="display: none;" />
                    </label>
                  </div>

                  <!-- URL Input -->
                  <div class="custom-url-row">
                    <span class="mono text-dim" style="font-size: 0.72rem;">3. ATAU URL KUSTOM:</span>
                    <input 
                      v-model="siteContent.hero.background_image" 
                      type="text" 
                      class="form-input mono" 
                      placeholder="Masukkan URL https://..."
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- 4. TAB: TICKER MARQUEE -->
        <section v-if="activeTab === 'ticker'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Ticker Marquee Teks Berjalan</h2>
              <p class="pane-desc">Kelola frasa tipografi kinetik yang bergulir tanpa henti di layar.</p>
            </div>
            <div class="pane-actions">
              <button class="btn btn-secondary btn-sm mono" @click="addTickerItem">
                ➕ Tambah Teks
              </button>
              <button class="btn btn-primary btn-sm mono" @click="saveTicker" :disabled="isSaving">
                💾 {{ isSaving ? 'Menyimpan...' : 'Simpan Ticker' }}
              </button>
            </div>
          </div>

          <div class="glass-panel content-form-card">
            <div class="ticker-preview-box">
              <span class="mono text-dim">PRATINJAU LANGSUNG:</span>
              <div class="ticker-live-demo">
                <span v-for="(item, idx) in siteContent.ticker" :key="idx" class="live-item mono">
                  {{ item }} <span class="dot-sep">•</span>
                </span>
              </div>
            </div>

            <div class="ticker-edit-list">
              <div 
                v-for="(item, idx) in siteContent.ticker" 
                :key="idx" 
                class="ticker-edit-row"
              >
                <span class="item-number mono">#0{{ idx + 1 }}</span>
                <input 
                  v-model="siteContent.ticker[idx]" 
                  type="text" 
                  class="form-input ticker-text-input" 
                  placeholder="Kata atau frasa ticker..."
                />
                <button 
                  class="btn btn-secondary btn-xs mono" 
                  @click="removeTickerItem(idx)"
                  title="Hapus kata ini"
                >
                  ✕
                </button>
              </div>
            </div>

            <button class="btn btn-secondary btn-sm mono" @click="addTickerItem" style="align-self: flex-start;">
              ➕ Tambah Kata / Frasa Baru
            </button>
          </div>
        </section>

        <!-- 5. TAB: KEAHLIAN & DISIPLIN -->
        <section v-if="activeTab === 'disciplines'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Keahlian &amp; Disiplin Rekayasa</h2>
              <p class="pane-desc">Daftar keahlian tipografi vertikal dengan showcase tile visual di sebelah kiri (Scene.mp4 Frame 00:03).</p>
            </div>
            <div class="pane-actions">
              <button class="btn btn-secondary btn-sm mono" @click="openDisciplineModal(null)">
                ➕ Tambah Disiplin
              </button>
              <button class="btn btn-primary btn-sm mono" @click="saveDisciplines" :disabled="isSaving">
                💾 {{ isSaving ? 'Menyimpan...' : 'Simpan Perubahan' }}
              </button>
            </div>
          </div>

          <div class="disciplines-cards-grid">
            <div 
              v-for="(d, idx) in siteContent.disciplines" 
              :key="d.id || idx"
              class="glass-panel discipline-item-card"
            >
              <div class="card-thumb-header">
                <img :src="resolveImage(d.image)" :alt="d.title" class="card-thumb-img" />
                <div class="card-header-overlay">
                  <span class="mono card-id">0{{ idx + 1 }}</span>
                  <span class="pill-metric mono">{{ d.benchmark }}</span>
                </div>
              </div>

              <h3 class="card-title">{{ d.title }}</h3>
              <p class="card-desc">{{ d.description }}</p>

              <div class="card-stack-row">
                <span v-for="(t, i) in d.stack" :key="i" class="stack-tag mono">{{ t }}</span>
              </div>

              <div class="card-footer-actions">
                <button class="btn btn-secondary btn-xs mono" @click="openDisciplineModal(d, idx)">
                  ✏️ Edit
                </button>
                <button class="btn btn-danger btn-xs mono" @click="removeDiscipline(idx)">
                  🗑️ Hapus
                </button>
              </div>
            </div>
          </div>
        </section>

        <!-- 6. TAB: KARYA TERPILIH (PROJECTS) -->
        <section v-if="activeTab === 'projects'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Karya Terpilih (Projects Carousel)</h2>
              <p class="pane-desc">Proyek sistem produksi dengan metrik benchmark, tech stack, tautan repositori, dan live demo.</p>
            </div>
            <button class="btn btn-primary btn-sm mono" @click="openProjectModal(null)">
              ➕ Tambah Karya Baru
            </button>
          </div>

          <div class="glass-panel projects-table-card">
            <table class="data-table">
              <thead>
                <tr class="mono">
                  <th>GAMBAR</th>
                  <th>JUDUL &amp; KATEGORI</th>
                  <th>BENCHMARK METRIK</th>
                  <th>RINGKASAN</th>
                  <th>TECH STACK</th>
                  <th>TAUTAN LINK</th>
                  <th>AKSI</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="p in projects" :key="p.id">
                  <td>
                    <img :src="resolveImage(p.image_url)" :alt="p.title" class="table-proj-thumb" />
                  </td>
                  <td>
                    <strong>{{ p.title }}</strong><br />
                    <span class="badge badge-cyan mono">{{ p.category }}</span>
                  </td>
                  <td>
                    <span class="mono green-text">{{ p.metrics || 'p99 < 1ms' }}</span>
                  </td>
                  <td>
                    <p class="table-desc-preview">{{ p.description }}</p>
                  </td>
                  <td>
                    <div class="stack-pill-group">
                      <span v-for="(t, idx) in p.tech_stack" :key="idx" class="stack-tag mono">{{ t }}</span>
                    </div>
                  </td>
                  <td>
                    <div class="link-pill-group">
                      <a v-if="p.github_url" :href="p.github_url" target="_blank" class="table-link mono">Repo &nearr;</a>
                      <a v-if="p.live_url" :href="p.live_url" target="_blank" class="table-link mono cyan-text">Demo &nearr;</a>
                    </div>
                  </td>
                  <td>
                    <div class="table-actions">
                      <button class="btn btn-secondary btn-xs mono" @click="openProjectModal(p)">Edit</button>
                      <button class="btn btn-danger btn-xs mono" @click="deleteProject(p.id)">Hapus</button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>

        <!-- 7. TAB: MANIFESTO & FILOSOFI -->
        <section v-if="activeTab === 'manifesto'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Manifesto &amp; Filosofi Pengembang</h2>
              <p class="pane-desc">Pernyataan visi arsitektur 2 kolom di bawah carousel proyek (Scene.mp4 00:04 &amp; 00:06).</p>
            </div>
            <button class="btn btn-primary btn-sm mono" @click="saveManifesto" :disabled="isSaving">
              💾 {{ isSaving ? 'Menyimpan...' : 'Simpan Manifesto' }}
            </button>
          </div>

          <div class="glass-panel content-form-card">
            <div class="form-group">
              <label class="form-label mono">KOLOM 1: PRINSIP REKAYASA KODE (TEKS UTAMA TERANG)</label>
              <textarea 
                v-model="siteContent.manifesto.text_col1" 
                class="form-textarea" 
                rows="4"
                placeholder="Perangkat lunak hebat tidak memiliki tanggal kedaluwarsa..."
              ></textarea>
            </div>

            <div class="form-group">
              <label class="form-label mono">KOLOM 2: FILOSOFI KEANDALAN SISTEM (TEKS REDUP / MUTED)</label>
              <textarea 
                v-model="siteContent.manifesto.text_col2" 
                class="form-textarea" 
                rows="4"
                placeholder="Menciptakan arsitektur yang berbobot, keandalan sistem tanpa akhir..."
              ></textarea>
            </div>

            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">LABEL TOMBOL AKSI</label>
                <input 
                  v-model="siteContent.manifesto.button_text" 
                  type="text" 
                  class="form-input mono" 
                  placeholder="Tentang Vadikara"
                />
              </div>

              <div class="form-group">
                <label class="form-label mono">TAUTAN TUJUAN TOMBOL (URL ATAU HASH)</label>
                <input 
                  v-model="siteContent.manifesto.button_url" 
                  type="text" 
                  class="form-input mono" 
                  placeholder="#contact"
                />
              </div>
            </div>
          </div>
        </section>

        <!-- 8. TAB: VADIKARA LABS -->
        <section v-if="activeTab === 'labs'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Vadikara Labs &amp; Prototipe</h2>
              <p class="pane-desc">Pengaturan kartu labs dan daftar prototipe eksperimental aktif pada panel drawer interaktif.</p>
            </div>
            <div class="pane-actions">
              <button class="btn btn-secondary btn-sm mono" @click="openLabModal(null)">
                ➕ Tambah Prototipe
              </button>
              <button class="btn btn-primary btn-sm mono" @click="saveLabs" :disabled="isSaving">
                💾 {{ isSaving ? 'Menyimpan...' : 'Simpan Labs' }}
              </button>
            </div>
          </div>

          <div class="glass-panel content-form-card">
            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">TAG HEADER LABS</label>
                <input 
                  v-model="siteContent.labs.tag" 
                  type="text" 
                  class="form-input mono" 
                  placeholder="EKSPERIMEN // LABS 0.9"
                />
              </div>

              <div class="form-group">
                <label class="form-label mono">JUDUL BESAR KARTU LABS</label>
                <input 
                  v-model="siteContent.labs.title" 
                  type="text" 
                  class="form-input" 
                  placeholder="Vadikara Labs — arena bermain kode kreatif..."
                />
              </div>
            </div>

            <!-- Labs Image Management -->
            <div class="form-group">
              <label class="form-label mono">GAMBAR LATAR KARTU LABS</label>
              <div class="image-manager-box">
                <div class="current-image-preview">
                  <img :src="resolveImage(siteContent.labs.background_image)" alt="Preview" class="preview-img" />
                  <span class="preview-badge mono">GAMBAR AKTIF</span>
                </div>
                <div class="image-options-col">
                  <div class="preset-picker-grid">
                    <div 
                      v-for="preset in PRESET_IMAGES" 
                      :key="preset.path"
                      class="preset-thumb-card"
                      :class="{ 'is-selected': siteContent.labs.background_image === preset.path }"
                      @click="siteContent.labs.background_image = preset.path"
                    >
                      <img :src="preset.value" :alt="preset.label" class="preset-img" />
                      <span class="preset-label mono">{{ preset.label.split('(')[0] }}</span>
                    </div>
                  </div>
                  <div class="upload-option-row">
                    <span class="mono text-dim" style="font-size: 0.72rem;">UNGGAH GAMBAR:</span>
                    <label class="upload-btn mono">
                      📁 Pilih File...
                      <input type="file" accept="image/*" @change="e => handleFileUpload(e, siteContent.labs, 'background_image')" style="display: none;" />
                    </label>
                  </div>
                  <div class="custom-url-row">
                    <span class="mono text-dim" style="font-size: 0.72rem;">URL KUSTOM:</span>
                    <input v-model="siteContent.labs.background_image" type="text" class="form-input mono" placeholder="https://..." />
                  </div>
                </div>
              </div>
            </div>

            <h3 class="subheading mono" style="margin-top: 30px;">// DAFTAR PROTOTIPE EKSPERIMENTAL ({{ siteContent.labs.experiments?.length || 0 }})</h3>

            <div class="experiments-admin-grid">
              <div 
                v-for="(exp, idx) in siteContent.labs.experiments" 
                :key="exp.id || idx"
                class="experiment-admin-card"
              >
                <div class="card-header-row">
                  <span class="mono card-id">0{{ idx + 1 }}</span>
                  <span class="badge badge-cyan mono">{{ exp.status }}</span>
                </div>
                <h4 class="card-title">{{ exp.name }}</h4>
                <p class="card-desc">{{ exp.desc }}</p>
                <div class="card-stack-row">
                  <span v-for="(t, i) in exp.tags" :key="i" class="stack-tag mono">{{ t }}</span>
                </div>
                <div class="card-footer-actions">
                  <button class="btn btn-secondary btn-xs mono" @click="openLabModal(exp, idx)">Edit</button>
                  <button class="btn btn-danger btn-xs mono" @click="removeLabExperiment(idx)">Hapus</button>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- 9. TAB: KONTAK & JEJARING -->
        <section v-if="activeTab === 'contact'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Kontak &amp; Media Sosial</h2>
              <p class="pane-desc">Kordinat pengembang, email langsung, tautan sosial, dan pilihan topik pada form website.</p>
            </div>
            <button class="btn btn-primary btn-sm mono" @click="saveContact" :disabled="isSaving">
              💾 {{ isSaving ? 'Menyimpan...' : 'Simpan Kontak' }}
            </button>
          </div>

          <div class="glass-panel content-form-card">
            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">TAG HEADER KONTAK</label>
                <input 
                  v-model="siteContent.contact.tag" 
                  type="text" 
                  class="form-input mono" 
                  placeholder="KOLABORASI // JALUR LANGSUNG"
                />
              </div>

              <div class="form-group">
                <label class="form-label mono">EMAIL RESMI LANGSUNG (TOMBOL SALIN)</label>
                <input 
                  v-model="siteContent.contact.direct_email" 
                  type="email" 
                  class="form-input mono" 
                  placeholder="vadikara.dev@gmail.com"
                />
              </div>
            </div>

            <div class="form-group">
              <label class="form-label mono">JUDUL BESAR AJAKAN KOLABORASI</label>
              <input 
                v-model="siteContent.contact.heading" 
                type="text" 
                class="form-input" 
                placeholder="Mari bangun sistem yang bermakna dan berdaya tahan tinggi."
              />
            </div>

            <div class="form-group">
              <label class="form-label mono">SUBTITLE / PENJELASAN KETERBUKAAN KONTRAK</label>
              <textarea 
                v-model="siteContent.contact.subheading" 
                class="form-textarea" 
                rows="3"
                placeholder="Terbuka untuk konsultasi arsitektur sistem prinsipal..."
              ></textarea>
            </div>

            <h3 class="subheading mono" style="margin-top: 24px;">// TAUTAN MEDIA SOSIAL</h3>

            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">GITHUB URL</label>
                <input 
                  v-model="siteContent.contact.github_url" 
                  type="url" 
                  class="form-input mono" 
                  placeholder="https://github.com/vadikara"
                />
              </div>

              <div class="form-group">
                <label class="form-label mono">X / TWITTER URL</label>
                <input 
                  v-model="siteContent.contact.twitter_url" 
                  type="url" 
                  class="form-input mono" 
                  placeholder="https://x.com/vadikara"
                />
              </div>

              <div class="form-group">
                <label class="form-label mono">LINKEDIN URL</label>
                <input 
                  v-model="siteContent.contact.linkedin_url" 
                  type="url" 
                  class="form-input mono" 
                  placeholder="https://linkedin.com"
                />
              </div>
            </div>

            <div class="form-group" style="margin-top: 20px;">
              <label class="form-label mono">PILIHAN TOPIK FOKUS SISTEM PADA FORMULIR (PISAHKAN DENGAN KOMA)</label>
              <input 
                v-model="topicsInput" 
                type="text" 
                class="form-input mono" 
                placeholder="Sistem Terdistribusi, Backend Golang 60k+ RPS, Web Kreatif 60 FPS, Arsitektur Cloud K8s"
              />
              <span class="field-hint mono">Topik-topik ini muncul sebagai tombol pilihan (pills) di form kontak publik.</span>
            </div>
          </div>
        </section>

        <!-- 10. TAB: FOOTER & AUDIO -->
        <section v-if="activeTab === 'footer'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Footer &amp; Ambient Player</h2>
              <p class="pane-desc">Atur teks footer, hak cipta, dan pemutar audio synthesizer ambient mengambang.</p>
            </div>
            <button class="btn btn-primary btn-sm mono" @click="saveFooterAndAmbient" :disabled="isSaving">
              💾 {{ isSaving ? 'Menyimpan...' : 'Simpan Footer & Audio' }}
            </button>
          </div>

          <div class="glass-panel content-form-card">
            <h3 class="subheading mono">// TEKS &amp; INFORMASI FOOTER</h3>

            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">TEKS LOGO FOOTER</label>
                <input v-model="siteContent.footer.logo_text" type="text" class="form-input mono" placeholder="Vadikara" />
              </div>
              <div class="form-group">
                <label class="form-label mono">SIMBOL SUP FOOTER</label>
                <input v-model="siteContent.footer.logo_sup" type="text" class="form-input mono" placeholder="®" />
              </div>
            </div>

            <div class="form-group">
              <label class="form-label mono">TAGLINE FOOTER</label>
              <textarea v-model="siteContent.footer.tagline" class="form-textarea" rows="2" placeholder="Rekayasawan sistem kreatif & pengembang..."></textarea>
            </div>

            <div class="form-row">
              <div class="form-group">
                <label class="form-label mono">TEKS HAK CIPTA (COPYRIGHT)</label>
                <input v-model="siteContent.footer.copyright_text" type="text" class="form-input mono" placeholder="© 2026 Vadikara. Seluruh hak cipta dilindungi." />
              </div>
              <div class="form-group">
                <label class="form-label mono">TEKS TEKNOLOGI (BUILT TEXT)</label>
                <input v-model="siteContent.footer.built_text" type="text" class="form-input mono" placeholder="Dirancang & Dibangun dengan Go 1.26 & Vue 3 • 60 FPS" />
              </div>
            </div>

            <!-- Comprehensive Ambient Audio Manager Suite -->
            <div class="audio-manager-card">
              <div class="audio-card-header">
                <div>
                  <h3 class="subheading mono" style="margin: 0;">// PEMUTAR AUDIO AMBIENT (SCENE.MP4 00:07)</h3>
                  <p class="table-sub mono" style="margin-top: 4px;">Kelola musik latar mengambang: unggah berkas MP3/WAV, masukkan URL audio online, atau gunakan generator synthesizer 432 Hz.</p>
                </div>
                <div class="toggle-switch-wrapper">
                  <label class="switch-container mono">
                    <input type="checkbox" v-model="siteContent.ambient.enabled" />
                    <span class="switch-slider"></span>
                    <span class="switch-label">{{ siteContent.ambient.enabled !== false ? 'DOCK AKTIF' : 'DOCK NONAKTIF' }}</span>
                  </label>
                </div>
              </div>

              <!-- Metadata fields -->
              <div class="form-row" style="margin-top: 18px;">
                <div class="form-group">
                  <label class="form-label mono">JUDUL TREK AUDIO</label>
                  <input v-model="siteContent.ambient.track_title" type="text" class="form-input mono" placeholder="01 // RESONANSI CAIRAN" />
                </div>
                <div class="form-group">
                  <label class="form-label mono">SUBTITLE / ARTIST / FREKUENSI</label>
                  <input v-model="siteContent.ambient.track_subtitle" type="text" class="form-input mono" placeholder="Sintesis Harmonik 432 Hz" />
                </div>
              </div>

              <!-- Audio Source Selector -->
              <div class="form-group">
                <label class="form-label mono">SUMBER SUARA AUDIO (PILIH SALAH SATU ATAU UNGGAH):</label>
                <div class="audio-source-manager">
                  <!-- Mode 1: File Upload -->
                  <div class="audio-mode-box">
                    <div class="mode-header">
                      <strong class="mono">1. UNGGAH BERKAS AUDIO LOKAL (MP3 / WAV / OGG)</strong>
                    </div>
                    <p class="mode-desc">Unggah file audio dari perangkat Anda. Berkas akan dikonversi dan disimpan ke database SQLite.</p>
                    <div class="upload-btn-row">
                      <label class="btn btn-primary btn-sm mono file-upload-label">
                        🎵 Unggah Berkas Audio...
                        <input type="file" accept="audio/*" @change="handleAudioUpload" style="display: none;" />
                      </label>
                      <span v-if="siteContent.ambient.audio_url && siteContent.ambient.audio_url.startsWith('data:audio')" class="badge badge-live mono">
                        ✓ BERKAS AUDIO LOKAL TERSIMPAN
                      </span>
                    </div>
                  </div>

                  <!-- Mode 2: External URL -->
                  <div class="audio-mode-box">
                    <div class="mode-header">
                      <strong class="mono">2. ATAU GUNAKAN URL AUDIO EKSTERNAL</strong>
                    </div>
                    <p class="mode-desc">Masukkan tautan langsung file audio online (format .mp3, .wav, audio stream).</p>
                    <input 
                      v-model="siteContent.ambient.audio_url" 
                      type="url" 
                      class="form-input mono" 
                      placeholder="https://domain.com/audio/ambient_track.mp3" 
                    />
                  </div>

                  <!-- Mode 3: Native Web Audio Synth Preset -->
                  <div class="audio-mode-box">
                    <div class="mode-header">
                      <strong class="mono">3. ATAU SYNTHESIZER GENERATIF BAWAAN</strong>
                    </div>
                    <p class="mode-desc">Gunakan generator gelombang drone analog harmonik 432 Hz murni tanpa berkas audio tambahan.</p>
                    <button type="button" class="btn btn-secondary btn-sm mono" @click="siteContent.ambient.audio_url = ''; siteContent.ambient.track_title = '01 // RESONANSI CAIRAN'; siteContent.ambient.track_subtitle = 'Sintesis Harmonik 432 Hz'">
                      ⚡ Setel ke Synthesizer 432 Hz
                    </button>
                  </div>
                </div>
              </div>

              <!-- Audio Player Test / Preview inside CMS -->
              <div class="cms-audio-preview-dock">
                <div class="preview-dock-left">
                  <button 
                    type="button" 
                    class="audio-test-btn" 
                    :class="{ 'is-playing': isAudioTesting }" 
                    @click="toggleAudioPreview"
                    title="Uji dengar audio di dashboard"
                  >
                    <span v-if="!isAudioTesting">▶</span>
                    <span v-else>❚❚</span>
                  </button>
                  <div class="preview-track-meta">
                    <strong class="track-title-preview mono">{{ siteContent.ambient.track_title || '01 // RESONANSI CAIRAN' }}</strong>
                    <span class="track-sub-preview mono">{{ siteContent.ambient.track_subtitle || 'Sintesis Harmonik 432 Hz' }}</span>
                    <span class="track-status-pill mono text-dim">
                      {{ siteContent.ambient.audio_url ? (siteContent.ambient.audio_url.startsWith('data:audio') ? 'Format: Berkas Audio Lokal' : 'Format: URL Eksternal') : 'Format: Generator Synth 432 Hz' }}
                    </span>
                  </div>
                </div>

                <div class="preview-dock-actions">
                  <div class="preview-spectrum" :class="{ 'spectrum-active': isAudioTesting }">
                    <span class="bar b1"></span>
                    <span class="bar b2"></span>
                    <span class="bar b3"></span>
                    <span class="bar b4"></span>
                    <span class="bar b5"></span>
                  </div>
                  <button 
                    v-if="siteContent.ambient.audio_url" 
                    type="button" 
                    class="btn btn-secondary btn-xs mono" 
                    @click="siteContent.ambient.audio_url = ''"
                    title="Hapus berkas audio dan kembali ke synth"
                  >
                    🗑️ Reset Audio
                  </button>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- 11. TAB: KOTAK MASUK (INBOX) -->
        <section v-if="activeTab === 'messages'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Kotak Masuk (Inquiries Inbox)</h2>
              <p class="pane-desc">Pesan dan proposal kolaborasi yang masuk langsung dari pengunjung website.</p>
            </div>
            <button class="btn btn-secondary btn-sm mono" @click="loadAllData">
              🔄 Segarkan Pesan
            </button>
          </div>

          <div class="glass-panel projects-table-card">
            <div v-if="messages.length === 0" class="empty-hint mono" style="padding: 40px; text-align: center;">
              Kotak masuk masih kosong. Formulir website aktif dan siap menerima pesan.
            </div>
            <table v-else class="data-table">
              <thead>
                <tr class="mono">
                  <th>PENGIRIM</th>
                  <th>TOPIK SISTEM</th>
                  <th>ISI PESAN</th>
                  <th>STATUS</th>
                  <th>AKSI</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="m in messages" :key="m.id">
                  <td>
                    <strong>{{ m.name }}</strong><br />
                    <span class="table-sub mono cyan-text">{{ m.email }}</span><br />
                    <span class="table-sub mono">{{ formatDate(m.created_at) }}</span>
                  </td>
                  <td>
                    <span class="badge badge-cyan mono">{{ m.service_interest || 'Dev Inquiry' }}</span>
                  </td>
                  <td>
                    <p class="table-desc-preview message-content-text">{{ m.message }}</p>
                  </td>
                  <td>
                    <span class="badge" :class="m.status === 'new' ? 'badge-amber' : 'badge-live'">
                      {{ m.status === 'new' ? 'BARU' : 'DITINJAU' }}
                    </span>
                  </td>
                  <td>
                    <div class="table-actions">
                      <a :href="`mailto:${m.email}?subject=Tanggapan Vadikara: ${encodeURIComponent(m.service_interest || 'Kolaborasi Rekayasa Sistem')}`" class="btn btn-secondary btn-xs mono">
                        ✉️ Balas
                      </a>
                      <button 
                        v-if="m.status === 'new'" 
                        class="btn btn-primary btn-xs mono" 
                        @click="markMessageReviewed(m.id)"
                      >
                        ✓ Ditinjau
                      </button>
                      <button class="btn btn-danger btn-xs mono" @click="deleteMessage(m.id)">
                        🗑️ Hapus
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>

        <!-- 12. TAB: PENGATURAN AKUN -->
        <section v-if="activeTab === 'account'" class="tab-pane">
          <div class="pane-header">
            <div>
              <h2 class="pane-title">Pengaturan Akun Administrator</h2>
              <p class="pane-desc">Kredensial login dan status keamanan portal CMS.</p>
            </div>
          </div>

          <div class="glass-panel content-form-card" style="max-width: 600px;">
            <div class="form-group">
              <label class="form-label mono">USERNAME</label>
              <input type="text" class="form-input mono" value="vadikara" disabled />
            </div>

            <div class="form-group">
              <label class="form-label mono">EMAIL LOGIN</label>
              <input type="email" class="form-input mono" value="admin@vadikara.com" disabled />
            </div>

            <div class="form-group">
              <label class="form-label mono">ROLE SISTEM</label>
              <span class="badge badge-live mono">SUPERADMIN (AKSES PENUH)</span>
            </div>

            <div style="margin-top: 24px; border-top: 1px solid rgba(255, 255, 255, 0.08); padding-top: 18px;">
              <p class="mono text-dim" style="font-size: 0.8rem;">
                Kredensial database tersimpan di <code>vadikara.db</code> menggunakan hashing bcrypt standar industri.
              </p>
            </div>
          </div>
        </section>
      </main>
    </div>

    <!-- MODAL: KEAHLIAN / DISCIPLINE -->
    <div class="modal-backdrop" v-if="showDisciplineModal" @click.self="showDisciplineModal = false">
      <div class="modal-box glass-panel">
        <div class="modal-header">
          <h3>{{ editingDisciplineIndex !== null ? 'Edit Keahlian' : 'Tambah Keahlian Baru' }}</h3>
          <button class="modal-close" @click="showDisciplineModal = false">&times;</button>
        </div>
        <form @submit.prevent="saveDisciplineItem">
          <div class="form-group">
            <label class="form-label mono">JUDUL DISIPLIN *</label>
            <input v-model="disciplineForm.title" type="text" class="form-input" required placeholder="contoh: Sistem Terdistribusi" />
          </div>

          <div class="form-group">
            <label class="form-label mono">BENCHMARK / METRIK UTAMA *</label>
            <input v-model="disciplineForm.benchmark" type="text" class="form-input mono" required placeholder="contoh: p99 < 1.2ms atau 350k+ RPS" />
          </div>

          <div class="form-group">
            <label class="form-label mono">DESKRIPSI LENGKAP *</label>
            <textarea v-model="disciplineForm.description" class="form-textarea" rows="3" required placeholder="Jelaskan arsitektur dan keunggulannya..."></textarea>
          </div>

          <div class="form-group">
            <label class="form-label mono">TECH STACK (PISAHKAN DENGAN KOMA)</label>
            <input v-model="disciplineStackInput" type="text" class="form-input mono" placeholder="Golang 1.26, gRPC, Raft, Protobuf v3" />
          </div>

          <!-- Image Picker with Upload -->
          <div class="form-group">
            <label class="form-label mono">GAMBAR SHOWCASE TILE</label>
            <div class="modal-image-selector">
              <div class="thumb-preview">
                <img :src="resolveImage(disciplineForm.image)" alt="Preview" class="modal-thumb-img" />
              </div>
              <div class="selector-options">
                <select v-model="disciplineForm.image" class="form-select mono">
                  <option v-for="p in PRESET_IMAGES" :key="p.path" :value="p.path">
                    Preset: {{ p.label.split('(')[0] }}
                  </option>
                </select>
                <label class="upload-btn mono">
                  📁 Unggah Gambar Sendiri...
                  <input type="file" accept="image/*" @change="e => handleFileUpload(e, disciplineForm, 'image')" style="display: none;" />
                </label>
              </div>
            </div>
          </div>

          <button type="submit" class="btn btn-primary btn-full mono">Simpan Disiplin</button>
        </form>
      </div>
    </div>

    <!-- MODAL: KARYA / PROJECT -->
    <div class="modal-backdrop" v-if="showProjectModal" @click.self="showProjectModal = false">
      <div class="modal-box glass-panel">
        <div class="modal-header">
          <h3>{{ editingProject.id ? 'Edit Karya Proyek' : 'Tambah Karya Proyek Baru' }}</h3>
          <button class="modal-close" @click="showProjectModal = false">&times;</button>
        </div>
        <form @submit.prevent="saveProjectItem">
          <div class="form-group">
            <label class="form-label mono">JUDUL KARYA *</label>
            <input v-model="editingProject.title" type="text" class="form-input" required placeholder="NovaMesh Inti Terdistribusi" />
          </div>

          <div class="form-row">
            <div class="form-group">
              <label class="form-label mono">KATEGORI // TAG *</label>
              <input v-model="editingProject.category" type="text" class="form-input mono" required placeholder="SISTEM // GOLANG" />
            </div>

            <div class="form-group">
              <label class="form-label mono">BENCHMARK METRIK TERUJI</label>
              <input v-model="editingProject.metrics" type="text" class="form-input mono" placeholder="185,000 req/s • p99 < 1.1ms" />
            </div>
          </div>

          <div class="form-group">
            <label class="form-label mono">RINGKASAN SINGKAT (KARTU SLIDER) *</label>
            <textarea v-model="editingProject.description" class="form-textarea" rows="2" required placeholder="Ringkasan 1-2 kalimat..."></textarea>
          </div>

          <div class="form-group">
            <label class="form-label mono">RINGKASAN ARSITEKTUR LENGKAP (MODAL POPUP)</label>
            <textarea v-model="editingProject.architecture_summary" class="form-textarea" rows="3" placeholder="Penjelasan runtime, konsensus, dan rincian teknis..."></textarea>
          </div>

          <div class="form-group">
            <label class="form-label mono">TECH STACK (PISAHKAN DENGAN KOMA)</label>
            <input v-model="projectStackInput" type="text" class="form-input mono" placeholder="Go 1.26, gRPC, Kafka, Redis Cluster, Prometheus" />
          </div>

          <div class="form-row">
            <div class="form-group">
              <label class="form-label mono">URL REPOSITORI GITHUB</label>
              <input v-model="editingProject.github_url" type="url" class="form-input mono" placeholder="https://github.com/vadikara/..." />
            </div>

            <div class="form-group">
              <label class="form-label mono">URL LIVE DEMO</label>
              <input v-model="editingProject.live_url" type="url" class="form-input mono" placeholder="https://..." />
            </div>
          </div>

          <div class="form-group">
            <label class="form-label mono">GAMBAR KARYA SHOWCASE</label>
            <div class="modal-image-selector">
              <div class="thumb-preview">
                <img :src="resolveImage(editingProject.image_url)" alt="Preview" class="modal-thumb-img" />
              </div>
              <div class="selector-options">
                <select v-model="editingProject.image_url" class="form-select mono">
                  <option v-for="p in PRESET_IMAGES" :key="p.path" :value="p.path">
                    Preset: {{ p.label.split('(')[0] }}
                  </option>
                </select>
                <label class="upload-btn mono">
                  📁 Unggah Gambar Sendiri...
                  <input type="file" accept="image/*" @change="e => handleFileUpload(e, editingProject, 'image_url')" style="display: none;" />
                </label>
                <input v-model="editingProject.image_url" type="text" class="form-input mono" placeholder="Atau ketik URL kustom..." />
              </div>
            </div>
          </div>

          <button type="submit" class="btn btn-primary btn-full mono">Simpan Karya Proyek</button>
        </form>
      </div>
    </div>

    <!-- MODAL: LABS EXPERIMENT -->
    <div class="modal-backdrop" v-if="showLabModal" @click.self="showLabModal = false">
      <div class="modal-box glass-panel">
        <div class="modal-header">
          <h3>{{ editingLabIndex !== null ? 'Edit Prototipe Labs' : 'Tambah Prototipe Labs Baru' }}</h3>
          <button class="modal-close" @click="showLabModal = false">&times;</button>
        </div>
        <form @submit.prevent="saveLabExperimentItem">
          <div class="form-group">
            <label class="form-label mono">NAMA PROTOTIPE *</label>
            <input v-model="labForm.name" type="text" class="form-input" required placeholder="Visualisator Ring-Buffer Lock-Free" />
          </div>

          <div class="form-group">
            <label class="form-label mono">STATUS PROTOTIPE *</label>
            <input v-model="labForm.status" type="text" class="form-input mono" required placeholder="RUNTIME AKTIF / SIMULASI KLASTER / WEBGL SHADER" />
          </div>

          <div class="form-group">
            <label class="form-label mono">DESKRIPSI EKSPERIMEN *</label>
            <textarea v-model="labForm.desc" class="form-textarea" rows="3" required placeholder="Penjelasan teknis eksperimen algoritma..."></textarea>
          </div>

          <div class="form-group">
            <label class="form-label mono">TAG TEKNOLOGI (PISAHKAN DENGAN KOMA)</label>
            <input v-model="labTagsInput" type="text" class="form-input mono" placeholder="Go Assembly, Tata Letak Memori, Bebas Kunci" />
          </div>

          <button type="submit" class="btn btn-primary btn-full mono">Simpan Prototipe</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { api } from '../../services/api'
import { useContent, PRESET_IMAGES, resolveImage } from '../../composables/useContent'
import { parseVideoUrl } from '../../utils/media'

const router = useRouter()
const { content: globalContent, fetchContent: refreshGlobalContent } = useContent()

const activeTab = ref('overview')
const isSaving = ref(false)
const successAlert = ref('')
const errorAlert = ref('')

const tabs = [
  { id: 'overview', label: 'Ringkasan', icon: '📊' },
  { id: 'navbar', label: 'Navigasi & Header', icon: '🧭' },
  { id: 'hero', label: 'Hero & Reel', icon: '🎬' },
  { id: 'ticker', label: 'Ticker Marquee', icon: '⚡' },
  { id: 'disciplines', label: 'Keahlian & Disiplin', icon: '💎' },
  { id: 'projects', label: 'Karya Terpilih', icon: '🚀' },
  { id: 'manifesto', label: 'Manifesto', icon: '📜' },
  { id: 'labs', label: 'Vadikara Labs', icon: '🧪' },
  { id: 'contact', label: 'Kontak & Sosial', icon: '📬' },
  { id: 'footer', label: 'Footer & Audio', icon: '⚓' },
  { id: 'messages', label: 'Kotak Masuk', icon: '✉️', badge: computed(() => newMessagesCount.value > 0 ? String(newMessagesCount.value) : null) },
  { id: 'account', label: 'Akun Admin', icon: '⚙️' },
]

// Local reactive editable copies of CMS sections
const siteContent = reactive({
  navbar: {
    brand_name: 'Vadikara',
    brand_sup: '®',
    action_label: 'GitHub',
    action_url: 'https://github.com/vadikara',
    links: []
  },
  hero: {
    title: '',
    reel_badge: '',
    reel_url: '',
    reel_telemetry: '',
    background_image: ''
  },
  ticker: [],
  disciplines: [],
  manifesto: {
    text_col1: '',
    text_col2: '',
    button_text: '',
    button_url: ''
  },
  labs: {
    tag: '',
    title: '',
    background_image: '',
    experiments: []
  },
  contact: {
    tag: '',
    heading: '',
    subheading: '',
    direct_email: '',
    github_url: '',
    twitter_url: '',
    linkedin_url: '',
    topics: []
  },
  footer: {
    logo_text: 'Vadikara',
    logo_sup: '®',
    tagline: '',
    copyright_text: '',
    built_text: ''
  },
  ambient: {
    track_title: '',
    track_subtitle: '',
    audio_url: '',
    enabled: true
  }
})

// Database collections
const projects = ref([])
const messages = ref([])

// Form string inputs for comma-separated fields
const topicsInput = ref('')
const disciplineStackInput = ref('')
const projectStackInput = ref('')
const labTagsInput = ref('')

// Modal states
const showDisciplineModal = ref(false)
const editingDisciplineIndex = ref(null)
const disciplineForm = reactive({
  title: '',
  benchmark: '',
  description: '',
  stack: [],
  image: '/src/assets/dev_workspace.jpg'
})

const showProjectModal = ref(false)
const editingProject = ref({})

const showLabModal = ref(false)
const editingLabIndex = ref(null)
const labForm = reactive({
  name: '',
  status: '',
  desc: '',
  tags: []
})

const newMessagesCount = computed(() => {
  return messages.value.filter(m => m.status === 'new').length
})

const triggerAlert = (msg, isError = false) => {
  if (isError) {
    errorAlert.value = msg
    setTimeout(() => { errorAlert.value = '' }, 4000)
  } else {
    successAlert.value = msg
    setTimeout(() => { successAlert.value = '' }, 3500)
  }
}

// ---------------- MEDIA & FILE UPLOAD HELPERS ----------------
const heroVideoMeta = computed(() => parseVideoUrl(siteContent.hero.reel_url))

const handleFileUpload = (event, targetObj, key) => {
  const file = event.target.files[0]
  if (!file) return
  if (file.size > 3 * 1024 * 1024) {
    alert('Ukuran file gambar maksimal 3MB untuk performa optimal.')
    return
  }
  const reader = new FileReader()
  reader.onload = (e) => {
    targetObj[key] = e.target.result
    triggerAlert('Gambar berhasil dimuat ke editor!')
  }
  reader.readAsDataURL(file)
}

const handleVideoUpload = (event) => {
  const file = event.target.files[0]
  if (!file) return
  if (file.size > 25 * 1024 * 1024) {
    alert('Ukuran video maksimal 25MB untuk penyimpanan database.')
    return
  }
  const reader = new FileReader()
  reader.onload = (e) => {
    siteContent.hero.reel_url = e.target.result
    triggerAlert('Video berhasil dimuat ke editor!')
  }
  reader.readAsDataURL(file)
}

const handleAudioUpload = (event) => {
  const file = event.target.files[0]
  if (!file) return
  if (file.size > 15 * 1024 * 1024) {
    alert('Ukuran file audio maksimal 15MB.')
    return
  }
  const reader = new FileReader()
  reader.onload = (e) => {
    siteContent.ambient.audio_url = e.target.result
    if (!siteContent.ambient.track_title || siteContent.ambient.track_title === '01 // RESONANSI CAIRAN') {
      const name = file.name.replace(/\.[^/.]+$/, '')
      siteContent.ambient.track_title = name.toUpperCase()
    }
    triggerAlert('Berkas audio berhasil diunggah ke editor!')
  }
  reader.readAsDataURL(file)
}

// ---------------- CMS AUDIO PREVIEW TESTER ----------------
const isAudioTesting = ref(false)
let previewAudioObj = null
let synthAudioCtx = null
let synthOscNode = null
let synthGainNode = null

const stopAudioPreview = () => {
  if (previewAudioObj) {
    try { previewAudioObj.pause() } catch (e) {}
    previewAudioObj = null
  }
  if (synthOscNode) {
    try { synthOscNode.stop() } catch (e) {}
    synthOscNode = null
  }
  isAudioTesting.value = false
}

const toggleAudioPreview = () => {
  if (isAudioTesting.value) {
    stopAudioPreview()
    return
  }

  const src = siteContent.ambient.audio_url
  if (!src) {
    // Generative Web Audio API drone preview
    playTestSynth()
    return
  }

  try {
    previewAudioObj = new Audio(src)
    previewAudioObj.onended = () => { isAudioTesting.value = false }
    previewAudioObj.onerror = () => {
      triggerAlert('Gagal memutar audio: periksa tautan atau format berkas.', true)
      isAudioTesting.value = false
    }
    previewAudioObj.play().then(() => {
      isAudioTesting.value = true
    }).catch(err => {
      triggerAlert('Pemutaran diblokir peramban: ' + err.message, true)
      isAudioTesting.value = false
    })
  } catch (err) {
    triggerAlert('Gagal inisialisasi audio: ' + err.message, true)
  }
}

const playTestSynth = () => {
  try {
    const AudioContext = window.AudioContext || window.webkitAudioContext
    synthAudioCtx = new AudioContext()
    synthOscNode = synthAudioCtx.createOscillator()
    synthGainNode = synthAudioCtx.createGain()

    synthOscNode.type = 'sine'
    synthOscNode.frequency.setValueAtTime(108, synthAudioCtx.currentTime) // 432 Hz harmonic
    synthGainNode.gain.setValueAtTime(0.06, synthAudioCtx.currentTime)

    synthOscNode.connect(synthGainNode)
    synthGainNode.connect(synthAudioCtx.destination)
    synthOscNode.start()
    isAudioTesting.value = true
    triggerAlert('Memutar pratinjau nada harmonik 432 Hz...')

    setTimeout(() => {
      if (synthGainNode && synthAudioCtx) {
        synthGainNode.gain.setTargetAtTime(0, synthAudioCtx.currentTime, 0.4)
        setTimeout(() => {
          stopAudioPreview()
        }, 500)
      }
    }, 4000)
  } catch (e) {
    triggerAlert('Gagal memutar generator synth: ' + e.message, true)
  }
}

// ---------------- LOAD ALL DATA ----------------
const loadAllData = async () => {
  try {
    const [cData, pData, mData] = await Promise.all([
      api.getSiteContent().catch(() => ({})),
      api.getProjects().catch(() => []),
      api.getAdminMessages().catch(() => [])
    ])

    if (cData) {
      if (cData.navbar) {
        Object.assign(siteContent.navbar, cData.navbar)
        if (Array.isArray(cData.navbar.links)) siteContent.navbar.links = JSON.parse(JSON.stringify(cData.navbar.links))
      }
      if (cData.hero) Object.assign(siteContent.hero, cData.hero)
      if (Array.isArray(cData.ticker)) siteContent.ticker = [...cData.ticker]
      if (Array.isArray(cData.disciplines)) siteContent.disciplines = JSON.parse(JSON.stringify(cData.disciplines))
      if (cData.manifesto) Object.assign(siteContent.manifesto, cData.manifesto)
      if (cData.labs) {
        siteContent.labs.tag = cData.labs.tag || ''
        siteContent.labs.title = cData.labs.title || ''
        siteContent.labs.background_image = cData.labs.background_image || ''
        siteContent.labs.experiments = Array.isArray(cData.labs.experiments) ? JSON.parse(JSON.stringify(cData.labs.experiments)) : []
      }
      if (cData.contact) {
        Object.assign(siteContent.contact, cData.contact)
        if (Array.isArray(cData.contact.topics)) {
          topicsInput.value = cData.contact.topics.join(', ')
        }
      }
      if (cData.footer) Object.assign(siteContent.footer, cData.footer)
      if (cData.ambient) Object.assign(siteContent.ambient, cData.ambient)
    }

    projects.value = Array.isArray(pData) ? pData : []
    messages.value = Array.isArray(mData) ? mData : []

    triggerAlert('Data CMS tersinkronisasi dengan server.')
  } catch (err) {
    triggerAlert('Gagal memuat data: ' + err.message, true)
  }
}

// ---------------- SAVE NAVBAR ----------------
const saveNavbar = async () => {
  isSaving.value = true
  try {
    await api.updateSiteContent({ navbar: siteContent.navbar })
    triggerAlert('Navigasi & Header berhasil disimpan!')
    await refreshGlobalContent()
  } catch (err) {
    triggerAlert('Gagal menyimpan Navigasi: ' + err.message, true)
  } finally {
    isSaving.value = false
  }
}

const addNavbarLink = () => {
  if (!Array.isArray(siteContent.navbar.links)) siteContent.navbar.links = []
  siteContent.navbar.links.push({ label: 'Menu Baru', url: '#section' })
}

const removeNavbarLink = (idx) => {
  siteContent.navbar.links.splice(idx, 1)
}

// ---------------- SAVE HERO ----------------
const saveHero = async () => {
  isSaving.value = true
  try {
    await api.updateSiteContent({ hero: siteContent.hero })
    triggerAlert('Hero Section & Showcase Reel berhasil disimpan!')
    await refreshGlobalContent()
  } catch (err) {
    triggerAlert('Gagal menyimpan Hero: ' + err.message, true)
  } finally {
    isSaving.value = false
  }
}

// ---------------- SAVE TICKER ----------------
const saveTicker = async () => {
  isSaving.value = true
  try {
    await api.updateSiteContent({ ticker: siteContent.ticker })
    triggerAlert('Daftar teks Ticker Marquee berhasil disimpan!')
    await refreshGlobalContent()
  } catch (err) {
    triggerAlert('Gagal menyimpan Ticker: ' + err.message, true)
  } finally {
    isSaving.value = false
  }
}

const addTickerItem = () => {
  siteContent.ticker.push('rekayasa sistem')
}

const removeTickerItem = (idx) => {
  siteContent.ticker.splice(idx, 1)
}

// ---------------- SAVE DISCIPLINES ----------------
const saveDisciplines = async () => {
  isSaving.value = true
  try {
    await api.updateSiteContent({ disciplines: siteContent.disciplines })
    triggerAlert('Daftar Keahlian & Disiplin berhasil disimpan!')
    await refreshGlobalContent()
  } catch (err) {
    triggerAlert('Gagal menyimpan Keahlian: ' + err.message, true)
  } finally {
    isSaving.value = false
  }
}

const openDisciplineModal = (item = null, idx = null) => {
  if (item) {
    editingDisciplineIndex.value = idx
    disciplineForm.title = item.title
    disciplineForm.benchmark = item.benchmark
    disciplineForm.description = item.description
    disciplineForm.stack = Array.isArray(item.stack) ? [...item.stack] : []
    disciplineForm.image = item.image || '/src/assets/dev_workspace.jpg'
    disciplineStackInput.value = disciplineForm.stack.join(', ')
  } else {
    editingDisciplineIndex.value = null
    disciplineForm.title = ''
    disciplineForm.benchmark = ''
    disciplineForm.description = ''
    disciplineForm.stack = []
    disciplineForm.image = '/src/assets/dev_workspace.jpg'
    disciplineStackInput.value = ''
  }
  showDisciplineModal.value = true
}

const saveDisciplineItem = async () => {
  disciplineForm.stack = disciplineStackInput.value
    .split(',')
    .map(s => s.trim())
    .filter(Boolean)

  if (editingDisciplineIndex.value !== null) {
    siteContent.disciplines[editingDisciplineIndex.value] = {
      ...siteContent.disciplines[editingDisciplineIndex.value],
      ...JSON.parse(JSON.stringify(disciplineForm))
    }
  } else {
    const newId = siteContent.disciplines.length + 1
    siteContent.disciplines.push({
      id: newId,
      ...JSON.parse(JSON.stringify(disciplineForm))
    })
  }

  showDisciplineModal.value = false
  await saveDisciplines()
}

const removeDiscipline = async (idx) => {
  if (!confirm(`Hapus disiplin "${siteContent.disciplines[idx].title}"?`)) return
  siteContent.disciplines.splice(idx, 1)
  await saveDisciplines()
}

// ---------------- SAVE MANIFESTO ----------------
const saveManifesto = async () => {
  isSaving.value = true
  try {
    await api.updateSiteContent({ manifesto: siteContent.manifesto })
    triggerAlert('Manifesto & Pernyataan Pengembang berhasil disimpan!')
    await refreshGlobalContent()
  } catch (err) {
    triggerAlert('Gagal menyimpan Manifesto: ' + err.message, true)
  } finally {
    isSaving.value = false
  }
}

// ---------------- SAVE LABS ----------------
const saveLabs = async () => {
  isSaving.value = true
  try {
    await api.updateSiteContent({ labs: siteContent.labs })
    triggerAlert('Vadikara Labs & Prototipe berhasil disimpan!')
    await refreshGlobalContent()
  } catch (err) {
    triggerAlert('Gagal menyimpan Labs: ' + err.message, true)
  } finally {
    isSaving.value = false
  }
}

const openLabModal = (item = null, idx = null) => {
  if (item) {
    editingLabIndex.value = idx
    labForm.name = item.name
    labForm.status = item.status
    labForm.desc = item.desc
    labForm.tags = Array.isArray(item.tags) ? [...item.tags] : []
    labTagsInput.value = labForm.tags.join(', ')
  } else {
    editingLabIndex.value = null
    labForm.name = ''
    labForm.status = 'RUNTIME AKTIF'
    labForm.desc = ''
    labForm.tags = []
    labTagsInput.value = ''
  }
  showLabModal.value = true
}

const saveLabExperimentItem = async () => {
  labForm.tags = labTagsInput.value
    .split(',')
    .map(s => s.trim())
    .filter(Boolean)

  if (editingLabIndex.value !== null) {
    siteContent.labs.experiments[editingLabIndex.value] = {
      ...siteContent.labs.experiments[editingLabIndex.value],
      ...JSON.parse(JSON.stringify(labForm))
    }
  } else {
    const newId = (siteContent.labs.experiments?.length || 0) + 1
    if (!siteContent.labs.experiments) siteContent.labs.experiments = []
    siteContent.labs.experiments.push({
      id: newId,
      ...JSON.parse(JSON.stringify(labForm))
    })
  }

  showLabModal.value = false
  await saveLabs()
}

const removeLabExperiment = async (idx) => {
  if (!confirm(`Hapus prototipe "${siteContent.labs.experiments[idx].name}"?`)) return
  siteContent.labs.experiments.splice(idx, 1)
  await saveLabs()
}

// ---------------- SAVE CONTACT ----------------
const saveContact = async () => {
  isSaving.value = true
  try {
    siteContent.contact.topics = topicsInput.value
      .split(',')
      .map(s => s.trim())
      .filter(Boolean)

    await api.updateSiteContent({ contact: siteContent.contact })
    triggerAlert('Kordinat Kontak & Media Sosial berhasil disimpan!')
    await refreshGlobalContent()
  } catch (err) {
    triggerAlert('Gagal menyimpan Kontak: ' + err.message, true)
  } finally {
    isSaving.value = false
  }
}

// ---------------- SAVE FOOTER & AMBIENT ----------------
const saveFooterAndAmbient = async () => {
  isSaving.value = true
  try {
    await api.updateSiteContent({
      footer: siteContent.footer,
      ambient: siteContent.ambient
    })
    triggerAlert('Footer & Pemutar Audio Ambient berhasil disimpan!')
    await refreshGlobalContent()
  } catch (err) {
    triggerAlert('Gagal menyimpan Footer & Audio: ' + err.message, true)
  } finally {
    isSaving.value = false
  }
}

// ---------------- PROJECT CRUD ----------------
const openProjectModal = (proj = null) => {
  if (proj) {
    editingProject.value = { ...proj }
    projectStackInput.value = Array.isArray(proj.tech_stack) ? proj.tech_stack.join(', ') : ''
  } else {
    editingProject.value = {
      title: '',
      slug: '',
      description: '',
      full_overview: '',
      category: 'SISTEM // GOLANG',
      metrics: '185,000 req/s • p99 < 1.1ms',
      live_url: '',
      github_url: 'https://github.com/vadikara',
      image_url: '/src/assets/creative_poster.jpg',
      featured: true,
      architecture_summary: ''
    }
    projectStackInput.value = 'Golang, gRPC, Kafka, Redis'
  }
  showProjectModal.value = true
}

const saveProjectItem = async () => {
  try {
    const stack = projectStackInput.value
      .split(',')
      .map(s => s.trim())
      .filter(Boolean)

    const payload = {
      ...editingProject.value,
      tech_stack: stack
    }

    if (!payload.slug) {
      payload.slug = payload.title.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/(^-|-$)/g, '')
    }

    if (payload.id) {
      await api.updateProject(payload.id, payload)
      triggerAlert('Karya proyek berhasil diperbarui!')
    } else {
      await api.createProject(payload)
      triggerAlert('Karya proyek baru berhasil ditambahkan!')
    }

    showProjectModal.value = false
    const updatedList = await api.getProjects()
    projects.value = updatedList
  } catch (err) {
    triggerAlert('Gagal menyimpan proyek: ' + err.message, true)
  }
}

const deleteProject = async (id) => {
  if (!confirm('Hapus karya proyek ini secara permanen?')) return
  try {
    await api.deleteProject(id)
    triggerAlert('Proyek telah dihapus.')
    projects.value = await api.getProjects()
  } catch (err) {
    triggerAlert('Gagal menghapus proyek: ' + err.message, true)
  }
}

// ---------------- MESSAGES INBOX ----------------
const markMessageReviewed = async (id) => {
  try {
    await api.updateMessageStatus(id, 'reviewed')
    triggerAlert('Status pesan ditandai telah ditinjau.')
    messages.value = await api.getAdminMessages()
  } catch (err) {
    triggerAlert(err.message, true)
  }
}

const deleteMessage = async (id) => {
  if (!confirm('Hapus pesan ini dari kotak masuk?')) return
  try {
    await api.deleteMessage(id)
    triggerAlert('Pesan dihapus.')
    messages.value = await api.getAdminMessages()
  } catch (err) {
    triggerAlert(err.message, true)
  }
}

const formatDate = (d) => {
  if (!d) return '-'
  return new Date(d).toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const handleLogout = () => {
  localStorage.removeItem('vk_admin_token')
  localStorage.removeItem('vk_admin_user')
  router.push('/admin/login')
}

onMounted(() => {
  loadAllData()
})

onUnmounted(() => {
  stopAudioPreview()
})
</script>

<style scoped>
.admin-wrapper {
  min-height: 100vh;
  width: 100%;
  background: #080808;
  color: #f4f4f5;
  display: flex;
  flex-direction: column;
}

.admin-topbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 36px;
  background: rgba(14, 14, 16, 0.95);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  position: sticky;
  top: 0;
  z-index: 100;
}

.topbar-brand {
  display: flex;
  align-items: center;
  gap: 16px;
}

.brand-link {
  display: flex;
  align-items: center;
  gap: 10px;
  text-decoration: none;
  color: #ffffff;
}

.brand-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #00f2fe;
  box-shadow: 0 0 10px #00f2fe;
}

.brand-title {
  font-size: 0.88rem;
  font-weight: 700;
  letter-spacing: 0.08em;
}

.topbar-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

.admin-main {
  display: grid;
  grid-template-columns: 270px 1fr;
  flex: 1;
}

/* Sidebar */
.admin-sidebar {
  background: rgba(12, 12, 14, 0.95);
  border-right: 1px solid rgba(255, 255, 255, 0.06);
  padding: 24px 16px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  height: calc(100vh - 65px);
  position: sticky;
  top: 65px;
}

.sidebar-user {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 14px;
  border-radius: var(--radius-sm);
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.06);
}

.user-avatar {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: linear-gradient(135deg, #27272a, #18181b);
  border: 1px solid rgba(255, 255, 255, 0.15);
  color: #00f2fe;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 1rem;
}

.user-meta {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.user-name {
  font-size: 0.88rem;
  font-weight: 700;
  color: #ffffff;
}

.user-role {
  font-size: 0.65rem;
  color: #a1a1aa;
}

.sidebar-nav {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
  overflow-y: auto;
}

.nav-tab-btn {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 11px 14px;
  border-radius: var(--radius-sm);
  background: transparent;
  border: 1px solid transparent;
  color: #a1a1aa;
  font-size: 0.78rem;
  cursor: pointer;
  text-align: left;
  transition: all 0.2s ease;
  width: 100%;
}

.nav-tab-btn:hover {
  background: rgba(255, 255, 255, 0.04);
  color: #ffffff;
}

.nav-tab-btn.active {
  background: rgba(0, 242, 254, 0.08);
  border-color: rgba(0, 242, 254, 0.25);
  color: #00f2fe;
  font-weight: 700;
}

.tab-icon {
  font-size: 1rem;
}

.tab-label {
  flex: 1;
}

.tab-badge {
  background: #ef4444;
  color: #ffffff;
  font-size: 0.62rem;
  font-weight: 700;
  padding: 2px 6px;
  border-radius: 10px;
}

.sidebar-footer-info {
  font-size: 0.68rem;
  padding-top: 14px;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}

.engine-badge {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #71717a;
}

.dot.green-pulse {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #22c55e;
  box-shadow: 0 0 8px #22c55e;
}

/* Content Area */
.admin-content {
  padding: 36px 44px;
  overflow-y: auto;
  max-width: 1350px;
}

.tab-pane {
  display: flex;
  flex-direction: column;
  gap: 26px;
}

.pane-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding-bottom: 22px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.pane-title {
  font-size: 1.85rem;
  font-weight: 800;
  color: #ffffff;
  margin: 0 0 6px;
  letter-spacing: -0.03em;
}

.pane-desc {
  font-size: 0.88rem;
  color: #a1a1aa;
  margin: 0;
}

.pane-actions {
  display: flex;
  gap: 12px;
}

.subheading {
  font-size: 0.78rem;
  color: #a1a1aa;
  letter-spacing: 0.08em;
  margin: 0 0 14px;
}

/* Stats Cards */
.stats-cards-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.stat-card {
  padding: 24px;
  border-radius: var(--radius-md);
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.08);
  cursor: pointer;
  transition: all 0.25s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
  border-color: rgba(0, 242, 254, 0.4);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
}

.stat-label {
  font-size: 0.68rem;
  color: #71717a;
  letter-spacing: 0.08em;
}

.stat-val {
  font-size: 2.6rem;
  font-weight: 800;
  margin: 6px 0;
  color: #ffffff;
}

.stat-sub {
  font-size: 0.76rem;
  color: #a1a1aa;
}

/* Overview Details */
.overview-details-grid {
  display: grid;
  grid-template-columns: 1.25fr 1fr;
  gap: 22px;
}

.overview-box {
  padding: 26px;
  border-radius: var(--radius-md);
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.box-title {
  font-size: 1.1rem;
  font-weight: 700;
  margin: 0 0 18px;
  color: #ffffff;
}

.box-header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 18px;
}

.module-status-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.module-status-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 13px 16px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all 0.2s ease;
}

.module-status-row:hover {
  background: rgba(255, 255, 255, 0.06);
  border-color: rgba(0, 242, 254, 0.3);
}

/* Content Form Card */
.content-form-card {
  padding: 32px;
  border-radius: var(--radius-md);
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.08);
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 20px;
}

.form-label {
  font-size: 0.72rem;
  font-weight: 700;
  color: #a1a1aa;
  letter-spacing: 0.08em;
}

.form-input, .form-textarea, .form-select {
  width: 100%;
  padding: 13px 16px;
  background: #09090b;
  border: 1px solid rgba(255, 255, 255, 0.14);
  border-radius: var(--radius-sm);
  color: #ffffff;
  font-size: 0.9rem;
  transition: all 0.2s ease;
}

.form-input:focus, .form-textarea:focus, .form-select:focus {
  outline: none;
  border-color: #00f2fe;
  background: #0d0d10;
  box-shadow: 0 0 0 3px rgba(0, 242, 254, 0.15);
}

.headline-input {
  font-size: 1.1rem;
  font-weight: 600;
  line-height: 1.45;
}

.field-hint {
  font-size: 0.7rem;
  color: #71717a;
}

/* Image Manager Box */
.image-manager-box {
  display: grid;
  grid-template-columns: 240px 1fr;
  gap: 22px;
  padding: 18px;
  background: #09090b;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-sm);
}

.current-image-preview {
  position: relative;
  border-radius: var(--radius-sm);
  overflow: hidden;
  height: 160px;
  background: #000000;
}

.preview-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.preview-badge {
  position: absolute;
  bottom: 8px;
  left: 8px;
  background: rgba(0, 0, 0, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.2);
  color: #00f2fe;
  padding: 3px 8px;
  font-size: 0.62rem;
  border-radius: 4px;
}

.image-options-col {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.upload-option-row {
  display: flex;
  align-items: center;
  gap: 14px;
}

.upload-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  background: rgba(255, 255, 255, 0.08);
  border: 1px dashed rgba(255, 255, 255, 0.3);
  color: #ffffff;
  border-radius: var(--radius-sm);
  font-size: 0.76rem;
  cursor: pointer;
  transition: all 0.2s ease;
}

.upload-btn:hover {
  background: rgba(255, 255, 255, 0.15);
  border-color: #00f2fe;
  color: #00f2fe;
}

/* Preset Picker */
.preset-picker-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 10px;
}

.preset-thumb-card {
  border-radius: var(--radius-sm);
  overflow: hidden;
  border: 2px solid rgba(255, 255, 255, 0.1);
  background: #000000;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  flex-direction: column;
}

.preset-thumb-card:hover {
  border-color: rgba(255, 255, 255, 0.4);
}

.preset-thumb-card.is-selected {
  border-color: #00f2fe;
  box-shadow: 0 0 14px rgba(0, 242, 254, 0.3);
}

.preset-img {
  width: 100%;
  height: 60px;
  object-fit: cover;
}

.preset-label {
  padding: 5px 8px;
  font-size: 0.6rem;
  color: #a1a1aa;
  background: #111113;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.custom-url-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* Ticker Marquee Tab */
.ticker-preview-box {
  padding: 18px;
  background: #09090b;
  border-radius: var(--radius-sm);
  border: 1px dashed rgba(255, 255, 255, 0.15);
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.ticker-live-demo {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  align-items: center;
}

.live-item {
  font-size: 1.15rem;
  font-weight: 700;
  color: #ffffff;
}

.dot-sep {
  color: #3f3f46;
  margin-left: 6px;
}

.ticker-edit-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.ticker-edit-row {
  display: flex;
  align-items: center;
  gap: 14px;
}

.item-number {
  font-size: 0.76rem;
  color: #71717a;
  width: 36px;
}

.ticker-text-input {
  font-weight: 600;
  font-size: 0.95rem;
}

/* Links Manage List */
.links-manage-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.link-manage-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* Disciplines Grid */
.disciplines-cards-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  gap: 22px;
}

.discipline-item-card {
  border-radius: var(--radius-md);
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.08);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  padding: 0 0 20px;
  gap: 14px;
}

.card-thumb-header {
  position: relative;
  width: 100%;
  height: 150px;
  background: #000000;
}

.card-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.card-header-overlay {
  position: absolute;
  top: 12px;
  left: 12px;
  right: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-id {
  background: rgba(0, 0, 0, 0.7);
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.7rem;
  color: #ffffff;
}

.pill-metric {
  padding: 4px 10px;
  background: rgba(34, 197, 94, 0.2);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(34, 197, 94, 0.4);
  color: #22c55e;
  border-radius: 20px;
  font-size: 0.7rem;
  font-weight: 700;
}

.discipline-item-card .card-title,
.discipline-item-card .card-desc,
.discipline-item-card .card-stack-row,
.discipline-item-card .card-footer-actions {
  padding: 0 20px;
}

.card-title {
  font-size: 1.2rem;
  font-weight: 700;
  color: #ffffff;
  margin: 0;
}

.card-desc {
  font-size: 0.84rem;
  color: #a1a1aa;
  line-height: 1.55;
  margin: 0;
  flex: 1;
}

.card-stack-row {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.stack-tag {
  font-size: 0.68rem;
  padding: 3px 8px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.06);
  color: #d4d4d8;
}

.card-footer-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  padding-top: 14px;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}

/* Labs Experiments */
.experiments-admin-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 18px;
}

.experiment-admin-card {
  padding: 20px;
  border-radius: var(--radius-sm);
  background: #09090b;
  border: 1px solid rgba(255, 255, 255, 0.08);
  display: flex;
  flex-direction: column;
  gap: 12px;
}

/* Data Table */
.projects-table-card {
  border-radius: var(--radius-md);
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.08);
  overflow-x: auto;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  text-align: left;
}

.data-table th {
  padding: 14px 20px;
  font-size: 0.7rem;
  color: #71717a;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(0, 0, 0, 0.25);
}

.data-table td {
  padding: 16px 20px;
  font-size: 0.86rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.04);
  vertical-align: middle;
}

.table-proj-thumb {
  width: 70px;
  height: 48px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.table-desc-preview {
  max-width: 260px;
  color: #a1a1aa;
  font-size: 0.78rem;
  margin: 0;
  line-height: 1.4;
}

.table-sub {
  font-size: 0.72rem;
  color: #71717a;
}

.table-link {
  color: #00f2fe;
  text-decoration: underline;
  font-size: 0.76rem;
}

.table-actions {
  display: flex;
  gap: 6px;
}

.link-pill-group {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.message-content-text {
  max-width: 380px;
  color: #f4f4f5;
  white-space: pre-wrap;
}

/* Modal Image Selector */
.modal-image-selector {
  display: grid;
  grid-template-columns: 140px 1fr;
  gap: 16px;
  padding: 12px;
  background: #09090b;
  border-radius: var(--radius-sm);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.thumb-preview {
  height: 90px;
  border-radius: 4px;
  overflow: hidden;
  background: #000000;
}

.modal-thumb-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.selector-options {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

/* Buttons */
.btn {
  padding: 9px 18px;
  border-radius: var(--radius-sm);
  font-size: 0.82rem;
  font-weight: 600;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  text-decoration: none;
  transition: all 0.2s ease;
  border: 1px solid transparent;
}

.btn-primary {
  background: #ffffff;
  color: #000000;
}

.btn-primary:hover:not(:disabled) {
  background: #00f2fe;
  box-shadow: 0 4px 16px rgba(0, 242, 254, 0.3);
}

.btn-secondary {
  background: rgba(255, 255, 255, 0.06);
  color: #ffffff;
  border-color: rgba(255, 255, 255, 0.12);
}

.btn-secondary:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.12);
  border-color: rgba(255, 255, 255, 0.25);
}

.btn-danger {
  background: rgba(239, 68, 68, 0.12);
  color: #ef4444;
  border-color: rgba(239, 68, 68, 0.25);
}

.btn-danger:hover {
  background: #ef4444;
  color: #ffffff;
}

.btn-sm {
  padding: 7px 14px;
  font-size: 0.76rem;
}

.btn-xs {
  padding: 4px 10px;
  font-size: 0.72rem;
}

.btn-full {
  width: 100%;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Badges */
.badge {
  display: inline-block;
  padding: 3px 8px;
  border-radius: 4px;
  font-size: 0.65rem;
  font-weight: 700;
  letter-spacing: 0.04em;
}

.badge-cyan {
  background: rgba(0, 242, 254, 0.1);
  color: #00f2fe;
  border: 1px solid rgba(0, 242, 254, 0.25);
}

.badge-live {
  background: rgba(34, 197, 94, 0.1);
  color: #22c55e;
  border: 1px solid rgba(34, 197, 94, 0.25);
}

.badge-amber {
  background: rgba(245, 158, 11, 0.1);
  color: #f59e0b;
  border: 1px solid rgba(245, 158, 11, 0.25);
}

.green-text { color: #22c55e; }
.cyan-text { color: #00f2fe; }
.text-dim { color: #71717a; }

/* Alerts */
.alert {
  padding: 14px 20px;
  border-radius: var(--radius-sm);
  font-size: 0.84rem;
  display: flex;
  align-items: center;
  gap: 12px;
  position: fixed;
  bottom: 24px;
  right: 24px;
  z-index: 1000;
  box-shadow: 0 16px 36px rgba(0, 0, 0, 0.8);
}

.alert-success {
  background: #14532d;
  color: #bbf7d0;
  border: 1px solid #22c55e;
}

.alert-error {
  background: #7f1d1d;
  color: #fecaca;
  border: 1px solid #ef4444;
}

/* Modals */
.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.85);
  backdrop-filter: blur(14px);
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}

.modal-box {
  width: 100%;
  max-width: 680px;
  max-height: 90vh;
  overflow-y: auto;
  background: #111113;
  border: 1px solid rgba(255, 255, 255, 0.15);
  border-radius: var(--radius-md);
  padding: 32px;
  box-shadow: 0 30px 90px rgba(0, 0, 0, 0.8);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 22px;
  padding-bottom: 14px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.modal-header h3 {
  font-size: 1.3rem;
  font-weight: 700;
  margin: 0;
  color: #ffffff;
}

.modal-close {
  background: transparent;
  border: none;
  font-size: 1.6rem;
  color: #a1a1aa;
  cursor: pointer;
  line-height: 1;
}

.modal-close:hover {
  color: #ffffff;
}

.modal-box form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

@media (max-width: 1024px) {
  .admin-main {
    grid-template-columns: 1fr;
  }
  .admin-sidebar {
    position: static;
    height: auto;
  }
  .stats-cards-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  .overview-details-grid {
    grid-template-columns: 1fr;
  }
  .preset-picker-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

/* ===================================================
   NEW: VIDEO PREVIEW & ADVANCED AUDIO MANAGER STYLES
=================================================== */
.label-with-badge {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.reel-quick-actions {
  display: flex;
  gap: 10px;
  margin-top: 8px;
  align-items: center;
}

.file-upload-label {
  cursor: pointer;
  display: inline-flex;
  align-items: center;
}

.video-preview-group {
  margin-top: 14px;
  padding: 16px;
  background: rgba(8, 8, 8, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: var(--radius-sm);
}

.video-preview-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.video-preview-player-box {
  width: 100%;
  aspect-ratio: 16 / 9;
  max-width: 680px;
  background: #000000;
  border-radius: var(--radius-sm);
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.12);
  margin-bottom: 10px;
}

.cms-video-frame,
.cms-video-player {
  width: 100%;
  height: 100%;
  border: none;
  display: block;
  object-fit: contain;
}

/* Audio Suite Styles */
.audio-manager-card {
  margin-top: 30px;
  padding: 24px;
  background: rgba(14, 14, 16, 0.85);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-md);
}

.audio-card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  padding-bottom: 16px;
}

/* Toggle Switch */
.switch-container {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  user-select: none;
  font-size: 0.76rem;
}

.switch-container input {
  display: none;
}

.switch-slider {
  width: 44px;
  height: 24px;
  background: #27272a;
  border-radius: 20px;
  position: relative;
  transition: background 0.25s ease;
}

.switch-slider::after {
  content: '';
  position: absolute;
  top: 3px;
  left: 3px;
  width: 18px;
  height: 18px;
  background: #ffffff;
  border-radius: 50%;
  transition: transform 0.25s ease;
}

.switch-container input:checked + .switch-slider {
  background: #00f2fe;
}

.switch-container input:checked + .switch-slider::after {
  transform: translateX(20px);
}

.switch-label {
  color: #e4e4e7;
  font-weight: 600;
  letter-spacing: 0.05em;
}

/* Audio Source Manager Grid */
.audio-source-manager {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  margin-top: 12px;
}

@media (max-width: 900px) {
  .audio-source-manager {
    grid-template-columns: 1fr;
  }
}

.audio-mode-box {
  padding: 16px;
  background: rgba(8, 8, 8, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: var(--radius-sm);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.mode-header {
  font-size: 0.74rem;
  color: #22d3ee;
  letter-spacing: 0.04em;
  margin-bottom: 8px;
}

.mode-desc {
  font-size: 0.78rem;
  color: #a1a1aa;
  line-height: 1.45;
  margin-bottom: 14px;
  flex-grow: 1;
}

.upload-btn-row {
  display: flex;
  flex-direction: column;
  gap: 8px;
  align-items: flex-start;
}

/* CMS Audio Test Dock */
.cms-audio-preview-dock {
  margin-top: 20px;
  padding: 14px 20px;
  background: rgba(0, 0, 0, 0.75);
  border: 1px solid rgba(34, 211, 238, 0.3);
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
}

.preview-dock-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.audio-test-btn {
  width: 42px;
  height: 42px;
  border-radius: 50%;
  background: #ffffff;
  color: #000000;
  border: none;
  font-size: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.2s, background 0.2s;
  flex-shrink: 0;
}

.audio-test-btn:hover {
  transform: scale(1.08);
  background: #22d3ee;
}

.audio-test-btn.is-playing {
  background: #22d3ee;
  box-shadow: 0 0 15px rgba(34, 211, 238, 0.6);
}

.preview-track-meta {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.track-title-preview {
  font-size: 0.82rem;
  color: #ffffff;
  letter-spacing: 0.05em;
}

.track-sub-preview {
  font-size: 0.72rem;
  color: #22d3ee;
}

.track-status-pill {
  font-size: 0.66rem;
  margin-top: 2px;
}

.preview-dock-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

/* Animated spectrum in preview dock */
.preview-spectrum {
  display: flex;
  align-items: flex-end;
  gap: 4px;
  height: 20px;
}

.preview-spectrum .bar {
  width: 3px;
  height: 4px;
  background: #22d3ee;
  border-radius: 1px;
  transition: height 0.2s ease;
}

.preview-spectrum.spectrum-active .b1 { animation: specBounce 0.6s infinite ease-in-out alternate; }
.preview-spectrum.spectrum-active .b2 { animation: specBounce 0.45s infinite ease-in-out alternate 0.1s; }
.preview-spectrum.spectrum-active .b3 { animation: specBounce 0.7s infinite ease-in-out alternate 0.2s; }
.preview-spectrum.spectrum-active .b4 { animation: specBounce 0.5s infinite ease-in-out alternate 0.15s; }
.preview-spectrum.spectrum-active .b5 { animation: specBounce 0.65s infinite ease-in-out alternate 0.25s; }

@keyframes specBounce {
  0% { height: 4px; opacity: 0.4; }
  100% { height: 18px; opacity: 1; }
}
</style>
