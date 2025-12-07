/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_3709660955")

  // update field
  collection.fields.addAt(1, new Field({
    "hidden": false,
    "id": "select1204587666",
    "maxSelect": 1,
    "name": "action",
    "presentable": true,
    "required": false,
    "system": false,
    "type": "select",
    "values": [
      "create",
      "edit",
      "delete"
    ]
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_3709660955")

  // update field
  collection.fields.addAt(1, new Field({
    "hidden": false,
    "id": "select1204587666",
    "maxSelect": 1,
    "name": "action",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "select",
    "values": [
      "create",
      "edit",
      "delete"
    ]
  }))

  return app.save(collection)
})
