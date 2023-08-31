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

import * as mshipAdmin from 'bazel-bin/tools/mothership/proto/admin/v1/mshipadminpb_ts_proto_gen';
import * as srpmArchiver from 'bazel-bin/tools/mothership/proto/v1/mothershippb_ts_proto_gen';

const archiverCfg = new srpmArchiver.Configuration({
  basePath: '/api',
})

export const srpmArchiverApi = new srpmArchiver.SrpmArchiverApi(archiverCfg);

const cfg = new mshipAdmin.Configuration({
  basePath: '/admin/api',
})

export const mshipAdminApi = new mshipAdmin.MshipAdminApi(cfg);
