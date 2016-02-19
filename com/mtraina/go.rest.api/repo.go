package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

var currentId int
var todos Todos
var db *bolt.DB

func init()  {
	fmt.Printf("start!!")

	db2, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	// todo: crap!
	db = db2

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//err := db.Update(func(tx *bolt.Tx) error {
	err2 := db.Update(func(tx *bolt.Tx) error {
		//return nil
		b, err := tx.CreateBucket([]byte("Todos"))
		if err != nil {
			return err
		}

		return b.Put([]byte("todo"), []byte("Write presentation2"))
	})

	if err2 != nil {
		log.Fatal(err2)
	}

  	//RepoCreateTodo(Todo{Name: "Write presentation"})
	//RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) (string, error) {
	var value string

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Todos"))
		v := b.Get([]byte("todo"))
		value = string(v)
		return nil
	})

//	if err != nil {
//		log.Fatal(err)
//	}

	return value, err


//	for _, t := range todos {
//		if t.Id == id {
//			return t
//		}
//	}
//	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i + 1:] ...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}