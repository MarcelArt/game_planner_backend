package repository

import (
	"game_planner_backend/model"
	"game_planner_backend/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepo interface {
	Create(userInput *model.UserInput) (uint, error)
	Read(page int, limit int) ([]*model.User, error)
	Update(id string, userInput *model.UserInput) error
	Delete(id string) (*model.User, error)
	GetByID(id string) (*model.User, error)
	GetByUsernameOrEmail(username string) (*model.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(userInput *model.UserInput) (uint, error) {
	err := r.db.Create(&userInput).Error

	return userInput.ID, err
}

func (r *UserRepo) Read(page int, limit int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Scopes(utils.Paginate(page, limit)).Find(&users).Error

	return users, err
}

func (r *UserRepo) Update(id string, userInput *model.UserInput) error {
	return r.db.Model(userInput).Where(id).Updates(userInput).Error
}

func (r *UserRepo) Delete(id string) (*model.User, error) {
	var user *model.User
	err := r.db.Clauses(clause.Returning{}).Delete(&user, id).Error

	return user, err
}

func (r *UserRepo) GetByID(id string) (*model.User, error) {
	var user *model.User
	err := r.db.First(&user, id).Error

	return user, err
}

func (r *UserRepo) GetByUsernameOrEmail(username string) (*model.User, error) {
	var user *model.User
	err := r.db.
		Where("username = ?", username).
		Or("email = ?", username).
		First(&user).Error

	return user, err
}
