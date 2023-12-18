package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

// This is just a test. and does not set any real value.

type Task struct {
	ID          int
	Title       string
	PersonEamil string
}

type dbTaskItem struct {
	PK string `dynamo:"PK"`
	SK string `dynamo:"SK"`
	Task
}

type Person struct {
	Email string
	Name  string
}

type dbPersonItem struct {
	PK string `dynamo:"PK"`
	SK string `dynamo:"SK"`
	Person
}

func main() {

	table := dynamo.New(session.Must(session.NewSession()), &aws.Config{Region: aws.String("us-west-2")}).Table("...")

	t := Task{ID: 1, Title: "Buy milk"}

	key := fmt.Sprintf("TASK#%d", t.ID)

	// item := dbItem{
	// 	Task: t,
	// 	PK: key,
	// 	SK: "Task",
	// }
	item := struct {
		PK string `dynamo:"PK"`
		SK string `dynamo:"SK"`
		Task
	}{
		Task: t,
		PK:   key,
		SK:   "Task",
	}

	err := table.Put(item).Run()

	if err != nil {
		panic(err)
	}

	listItem := dbTaskItem{
		Task: t,
		PK:   "AllTasks",
		SK:   key,
	}

	err = table.Put(listItem).Run()

	if err != nil {
		panic(err)
	}

	// get a single task by task ID
	var result Task
	err = table.Get("PK", key).
		Range("SK", dynamo.Equal, "Task").
		One(&result)

	// get all tasks
	var result1 []Task
	err = table.Get("PK", "AllTasks").All(&result1)
}
