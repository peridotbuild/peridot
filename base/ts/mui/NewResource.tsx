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
import { useNavigate } from 'react-router-dom';

import Box from '@mui/material/Box';
import FormControl from '@mui/material/FormControl';
import TextField from '@mui/material/TextField';
import FormHelperText from '@mui/material/FormHelperText';
import Alert from '@mui/material/Alert';
import Dialog from '@mui/material/Dialog';

import LoadingButton from '@mui/lab/LoadingButton';
import DialogActions from '@mui/material/DialogActions';
import Button from '@mui/material/Button';

import { StandardResource } from '../resource';

export interface NewResourceField {
  key: string;
  label: string;
  type: 'text' | 'number' | 'date';
  subtitle?: string;
  required?: boolean;
}

export interface NewResourceProps<T> {
  fields: NewResourceField[];

  save(x: T): Promise<any>;

  // show a dialog before redirecting to the resource page
  // the caller gets the resource as a parameter and can
  // return a React node to show in the modal
  showDialog?(resource: T): React.ReactNode;
}

export const NewResource = <T extends StandardResource>(
  props: NewResourceProps<T>,
) => {
  const navigate = useNavigate();
  const [loading, setLoading] = React.useState<boolean>(false);
  const [error, setError] = React.useState<string | undefined>(undefined);
  const [res, setRes] = React.useState<T | undefined>(undefined);

  const onSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    setError(undefined);
    setLoading(true);
    const data = new FormData(event.currentTarget);
    const obj = Object.fromEntries(data.entries());
    const [res, err] = await props.save(obj as T);
    setLoading(false);
    if (err) {
      console.log(err);
      setError(err.message);
      return;
    }

    if (!res) {
      setError('Unknown error');
      return;
    }

    setError(undefined);

    if (props.showDialog) {
      setRes(res);
      return;
    }

    // res should always have "name" so we can redirect to the resource page.
    navigate(`/${res.name}`);
  };

  const navigateToResource = () => {
    if (!res) {
      return;
    }

    navigate(`/${res.name}`);
  };

  return (
    <>
      {props.showDialog && res && (
        <Dialog
          disableEscapeKeyDown
          disableScrollLock
          open>
          {props.showDialog(res)}
          <DialogActions>
            <Button onClick={navigateToResource}>
              Go to resource
            </Button>
          </DialogActions>
        </Dialog>
      )}
      <Box component="form" onSubmit={onSubmit}>
        {error && (
          <Alert severity="error" sx={{ mb: 4 }}>
            {error}
          </Alert>
        )}
        {props.fields.map((field: NewResourceField) => (
          <FormControl key={field.key} sx={{ width: '100%', display: 'block' }}>
            <TextField
              name={field.key}
              sx={{ width: '100%' }}
              size="small"
              required={field.required}
              label={field.label}
              type={field.type}
            />
            {field.subtitle && (
              <FormHelperText>{field.subtitle}</FormHelperText>
            )}
          </FormControl>
        ))}
        <LoadingButton
          type="submit"
          variant="contained"
          size="small"
          sx={{ mt: 2 }}
          disabled={loading}
          loading={loading}
        >
          <span>Save</span>
        </LoadingButton>
      </Box>
    </>
  );
};
