package db

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlConnector ConnectionFunc = ConnectorDB("mysql")
