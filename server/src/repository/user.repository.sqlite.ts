import { db } from "@/shared/db/db";
import type { CreateUserDTO, UserDTO } from "@/dto/user.dto";
import type { UserRepository } from "./user.repository";
import { getPasswordHashTx, getUserByEmailTx, getUserByIDTx, insertUserTx } from "@/shared/db/user";

export const createSQLiteUserRepository = (): UserRepository => {
  const insertUser = async (user: CreateUserDTO): Promise<UserDTO> => {
    return await db.transaction(async (tx) => {
      const insertedUser = await insertUserTx(tx, user)
      return insertedUser
    })
  }

  const getUserByID = async (id: number): Promise<UserDTO | null> => {
    return await db.transaction(async (tx) => {
      const user = await getUserByIDTx(tx, id)
      return user
    })
  }

  const getUserByEmail = async (email: string): Promise<UserDTO | null> => {
    return await db.transaction(async (tx) => {
      const user = await getUserByEmailTx(tx, email)
      return user
    })
  }

  const getUserPasswordHash = async (userID: number): Promise<string | null> => {
    return await db.transaction(async (tx) => {
      const passwordHash = await getPasswordHashTx(tx, userID)
      return passwordHash
    })
  }

  return {
    insertUser,
    getUserByID,
    getUserByEmail,
    getUserPasswordHash
  }
}
