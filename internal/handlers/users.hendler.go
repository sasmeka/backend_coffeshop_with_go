package handlers

import (
	"math"
	"net/http"
	"sasmeka/coffeeshop/internal/models"
	"sasmeka/coffeeshop/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler_Users struct {
	*repositories.Repo_Users
}

func New_Users(r *repositories.Repo_Users) *Handler_Users {
	return &Handler_Users{r}
}

func (h *Handler_Users) Get_Data(ctx *gin.Context) {
	var user models.Users
	var meta_user models.Meta_Users
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	var offset int = 0
	var page_int, _ = strconv.Atoi(page)
	var limit_int, _ = strconv.Atoi(limit)
	if limit == "" {
		limit_int = 100
	}
	if page == "" {
		page_int = 1
	}
	if page_int > 0 {
		offset = (page_int - 1) * limit_int
	} else {
		offset = 0
	}

	count_data := h.Get_Count_Users()

	if count_data <= 0 {
		meta_user.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_user.Next = ""
		} else {
			meta_user.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_user.Prev = ""
	} else {
		meta_user.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_user.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_user.Last_page = ""
	}

	if count_data != 0 {
		meta_user.Total_data = strconv.Itoa(count_data)
	} else {
		meta_user.Last_page = ""
	}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Get_Users(&user, offset, limit_int)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
		"meta":        meta_user,
	})
}

func (h *Handler_Users) Post_Data(ctx *gin.Context) {
	var user models.Users

	if err := ctx.ShouldBind(&user); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Insert_User(&user)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"message":     response,
	})
}
func (h *Handler_Users) Put_Data(ctx *gin.Context) {
	var user models.Users
	user.Id_user = ctx.Param("id")
	if err := ctx.ShouldBind(&user); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Update_User(&user)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"message":     response,
	})
}

func (h *Handler_Users) Delete_Data(ctx *gin.Context) {
	var user models.Users
	user.Id_user = ctx.Param("id")
	if err := ctx.ShouldBind(&user); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Delete_User(&user)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"message":     response,
	})
}
