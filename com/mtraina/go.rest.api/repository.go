package main

import (
	"log"
	"time"
	"github.com/boltdb/bolt"
	"fmt"
)

var db *bolt.DB

func init(){
	fmt.Print("start!")


	db2, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	// todo: crap!
	db = db2

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		//return nil
		b, err := tx.CreateBucket([]byte("todos"))
		if err != nil {
			log.Fatalf("error! %s", err)
			return err
		}
		return b.Put([]byte("1"), []byte("Write presentation"))
	})
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

func FindTodo() string {
	log.Fatalf("find")

	var todo string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		v := b.Get([]byte("1"))
		fmt.Printf("The todo is: %s\n", v)
		todo = string(v[:])
		return nil
	})

	return todo
}


