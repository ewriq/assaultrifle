import { writable } from 'svelte/store';
import axios from 'axios';
import Cookies from "js-cookie";

export function useFileManager(containerId: string) {
    const user = Cookies.get('token');
    const path = writable("/");
    const files = writable<string[]>([]);
    const loading = writable(false);
    const error = writable<string | null>(null);
    const message = writable<string | null>(null);
    const selectedFile = writable<File | null>(null);

    function fileNameFromLine(file: string): string {
        const parts = file.trim().split(/\s+/);
        return parts[parts.length - 1];
    }

    async function fetchFiles(currentPath: string) {
        loading.set(true);
        error.set(null);
        message.set(null);
        try {
            const res = await axios.post("http://localhost:3000/api/container/file/list", {
                token: containerId,
                path: currentPath,
                user
            });
            if (res.data.status === "OK") {
                files.set(res.data.data);
            } else {
                error.set(res.data.error || "Dosyalar alınamadı");
            }
        } catch (e: any) {
            error.set("Sunucu hatası: " + e.message);
        }
        loading.set(false);
    }

    async function deleteFile(currentPath: string, fileName: string) {
        const filePath = currentPath.endsWith("/") ? currentPath + fileName : currentPath + "/" + fileName;
        try {
            const res = await axios.post("http://localhost:3000/api/container/file/del", {
                token: containerId,
                path: filePath,
                user
            });
            if (res.data.status === "OK") {
                message.set(res.data.message);
                await fetchFiles(currentPath);
            } else {
                error.set(res.data.error || "Dosya silinemedi");
            }
        } catch (e: any) {
            error.set("Sunucu hatası: " + e.message);
        }
    }

    async function uploadFile(currentPath: string, file: File) {
        const formData = new FormData();
        formData.append("file", file);
        formData.append("token", containerId);
        formData.append("target", currentPath);
        formData.append("user", user);

        try {
            const res = await axios.post("http://localhost:3000/api/container/file/add", formData, {
                headers: { "Content-Type": "multipart/form-data" }
            });
            if (res.data.status === "OK") {
                message.set(res.data.message);
                await fetchFiles(currentPath);
            } else {
                error.set(res.data.error || "Dosya yüklenemedi");
            }
        } catch (e: any) {
            error.set("Sunucu hatası: " + e.message);
        }
    }

    return {
        path,
        files,
        loading,
        error,
        message,
        selectedFile,
        fetchFiles,
        deleteFile,
        uploadFile,
        fileNameFromLine
    };
}
