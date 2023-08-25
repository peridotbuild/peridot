import React from 'react';

import Box from '@mui/material/Box';
import Divider from '@mui/material/Divider';

import { NewResource } from 'base/ts/mui/NewResource';
import { reqap } from 'base/ts/reqap';

import { V1Worker } from 'bazel-bin/tools/mothership/proto/admin/v1/mshipadminpb_ts_proto_gen';
import { mshipAdminApi } from 'tools/mothership/admin/ui/api';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';

export const CreateWorker = () => {
  return (
    <Box>
      <Box sx={{ p: 1.5 }}>Create a new worker</Box>
      <Divider />
      <Box sx={{ p: 1.5, mt: 2 }}>
        <NewResource
          fields={[
            {
              key: 'workerId',
              label: 'Worker ID',
              subtitle:
                'This ID will be used to uniquely identify this worker.',
              type: 'text',
              required: true,
            },
          ]}
          save={(x: V1Worker) =>
            reqap(
              mshipAdminApi.createWorker({ body: { workerId: x.workerId!! } }),
            )
          }
          showDialog={(resource: V1Worker) => (
            <>
              <DialogTitle>Worker created</DialogTitle>
              <DialogContent>
                <DialogContentText>
                  The worker API secret is <code>{resource.apiSecret}</code>.
                  <br /><br />
                  This is the only time it will be shown, so make sure to save
                  it somewhere safe.
                </DialogContentText>
              </DialogContent>
            </>
          )}
        />
      </Box>
    </Box>
  );
};
