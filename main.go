package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"github.com/turnon/imdbtsv/tsv"
	_ "github.com/turnon/impb/migrations"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		return loadMovies(app)
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func loadMovies(app *pocketbase.PocketBase) error {
	collection, err := app.Dao().FindCollectionByNameOrId("movies")
	if err != nil {
		return err
	}

	// load
	limit := 1000
	return tsv.IterateTitleBasic("imdb_tsv/title.basics.tsv", func(tbr *tsv.TitleBasicRow) error {
		if limit == 0 {
			return nil
		}

		if tbr.TitleType != "movie" {
			return nil
		}

		limit -= 1

		record := models.NewRecord(collection)
		newRecord := forms.NewRecordUpsert(app, record)
		newRecord.LoadData(map[string]any{
			"title":      tbr.PrimaryTitle,
			"start_year": tbr.StartYear,
			"end_year":   tbr.EndYear,
			"minutes":    tbr.RuntimeMinutes,
			"type":       tbr.TitleType,
		})
		return newRecord.Submit()
	})
}
