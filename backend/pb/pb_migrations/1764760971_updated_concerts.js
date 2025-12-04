/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_860379368")

  // add field
  collection.fields.addAt(4, new Field({
    "cascadeDelete": false,
    "collectionId": "pbc_1687431684",
    "hidden": false,
    "id": "relation1001261735",
    "maxSelect": 1,
    "minSelect": 0,
    "name": "event",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "relation"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_860379368")

  // remove field
  collection.fields.removeById("relation1001261735")

  return app.save(collection)
})
