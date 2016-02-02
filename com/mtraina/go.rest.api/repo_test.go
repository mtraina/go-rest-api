package main

import "testing"

func TestRepoFindTodo(t *testing.T){

	todo := RepoFindTodo(1)

	if todo.Id != 1 {
		t.Error("expected id = 1, got ", todo.Id)
	}
}
