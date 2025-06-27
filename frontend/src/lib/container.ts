import axios from "axios";
import Cookies from "js-cookie";

const user = Cookies.get("token");
const API = "http://localhost:3000";

export type Container = {
  Name: string;
  Password: string;
  Port: number;
  Token: string;
  Types: string;
};

export let fileStates: Record<string, string[]> = {};
export let fileInputs: Record<string, File | null> = {};

export const fetchContainers = async (): Promise<Container[]> => {
  const res = await axios.post(`${API}/api/container/list`, { user });
  return res.data.data;
};

export const containerActions = {
  delete: (token: string) => axios.post(`${API}/api/container/del`, { token }),
  start: (token: string) => axios.post(`${API}/api/container/start`, { token }),
  stop: (token: string) => axios.post(`${API}/api/container/stop`, { token }),
  status: (token: string) => axios.post(`${API}/api/container/status`, { token }),
};

export const handleFileChange = (token: string, file: File | null) => {
  fileInputs[token] = file;
};

export const uploadFile = async (token: string, file: File | null) => {
  if (!file) return;
  const formData = new FormData();
  formData.append("token", token);
  formData.append("target", "/home/");
  formData.append("user", user);
  formData.append("file", file);

  await axios.post(`${API}/api/container/file/add`, formData, {
    headers: { "Content-Type": "multipart/form-data" }
  });

  await fetchFiles(token);
};

export const fetchFiles = async (token: string) => {
  try {
    const res = await axios.post(`${API}/api/container/file/list`, {
      token,
      path: "/",
      user
    });

    const lines: string[] = res.data.data;
    fileStates[token] = lines
      .filter(line => !line.startsWith("total"))
      .map(line => line.trim().split(/\s+/).slice(8).join(" "));
  } catch {
    fileStates[token] = [];
  }
};

export const deleteFile = async (token: string, path: string) => {
  await axios.post(`${API}/api/container/file/del`, {
    token,
    path: `/home/${path}`,
    user
  });
  await fetchFiles(token);
};
