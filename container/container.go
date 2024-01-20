package container

import (
	"database/sql"
	"go-emqx-ex-hook-to-tdengine/model"
)

var AppConfig model.Config
var InsertSqlTemplate string
var SqlDb *sql.DB
