package category_repo

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
)

type Repository interface {
	CreateCategories(categoryPayLoad *entity.Category) (*dto.NewCategoryRequest, errs.MessageErr)
}