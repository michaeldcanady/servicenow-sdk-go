// Tracks clicks on any code-block copy button as a Plausible "Copy Code"
// event. The button has no stable class (its styling comes from a CSS
// module), so this keys off its aria-label, which is stable English —
// the site is single-locale.
import {trackEvent} from '../utils/analytics';

if (typeof document !== 'undefined') {
  document.addEventListener('click', (event) => {
    const target = event.target as Element | null;
    if (target?.closest?.('button[aria-label="Copy code to clipboard"]')) {
      trackEvent('Copy Code', {page: window.location.pathname});
    }
  });
}
