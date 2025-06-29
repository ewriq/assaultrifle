import axios from "axios";
import { goto } from "$app/navigation";

const API = "http://localhost:3000";

export async function getUser(token: string): Promise<any | null> {
  try {
    const res = await axios.post(`${API}/api/auth/user`, { token });

    if (res.data.status === "OK") {
      return res.data.data;
    } else {
      goto("/login");
      return null;
    }
  } catch (e) {
    console.error("Kullanıcı alınamadı:", e);
    goto("/login");
    return null;
  }
}
