import type { AuthService } from "@/auth/auth.service";
import { SESSION_COOKIE_NAME } from "@/auth/constant";
import type { RoleService } from "@/auth/role.service";
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

    console.log("Before")
    const roles = await roleService.getUserRoles(user.id)
    console.log("After")


    const rolesPermissions = await Promise.all(
      roles.map((role) => roleService.getRolePermissions(role.id))
    )

    const userPermissions = rolesPermissions.flat().map(p => p.name)
    console.log(`${user.email}\nroles: ${roles}\npermissions: ${userPermissions}`)

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
