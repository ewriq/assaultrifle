import axios from "axios";
import Cookies from "js-cookie";

const API = "http://localhost:3000";

export async function getUser(
  onSuccess: (user: any) => void,
  onFailure: () => void
) {
  const token = Cookies.get("token");

  if (!token) {
    onFailure();
    return;
  }

  try {
    const res = await axios.post(`${API}/api/auth/user`, { token });

    if (res.data.status === "OK") {
      onSuccess(res.data.data);
    } else {
      onFailure();
    }
  } catch (err) {
    console.error("Kullanıcı verisi alınamadı:", err);
    onFailure();
  }
}
