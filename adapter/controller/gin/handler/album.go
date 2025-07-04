package handler

import (
	"go-api-arch-clean/adapter/controller/gin/presenter"
	"go-api-arch-clean/api"
	"go-api-arch-clean/entity"
	"go-api-arch-clean/pkg/logger"
	"go-api-arch-clean/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	albumUseCase usecase.AlbumUseCase
}

func NewAlbumHandler(albumUseCase usecase.AlbumUseCase) *AlbumHandler {
	return &AlbumHandler{
		albumUseCase: albumUseCase,
	}
}

func (a *AlbumHandler) albumToResponse(album *entity.Album) *presenter.AlbumResponse {
	return &presenter.AlbumResponse{
		ApiVersion: api.Version,
		Data: presenter.Album{
			Kind:        "album",
			Id:          album.ID,
			Title:       album.Title,
			ReleaseDate: presenter.ReleaseDate{Time: album.ReleaseDate},
			Category: presenter.Category{
				Id:   &album.Category.ID,
				Name: presenter.CategoryName(album.Category.Name),
			},
		},
	}
}

func (a *AlbumHandler) CreateAlbum(c *gin.Context) {
	var requestBody presenter.CreateAlbumJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	category, err := entity.NewCategory(string(requestBody.Category.Name))
	if err != nil {
		logger.Warn(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	album := &entity.Album{
		Title:       requestBody.Title,
		ReleaseDate: requestBody.ReleaseDate.Time,
		Category:    *category,
	}

	createdAlbum, err := a.albumUseCase.Create(album)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, a.albumToResponse(createdAlbum))
}

func (a *AlbumHandler) GetAlbumById(c *gin.Context, ID int) {
	album, err := a.albumUseCase.Get(ID)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, a.albumToResponse(album))
}

func (a *AlbumHandler) UpdateAlbumById(c *gin.Context, ID int) {
	var requestBody presenter.UpdateAlbumByIdJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	category, err := entity.NewCategory(string(requestBody.Category.Name))
	if err != nil {
		logger.Warn(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	album := &entity.Album{ID: ID, Title: *requestBody.Title, Category: *category}

	updatedAlbum, err := a.albumUseCase.Save(album)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, a.albumToResponse(updatedAlbum))
}

func (a *AlbumHandler) DeleteAlbumById(c *gin.Context, ID int) {
	if err := a.albumUseCase.Delete(ID); err != nil {
		logger.Error(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
