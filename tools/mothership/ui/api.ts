import * as srpmArchiver from 'bazel-bin/tools/mothership/proto/v1/mothershippb_ts_proto_gen';

const cfg = new srpmArchiver.Configuration({
  basePath: '/api',
})

export const srpmArchiverApi = new srpmArchiver.SrpmArchiverApi(cfg);
