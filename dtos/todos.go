package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/melardev/gogonic_gorm_api_crud/models"
	"net/http"
)

type CreateTodo struct {
	Title       string `form:"title" json:"title" xml:"title"  binding:"required"`
	Description string `form:"description" json:"description" xml:"description"`
	Completed   bool   `form:"completed" json:"completed" xml:"completed"`
}

func CreateTodoPagedResponse(request *http.Request, todos []models.Todo, page, pageSize, count int) gin.H {
	var resources = make([]interface{}, len(todos))
	for index, todo := range todos {
		resources[index] = GetTodoDto(&todo)
	}
	return CreatePagedResponse(request, resources, "todos", page, pageSize, count)
}

func GetTodoDto(todo *models.Todo) map[string]interface{} {
	dto := map[string]interface{}{
		"id":         todo.ID,
		"title":      todo.Title,
		"completed":  todo.Completed,
		"created_at": todo.CreatedAt,
		"updated_at": todo.UpdatedAt,
	}

	return dto
}

func CreateTodoCreatedDto(todo *models.Todo) interface{} {
	return CreateSuccessWithDtoAndMessageDto(GetTodoDto(todo), "Todo created successfully")
}

func CreateTodoUpdatedDto(todo *models.Todo) interface{} {
	return CreateSuccessWithDtoAndMessageDto(GetTodoDto(todo), "Todo updated successfully")
}
