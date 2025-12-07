/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_3709660955")

  // add field
  collection.fields.addAt(1, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text1579384326",
    "max": 0,
    "min": 0,
    "name": "name",
    "pattern": "",
    "presentable": true,
    "primaryKey": false,
    "required": true,
    "system": false,
    "type": "text"
  }))

  // update field
  collection.fields.addAt(2, new Field({
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
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_3709660955")

  // remove field
  collection.fields.removeById("text1579384326")

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
})
