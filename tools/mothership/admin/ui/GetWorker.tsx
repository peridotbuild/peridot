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

import { mshipAdminApi } from 'tools/mothership/admin/ui/api';

import { V1Worker } from 'bazel-bin/tools/mothership/proto/admin/v1/mshipadminpb_ts_proto_gen';
import Box from '@mui/material/Box';
import Divider from '@mui/material/Divider';

export const GetWorker = () => {
  const params = useParams();
  const [resource, setResource] = React.useState<V1Worker | undefined | null>(
    undefined,
  );

  // Load the resource
  React.useEffect(() => {
    (async () => {
      const [res, err] = await reqap(
        mshipAdminApi.getWorker({
          name: `workers/${params.name}`,
        }),
      );

      if (err) {
        setResource(null);
        return;
      }

      setResource(res);
    })().then();
  }, []);

  return (
    <Box>
      <Box sx={{ p: 1.5 }}>workers/{params.name}</Box>
      <Divider />
      <Box sx={{ p: 1.5 }}>
        <ResourceView
          resource={resource}
          fields={[
            { key: 'workerId', label: 'Worker ID' },
            { key: 'createTime', label: 'Created' },
            { key: 'lastCheckinTime', label: 'Last active' },
          ]}
        />
      </Box>
    </Box>
  );
};
