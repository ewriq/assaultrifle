import axios from "axios";
import Cookies from "js-cookie";

export interface Container {
  Name: string;
  Password: string;
  Port: string;
  Token: string;
  Types: string;
}

interface FetchResult {
  data: Container[];
  error: string | null;
}

export async function fetchAllContainers(): Promise<FetchResult> {
  const token = Cookies.get("token");
  if (!token) {
    return { data: [], error: "Kimlik doğrulaması başarısız." };
  }

  try {
    const res = await axios.post("http://localhost:3000/api/container/listmy", {
      user: token,
    });

    if (res.data.status === "OK") {
      return {
        data: res.data.data as Container[],
        error: null,
      };
    }

    return {
      data: [],
      error: res.data.error || "Sunucu hatası oluştu.",
    };
  } catch (e: any) {
    return {
      data: [],
      error: e.message || "Bilinmeyen hata oluştu.",
    };
  }
}
