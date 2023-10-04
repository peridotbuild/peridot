import * as React from 'react';

import { ResourceTable } from 'base/ts/mui/ResourceTable';
import {
  V1ListKernelsResponse,
  V1Kernel,
} from 'bazel-bin/tools/kernelmanager/proto/v1/kernelmanagerpb_ts_proto_gen';
import { reqap } from 'base/ts/reqap';
import { kernelManagerApi } from 'tools/kernelmanager/ui/api';

export const Kernels = () => {
  return (
    <ResourceTable<V1Kernel>
      hideFilter
      load={(pageSize: number, pageToken?: string, filter?: string) => reqap(kernelManagerApi.listKernels({
        pageSize,
        pageToken,
      }))}
      transform={((response: V1ListKernelsResponse) => response.kernels || [])}
      fields={[
        { key: 'name', label: 'Kernel Name' },
      ]}
    />
  );
};
