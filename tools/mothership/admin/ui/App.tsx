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
import { Navigate, Route, Routes } from 'react-router-dom';

import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import { Theme } from '@mui/material/styles';

import DashboardIcon from '@mui/icons-material/Dashboard';
import EngineeringIcon from '@mui/icons-material/Engineering';
import ListAltIcon from '@mui/icons-material/ListAlt';

import { Workers } from './Workers';
import { Drawer } from 'base/ts/mui/Drawer';
import { CreateWorker } from 'tools/mothership/admin/ui/CreateWorker';
import { GetWorker } from 'tools/mothership/admin/ui/GetWorker';

export const App = () => {
  return (
    <Box sx={{ display: 'flex' }}>
      <AppBar
        elevation={5}
        position="fixed"
        sx={{ zIndex: (theme: Theme) => theme.zIndex.drawer + 1 }}
      >
        <Toolbar variant="dense">
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Mship Admin
          </Typography>
          <Box sx={{ flexGrow: 1, textAlign: 'right' }}>
            <Button className="native-link" href="/admin/auth/oidc/logout" variant="primary">
              Logout
            </Button>
            <Button className="native-link" href="/" variant="primary">
              Go back to Mship
            </Button>
          </Box>
        </Toolbar>
      </AppBar>
      <Drawer
        sections={[
          {
            links: [
              { text: 'Workers', href: '/workers', icon: <EngineeringIcon /> },
            ],
          },
        ]}
      />
      <Box component="main" sx={{ flexGrow: 1 }}>
        <Toolbar variant="dense" />
        <Routes>
          <Route index element={<Navigate to="/workers" replace />} />
          <Route path="/workers">
            <Route index element={<Workers />} />
            <Route path="create" element={<CreateWorker />} />
            <Route path=":name" element={<GetWorker />} />
          </Route>
        </Routes>
      </Box>
    </Box>
  );
};
