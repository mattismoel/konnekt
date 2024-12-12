import type { UserDTO } from "@/dto/user.dto";
import type { UserRepository } from "@/repository/user.repository";

export class UserService {
  constructor(
    private readonly userRepo: UserRepository,
  ) { }

  getByID = async (id: number): Promise<UserDTO | null> => {
    return await this.userRepo.getByID(id)
  }
}
