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
import Paper from '@mui/material/Paper';
import Box from '@mui/material/Box';
import LinearProgress from '@mui/material/LinearProgress';
import Alert from '@mui/material/Alert';
import Typography from '@mui/material/Typography';

export interface ResourceViewField {
  key: string;
  label: string;
  type?: 'text' | 'number' | 'date';
}

export interface ResourceViewProps<T> {
  // null means error, undefined means loading
  resource: T | null | undefined;
  fields: ResourceViewField[];
}

export function ResourceView<T>(props: ResourceViewProps<T>) {
  return (
    <Box>
      {props.resource === undefined && <LinearProgress />}
      {props.resource === null && (
        <Alert severity="error">Error loading resource</Alert>
      )}
      {props.resource && (
        <Paper elevation={2} sx={{ px: 2, py: 1 }}>
          {props.fields.map((field: ResourceViewField) => (
            <Box sx={{ my: 2.5 }}>
              <Typography variant="subtitle2">{field.label}</Typography>
              <Typography variant="body1">
                {(props.resource as any)[field.key]?.toString() || '--'}
              </Typography>
            </Box>
          ))}
        </Paper>
      )}
    </Box>
  );
}
