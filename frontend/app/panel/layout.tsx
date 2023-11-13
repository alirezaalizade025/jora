'use client';

import { Container, IconButton, Paper } from '@mui/material';
import Sidebar from 'src/Components/PageLayout/Sidebar';
import React, { useState } from 'react';
import MenuIcon from '@mui/icons-material/Menu';

const Layout = ({ children }) => {
  const [isSidebarOpen, setSidebarOpen] = useState(true);

  const handleToggleSidebar = () => {
    setSidebarOpen(!isSidebarOpen);
  };

  return (
    <div style={{ display: 'flex' }}>
      <IconButton
        onClick={handleToggleSidebar}
        style={{ marginRight: '16px', marginLeft: '16px', marginTop: '16px' }}>
        <MenuIcon color='info' />
      </IconButton>
      <Container component="main" maxWidth="xs">
        <Paper elevation={3} style={{ display: 'flex' }}>
          <Sidebar open={isSidebarOpen} onClose={handleToggleSidebar} />
          <div style={{ flexGrow: 1, padding: '16px' }}>{children}</div>
        </Paper>
      </Container>
    </div>
  );
};

export default Layout;
