
export const apiBaseUrl: string = process.env.NEXT_PUBLIC_API_BASE_URL;
export const appVersion: string = process.env.NEXT_PUBLIC_APP_VERSION || 'local';

export const COOKIES = {
  jwt: `jora_auth_token`,
};

export const cookieDomain = '.jora.com';
export const ONE_DAY_SECONDS = 86_400;
export const ONE_Min_SECONDS = 60;
export const COOKIE_LONG_TIME = 100_000_000;

export default {};
