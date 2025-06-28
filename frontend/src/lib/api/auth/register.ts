import axios from 'axios';
import Cookies from 'js-cookie';
import { goto } from '$app/navigation';

export async function registerUser(email: string, password: string, username: string, onSuccess: () => void, onError: (msg: string) => void) {
  const form = { email, password, username };

  try {
    const response = await axios.post('http://localhost:3000/api/auth/register', form);

    if (response.data.status === 'OK') {
      const token = response.data.token;
      Cookies.set('token', token);
      onSuccess();
    } else {
      onError('Bir hata oluştu. Lütfen bilgilerinizi kontrol edin.');
    }
  } catch (err) {
    console.error('İstek hatası:', err);
    onError('Sunucuya bağlanılamadı.');
  }
}
