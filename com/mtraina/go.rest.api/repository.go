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

type Result struct {
	Data []byte
}

type DBOperation func(db *bolt.DB) Result

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

	v := execute(selectTodo)
	return string(v.Data[:])
}

func selectTodo(db *bolt.DB) Result {
	var todo []byte

	err := db.View(func(tx *bolt.Tx) error {
		log.Print("get bucket")
		b := tx.Bucket([]byte(bucketName))
		todo = b.Get([]byte("1"))
		log.Print("got value")
		return nil
	})

	if err != nil {
		log.Fatalf("error on get %v", err)
	}

	r := Result{}
	r.Data = make([]byte, len(todo))
	copy(r.Data, todo)
	return r
}

func execute(fn DBOperation) Result {
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: true})
	if err != nil {
		log.Printf("error opening db %v", err)
	}
	defer db.Close()
	return fn(db)
}


