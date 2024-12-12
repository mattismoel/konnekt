import { createContext, useContext, useEffect, useState } from "react";
import { User } from "../dto/user.dto";

type UserContextProps = {
  user: User | null;
  setUser: (user: User) => void;
}

const UserContext = createContext<UserContextProps | null>(null)

export const UserProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<User | null>(null)

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}
    </UserContext.Provider>
  )
}

export const useUser = () => {
  const context = useContext(UserContext)

  if (!context) {
    throw new Error("useUser must be used within <UserContext.Provider>")
  }

  return context;
}
