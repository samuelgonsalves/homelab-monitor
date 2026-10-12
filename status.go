package main

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"encoding/json"
)

func StatusDbRecord(key string) StatusRequest{
	var v []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME))
		v = b.Get([]byte(key))
		return nil
	})

	if v != nil {
		fmt.Printf("Found value for key: %s\n", v)
		var req StatusRequest
		json.Unmarshal(v, &req)
		return req
	}

	return StatusRequest{}
}
