package daos

import (
	"armeris/gorest/domains"
	"fmt"
)

var currentId int
var TodosList domains.Todos

func init() {
	RepoCreateTodo(domains.Todo{Name: "Write presentation"})
	RepoCreateTodo(domains.Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) domains.Todo {
	for _, t := range TodosList {
		if t.Id == id {
			return t
		}
	}

	return domains.Todo{}
}

func RepoCreateTodo(t domains.Todo) domains.Todo {
	currentId += 1
	t.Id = currentId
	TodosList = append(TodosList, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range TodosList {
		if t.Id == id {
			TodosList = append(TodosList[:i], TodosList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
