import React from 'react';

import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import MuiDrawer from '@mui/material/Drawer';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';

export interface DrawerLink {
  text: string;
  href: string;
  external?: boolean;
  icon?: React.ReactNode;
}

export interface DrawerSection {
  links: DrawerLink[];
  title?: string;
}

export interface DrawerProps {
  sections: DrawerSection[];
}

const drawerWidth = 240;

export const Drawer = (props: DrawerProps) => {
  return (
    <MuiDrawer
      variant='permanent'
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' }
      }}
    >
      <Toolbar variant='dense' />
      <Box sx={{ overflow: 'auto' }}>
        <List>
          {props.sections.map((section: DrawerSection) => (
            <>
              {section.title && <ListSubheader>{section.title}</ListSubheader>}
              {section.links.map((link: DrawerLink) => (
                <ListItem key={link.href} disablePadding sx={{ display: 'block' }} dense>
                  <ListItemButton
                    href={link.href}
                    sx={{
                      justifyContent: 'center',
                      px: 2.5
                    }}
                  >
                    {link.icon && (
                      <ListItemIcon>
                        {link.icon}
                      </ListItemIcon>
                    )}
                    <ListItemText primary={link.text} />
                  </ListItemButton>
                </ListItem>
              ))}
            </>
          ))}
        </List>
      </Box>
    </MuiDrawer>
  );
};
