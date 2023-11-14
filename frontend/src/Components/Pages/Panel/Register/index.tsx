import React, { useState } from 'react';
import { Button, TextField, Typography, Container, Box, Paper } from '@mui/material';
import { cookieSetter } from 'utils/cookieUtils';
import AuthApi from 'src/Api/Auth';
import { appVersion } from 'utils/consts';
import Link from 'next/link';
import MuiLink from '@mui/material/Link';

const Register = () => {
  const [title, setTitle] = useState('');
  const [phone, setPhone] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();
    setIsLoading(true);
    AuthApi.register({ confirmPassword, password, phone, title })
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
          <Typography variant="h6" align="center" gutterBottom>
            ثبت شرکت جدید
          </Typography>

          <form onSubmit={handleSubmit}>
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              id="username"
              label="نام شرکت"
              name="username"
              autoComplete="username"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              name="password"
              label="تلفن"
              id="password"
              autoComplete="current-password"
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
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              name="password"
              label="تکرار رمز عبور"
              type="password"
              id="password"
              autoComplete="current-password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              disabled={isLoading}>
              {isLoading ? 'در حال ثبت نام...' : 'ثبت نام'}
            </Button>
          </form>
          <Typography variant="body2" align="center">
            {appVersion}
          </Typography>
          <Typography variant="body2" align="center">
            اگر قبلا ثبت نام کرده اید برای ورود
            <Link href={'/panel/login'}>
              <MuiLink>اینجا</MuiLink>
            </Link>
            کلیک کنید
          </Typography>
        </Box>
      </Paper>
    </Container>
  );
};

export default Register;
