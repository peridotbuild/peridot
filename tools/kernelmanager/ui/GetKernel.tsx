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

import Box from '@mui/material/Box';
import Divider from '@mui/material/Divider';

import { ResourceView } from 'base/ts/mui/ResourceView';
import { reqap } from 'base/ts/reqap';

import { V1Kernel } from 'bazel-bin/tools/kernelmanager/proto/v1/kernelmanagerpb_ts_proto_gen';
import { kernelManagerApi } from 'tools/kernelmanager/ui/api';

export const GetKernel = () => {
  const params = useParams();
  const [resource, setResource] = React.useState<V1Kernel | undefined | null>(
    undefined,
  );

  // Load the resource
  React.useEffect(() => {
    (async () => {
      const [res, err] = await reqap(
        kernelManagerApi.getKernel({
          name: location.pathname.substring(1),
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
      <Box
        sx={{
          px: 1.5,
          height: '48px',
          display: 'flex',
          justifyContent: 'justify-between',
          alignItems: 'center',
        }}
      >
        <span>{location.pathname.substring(1)}</span>
      </Box>
      <Divider />
      <Box sx={{ p: 1.5 }}>
        <pre>
          {resource ? JSON.stringify(resource, null, 2) : 'Loading'}
        </pre>
      </Box>
    </Box>
  );
};
