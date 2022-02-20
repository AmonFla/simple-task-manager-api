package daoConnection

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func getPQInstance(uri string) *sql.DB {
	if SingleInstance == nil {
		Once.Do(
			func() {
				var err error
				SingleInstance, err = sql.Open("postgres", uri)
				if err != nil {
					//TODO
					panic(err)
				}
			})
	}
	return SingleInstance
}
