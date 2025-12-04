/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_860379368")

  // update collection data
  unmarshal({
    "listRule": ""
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_860379368")

  // update collection data
  unmarshal({
    "listRule": null
  }, collection)

  return app.save(collection)
})
