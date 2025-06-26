import axios from "axios";

/**
 * Giriş yapan kullanıcıdan gelen email ve şifreyi sunucuya gönderir.
 * Başarılıysa token döner, değilse hata fırlatır.
 * 
 * @param email Kullanıcı email adresi
 * @param password Kullanıcı şifresi
 * @returns token (string)
 **/

export async function loginUser(email: string, password: string): Promise<string> {
  const API_URL = "http://localhost:3000/api/auth/login";

  try {
    const response = await axios.post(API_URL, { email, password });

    if (response.data.status === "OK") {
      return response.data.data;
    } else {
      throw new Error("Email ya da parola hatalı");
    }
  } catch (err: any) {
    console.error("Giriş hatası:", err);
    throw new Error("Sunucu hatası oluştu.");
  }
}
