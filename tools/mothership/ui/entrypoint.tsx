import React from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import CssBaseline from '@mui/material/CssBaseline';
import ThemeProvider from '@mui/material/styles/ThemeProvider';

import { App } from './App';
import { peridotTheme } from 'base/ts/mui/theme';

const root = createRoot(document.getElementById('app') || document.body);
root.render(
  <BrowserRouter>
    <ThemeProvider theme={peridotTheme}>
      <CssBaseline />
      <App />
    </ThemeProvider>
  </BrowserRouter>
);
