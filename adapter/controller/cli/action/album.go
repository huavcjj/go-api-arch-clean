package action

import (
	"go-api-arch-clean/entity"
	"go-api-arch-clean/pkg/logger"
	"go-api-arch-clean/usecase"
	"time"
)

var AlbumName string

type AlbumAction struct {
	albumUseCase usecase.AlbumUseCase
}

func NewAlbumAction(albumUseCase usecase.AlbumUseCase) *AlbumAction {
	return &AlbumAction{
		albumUseCase: albumUseCase,
	}
}

func (a *AlbumAction) CreateAlbum(title, categoryName string) (*entity.Album, error) {
	category, err := entity.NewCategory(categoryName)
	if err != nil {
		return nil, err
	}
	album := &entity.Album{
		Title:       title,
		ReleaseDate: time.Now(),
		Category:    *category,
	}

	createdAlbum, err := a.albumUseCase.Create(album)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return createdAlbum, nil
}
