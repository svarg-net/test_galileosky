package usecase

import (
	"test_galileosky/internal/entity"
	"test_galileosky/internal/gateway"
)

type TaskUsecase interface {
	AddTask(task *entity.Task) (*entity.Task, error)
	GetTasks(sortBy string, filterDate string, page int, pageSize int) ([]*entity.Task, error)
	GetTotalTasks(filterDate string) (int, error)
	ExportToXLSX(sortBy string, filterDate string, page int, pageSize int) ([]byte, error)
}

type taskUsecase struct {
	dbGateway  gateway.DBGateway
	xlsxWriter gateway.XLSXExporter
}

func NewTaskUsecase(db gateway.DBGateway, xlsxWriter gateway.XLSXExporter) TaskUsecase {
	return &taskUsecase{
		dbGateway:  db,
		xlsxWriter: xlsxWriter,
	}
}

func (u *taskUsecase) AddTask(task *entity.Task) (*entity.Task, error) {
	return u.dbGateway.CreateTask(task)
}

func (u *taskUsecase) GetTasks(sortBy string, filterDate string, page int, pageSize int) ([]*entity.Task, error) {
	filteredTasks := []*entity.Task{}
	return filteredTasks, nil
}

func (u *taskUsecase) ExportToXLSX(sortBy string, filterDate string, page int, pageSize int) ([]byte, error) {
	return []byte{}, nil
}

func (u *taskUsecase) GetTotalTasks(filterDate string) (int, error) {

	return 0, nil
}
