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
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import { Theme } from '@mui/material/styles';

import DashboardIcon from '@mui/icons-material/Dashboard';
import ListAltIcon from '@mui/icons-material/ListAlt';

import { Dashboard } from './Dashboard';
import { Drawer } from 'base/ts/mui/Drawer';
import { Route, Routes } from 'react-router-dom';

export const App = () => {
  return (
    <Box sx={{ display: 'flex' }}>
      <AppBar
        position="fixed"
        elevation={0}
        sx={{ zIndex: (theme: Theme) => theme.zIndex.drawer + 1 }}
      >
        <Toolbar variant="dense">
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Mship
          </Typography>
          <Button className="native-link" href="/auth/oidc" color="inherit">
            Login
          </Button>
        </Toolbar>
      </AppBar>
      <Drawer
        sections={[
          {
            links: [
              { text: 'Dashboard', href: '/', icon: <DashboardIcon /> },
              { text: 'Packages', href: '/packages', icon: <ListAltIcon /> },
            ],
          },
        ]}
      />
      <Box component="main" sx={{ p: 3, flexGrow: 1 }}>
        <Toolbar variant="dense" />
        <Routes>
          <Route index element={<Dashboard />} />
        </Routes>
      </Box>
    </Box>
  );
};
