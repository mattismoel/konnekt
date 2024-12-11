import type { PermissionDTO, RoleDTO } from "@/dto/role.dto"
import type { TX } from "./db"
import { permissionsTable, rolesPermissionsTable, rolesTable, usersRoles } from "./schema/permission"
import { eq } from "drizzle-orm"

export const getUserRolesTx = async (tx: TX, userID: number): Promise<RoleDTO[]> => {
  const results = await tx
    .select({ role: rolesTable })
    .from(rolesTable)
    .innerJoin(usersRoles, eq(rolesTable.id, usersRoles.roleID))
    .where(eq(usersRoles.userID, userID))

  return results.map(result => result.role)
}

export const getRolePermissionsTx = async (tx: TX, roleID: number): Promise<PermissionDTO[]> => {
  const results = await tx
    .select({
      permission: permissionsTable
    })
    .from(rolesPermissionsTable)
    .innerJoin(permissionsTable, eq(rolesPermissionsTable.permissionID, permissionsTable.id))
    .where(eq(rolesPermissionsTable.roleID, roleID))

  return results.map(result => result.permission)
}
