package main

import (
	"log"
	"time"
	"github.com/boltdb/bolt"
	"fmt"
)

var (
	dbName string = "my.db"
	bucketName string = "todos"
)

func init(){
	fmt.Print("start!")

	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		//return nil
		b, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			log.Fatalf("create bucket error %v", err)
			return err
		}
		return b.Put([]byte("1"), []byte("Write presentation"))
	})
	if err != nil {
		log.Fatalf("update error %v", err)
	}
}

func FindTodo() string {
	log.Print("find todo")

	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: true})
	if err != nil {
		log.Printf("error opening db %v", err)
	}
	defer db.Close()

	var todo string

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte("1"))
		todo = string(v[:])
		return nil
	})

	return todo
}


