-- Copyright 2023 Peridot Authors
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

CREATE TABLE workers
(
  name              VARCHAR(255) PRIMARY KEY,
  worker_id         VARCHAR(255) UNIQUE NOT NULL,
  create_time       TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
  last_checkin_time TIMESTAMPTZ,
  api_secret        VARCHAR(255)        NOT NULL
);

CREATE TABLE entries
(
  name            VARCHAR(255) PRIMARY KEY,
  entry_id        VARCHAR(255) UNIQUE NOT NULL,
  create_time     TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
  os_release      TEXT                NOT NULL,
  sha256_sum      VARCHAR(255)        NOT NULL,
  repository_name VARCHAR(255)        NOT NULL,
  worker_id       VARCHAR(255) REFERENCES workers (worker_id),
  batch_name      VARCHAR(255),
  user_email      TEXT,
  commit_uri      TEXT                NOT NULL,
  commit_hash     TEXT                NOT NULL
);

CREATE TABLE batches
(
  name           VARCHAR(255) PRIMARY KEY,
  batch_id       VARCHAR(255) UNIQUE,
  create_time    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  update_time    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  seal_time      TIMESTAMPTZ,
  bugtracker_url TEXT
);

CREATE TABLE bugtracker_configs
(
  name        VARCHAR(255) PRIMARY KEY,
  create_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  update_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  config      JSONB       NOT NULL
);
