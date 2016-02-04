package main

import (
	"testing"
)

func TestRepoFindTodo(t *testing.T){
	todo := RepoFindTodo(1)

	if todo.Id != 1 {
		t.Error("expected id = 1, got ", todo.Id)
	}
}

func TestGetNilTodoWithNotExpectedId(t *testing.T){
	todo := RepoFindTodo(999)

	if todo.Id != 0 {
		t.Error("expected todo with id 0, got one with id  ", todo.Id)
	}
}

func TestDestroyExistingTodo(t *testing.T){
	err := RepoDestroyTodo(1)

	if err != nil {
		t.Errorf("error destroying todo with id %d", 1)
	}
}

func TestDestroyNonExistingTodo(t *testing.T){
	err := RepoDestroyTodo(999)

	if err == nil {
		t.Errorf("destroying a non existing todo should generate an error")
	}
}
