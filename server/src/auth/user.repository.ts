import type { CreateUserDTO, UserDTO } from "./user.dto";

export interface UserRepository {
  getByEmail(email: string): Promise<UserDTO | null>
  getByID(id: number): Promise<UserDTO | null>
  getPasswordHash(userID: number): Promise<string | null>
  insert(user: CreateUserDTO): Promise<UserDTO>
}
