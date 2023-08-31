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
import { useParams } from 'react-router-dom';

import { ResourceView } from 'base/ts/mui/ResourceView';
import { reqap } from 'base/ts/reqap';

import {
  EntryState,
  V1Entry,
} from 'bazel-bin/tools/mothership/proto/v1/mothershippb_ts_proto_gen';
import Box from '@mui/material/Box';
import Divider from '@mui/material/Divider';
import { mshipAdminApi, srpmArchiverApi } from 'tools/mothership/admin/ui/api';
import Button from '@mui/material/Button';

export const GetEntry = () => {
  const params = useParams();
  const [resource, setResource] = React.useState<V1Entry | undefined | null>(
    undefined,
  );

  // Load the resource
  React.useEffect(() => {
    (async () => {
      const [res, err] = await reqap(
        srpmArchiverApi.getEntry({
          name1: `entries/${params.name}`,
        }),
      );

      if (err) {
        setResource(null);
        return;
      }

      setResource(res);
    })().then();
  }, []);

  // Rescue the entry (call API)
  const rescueEntry = async () => {
    const [res, err] = await reqap(
      mshipAdminApi.rescueEntryImport({
        name: `entries/${params.name}`,
      }),
    );

    if (err) {
      return;
    }

    window.location.reload();
  };

  return (
    <Box>
      <Box
        sx={{
          px: 1.5,
          height: '48px',
          display: 'flex',
          justifyContent: 'justify-between',
          alignItems: 'center',
        }}
      >
        <span>entries/{params.name}</span>
        {resource && resource.state == EntryState.OnHold && (
          <Button
            sx={{ ml: 'auto', textAlign: 'right' }}
            variant="outlined"
            onClick={rescueEntry}
          >
            Rescue
          </Button>
        )}
      </Box>
      <Divider />
      <Box sx={{ p: 1.5 }}>
        <ResourceView
          resource={resource}
          fields={[
            { key: 'entryId', label: 'Entry ID' },
            { key: 'createTime', label: 'Created' },
            { key: 'osRelease', label: 'OS Release' },
            { key: 'sha256Sum', label: 'SHA256 Sum' },
            { key: 'repository', label: 'Repository' },
            { key: 'workerId', label: 'Worker ID' },
            { key: 'commitUri', label: 'Commit URI', linkToSelf: true },
            { key: 'commitHash', label: 'Commit Hash' },
            { key: 'state', label: 'State' },
          ]}
        />
      </Box>
      {resource && resource.state == EntryState.OnHold && (
        <Box sx={{ p: 1.5 }}>
          <code>{resource.errorMessage}</code>
        </Box>
      )}
    </Box>
  );
};
