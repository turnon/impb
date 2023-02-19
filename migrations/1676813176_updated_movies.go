package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("fech3zcely4ejkf")
		if err != nil {
			return err
		}

		// add
		new_end_year := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "g1zt85qk",
			"name": "end_year",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_end_year)
		collection.Schema.AddField(new_end_year)

		// add
		new_minutes := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "goxrzmfw",
			"name": "minutes",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_minutes)
		collection.Schema.AddField(new_minutes)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("fech3zcely4ejkf")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("g1zt85qk")

		// remove
		collection.Schema.RemoveField("goxrzmfw")

		return dao.SaveCollection(collection)
	})
}
