package query

import (
	"database/sql"
	"github.com/seichewarning/mysql-restful-server/conf"
	"github.com/seichewarning/mysql-restful-server/connection"
	"log"
)

func CheckUser(username string, passwd string) bool {
	sqlstring := "select id from " + conf.GetAuthTableName() + " where " + conf.GetAuthName() + " = '" + username + "' and " + conf.GetAuthPwd() + " = '" + passwd + "' ;"
	log.Println(sqlstring)
	rows, err := connection.GetConnection().Query(sqlstring)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		return false
	}
	values := make([]sql.RawBytes, len(columns))
	if len(values) > 0 {
		return true
	}
	return false
}
