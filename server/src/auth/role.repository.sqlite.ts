import { db } from "@/shared/db/db";
import type { PermissionDTO, RoleDTO } from "./role.dto";
import type { RoleRepository } from "./role.repository";
import { permissionsTable, rolesPermissionsTable, rolesTable, usersRoles } from "@/shared/db/schema/permission";
import { eq } from "drizzle-orm";

export class SQLiteRoleRepository implements RoleRepository {
  getUserRoles = async (userID: number): Promise<RoleDTO[]> => {
    const results = await db
      .select({ role: rolesTable })
      .from(rolesTable)
      .innerJoin(usersRoles, eq(rolesTable.id, usersRoles.roleID))
      .where(eq(usersRoles.userID, userID))

    return results.map(result => result.role)
  }

  getRolePermissions = async (roleID: number): Promise<PermissionDTO[]> => {
    const results = await db
      .select({
        permission: permissionsTable
      })
      .from(rolesPermissionsTable)
      .innerJoin(permissionsTable, eq(rolesPermissionsTable.permissionID, permissionsTable.id))
      .where(eq(rolesPermissionsTable.roleID, roleID))

    return results.map(result => result.permission)
  }
}
