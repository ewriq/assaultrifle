import { writable } from 'svelte/store';
import axios from 'axios';
import Cookies from 'js-cookie';

const API = 'http://localhost:3000';
const user = Cookies.get('token');

export type Container = {
  Name: string;
  Password: string;
  Port: number;
  Token: string;
  Types: string;
};

export const containers = writable<Container[]>([]);
export const loading = writable(true);
export const error = writable<string | null>(null);

export const fileStates = writable<Record<string, string[]>>({});
export const fileInputs = writable<Record<string, File | null>>({});

export async function fetchContainers() {
  loading.set(true);
  try {
    const res = await axios.post(`${API}/api/container/list`, { user });
    containers.set(res.data.data ?? []);
    error.set(null);
  } catch (err: any) {
    error.set(err.response?.data?.error || 'Bir hata oluştu');
  } finally {
    loading.set(false);
  }
}

export async function fetchFiles(token: string) {
  try {
    const res = await axios.post(`${API}/api/container/file/list`, {
      token,
      path: '/home/',
      user,
    });

    if (res.data.status === 'OK') {
      const lines: string[] = res.data.data ?? [];
      const files = lines
        .filter(line => !line.startsWith('total'))
        .map(line => {
          const parts = line.trim().split(/\s+/);
          return parts.slice(8).join(' ');
        });
      fileStates.update(fs => ({ ...fs, [token]: files }));
    } else {
      fileStates.update(fs => ({ ...fs, [token]: [] }));
    }
  } catch (err) {
    console.error('Dosya listesi alınamadı:', err);
    fileStates.update(fs => ({ ...fs, [token]: [] }));
  }
}

export async function startContainer(token: string) {
  await axios.post(`${API}/api/container/start`, { token });
  await fetchContainers();
}

export async function stopContainer(token: string) {
  await axios.post(`${API}/api/container/stop`, { token });
  await fetchContainers();
}

export async function deleteContainer(token: string) {
  await axios.post(`${API}/api/container/del`, { token });
  await fetchContainers();
}

export async function statusContainer(token: string) {
  await axios.post(`${API}/api/container/status`, { token });
  await fetchContainers();
}

export async function uploadFile(token: string, file: File | null) {
  if (!file) return;

  const formData = new FormData();
  formData.append('token', token);
  formData.append('target', '/home/');
  formData.append('user', user || '');
  formData.append('file', file);

  try {
    await axios.post(`${API}/api/container/file/add`, formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    });
    await fetchFiles(token);
  } catch (err) {
    console.error('Yükleme hatası:', err);
  }
}

export async function deleteFile(token: string, filename: string) {
  await axios.post(`${API}/api/container/file/del`, {
    token,
    path: '/home/' + filename,
    user,
  });
  await fetchFiles(token);
}

export function handleFileChange(token: string, file: File | null) {
  fileInputs.update(fi => ({ ...fi, [token]: file }));
}
