package handlers

import (
	"math"
	"net/http"
	"sasmeka/coffeeshop/internal/models"
	"sasmeka/coffeeshop/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler_Delivery_Methods struct {
	*repositories.Repo_Delivery_Methods
}

func New_Delivery_Methods(r *repositories.Repo_Delivery_Methods) *Handler_Delivery_Methods {
	return &Handler_Delivery_Methods{r}
}

func (h *Handler_Delivery_Methods) Get_Data_Delivery_Methods(ctx *gin.Context) {
	var deliver_method models.Delivery_Methods
	var meta_delivery_method models.Meta_Sizes
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
		meta_delivery_method.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_delivery_method.Next = ""
		} else {
			meta_delivery_method.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_delivery_method.Prev = ""
	} else {
		meta_delivery_method.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_delivery_method.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_delivery_method.Last_page = ""
	}

	if count_data != 0 {
		meta_delivery_method.Total_data = strconv.Itoa(count_data)
	} else {
		meta_delivery_method.Last_page = ""
	}

	if err := ctx.ShouldBind(&deliver_method); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Get_Data(&deliver_method, offset, limit_int)
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
		"meta":        meta_delivery_method,
	})
}

func (h *Handler_Delivery_Methods) Post_Data_Delivery_Method(ctx *gin.Context) {
	var deliver_method models.Delivery_Methods

	if err := ctx.ShouldBind(&deliver_method); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Insert_Data(&deliver_method)
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
func (h *Handler_Delivery_Methods) Put_Data_Delivery_Method(ctx *gin.Context) {
	var deliver_method models.Delivery_Methods
	deliver_method.Id_dm = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(deliver_method.Id_dm)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     "data not found.",
		})
		return
	}

	if err := ctx.ShouldBind(&deliver_method); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Update_Data(&deliver_method)
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

func (h *Handler_Delivery_Methods) Delete_Data_Delivery_Method(ctx *gin.Context) {
	var deliver_method models.Delivery_Methods
	deliver_method.Id_dm = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(deliver_method.Id_dm)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     "data not found.",
		})
		return
	}

	if err := ctx.ShouldBind(&deliver_method); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Delete_Data(&deliver_method)
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
