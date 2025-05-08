package controllers

import (
	"fmt"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

type TodosController struct{}

// hello
func (con TodosController) Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello TodoList!")
}

// 增-创建待办
func (con TodosController) CreateEvent(ctx *gin.Context) {
	// // // 从请求中获取数据
	var result models.Todos
	ctx.BindJSON(&result)
	fmt.Printf("------result:%#v-------", result)

	// data := models.Todos{
	// 	Title:  "code",
	// 	Status: false,
	// }
	// fmt.Println("data", &data)

	if err := models.DB.Create(&result).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": result,
		})
	}
}

// 删-删除待办
func (con TodosController) DeleteEvent(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "invalid id",
		})
		return
	}
	if err := models.DB.Where("id=?", id).Delete(&models.Todos{}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

// 改-更新待办
func (con TodosController) EditEvent(ctx *gin.Context) {
	var todo models.Todos
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "invalid id",
		})
		return
	}
	if err := models.DB.Where("id=?", id).First(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.BindJSON(&todo)
	if err := models.DB.Save(&todo).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todo,
		})
	}
}

// 查-查询待办
func (con TodosController) GetEvent(ctx *gin.Context) {
	var todoList []models.Todos
	if err := models.DB.Find(&todoList).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"code": 200,
			"data": todoList,
		})
	}
}
