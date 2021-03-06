package main

import (
	"auth/server/db"
	"auth/server/router"
	_ "database/sql"
	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()

	// Listen and Serve in 0.0.0.0:8081
	r.Run(":8081")
}
