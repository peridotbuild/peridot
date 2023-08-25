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
  name        VARCHAR(255) PRIMARY KEY,
  entry_id    VARCHAR(255) UNIQUE NOT NULL,
  create_time TIMESTAMPTZ         NOT NULL DEFAULT NOW(),
  os_release  TEXT                NOT NULL,
  sha256_sum  VARCHAR(255)        NOT NULL,
  worker_id   VARCHAR(255) REFERENCES workers (worker_id),
  batch_name  VARCHAR(255),
  user_email  TEXT
)
