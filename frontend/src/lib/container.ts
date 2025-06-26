import axios from "axios";
import Cookies from "js-cookie";

const user = Cookies.get("token") || "";
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

export async function fetchContainers(): Promise<Container[]> {
  const res = await axios.post(`${API}/api/container/list`, { user });
  return res.data.data;
}

export async function deleteContainer(token: string) {
  await axios.post(`${API}/api/container/del`, { token });
}

export async function startContainer(token: string) {
  await axios.post(`${API}/api/container/start`, { token });
}

export async function stopContainer(token: string) {
  await axios.post(`${API}/api/container/stop`, { token });
}

export async function statusContainer(token: string) {
  await axios.post(`${API}/api/container/status`, { token });
}

export function handleFileChange(token: string, file: File | null) {
  fileInputs[token] = file;
}

export async function uploadFile(token: string, file: File | null) {
  if (!file) return;

  const formData = new FormData();
  formData.append("token", token);
  formData.append("target", "/home/");
  formData.append("user", user);
  formData.append("file", file);

  await axios.post(`${API}/api/container/file/add`, formData, {
    headers: {
      "Content-Type": "multipart/form-data"
    }
  });

  await fetchFiles(token);
}

export async function fetchFiles(token: string) {
  try {
    const res = await axios.post(`${API}/api/container/file/list`, {
      token,
      path: "/home/",
      user
    });

    if (res.data.status === "OK") {
      const lines: string[] = res.data.data;
      const files = lines
        .filter(line => !line.startsWith("total"))
        .map(line => line.trim().split(/\s+/).slice(8).join(" "));
      fileStates[token] = files;
    } else {
      fileStates[token] = [];
    }
  } catch {
    fileStates[token] = [];
  }
}

export async function deleteFile(token: string, paths: string) {
  await axios.post(`${API}/api/container/file/del`, {
    token,
    path: "/home/" + paths,
    user
  });
  await fetchFiles(token);
}
