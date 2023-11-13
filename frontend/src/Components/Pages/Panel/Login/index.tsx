import React, { useEffect, useState } from 'react';
import { Button, TextField, Typography, Container, Box, Paper } from '@mui/material';
import { cookieGetter, cookieSetter } from 'utils/cookieUtils';
import { useRouter } from 'next/navigation';
import AuthApi from 'src/Api/Auth';
import { appVersion } from 'utils/consts';
import Link from 'next/link';
import MuiLink from '@mui/material/Link';

const Auth = () => {
  const [phone, setPhone] = useState('');
  const [password, setPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

  useEffect(() => {
    let timer;
    if (cookieGetter({ name: 'jwt' })) {
      timer = setTimeout(() => {
        router.push('/');
      }, 1000);
    } else {
      setIsLoading(false);
    }
    return () => {
      if (timer) {
        clearTimeout(timer);
      }
    };
  }, [router]);

  const handleSubmit = (e) => {
    e.preventDefault();
    setIsLoading(true);
    AuthApi.AdminLogin({ phone, password })
      .then((res) => {
        if (res.data.jwtToken) {
          cookieSetter({ name: 'jwt', content: res.data.jwtToken, maxAge: 'oneDay' });
          window.location.replace('/panel');
        }
        setIsLoading(false);
      })
      .catch((error) => {
        setIsLoading(false);
        console.log(error);
      });
  };

  return (
    <Container component="main" maxWidth="xs">
      <Paper elevation={3}>
        <Box p={3}>
          <Typography variant="h4" align="center" gutterBottom>
            JORA
          </Typography>
          <form onSubmit={handleSubmit}>
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              id="phone"
              label="تلفن"
              name="phone"
              autoComplete="phone"
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
            />
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              name="password"
              label="رمز عبور"
              type="password"
              id="password"
              autoComplete="current-password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              disabled={isLoading}>
              {isLoading ? 'در حال بارگذاری...' : 'ورود'}
            </Button>
          </form>
          <Typography variant="body2" align="center">
            {appVersion}
          </Typography>
          <Typography variant="body2" align="center">
            برای ثبت نام شرکت جدید{' '}
            <Link href={'/panel/register'}>
              <MuiLink>اینجا</MuiLink>
            </Link>
            کلیک کنید
          </Typography>
        </Box>
      </Paper>
    </Container>
  );
};

export default Auth;
