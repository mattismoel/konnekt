/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1568971955")

  // add field
  collection.fields.addAt(4, new Field({
    "cascadeDelete": false,
    "collectionId": "pbc_3709660955",
    "hidden": false,
    "id": "relation770559087",
    "maxSelect": 999,
    "minSelect": 0,
    "name": "permissions",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "relation"
  }))

  // add field
  collection.fields.addAt(5, new Field({
    "cascadeDelete": false,
    "collectionId": "_pb_users_auth_",
    "hidden": false,
    "id": "relation1168167679",
    "maxSelect": 999,
    "minSelect": 0,
    "name": "members",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "relation"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1568971955")

  // remove field
  collection.fields.removeById("relation770559087")

  // remove field
  collection.fields.removeById("relation1168167679")

  return app.save(collection)
})
