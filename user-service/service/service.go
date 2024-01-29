package service

import (
	"context"
	"user-service/domain"
	"user-service/repository"

	"github.com/gin-gonic/gin"
)

func CreateUser(c context.Context, user domain.User) (int, error) {
	return repo.Create(c, user)
}
func DeleteUser(c context.Context, userId int) error {
	return repo.DeleteUser(c, userId)

}
func GetUserInfo(c *gin.Context, userId int) (*domain.User, error) {
	return repo.GetUserInfo(c, userId)
}

func IsAdmin(c context.Context, userId int) (bool, error) {
	user, err := repo.GetUserInfo(c, userId)
	if err != nil {
		return false, err
	}
	return user.Type == domain.Admin, nil
}
