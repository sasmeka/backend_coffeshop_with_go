package handlers

import (
	"sasmeka/coffeeshop/config"
	"sasmeka/coffeeshop/internal/models"
	"sasmeka/coffeeshop/internal/repositories"
	"sasmeka/coffeeshop/pkg"

	"github.com/asaskevich/govalidator"
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

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	search := ctx.Query("search")
	orderby := ctx.Query("order_by")

	if err := ctx.ShouldBind(&product); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&product, page, limit, search, orderby)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, response).Send(ctx)
}

func (h *Handler_Product) Post_Data_Product(ctx *gin.Context) {
	var productset models.Productsset

	if err := ctx.Bind(&productset); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	productset.Image = ctx.MustGet("image").(string)

	var err_val error
	_, err_val = govalidator.ValidateStruct(&productset)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.Insert_Data(&productset)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Product) Put_Data_Product(ctx *gin.Context) {
	var product models.Productsset

	product.Id_product = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(product.Id_product)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	product.Image = ctx.MustGet("image").(string)

	if err := ctx.ShouldBind(&product); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	var err_val error
	_, err_val = govalidator.ValidateStruct(&product)
	if err_val != nil {
		pkg.Responses(400, &config.Result{Message: err_val.Error()}).Send(ctx)
		return
	}

	response, err := h.Update_Data(&product)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
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
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&product); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&product, &products_sizes, &products_delivery_methods)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
