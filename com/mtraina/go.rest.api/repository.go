package main

import (
	"log"
	"time"
	"github.com/boltdb/bolt"
	"fmt"
)

//var db *bolt.DB

func init(){
	fmt.Print("start!")


	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	// todo: crap!
	//db = db2

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		//return nil
		b, err := tx.CreateBucket([]byte("todos"))
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
	log.Print("find")

	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: true})
	if err != nil {
		log.Printf("error opening db %v", err)
	}
	defer db.Close()

	//log.Fatalf("db opened")
	var todo string

	//err = db.View(func(tx *bolt.Tx) error {
	//
	//	log.Fatalf("open bucket")
	//	b := tx.Bucket([]byte("todo"))
	//	log.Fatalf("bucket opened")
	//
	//	v := b.Get([]byte("1"))
	//	log.Fatalf("got v")
	//	//fmt.Printf("The todo is: %s\n", v)
	//	todo = string(v[:])
	//	log.Fatalf("todo is: %s", todo)
	//
	//	return nil
	//})
	//
	//if err != nil {
	//	log.Fatalf("find todo error %v", err)
	//}

	todo = "Write presentation"

	return todo
}


