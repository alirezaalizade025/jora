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
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import IconButton from '@mui/material/IconButton';
import { useRouter } from 'next/navigation';

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

const Sidebar = ({ open, onClose }) => {
  const router = useRouter();

  const handleItemClick = (link) => {
    router.push(link);
    onClose(); // Close the sidebar when an item is clicked
  };

  return (
    <Drawer anchor="right" open={open} onClose={onClose}>
      <Box sx={{ width: 250 }} role="presentation">
        <IconButton onClick={onClose} style={{ textAlign: 'left', marginLeft: 'auto' }}>
          <ChevronRightIcon />
        </IconButton>
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
    </Drawer>
  );
};

export default Sidebar;