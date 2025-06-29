import axios from 'axios';

export async function getStatus(token: string): Promise<string> {
  const res = await axios.post('http://localhost:3000/api/container/status', { token });
  return res.data.data.mem;
}

export async function getLogs(token: string): Promise<string> {
  const res = await axios.post('http://localhost:3000/api/container/log', { token });
  return res.data.data;
}

export async function startContainer(token: string): Promise<string> {
  const res = await axios.post('http://localhost:3000/api/container/start', { token });
  return res.data.message;
}

export async function stopContainer(token: string): Promise<string> {
  const res = await axios.post('http://localhost:3000/api/container/stop', { token });
  return res.data.message;
}

export async function restartContainer(token: string): Promise<string> {
  const res = await axios.post('http://localhost:3000/api/container/restart', { token });
  return res.data.message;
}

export async function deleteContainer(token: string): Promise<string> {
  const res = await axios.post('http://localhost:3000/api/container/del', { token });
  return res.data.message;
}

export function cleanLogs(text: string): string {
  text = text.replace(/\x1b\[[0-9;]*[a-zA-Z]/g, "");
  text = text.replace(/\[\d+[GKHf]/g, "");
  const lines = text.split('\n').map(line => {
    const parts = line.split('\r');
    return parts[parts.length - 1];
  });
  return lines.map(l => l.trimEnd()).join('\n');
}
