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

import * as React from 'react';

import Button from '@mui/material/Button';
import Box from '@mui/material/Box';

import { ResourceTable } from 'base/ts/mui/ResourceTable';
import { reqap } from 'base/ts/reqap';
import { mshipAdminApi } from 'tools/mothership/admin/ui/api';
import {
  V1ListWorkersResponse,
  V1Worker,
} from 'bazel-bin/tools/mothership/proto/admin/v1/mshipadminpb_ts_proto_gen';

export const Workers = () => {
  return (
    <Box sx={{p: 1.5, px: 3, width: '100%'}}>
      <Box sx={{ mb: 2, ml: 'auto', textAlign: 'right' }}>
        <Button href="/workers/create" variant="outlined" size="small">
          Create a new worker
        </Button>
      </Box>
      <ResourceTable<V1Worker>
        load={(pageSize: number, pageToken?: string) =>
          reqap(
            mshipAdminApi.listWorkers({
              pageSize: pageSize,
              pageToken: pageToken,
            }),
          )
        }
        transform={(response: V1ListWorkersResponse) => response.workers || []}
        fields={[
          { key: 'name', label: 'Worker Name' },
          { key: 'workerId', label: 'Worker ID' },
          { key: 'createTime', label: 'Created' },
        ]}
      />
    </Box>
  );
};
