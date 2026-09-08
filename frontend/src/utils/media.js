/**
 * Media Helper Utilities for Vadikara Studio & CMS
 * Handles YouTube, Vimeo, direct MP4/WebM videos, and audio streams.
 */

/**
 * Parses any video URL (YouTube, Vimeo, direct MP4/WebM, or local file)
 * and returns embeddable URLs and metadata.
 */
export function parseVideoUrl(url) {
  if (!url || typeof url !== 'string') {
    return {
      type: 'video',
      id: null,
      embedUrl: null,
      previewEmbedUrl: null,
      directUrl: '/Scene.mp4',
      isYoutube: false,
      isVimeo: false,
      label: 'Video Bawaan (/Scene.mp4)'
    }
  }

  const cleanUrl = url.trim()
  if (!cleanUrl) {
    return {
      type: 'video',
      id: null,
      embedUrl: null,
      previewEmbedUrl: null,
      directUrl: '/Scene.mp4',
      isYoutube: false,
      isVimeo: false,
      label: 'Video Bawaan (/Scene.mp4)'
    }
  }

  // YouTube match:
  // Supports youtu.be/ID, youtube.com/watch?v=ID, youtube.com/embed/ID, youtube.com/shorts/ID
  const ytMatch = cleanUrl.match(/(?:youtu\.be\/|youtube\.com\/(?:embed\/|v\/|shorts\/|watch\?v=|watch\?.+&v=))([\w-]{11})/)
  if (ytMatch && ytMatch[1]) {
    const videoId = ytMatch[1]
    return {
      type: 'youtube',
      id: videoId,
      embedUrl: `https://www.youtube.com/embed/${videoId}?autoplay=1&rel=0&enablejsapi=1`,
      previewEmbedUrl: `https://www.youtube.com/embed/${videoId}?rel=0`,
      directUrl: null,
      isYoutube: true,
      isVimeo: false,
      label: `YouTube Video (ID: ${videoId})`
    }
  }

  // Vimeo match:
  // Supports vimeo.com/123456789
  const vimeoMatch = cleanUrl.match(/vimeo\.com\/(?:channels\/(?:\w+\/)?|groups\/([^\/]*)\/videos\/|album\/(\d+)\/video\/|video\/|)(\d+)/)
  if (vimeoMatch && vimeoMatch[3]) {
    const vimeoId = vimeoMatch[3]
    return {
      type: 'vimeo',
      id: vimeoId,
      embedUrl: `https://player.vimeo.com/video/${vimeoId}?autoplay=1`,
      previewEmbedUrl: `https://player.vimeo.com/video/${vimeoId}`,
      directUrl: null,
      isYoutube: false,
      isVimeo: true,
      label: `Vimeo Video (ID: ${vimeoId})`
    }
  }

  // Direct video file (.mp4, .webm, data:video, etc.)
  const isDataUrl = cleanUrl.startsWith('data:video')
  const isLocalAsset = cleanUrl.startsWith('/')
  return {
    type: 'video',
    id: null,
    embedUrl: null,
    previewEmbedUrl: null,
    directUrl: cleanUrl,
    isYoutube: false,
    isVimeo: false,
    label: isDataUrl ? 'Berkas Video Terunggah (Lokal)' : (isLocalAsset ? `Berkas Lokal (${cleanUrl})` : 'URL Video Langsung (MP4/WebM)')
  }
}
