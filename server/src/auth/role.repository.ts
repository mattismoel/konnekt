import type { PermissionDTO, RoleDTO } from "./role.dto";

export interface RoleRepository {
  getUserRoles(userID: number): Promise<RoleDTO[]>
  getRolePermissions(roleID: number): Promise<PermissionDTO[]>
}
