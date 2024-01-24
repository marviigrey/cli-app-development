In this project we will be creating a todo application using go CLI tooling.

This tool will receive users input through the STDIN(Standard input) and give back result using the STDOUT(Standard output) to users. We will also add CLI parameters using flags.
Environmental varibles will be added to help us modify how our programs work. We will use the STDERR for error handling incase users encounter errors while running commands.

For building standard CLI go code structure, we need to seperate the business logic from the CMD line interface. This means that in your code repository, you have double sub-directories, which contains a directory handling the command line interface and the other handling the business logic.
   directory structure: Todo-app_
                                 |_cmd--
                                      |_todo
                                           |_main.go,main_test.go
                        --go.mod
                        --go.sum
                        --todo.go
                        --todo_test.go


Defining Todo API:
Creating a business logic and API to handle todo items.
API:
    - items: A struct to represent a single todo item.
    - lists: A slice to hold values of every item in our struct item.

These custom items will represent the data on our todo items, we make use of methods which are used to execute functions directly on type's data.
When creating an application of any type, hold it in mind that your application will have functions that will:
    - add items to the todo list.
    - delete items 
    - save items
    - update items
    - get or obtain an item from the todo list.


Note: Go Exported Types - In Go, the visibility of a type of function is controlled by the case of the first character of its name. Name that starts with an uppercase are exported while names wstarting with lowercase are considered private to the package.