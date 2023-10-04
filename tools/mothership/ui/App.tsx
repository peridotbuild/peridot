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
import { Link, Navigate, Route, Routes } from 'react-router-dom';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Button from '@mui/material/Button';
import { Theme } from '@mui/material/styles';

import { GetEntry } from 'tools/mothership/ui/GetEntry';
import { Entries } from 'tools/mothership/ui/Entries';

export const App = () => {
  return (
    <Box sx={{ display: 'flex' }}>
      <AppBar
        elevation={5}
        position="fixed"
        sx={{ zIndex: (theme: Theme) => theme.zIndex.drawer + 1 }}
      >
        <Toolbar variant="dense">
          <Link to="/">
            <img
              alt="Mothership logo"
              src={
                window.__beta__
                  ? '/_ga/mship_gopher_beta.png'
                  : '/_ga/mship_gopher.png'
              }
              height="41.5px"
            />
          </Link>
          <Box sx={{ flexGrow: 1, textAlign: 'right' }}>
            <Button className="native-link" href="/admin" variant="primary">
              Admin
            </Button>
          </Box>
        </Toolbar>
      </AppBar>
      <Box component="main" sx={{ p: 3, flexGrow: 1 }}>
        <Toolbar variant="dense" />
        <Routes>
          <Route index element={<Navigate to="/entries" replace />} />
          <Route path="/entries">
            <Route index element={<Entries />} />
            <Route path=":name" element={<GetEntry />} />
          </Route>
        </Routes>
      </Box>
    </Box>
  );
};
