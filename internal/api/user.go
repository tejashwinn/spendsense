package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tejashwinn/sependsense/internal/model"
)

type UserDatabase interface {
	CreateUser(ctx *gin.Context, user *model.User) error
	GetById(ctx *gin.Context, id uint64) (*model.User, error)
}

type UserApi struct {
	DB UserDatabase
}

func (api *UserApi) CreateUser(ctx *gin.Context) {
	user := model.CreateUserReq{}
	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	internal := &model.User{
		Name:  user.Name,
		Email: user.Email,
	}
	err := api.DB.CreateUser(ctx, internal)
	if err != nil {
		panic("Error during creation")
	}
	ctx.JSON(http.StatusCreated, toUserResponse(internal))
}

func (api *UserApi) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := api.DB.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, toUserResponse(user))

}

func toUserResponse(user *model.User) *model.UserRes {
	return &model.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
