// main_test.go
package main_test

import (
	"log"
	"os"
	"testing"

	"github.com/AmonFla/simple-task-manager-api/server"
	"github.com/AmonFla/simple-task-manager-api/utils"
)

var a server.App

func TestMain(m *testing.M) {
	a = server.App{}
	a.Initialize(utils.GoDotEnvVariable("DATABASE_URI"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS public.users (
    id bigint NOT NULL,
    name character varying(150) NOT NULL,
    username character varying(150) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(150) NOT NULL,
    created_at time without time zone DEFAULT now(),
    updated_at time without time zone NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT username_unique UNIQUE (username)
)`
