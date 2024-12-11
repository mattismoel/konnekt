import type { CreateUserDTO, UserDTO } from "@/dto/user.dto"
import type { TX } from "./db"
import { usersTable } from "./schema/user"
import { eq } from "drizzle-orm"

/**
 * @description Gets a user by its email address. If not found, null is returned.
 */
export const getUserByEmailTx = async (tx: TX, email: string): Promise<UserDTO | null> => {
  const results = await tx
    .select()
    .from(usersTable)
    .where(eq(usersTable.email, email))

  if (results.length <= 0) {
    return null
  }

  return results[0]
}

/**
 * @description Gets a user by its ID. If not found, null is returned.
 */
export const getUserByIDTx = async (tx: TX, id: number): Promise<UserDTO | null> => {
  const results = await tx
    .select()
    .from(usersTable)
    .where(eq(usersTable.id, id))

  if (results.length <= 0) {
    return null
  }

  return results[0]
}

/**
 * @description Insert a user into the database, returning the inserted user.
 */
export const insertUserTx = async (tx: TX, user: CreateUserDTO): Promise<UserDTO> => {
  const result = await tx
    .insert(usersTable)
    .values(user)
    .returning()

  return result[0]
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
