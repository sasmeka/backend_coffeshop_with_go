package handlers

import (
	"sasmeka/coffeeshop/config"
	"sasmeka/coffeeshop/internal/models"
	"sasmeka/coffeeshop/internal/repositories"
	"sasmeka/coffeeshop/pkg"

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

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if err := ctx.ShouldBind(&size); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&size, page, limit)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, response).Send(ctx)
}

func (h *Handler_Sizes) Post_Data_Size(ctx *gin.Context) {
	var size models.Sizes

	if err := ctx.ShouldBind(&size); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Insert_Data(&size)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Sizes) Put_Data_Size(ctx *gin.Context) {
	var size models.Sizes
	size.Id_size = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(size.Id_size)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&size); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Update_Data(&size)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Sizes) Delete_Data_Size(ctx *gin.Context) {
	var size models.Sizes
	size.Id_size = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(size.Id_size)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&size); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&size)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
