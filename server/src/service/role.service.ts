import type { PermissionDTO, RoleDTO } from "@/dto/role.dto";
import type { RoleRepository } from "@/repository/role.repository";

export type RoleService = {
  getUserRoles(userID: number): Promise<RoleDTO[]>
  getRolePermissions(roleID: number): Promise<PermissionDTO[]>
}

export const createRoleService = (roleRepo: RoleRepository): RoleService => {
  const getUserRoles = async (userID: number): Promise<RoleDTO[]> => {
    const roles = await roleRepo.getUserRoles(userID)
    return roles
  }

  const getRolePermissions = async (roleID: number): Promise<PermissionDTO[]> => {
    const permissions = await roleRepo.getRolePermissions(roleID)
    return permissions
  }

  return {
    getUserRoles,
    getRolePermissions,
  }
}
