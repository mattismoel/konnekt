import type { AuthService } from "@/service/auth.service";
import type { RoleService } from "@/service/role.service";
import { SESSION_COOKIE_NAME } from "@/shared/auth/constant";
import type { RequestHandler } from "express";

export const checkPermissions = (
  authService: AuthService,
  roleService: RoleService,
  permissions: string[],
): RequestHandler => {
  return async (req, res, next) => {
    const token = req.cookies[SESSION_COOKIE_NAME] as string || null

    if (!token) {
      res.sendStatus(401)
      return
    }

    const { session, user } = await authService.validateSessionToken(token)

    if (!session) {
      res.sendStatus(401)
      return
    }

    const roles = await roleService.getUserRoles(user.id)

    const rolesPermissions = await Promise.all(
      roles.map((role) => roleService.getRolePermissions(role.id))
    )

    const userPermissions = rolesPermissions.flat().map(p => p.name)

    if (!permissionsMatch(permissions, userPermissions)) {
      res.sendStatus(401)
      return
    }

    next()
  }
}

const permissionsMatch = (permissions: string[], userPermissions: string[]): boolean => {
  return permissions.every(permission => userPermissions.includes(permission))
}
