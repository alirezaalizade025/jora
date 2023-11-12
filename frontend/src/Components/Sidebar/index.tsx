import * as React from 'react';
import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import DashboardIcon from '@mui/icons-material/Dashboard';
import PeopleAltIcon from '@mui/icons-material/PeopleAlt';
import SettingsIcon from '@mui/icons-material/Settings';

const Sidebar = () => {
  const listItems = [
    {
      title: 'داشبورد',
      icon: <DashboardIcon />,
      link: '/panel',
    },
    {
      title: 'پرسنل',
      icon: <PeopleAltIcon />,
      link: '/panel/users',
    },
    {
      title: 'تنظیمات',
      icon: <SettingsIcon />,
      link: '/panel/settings',
    },
  ];

  const handleItemClick = (link) => {
    window.location.replace(link);
  };

  const list = () => (
    <Box sx={{ width: 250 }} role="presentation">
      <List>
        {listItems.map((item) => (
          <ListItem key={item.title} disablePadding>
            <ListItemButton onClick={() => handleItemClick(item.link)}>
              <ListItemText primary={item.title} style={{ textAlign: 'right' }} />
              <ListItemIcon>{item.icon}</ListItemIcon>
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Box>
  );

  return (
    <Drawer anchor="right" open={true}>
      {list()}
    </Drawer>
  );
};

export default Sidebar;
