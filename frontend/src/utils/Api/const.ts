export const NETWORK_ERROR_MESSAGE = new Set([
  'Network Error',
  'Failed to fetch', // Chrome
  'NetworkError when attempting to fetch resource.', // Firefox
  'The Internet connection appears to be offline.', // Safari
  'Network request failed', // `cross-fetch`
]);
