package presenter

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/moeghifar/libgo/pkg/xttp"
	"github.com/moeghifar/mariana/config"
	"github.com/moeghifar/mariana/module/account/entity"
	"github.com/moeghifar/mariana/module/account/usecase"
)

var (
	errInvalidUserID = xttp.Error{Status: fiber.StatusBadRequest, Code: "ACC-400", Title: "Invalid user id"}
	errUserNotFound  = xttp.Error{Status: fiber.StatusNotFound, Code: "ACC-404", Title: "User not found"}
)

type userResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func startHTTP(ctx context.Context, conf config.Config, ucs usecase.Usecases) error {
	return newAPI(conf, ucs).Run(ctx, fmt.Sprintf(":%d", conf.App.Port))
}

func mapError(err error) error {
	if errors.Is(err, entity.ErrUserNotFound) {
		return errUserNotFound.With("no user with this id", err)
	}
	return err
}

func newAPI(conf config.Config, ucs usecase.Usecases) *xttp.API {
	api := xttp.New(xttp.Config{
		Info:          xttp.Info{Title: conf.App.Name, Version: conf.AppVersion},
		DefaultErrors: []int{fiber.StatusInternalServerError},
		MapError:      mapError,
	})

	api.Group("/v1").Add(xttp.Route{
		Method:  fiber.MethodGet,
		Path:    "/users/:id",
		Summary: "Get a user",
		Tags:    []string{"account"},
		Responses: map[int]any{
			fiber.StatusOK:         userResponse{},
			fiber.StatusBadRequest: xttp.ErrorDocument{},
			fiber.StatusNotFound:   xttp.ErrorDocument{},
		},
		Handler: func(c fiber.Ctx) error {
			id, err := strconv.ParseInt(c.Params("id"), 10, 64)
			if err != nil {
				return errInvalidUserID.With("id must be an integer", err)
			}

			u, err := ucs.MyUsecase.GetByID(c.Context(), id)
			if err != nil {
				return err
			}

			return c.JSON(userResponse{ID: u.ID, Name: u.Name, Email: u.Email})
		},
	})

	return api
}
