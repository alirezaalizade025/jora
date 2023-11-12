import React, { useEffect, useState } from 'react';
import { Button, TextField, Typography, Container, Box, Paper } from '@mui/material';
import { cookieGetter, cookieSetter } from 'utils/cookieUtils';
import { useRouter } from 'next/navigation';
import AuthApi from 'src/Api/Auth';
import { appVersion } from 'utils/consts';

import Sidebar from 'src/Components/Sidebar';

const Dashboard = () => {
  return (
    <Container component="main" maxWidth="xs">
      <Paper elevation={3}>
        <Sidebar />
      </Paper>
    </Container>
  );
};

export default Dashboard;
