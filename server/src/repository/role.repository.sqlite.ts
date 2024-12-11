import { db } from "@/shared/db/db";
import type { PermissionDTO, RoleDTO } from "@/dto/role.dto";
import type { RoleRepository } from "./role.repository";
import { getRolePermissionsTx, getUserRolesTx } from "@/shared/db/role";

export class SQLiteRoleRepository implements RoleRepository {
  getUserRoles = async (userID: number): Promise<RoleDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getUserRolesTx(tx, userID)
    })
  }

  getRolePermissions = async (roleID: number): Promise<PermissionDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getRolePermissionsTx(tx, roleID)
    })
  }
}
