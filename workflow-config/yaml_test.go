package config

import (
	"fmt"
	"testing"
)

func TestGetString(t *testing.T) {
	fmt.Println(GetString("application.name"))
	fmt.Println(GetString("application.version"))
	fmt.Println(GetString("database.postgres-dsn"))
}
