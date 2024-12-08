import { db } from "@/shared/db/db";
import type { CreateUserDTO, UserDTO } from "./user.dto";
import type { UserRepository } from "./user.repository";
import { usersTable } from "@/shared/db/schema/user";
import { eq } from "drizzle-orm";

export class SQLiteUserRepository implements UserRepository {
  getByEmail = async (email: string): Promise<UserDTO | null> => {
    const results = await db
      .select()
      .from(usersTable)
      .where(eq(usersTable.email, email))

    if (results.length <= 0) {
      return null
    }

    return results[0]
  }

  getByID = async (id: number): Promise<UserDTO | null> => {
    const results = await db
      .select()
      .from(usersTable)
      .where(eq(usersTable.id, id))

    if (results.length <= 0) {
      return null
    }

    return results[0]
  }

  insert = async (user: CreateUserDTO): Promise<UserDTO> => {
    const result = await db
      .insert(usersTable)
      .values(user)
      .returning()

    return result[0]
  }
}
