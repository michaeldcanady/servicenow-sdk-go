// Minimal Plausible custom-event client. Safe to call from anywhere: when
// the Plausible script is absent (local dev, PR previews) events queue
// in-memory and are dropped on page unload, so nothing is ever sent.

type PlausibleOptions = {props?: Record<string, string>};

type PlausibleFn = ((event: string, options?: PlausibleOptions) => void) & {
  q?: unknown[];
};

declare global {
  interface Window {
    plausible?: PlausibleFn;
  }
}

export function trackEvent(name: string, props?: Record<string, string>): void {
  if (typeof window === 'undefined') {
    return;
  }
  if (!window.plausible) {
    // Queue shim per https://plausible.io/docs/custom-events — the real
    // script drains the queue once it loads, then replaces window.plausible.
    const queue: unknown[] = [];
    const shim = ((event: string, options?: PlausibleOptions) => {
      queue.push([event, options]);
    }) as PlausibleFn;
    shim.q = queue;
    window.plausible = shim;
  }
  window.plausible(name, props ? {props} : undefined);
}
