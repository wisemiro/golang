package main

import (
	"auth/server/db"
	"auth/server/router"
	_ "database/sql"
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	b := &db.User{}
	data, err := proto.Marshal(b)
	if err != nil {
		log.Fatal("Marsh err", err)
	}
	fmt.Println(data)
	r := router.SetupRouter()

	// Listen and Serve in 0.0.0.0:8081
	r.Run(":8081")
}
