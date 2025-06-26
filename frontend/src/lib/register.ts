import axios from 'axios';

interface RegisterBody {
  email: string;
  password: string;
  username: string;
}

export async function registerUser(form: RegisterBody): Promise<string> {
  const API = 'http://localhost:3000/api/auth/register';

  try {
    const res = await axios.post(API, form);
    if (res.data.status === 'OK') {
      return res.data.token;
    } else {
      throw new Error(res.data.message || 'Kayıt başarısız.');
    }
  } catch (err: any) {
    console.error("Register error:", err);
    throw new Error('Sunucuya ulaşılamadı.');
  }
}
