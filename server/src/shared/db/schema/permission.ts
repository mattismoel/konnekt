import { integer, primaryKey, sqliteTable, text } from "drizzle-orm/sqlite-core";
import { usersTable } from "./user";

export const permissionsTable = sqliteTable("permission", {
  id: integer().primaryKey({ autoIncrement: true }),
  name: text().notNull(),
  description: text().notNull()
})

export const rolesTable = sqliteTable("role", {
  id: integer().primaryKey({ autoIncrement: true }),
  name: text().notNull(),
  description: text().notNull()
})

export const rolesPermissionsTable = sqliteTable("roles_permissions", {
  roleID: integer().notNull().references(() => rolesTable.id),
  permissionID: integer().notNull().references(() => permissionsTable.id),
}, (table) => ({
  pk: primaryKey({ columns: [table.roleID, table.permissionID] })
}))

export const usersRoles = sqliteTable("users_roles", {
  userID: integer().notNull().references(() => usersTable.id),
  roleID: integer().notNull().references(() => rolesTable.id),
}, (table) => ({
  pk: primaryKey({ columns: [table.userID, table.roleID] })
}))
