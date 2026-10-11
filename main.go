package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	bolt "go.etcd.io/bbolt"
	"strconv"
)

var db *bolt.DB

type StatusRequest struct {
	Status bool `json:"status"`
}

func status(w http.ResponseWriter, r *http.Request) {
	// Get the latest status
	var reqData StatusRequest
	json.NewDecoder(r.Body).Decode(&reqData)

	t := time.Now().Format("2006-01-02")

	var v []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("StatusBucket"))
		v = b.Get([]byte(t))
		return nil
	})

	if v != nil {
		fmt.Printf("Found key value for today: %s\n", v)
	}
}

func statusUpdate(w http.ResponseWriter, r *http.Request) {
	// Write the status to the db
	var reqData StatusRequest
	json.NewDecoder(r.Body).Decode(&reqData)

	t := time.Now().Format("2006-01-02")

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("StatusBucket"))
		err := b.Put([]byte(t), []byte(strconv.FormatBool(reqData.Status)))
		return err
	})

	fmt.Println(reqData)
}

func main() {
	var err error
	db, err = bolt.Open("./state/monitor.db", 0600, nil)
	if err != nil {
  		panic(fmt.Sprintf("Failed to open db: %v", err))
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("StatusBucket"))
		if (err != nil) {
			return err
		}
		return nil
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to create bucket: %v", err))
	}

	fmt.Println("DB and bucket finished setup")
	defer db.Close()
	http.HandleFunc("/backup-status-update", statusUpdate)
	http.HandleFunc("/backup-status", status)

	http.ListenAndServe(":8089", nil)
}
