import axios from "axios";
import Cookies from "js-cookie";
import { goto } from "$app/navigation";

export async function loginUser(
  email: string,
  password: string,
  onSuccess: () => void,
  onError: (msg: string) => void
) {
  try {
    const form = { email, password };
    const response = await axios.post("http://localhost:3000/api/auth/login", form);

    if (response.data.status === "OK") {
      Cookies.set("token", response.data.data);
      onSuccess();
    } else {
      onError("Email ya da parola hatalı");
    }
  } catch (err) {
    console.error("Hata oluştu:", err);
    onError("Sunucu hatası oluştu.");
  }
}
