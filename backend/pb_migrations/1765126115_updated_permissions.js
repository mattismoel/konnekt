/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_3709660955")

  // update collection data
  unmarshal({
    "indexes": [
      "CREATE INDEX `idx_2vvNDyjVBk` ON `permissions` (`name`)"
    ]
  }, collection)

  // remove field
  collection.fields.removeById("select1204587666")

  // remove field
  collection.fields.removeById("select4232930610")

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_3709660955")

  // update collection data
  unmarshal({
    "indexes": [
      "CREATE UNIQUE INDEX `idx_2vvNDyjVBk` ON `permissions` (\n  `action`,\n  `collection`\n)"
    ]
  }, collection)

  // add field
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

  // add field
  collection.fields.addAt(3, new Field({
    "hidden": false,
    "id": "select4232930610",
    "maxSelect": 1,
    "name": "collection",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "select",
    "values": [
      "artists",
      "events",
      "content",
      "venues"
    ]
  }))

  return app.save(collection)
})
