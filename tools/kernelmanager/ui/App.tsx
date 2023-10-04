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
import { Kernels } from './Kernels';

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
            RESF KernelManager{window.__beta__ ? ' (beta)' : ''}
          </Typography>
          <Box sx={{ flexGrow: 1, textAlign: 'right' }}>
            {window.__peridot_user__ ? (
              <>
                <Button variant="primary">
                  {window.__peridot_user__.email}
                </Button>
                <Button
                  className="native-link"
                  href="/auth/oidc/logout"
                  variant="primary"
                >
                  Logout
                </Button>
              </>
            ) : (
              <Button
                className="native-link"
                href="/auth/oidc/login"
                variant="primary"
              >
                Login
              </Button>
            )}
          </Box>
        </Toolbar>
      </AppBar>
      <Box component="main" sx={{ p: 3, flexGrow: 1 }}>
        <Toolbar variant="dense" />
        <Routes>
          <Route index element={<Navigate to="/kernels" replace />} />
          <Route path="/kernels">
            <Route index element={<Kernels />} />
          </Route>
        </Routes>
      </Box>
    </Box>
  );
};
