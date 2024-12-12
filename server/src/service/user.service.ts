import type { UserDTO } from "@/dto/user.dto";
import type { UserRepository } from "@/repository/user.repository";

export type UserService = {
  getUserByID(id: number): Promise<UserDTO | null>
}

export const createUserService = (userRepo: UserRepository): UserService => {
  const getUserByID = async (id: number): Promise<UserDTO | null> => {
    return await userRepo.getUserByID(id)
  }

  return {
    getUserByID,
  }
}
