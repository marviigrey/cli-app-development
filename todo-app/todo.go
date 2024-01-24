package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type items struct { //created a struct to take in values of our todo-app
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type lists []items //a type alias of the struct "items", this will help us to attach new  methods to it.

func (l *lists) Add(task string) {
	//Add creates a new value and append to the list.
	t := items{
		Task:        task,       //takes in a string
		Done:        false,      //initial value is false
		CreatedAt:   time.Now(), //use time function to declare time a task is created.
		CompletedAt: time.Time{},
	}
	*l = append(*l, t) //append the stdin into the slice  type alias of items,with the use of pointer l.

}

// Complete method marks a todo item as completed by
// setting Done = true and CompletedAt to current time
func (l *lists) Complete(i int) error {
	ls := *l

	if i <= 0 || i > (len(ls)) { //getting the index of the item to mark as complete.
		//returns error if index is greater than length of lists or index is less than 0.
		return fmt.Errorf("Item %v doesnt exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

func (l *lists) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %v doesnt exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

func (l *lists) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (l *lists) Get(filename string) error {
	file, err := os.ReadFile(filename) //takes in name of a file to unmarshal,redads the content of the file.
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}
		return nil
	}
	if len(file) == 0 { //if there's no content inside the file specified return error
		return nil
	}

	return json.Unmarshal(file, l)
}
