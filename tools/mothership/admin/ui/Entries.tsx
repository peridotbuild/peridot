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

import Box from '@mui/material/Box';

import { ResourceTable } from 'base/ts/mui/ResourceTable';
import { srpmArchiverApi } from 'tools/mothership/admin/ui/api';
import {
  V1ListEntriesResponse,
  V1Entry,
} from 'bazel-bin/tools/mothership/proto/v1/mothershippb_ts_proto_gen';
import { reqap } from 'base/ts/reqap';

export const Entries = () => {
  return (
    <Box sx={{p: 1.5, px: 3, width: '100%'}}>
    <ResourceTable<V1Entry>
      load={(pageSize: number, pageToken?: string) => reqap(srpmArchiverApi.listEntries({
        pageSize: pageSize,
        pageToken: pageToken,
      }))}
      transform={((response: V1ListEntriesResponse) => response.entries || [])}
      fields={[
        { key: 'name', label: 'Entry Name' },
        { key: 'entryId', label: 'Entry ID' },
        { key: 'state', label: 'State' },
      ]}
    />
    </Box>
  );
};
