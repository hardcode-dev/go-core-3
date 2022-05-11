package main

import (
	"fmt"
	"sf/32_NoSQL/KVStore/pkg/kvdb"
)

func main() {
	db := kvdb.New()
	db.SET("1", "test record")
	fmt.Println(db.GET("1"), db.GET("2"))
}
