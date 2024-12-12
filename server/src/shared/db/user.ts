import type { CreateUserDTO, UserDTO } from "@/dto/user.dto"
import type { TX } from "./db"
import { usersTable } from "./schema/user"
import { eq, getTableColumns } from "drizzle-orm"
import { getUserRolesTx, setUserRolesTx } from "./role"

/**
 * @description Gets a user by its email address. If not found, null is returned.
 */
export const getUserByEmailTx = async (tx: TX, email: string): Promise<UserDTO | null> => {
  const { passwordHash, ...rest } = getTableColumns(usersTable)

  const results = await tx
    .select({ ...rest })
    .from(usersTable)
    .where(eq(usersTable.email, email))

  if (results.length <= 0) {
    return null
  }

  const roles = await getUserRolesTx(tx, results[0].id)

  return { ...results[0], roles }
}

/**
 * @description Gets a user by its ID. If not found, null is returned.
 */
export const getUserByIDTx = async (tx: TX, id: number): Promise<UserDTO | null> => {
  const { passwordHash, ...rest } = getTableColumns(usersTable)

  const results = await tx
    .select({ ...rest })
    .from(usersTable)
    .where(eq(usersTable.id, id))

  if (results.length <= 0) {
    return null
  }

  const roles = await getUserRolesTx(tx, results[0].id)

  return {
    ...results[0],
    roles
  }
}

/**
 * @description Insert a user into the database, returning the inserted user.
 */
export const insertUserTx = async (tx: TX, data: CreateUserDTO): Promise<UserDTO> => {
  const result = await tx
    .insert(usersTable)
    .values(data)
    .returning()

  const user = result[0]

  await setUserRolesTx(tx, user.id, data.roles)

  const roles = await getUserRolesTx(tx, user.id)

  return { ...user, roles }
}

/**
 * @description Gets a users password hash, if exists.
 */
export const getPasswordHashTx = async (tx: TX, userID: number): Promise<string | null> => {
  const results = await tx
    .select({
      passwordHash: usersTable.passwordHash,
    })
    .from(usersTable)
    .where(eq(usersTable.id, userID))

  if (results.length <= 0) {
    return null
  }

  return results[0].passwordHash
}
