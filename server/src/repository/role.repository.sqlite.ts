import { db } from "@/shared/db/db";
import type { PermissionDTO, RoleDTO } from "@/dto/role.dto";
import type { RoleRepository } from "./role.repository";
import { getRolePermissionsTx, getUserRolesTx, setUserRolesTx } from "@/shared/db/role";

export const createSQLiteRoleRepository = (): RoleRepository => {
  const getUserRoles = async (userID: number): Promise<RoleDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getUserRolesTx(tx, userID)
    })
  }

  const getRolePermissions = async (roleID: number): Promise<PermissionDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getRolePermissionsTx(tx, roleID)
    })
  }

  const setUserRoles = async (userID: number, roles: string[]) => {
    return await db.transaction(async (tx) => {
      return await setUserRolesTx(tx, userID, roles)
    })
  }

  return {
    getRolePermissions,
    getUserRoles,
    setUserRoles
  }
}
