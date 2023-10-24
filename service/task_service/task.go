package task_service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helper"
	"final-project/repository/category_repo"
	"final-project/repository/task_repo"
	"final-project/repository/user_repo"
	"net/http"
)

type TaskService interface {
	Create(userId int, taskPayLoad *dto.NewTasksRequest) (*dto.NewTasksResponse, errs.MessageErr)
	Get() (*dto.GetResponseTasks, errs.MessageErr)
}

type taskService struct {
	taskRepo     task_repo.Repository
	categoryRepo category_repo.Repository
	userRepo     user_repo.Repository
}

func NewTaskService(taskRepo task_repo.Repository, categoryRepo category_repo.Repository, userRepo user_repo.Repository) TaskService {
	return &taskService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (ts *taskService) Create(userId int, taskPayLoad *dto.NewTasksRequest) (*dto.NewTasksResponse, errs.MessageErr) {
	err := helper.ValidateStruct(taskPayLoad)

	if err != nil {
		return nil, err
	}

	_, err = ts.categoryRepo.CheckCategoryId(taskPayLoad.CategoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	task := &entity.Task{
		UserId:      userId,
		Title:       taskPayLoad.Title,
		Description: taskPayLoad.Description,
		CategoryId:  taskPayLoad.CategoryId,
	}
	response, err := ts.taskRepo.CreateNewTask(task)

	if err != nil {
		return nil, err
	}

	response = &dto.NewTasksResponse{
		Id:          response.Id,
		Title:       response.Title,
		Description: response.Description,
		Status:      response.Status,
		UserId:      response.UserId,
		CategoryId:  response.CategoryId,
		CreatedAt:   response.CreatedAt,
	}

	return response, nil
}

func (ts *taskService) Get() (*dto.GetResponseTasks, errs.MessageErr) {
	tasks, err := ts.taskRepo.GetTask()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	response := dto.GetResponseTasks{
		StatusCode: http.StatusOK,
		Message: "successfully fetching task",
		Data: tasks,
	}

	return &response, nil
}


