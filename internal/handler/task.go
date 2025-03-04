package handler

import (
	"net/http"
	"strconv"
	"test_galileosky/internal/entity"
	"test_galileosky/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(uc usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: uc}
}

func (h *TaskHandler) AddTask(ctx *gin.Context) {
	var task entity.Task

	newTask, _ := h.taskUsecase.AddTask(&task)
	ctx.JSON(http.StatusCreated, newTask)
}

func (h *TaskHandler) GetTasks(ctx *gin.Context) {
	sortBy := ctx.DefaultQuery("sort_by", "")
	filterDate := ctx.DefaultQuery("filter_date", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	tasks, err := h.taskUsecase.GetTasks(sortBy, filterDate, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tasks":     tasks,
		"total":     0,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *TaskHandler) ExportToXLSX(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("Content-Disposition", "attachment; filename=tasks.xlsx")
	ctx.Data(http.StatusOK, "application/octet-stream", []byte{})
}
