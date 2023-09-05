package handlers

import (
	"net/http"
	"net/http/httptest"
	"sasmeka/coffeeshop/config"
	"sasmeka/coffeeshop/internal/repositories"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	_ "image/png"
)

var repoUserMock = repositories.RepoUserMock{}
var reqBody = `{
	"displayname":"sasmeka",
	"first_name":"verdi",
	"last_name":"sasmeka",
	"gender":"male",
	"phone":"087734768584",
	"email":"verdysas@gmail.com",
	"pass":"123456",
	"birth_date":"2006-01-02",
	"image":"file.jpg"
}`

func TestGet_Data_Users(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Users(&repoUserMock)
	expect := &config.Result{
		Data: []interface{}{map[string]interface{}{"birth_date": "2006-01-02T00:00:00Z", "create_at": "2023-09-01T20:07:30.662605Z", "displayname": "julian", "email": "julian@gmail.com", "first_name": "julian", "gender": "male", "id_user": "e9c02c27-17b0-4fc2-9048-9cae88663f48", "image": "/static/img/Default_Profile.png", "last_name": "mindria", "phone": "08737645", "role": "user", "status_verification": "0", "update_at": interface{}(nil)}},
		Meta: map[string]interface{}{"last_page": "1", "next": "", "prev": "", "total_data": "1"}}
	repoUserMock.On("Get_Users", mock.Anything).Return(expect, nil)

	r.GET("/user", handler.Get_Data_Users)
	req := httptest.NewRequest("GET", "/user?limit=1&page=1", nil)
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"code":200, "data": [
        {
            "id_user": "e9c02c27-17b0-4fc2-9048-9cae88663f48",
            "displayname": "julian",
            "first_name": "julian",
            "last_name": "mindria",
            "gender": "male",
            "phone": "08737645",
            "email": "julian@gmail.com",
            "birth_date": "2006-01-02T00:00:00Z",
            "status_verification": "0",
            "role": "user",
            "image": "/static/img/Default_Profile.png",
            "create_at": "2023-09-01T20:07:30.662605Z",
            "update_at": null
		}
    ], "meta": {
        "next": "",
        "prev": "",
        "last_page": "1",
        "total_data": "1"
    } ,"status":"OK"}`, w.Body.String())
}

func TestPost_Data_User(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	c.Set("image", "")

	handler := New_Users(&repoUserMock)
	count_id := 1
	repoUserMock.On("Get_Count_by_Email", mock.Anything).Return(count_id)
	repoUserMock.On("Insert_User", mock.Anything).Return("add user data successful.", nil)

	r.POST("/create_user", handler.Post_Data_User)
	req := httptest.NewRequest("POST", "/create_user", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if count_id == 0 {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "add user data successful.", "status":"OK"}`, w.Body.String())
	} else {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "e-mail already registered.", "status":"Bad Request"}`, w.Body.String())
	}
}

func TestPut_Data_User(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	c.Set("image", "image.jpg")

	handler := New_Users(&repoUserMock)
	count_id := 1
	repoUserMock.On("Get_Count_by_Id", mock.Anything).Return(count_id)
	repoUserMock.On("Update_User", mock.Anything).Return("update user data successful", nil)

	r.PUT("/update_user/:id", handler.Put_Data_User)
	req := httptest.NewRequest("PUT", "/update_user/asdg8awgd8wtd6", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if count_id == 0 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "data not found.", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "update user data successful", "status":"OK"}`, w.Body.String())
	}
}

func TestDelete_Data_User(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Users(&repoUserMock)
	count_id := 0
	repoUserMock.On("Get_Count_by_Id", mock.Anything).Return(count_id)
	repoUserMock.On("Delete_User", mock.Anything).Return("delete user data successful", nil)

	r.DELETE("/delete_user/:id", handler.Delete_Data_User)
	req := httptest.NewRequest("DELETE", "/delete_user/asdg8awgd8wtd6", strings.NewReader("{}"))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if count_id == 0 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "data not found.", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "delete user data successful", "status":"OK"}`, w.Body.String())
	}
}
