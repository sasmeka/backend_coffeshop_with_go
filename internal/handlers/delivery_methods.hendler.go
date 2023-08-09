package handlers

import (
	"sasmeka/coffeeshop/config"
	"sasmeka/coffeeshop/internal/models"
	"sasmeka/coffeeshop/internal/repositories"
	"sasmeka/coffeeshop/pkg"

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
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	if err := ctx.ShouldBind(&deliver_method); err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Get_Data(&deliver_method, page, limit)
	if err != nil {
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	pkg.Responses(200, response).Send(ctx)
}

func (h *Handler_Delivery_Methods) Post_Data_Delivery_Method(ctx *gin.Context) {
	var deliver_method models.Delivery_Methods

	if err := ctx.ShouldBind(&deliver_method); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Insert_Data(&deliver_method)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
func (h *Handler_Delivery_Methods) Put_Data_Delivery_Method(ctx *gin.Context) {
	var deliver_method models.Delivery_Methods
	deliver_method.Id_dm = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(deliver_method.Id_dm)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&deliver_method); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Update_Data(&deliver_method)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}

func (h *Handler_Delivery_Methods) Delete_Data_Delivery_Method(ctx *gin.Context) {
	var deliver_method models.Delivery_Methods
	deliver_method.Id_dm = ctx.Param("id")

	count_by_id := h.Get_Count_by_Id(deliver_method.Id_dm)
	if count_by_id == 0 {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: "data not found."}).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&deliver_method); err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}

	response, err := h.Delete_Data(&deliver_method)
	if err != nil {
		// ctx.AbortWithError(http.StatusBadRequest, err)
		pkg.Responses(400, &config.Result{Message: err.Error()}).Send(ctx)
		return
	}
	pkg.Responses(200, &config.Result{Message: response}).Send(ctx)
}
