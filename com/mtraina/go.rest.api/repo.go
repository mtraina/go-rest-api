package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

var currentId int
var todos Todos

func init()  {
  	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
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