package daoConnection

import (
	"database/sql"
	"sync"

	"github.com/AmonFla/simple-task-manager-api/utils"
)

var Once sync.Once

var SingleInstance *sql.DB

func FactoryDao() *sql.DB {
	var i *sql.DB
	switch utils.GoDotEnvVariable("DATABASE_TPYE") {

	/*case "Sqlite":
	i = getSqliteInstance(utils.GoDotEnvVariable("DATABASE_URI"))
	break;
	*/
	default:
		// log.Fatalf("El motor %s no esta implementado", e)
		i = getPQInstance(utils.GoDotEnvVariable("DATABASE_URI"))
	}

	return i
}
