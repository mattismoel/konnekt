import { integer, primaryKey, sqliteTable, text } from "drizzle-orm/sqlite-core";
import { usersTable } from "./user";

export const permissionsTable = sqliteTable("permission", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  name: text("name").notNull(),
  description: text("description").notNull()
})

export const rolesTable = sqliteTable("role", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  name: text("name").notNull(),
  description: text("description").notNull()
})

export const rolesPermissionsTable = sqliteTable("roles_permissions", {
  roleID: integer("role_id").notNull().references(() => rolesTable.id),
  permissionID: integer("permission_id").notNull().references(() => permissionsTable.id),
}, (table) => ({
  pk: primaryKey({ columns: [table.roleID, table.permissionID] })
}))

export const usersRoles = sqliteTable("users_roles", {
  userID: integer("user_id").notNull().references(() => usersTable.id),
  roleID: integer("role_id").notNull().references(() => rolesTable.id),
}, (table) => ({
  pk: primaryKey({ columns: [table.userID, table.roleID] })
}))
