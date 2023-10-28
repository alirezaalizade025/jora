import Cookies from 'universal-cookie';
import { COOKIES, COOKIE_LONG_TIME, ONE_DAY_SECONDS, cookieDomain } from './consts';

type CookieMaxAgeType = 'oneDay' | 'season' | 'longTime' | '30min' | '15sec';

type CookieSetterType = {
  name: keyof typeof COOKIES;
  content: string | number;
  maxAge: CookieMaxAgeType;
};

const getMaxAge = (maxAge: CookieMaxAgeType) => {
  switch (maxAge) {
    case 'oneDay': {
      return ONE_DAY_SECONDS;
    }
    case 'season': {
      return null;
    }
    case 'longTime': {
      return COOKIE_LONG_TIME;
    }
    case '30min': {
      return 1000 * 60 * 30;
    }
    case '15sec': {
      return 1000 * 15;
    }
    default: {
      return null;
    }
  }
};

export const cookieSetter = ({
  name,
  content,
  maxAge,
  isShowError,
  onError,
}: CookieSetterType & {
  isShowError?: boolean;
  onError?: () => void;
}) => {
  const cookies = new Cookies();
  cookies.set(COOKIES[name], content, {
    path: '/',
    maxAge: getMaxAge(maxAge),
    domain: cookieDomain,
    secure: true,
    sameSite: 'lax',
  });
  if (isShowError) {
    setTimeout(() => {
      if (!cookies.get(COOKIES[name])) {
        console.log('COOKIESKey', COOKIES[name]);
      }
    }, 100);
  }
};

export const cookieGetter = ({ name }: { name: keyof typeof COOKIES }) => {
  const cookies = new Cookies();
  return cookies.get(COOKIES[name]);
};

export const cookieRemover = ({ name }: { name: keyof typeof COOKIES }) => {
  const cookies = new Cookies();
  cookies.remove(COOKIES[name], {
    path: '/',
    domain: cookieDomain,
    secure: true,
    sameSite: 'lax',
  });
};

export default {};
