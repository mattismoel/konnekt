import type { CreateUserDTO, UserDTO } from "@/dto/user.dto";

export interface UserRepository {
  getUserByEmail(email: string): Promise<UserDTO | null>
  getUserByID(id: number): Promise<UserDTO | null>
  getUserPasswordHash(userID: number): Promise<string | null>
  insertUser(user: CreateUserDTO): Promise<UserDTO>
}
