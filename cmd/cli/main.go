package main

import (
	"go-api-arch-clean/adapter/controller/cli/action"
	"go-api-arch-clean/adapter/controller/cli/command"
	"go-api-arch-clean/adapter/gateway"
	"go-api-arch-clean/infrastructure/database"
	"go-api-arch-clean/pkg/logger"
	"go-api-arch-clean/usecase"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	db, err := database.NewDatabaseSQLFactory(database.InstanceMySQL)
	if err != nil {
		logger.Fatal(err.Error())
	}

	albumRepository := gateway.NewAlbumRepository(db)
	albumUseCase := usecase.NewAlbumUseCase(albumRepository)
	albumAction := action.NewAlbumAction(albumUseCase)

	categoryRepository := gateway.NewCategoryRepository(db)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryAction := action.NewCategoryAction(categoryUseCase)

	app := &cli.App{}
	command.SetAlbumCommand(app, albumAction)
	command.SetCategoryCommand(app, categoryAction)
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))
	err = app.Run(os.Args)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
