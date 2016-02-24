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

type Execute func(db *bolt.DB) []byte

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

	v := exe(selectTodo)
	log.Printf("value is %v", v)
	return v
}

func selectTodo(db *bolt.DB) []byte {
	var todo []byte

	err := db.View(func(tx *bolt.Tx) error {
		log.Print("get bucket")
		b := tx.Bucket([]byte(bucketName))
		todo = b.Get([]byte("1"))
		log.Print("got value")
		//todo = string(v[:])
		return nil
	})

	if err != nil {
		log.Fatalf("error on get %v", err)
	}

	return todo
}

func exe(fn Execute) string {
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: true})
	if err != nil {
		log.Printf("error opening db %v", err)
	}
	defer db.Close()
	var r []byte = fn(db)
	log.Printf("got %v", r)
	return string(r[:])
}


