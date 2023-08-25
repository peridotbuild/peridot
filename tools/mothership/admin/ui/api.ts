import * as mshipAdmin from 'bazel-bin/tools/mothership/proto/admin/v1/mshipadminpb_ts_proto_gen';

const cfg = new mshipAdmin.Configuration({
  basePath: '/admin/api',
})

export const mshipAdminApi = new mshipAdmin.MshipAdminApi(cfg);
