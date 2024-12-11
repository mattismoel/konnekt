import type { PermissionDTO, RoleDTO } from "@/dto/role.dto";

export interface RoleRepository {
  getUserRoles(userID: number): Promise<RoleDTO[]>
  getRolePermissions(roleID: number): Promise<PermissionDTO[]>
}
