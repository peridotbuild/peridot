package migrations

import "embed"

//go:embed *.up.sql
var UpSQLs embed.FS
