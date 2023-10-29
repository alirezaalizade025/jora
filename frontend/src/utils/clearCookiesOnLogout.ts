import { cookieRemover } from './cookieUtils';

const clearCookiesOnLogout = () => {
  cookieRemover({ name: 'jwt' });
};

export default clearCookiesOnLogout;
