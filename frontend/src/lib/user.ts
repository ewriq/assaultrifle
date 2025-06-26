import axios from "axios";

interface User {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: null | string;
    Username: string;
    Password: string;
    Email: string;
    Token: string;
    Perm: string;
  }

export async function fetchUser(token: string): Promise<User> {
  const API = "http://localhost:3000/api/auth/user";

  const res = await axios.post(API, { token });

  if (res.data.status === "OK" && res.data.data?.length > 0) {
    return res.data.data[0];
  } else {
    throw new Error("Kullanıcı bulunamadı.");
  }
}
