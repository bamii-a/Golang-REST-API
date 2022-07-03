package main

// Importing the packages that we need to use in our code.
import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// A todo is a struct with an ID, an Item, and a Completed field.
// @property {string} ID - This is the unique identifier for the todo item.
// @property {string} Item - The item that the user wants to add to the todo list
// @property {bool} Completed - This is a boolean property that will be used to determine whether a
// todo item is completed or not.
type todo struct {
	ID		string  `json: "id"`
	Item	string	`json: "itemn"`
	Completed bool  `json: "completed"`
}

// This is a slice of todo structs.
var todos = []todo{
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Buy groceries", Completed: true},
	{ID: "3", Item: "Play fifa", Completed: false},
}

// It takes a pointer to a gin.Context object, and returns a JSON response with a status code of 200
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

// We're binding the JSON body of the request to a new todo struct, appending it to the todos slice,
// and returning the new todo as JSON
func addTodo(context *gin.Context){
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)

}

// It gets the id from the URL, gets the todo from the database, and returns the todo as JSON
func getTodo (context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

// It loops through the todos slice and returns the todo with the matching ID
func getTodoById (id string) (*todo, error) {  
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

// It takes the id of a todo from the URL, finds the todo in the database, toggles the completed status
// of the todo, and returns the todo
func toggleTodoStatus (context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

// It creates a new router, adds a few routes, and starts the server
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}