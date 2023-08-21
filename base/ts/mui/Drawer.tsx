/**
 * Copyright 2023 Peridot Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
