package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/melardev/gogonic_gorm_api_crud/dtos"
	"github.com/melardev/gogonic_gorm_api_crud/services"
	"net/http"
	"strconv"
)

func GetAllTodos(c *gin.Context) {
	page, pageSize := getPagingParams(c)

	todos, totalTodoCount := services.FetchTodos(page, pageSize)

	c.JSON(http.StatusOK, dtos.CreateTodoPagedResponse(c.Request, todos, page, pageSize, totalTodoCount))
}

func GetAllPendingTodos(c *gin.Context) {
	page, pageSize := getPagingParams(c)
	todos, totalTodoCount := services.FetchPendingTodos(page, pageSize, false)

	c.JSON(http.StatusOK, dtos.CreateTodoPagedResponse(c.Request, todos, page, pageSize, totalTodoCount))
}
func GetAllCompletedTodos(c *gin.Context) {
	page, pageSize := getPagingParams(c)
	todos, totalTodoCount := services.FetchPendingTodos(page, pageSize, true)

	c.JSON(http.StatusOK, dtos.CreateTodoPagedResponse(c.Request, todos, page, pageSize, totalTodoCount))
}
func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	if id == "completed" {
		GetAllCompletedTodos(c)
		return
	} else if id == "pending" {
		GetAllPendingTodos(c)
		return
	}
	id64, _ := strconv.ParseUint(id, 10, 32)
	todo, err := services.FetchById(uint(id64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not find Todo"))
		return
	}

	c.JSON(http.StatusOK, dtos.GetSuccessTodoDto(&todo))
}
func CreateTodo(c *gin.Context) {
	var json dtos.CreateTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	todo, err := services.CreateTodo(json.Title, json.Description, json.Completed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
	}

	c.JSON(http.StatusOK, dtos.CreateTodoCreatedDto(&todo))
}

func UpdateTodo(c *gin.Context) {
	idStr := (c.Param("id"))
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must set an ID"))
		return
	}

	var json dtos.CreateTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	todo, err := services.UpdateTodo(uint(id), json.Title, json.Description, json.Completed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.CreateTodoUpdatedDto(&todo))

}

func DeleteTodo(c *gin.Context) {
	idStr := (c.Param("id"))
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must set an ID"))
		return
	}
	todo, err := services.FetchById(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, dtos.CreateErrorDtoWithMessage("todo not found"))
		return
	}

	err = services.DeleteTodo(&todo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not delete Todo"))
		return
	}

	c.JSON(http.StatusOK, dtos.CreateSuccessWithMessageDto("todo deleted successfully"))
}

func DeleteAllTodos(c *gin.Context) {
	services.DeleteAllTodos()
	c.JSON(http.StatusOK, dtos.CreateErrorDtoWithMessage("All todos deleted successfully"))
}
