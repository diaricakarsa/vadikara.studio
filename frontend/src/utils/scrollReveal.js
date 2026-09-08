/**
 * High-Performance Scroll Reveal Engine using IntersectionObserver
 */
export function initScrollReveal() {
  if (typeof window === 'undefined' || !('IntersectionObserver' in window)) return

  const observerCallback = (entries, observer) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('is-revealed')
        observer.unobserve(entry.target)
      }
    })
  }

  const observerOptions = {
    root: null,
    rootMargin: '0px 0px -60px 0px', // triggers slightly before full view for smooth anticipation
    threshold: 0.12
  }

  const observer = new IntersectionObserver(observerCallback, observerOptions)

  const scanElements = () => {
    const targets = document.querySelectorAll('.reveal-init, .reveal-scale-init')
    targets.forEach(el => {
      if (!el.classList.contains('is-revealed')) {
        observer.observe(el)
      }
    })
  }

  // Scan immediately and after DOM updates
  scanElements()
  setTimeout(scanElements, 250)
  setTimeout(scanElements, 750)

  // Listen to mutation changes (dynamic cards, tab changes)
  const mutationObserver = new MutationObserver(() => {
    scanElements()
  })

  mutationObserver.observe(document.body, {
    childList: true,
    subtree: true
  })

  return {
    scan: scanElements
  }
}
