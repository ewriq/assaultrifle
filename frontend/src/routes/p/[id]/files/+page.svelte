
<script lang="ts">
    import { page } from '$app/stores';
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';
    import axios from 'axios';
    import Cookies from "js-cookie";
  
    const id = get(page).params.id;
    const user = Cookies.get('token');
  
    let path = "/";
    let files: string[] = [];
    let loading = false;
    let error: string | null = null;
    let message: string | null = null;
    let selectedFile: File | null = null;
  
    function fileNameFromLine(file: string): string {
      const parts = file.trim().split(/\s+/);
      return parts[parts.length - 1];
    }
  
    async function fetchFiles() {
      loading = true;
      error = null;
      message = null;
      try {
        const res = await axios.post("http://localhost:3000/api/container/file/list", { token: id, path, user });
        if (res.data.status === "OK") {
          files = res.data.data;
        } else {
          error = res.data.error || "Dosyalar alınamadı";
        }
      } catch (e: any) {
        error = "Sunucu hatası: " + e.message;
      }
      loading = false;
    }
  
    async function deleteFile(fileName: string) {
      if (!confirm(`${fileName} dosyasını silmek istediğinize emin misiniz?`)) return;
      error = null;
      message = null;
      try {
        const filePath = path.endsWith("/") ? path + fileName : path + "/" + fileName;
        const res = await axios.post("http://localhost:3000/api/container/file/del", {
          token: id,
          path: filePath,
          user: user,
        });
        if (res.data.status === "OK") {
          message = res.data.message;
          await fetchFiles();
        } else {
          error = res.data.error || "Dosya silinemedi";
        }
      } catch (e: any) {
        error = "Sunucu hatası: " + e.message;
      }
    }
  
    async function uploadFile() {
      if (!selectedFile) {
        alert("Lütfen yüklemek için bir dosya seçin");
        return;
      }
      error = null;
      message = null;
  
      const formData = new FormData();
      formData.append("file", selectedFile);
      formData.append("token", id);
      formData.append("target", path);
      formData.append("user", user);
  
      try {
        const res = await axios.post("http://localhost:3000/api/container/file/add", formData, {
          headers: {
            "Content-Type": "multipart/form-data"
          }
        });
        if (res.data.status === "OK") {
          message = res.data.message;
          selectedFile = null;
          await fetchFiles();
        } else {
          error = res.data.error || "Dosya yüklenemedi";
        }
      } catch (e: any) {
        error = "Sunucu hatası: " + e.message;
      }
    }
  
    $: if (path) {
      fetchFiles();
    }
  
    onMount(() => {
      fetchFiles();
    });
  </script>
  
  <div class="p-4 max-w-xl mx-auto border rounded shadow">
    <h1 class="text-xl font-semibold mb-3">Container Dosyaları - {id}</h1>
  
    <label class="block font-medium mb-2" for="path-input">Path Seçin:</label>
    <input
      id="path-input"
      type="text"
      bind:value={path}
      placeholder="/"
      class="w-full mb-4 p-2 border rounded"
    />
  
    {#if loading}
      <p>Yükleniyor...</p>
    {:else}
      {#if error}
        <p class="text-red-600 mb-2">{error}</p>
      {/if}
  
      {#if message}
        <p class="text-green-600 mb-2">{message}</p>
      {/if}
  
      <ul class="mb-4">
        {#each files as file}
          <li class="flex justify-between items-center border-b py-1">
            <span>{fileNameFromLine(file)}</span>
            <button
              class="text-red-600 hover:text-red-800"
              on:click={() => deleteFile(fileNameFromLine(file))}
              title="Dosyayı Sil"
            >
              Sil
            </button>
          </li>
        {/each}
        {#if files.length === 0}
          <li>Dosya bulunamadı.</li>
        {/if}
      </ul>
    {/if}
  
    <div class="mb-2">
      <label class="block mb-1 font-medium" for="file-upload">Dosya Yükle</label>
      <input
        id="file-upload"
        type="file"
        on:change={(e) => (selectedFile = e.target.files?.[0] ?? null)}
      />
    </div>
  
    <button
      class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700"
      on:click={uploadFile}
      disabled={!selectedFile}
    >
      Yükle
    </button>


  </div>
  