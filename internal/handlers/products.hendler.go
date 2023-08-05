package handlers

import (
	"math"
	"net/http"
	"sasmeka/coffeeshop/internal/models"
	"sasmeka/coffeeshop/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler_Product struct {
	*repositories.Repo_Products
}

func New_Products(r *repositories.Repo_Products) *Handler_Product {
	return &Handler_Product{r}
}

func (h *Handler_Product) Get_Data_Products(ctx *gin.Context) {
	var product models.Products
	var meta_product models.Meta_Products
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	search := ctx.Query("search")
	orderby := ctx.Query("order_by")
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

	count_data := h.Get_Count_Data(search)

	if count_data <= 0 {
		meta_product.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_product.Next = ""
		} else {
			meta_product.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_product.Prev = ""
	} else {
		meta_product.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_product.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_product.Last_page = ""
	}

	if count_data != 0 {
		meta_product.Total_data = strconv.Itoa(count_data)
	} else {
		meta_product.Last_page = ""
	}

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Get_Data(&product, offset, limit_int, search, orderby)
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
		"meta":        meta_product,
	})
}

func (h *Handler_Product) Post_Data_Product(ctx *gin.Context) {
	var productset models.Productsset

	if err := ctx.ShouldBind(&productset); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Requestt",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Insert_Data(&productset)
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
func (h *Handler_Product) Put_Data_Product(ctx *gin.Context) {
	var product models.Productsset

	product.Id_product = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(product.Id_product)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     "data not found.",
		})
		return
	}

	if err := ctx.ShouldBind(&product); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Update_Data(&product)
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

func (h *Handler_Product) Delete_Data_Product(ctx *gin.Context) {
	var product models.Products
	var products_sizes models.Products_Sizes
	var products_delivery_methods models.Products_Delivery_Methods
	product.Id_product = ctx.Param("id")
	products_sizes.Id_product = ctx.Param("id")
	products_delivery_methods.Id_product = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(product.Id_product)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     "data not found.",
		})
		return
	}

	if err := ctx.ShouldBind(&product); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Delete_Data(&product, &products_sizes, &products_delivery_methods)
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
