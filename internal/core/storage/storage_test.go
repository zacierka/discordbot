package storage_test

import (
	"discordbot/internal/core/storage"
	"discordbot/internal/env"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {

	env.LoadEnv()

	db, _ := storage.New()

	ver := db.GetVersion()
	if ver == "NULL" {
		t.Error("CANNOT GET VERSION")
	} else {
		fmt.Println(ver)
	}

}
