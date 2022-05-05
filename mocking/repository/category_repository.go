package repository

import "go-testing/mocking/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
