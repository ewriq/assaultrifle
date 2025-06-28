import axios from "axios";

const API_URL = "http://localhost:3000/api/container/add";

export interface AddContainerRequest {
  name: string;
  password: string;
  port: string;
  user: string;
  type: string;
}

export async function addContainer(
  data: AddContainerRequest,
  onSuccess: (token: string) => void,
  onError: (msg: string) => void,
  onFinally?: () => void
) {
  try {
    const res = await axios.post(API_URL, data);

    if (res.data.status === "OK") {
      const token = res.data.token;
      onSuccess(token);
    } else {
      onError(res.data.error || "Bilinmeyen bir hata oluştu.");
    }
  } catch (error: any) {
    onError("Sunucu hatası: " + (error.message || "Bilinmiyor"));
  } finally {
    onFinally?.();
  }
}
