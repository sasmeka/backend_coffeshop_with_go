package repositories

import (
	"sasmeka/coffeeshop/config"
	"sasmeka/coffeeshop/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoUserMock struct {
	mock.Mock
}

func (rp *RepoUserMock) Get_Users(data *models.Users, page string, limit string) (*config.Result, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}
func (rp *RepoUserMock) Get_Count_by_Id(id string) int {
	args := rp.Mock.Called(id)
	return args.Get(0).(int)
}
func (rp *RepoUserMock) Get_Count_by_Email(email string) int {
	args := rp.Mock.Called(email)
	return args.Get(0).(int)
}
func (rp *RepoUserMock) Get_Count_Users() int {
	args := rp.Mock.Called()
	return args.Get(0).(int)
}
func (rp *RepoUserMock) Insert_User(data *models.Users) (string, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
func (rp *RepoUserMock) Update_User(data *models.Users) (string, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
func (rp *RepoUserMock) Delete_User(data *models.Users) (string, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
