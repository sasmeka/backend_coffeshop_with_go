package handlers

import (
	"math"
	"net/http"
	"sasmeka/coffeeshop/internal/models"
	"sasmeka/coffeeshop/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler_Sizes struct {
	*repositories.Repo_Sizes
}

func New_Sizes(r *repositories.Repo_Sizes) *Handler_Sizes {
	return &Handler_Sizes{r}
}

func (h *Handler_Sizes) Get_Data_Sizes(ctx *gin.Context) {
	var size models.Sizes
	var meta_size models.Meta_Sizes
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	var offset int = 0
	var page_int, _ = strconv.Atoi(page)
	var limit_int, _ = strconv.Atoi(limit)
	if limit == "" {
		limit_int = 5
	}
	if page == "" {
		page_int = 1
	}
	if page_int > 0 {
		offset = (page_int - 1) * limit_int
	} else {
		offset = 0
	}

	count_data := h.Get_Count_Data()

	if count_data <= 0 {
		meta_size.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_size.Next = ""
		} else {
			meta_size.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_size.Prev = ""
	} else {
		meta_size.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_size.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_size.Last_page = ""
	}

	if count_data != 0 {
		meta_size.Total_data = strconv.Itoa(count_data)
	} else {
		meta_size.Last_page = ""
	}

	if err := ctx.ShouldBind(&size); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Get_Data(&size, offset, limit_int)
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
		"meta":        meta_size,
	})
}

func (h *Handler_Sizes) Post_Data_Size(ctx *gin.Context) {
	var size models.Sizes

	if err := ctx.ShouldBind(&size); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Insert_Data(&size)
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
func (h *Handler_Sizes) Put_Data_Size(ctx *gin.Context) {
	var size models.Sizes
	size.Id_size = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(size.Id_size)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     "data not found.",
		})
		return
	}

	if err := ctx.ShouldBind(&size); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Update_Data(&size)
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

func (h *Handler_Sizes) Delete_Data_Size(ctx *gin.Context) {
	var size models.Sizes
	size.Id_size = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(size.Id_size)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     "data not found.",
		})
		return
	}

	if err := ctx.ShouldBind(&size); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Delete_Data(&size)
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
