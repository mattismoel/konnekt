/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1687431684")

  // update collection data
  unmarshal({
    "createRule": "@request.auth.approved = true",
    "deleteRule": "@request.auth.approved = true",
    "updateRule": "@request.auth.approved = true"
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1687431684")

  // update collection data
  unmarshal({
    "createRule": "",
    "deleteRule": null,
    "updateRule": null
  }, collection)

  return app.save(collection)
})
