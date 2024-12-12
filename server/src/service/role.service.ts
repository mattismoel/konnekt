import type { PermissionDTO, RoleDTO } from "@/dto/role.dto";
import type { RoleRepository } from "@/repository/role.repository";

export class RoleService {
  constructor(
    private readonly roleRepository: RoleRepository,
  ) { }

  getUserRoles = async (userID: number): Promise<RoleDTO[]> => {
    const roles = await this.roleRepository.getUserRoles(userID)
    return roles
  }

  getRolePermissions = async (roleID: number): Promise<PermissionDTO[]> => {
    const permissions = await this.roleRepository.getRolePermissions(roleID)
    return permissions
  }
}
