  <script lang="ts">
    import { onMount } from "svelte";
    import axios from "axios";
    import Cookies from "js-cookie";
    import Terminal from "./Cmd.svelte";

    const user = Cookies.get("token");

    type Container = {
      Name: string;
      Password: string;
      Port: number;
      Token: string;
      Types: string;
    };

    let containers: Container[] = [];
    let loading = true;
    let error: string | null = null;

    const API = "http://localhost:3000";

    let showModal = false;
    let activeContainerToken = "";
    let path = "/home/";
    let fileStates: Record<string, string[]> = {};
    let fileInputs: Record<string, File | null> = {};

    async function fetchContainers() {
      loading = true;
      try {
        const res = await axios.post(`${API}/api/container/list`, { user });
        containers = res.data.data ?? [];
      } catch (err: any) {
        error = err.response?.data?.error || "Bir hata oluştu";
      } finally {
        loading = false;
      }
    }

    async function deleteContainer(token: string) {
      await axios.post(`${API}/api/container/del`, { token });
      await fetchContainers();
    }

    async function startContainer(token: string) {
      await axios.post(`${API}/api/container/start`, { token });
      await fetchContainers();
    }

    async function stopContainer(token: string) {
      await axios.post(`${API}/api/container/stop`, { token });
      await fetchContainers();
    }

    async function statusContainer(token: string) {
      await axios.post(`${API}/api/container/status`, { token });
      await fetchContainers();
    }

    function openTerminal(token: string) {
      activeContainerToken = token;
      showModal = true;
    }

    function closeModal() {
      showModal = false;
      activeContainerToken = "";
    }

    async function uploadFile(token: string, file: File | null) {
      if (!file) return;

      const formData = new FormData();
      formData.append("token", token);
      formData.append("target", path);
      formData.append("user", user || "");
      formData.append("file", file);
      try {
        await axios.post(`${API}/api/container/file/add`, formData, {
          headers: {
            "Content-Type": "multipart/form-data"
          }
        });
        await fetchFiles(token);
      } catch (err) {
        console.error("Yükleme hatası:", err);
      }
    }

    async function fetchFiles(token: string) {
      try {
        const res = await axios.post(`${API}/api/container/file/list`, {
          token,
          path,
          user
        });

        if (res.data.status === "OK") {
          const lines: string[] = res.data.data ?? [];
          const files = lines
            .filter(line => !line.startsWith("total"))
            .map(line => {
              const parts = line.trim().split(/\s+/);
              return parts.slice(8).join(" ");
            });
          fileStates[token] = files ?? [];
        } else {
          fileStates[token] = [];
        }
      } catch (err) {
        console.error("Dosya listesi alınamadı:", err);
        fileStates[token] = [];
      }
    }

    async function deleteFile(token: string, paths: string) {
      await axios.post(`${API}/api/container/file/del`, {
        token,
        path: path + paths,
        user
      });
      await fetchFiles(token);
    }

    function handleFileChange(token: string, file: File | null) {
      fileInputs[token] = file;
    }

    onMount(async () => {
      await fetchContainers();
      for (const c of containers) {
        await fetchFiles(c.Token);
      }
    });
  </script>

  {#if loading}
    <p>Yükleniyor...</p>
  {:else if error}
    <p class="text-red-500">Hata: {error}</p>
  {:else if containers.length === 0}
    <p>Hiç konteyner bulunamadı.</p>
  {:else}
    <ul class="space-y-4">
      {#each containers as container}
        <li class="p-4 rounded-lg border shadow bg-white">
          <div class="flex justify-between items-center">
            <div>
              <h2 class="font-bold text-lg">{container.Name}</h2>
              <p class="text-sm text-gray-600">Port: {container.Port}</p>
              <p class="text-sm text-gray-600">Token: {container.Token}</p>
              <p class="text-sm text-gray-600">Tür: {container.Types}</p>
            </div>
            <div class="space-x-2 flex items-center">
              <button on:click={() => startContainer(container.Token)} class="bg-green-500 text-white px-3 py-1 rounded hover:bg-green-600">Başlat</button>
              <button on:click={() => stopContainer(container.Token)} class="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600">Durdur</button>
              <button on:click={() => deleteContainer(container.Token)} class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600">Sil</button>
              <button on:click={() => statusContainer(container.Token)} class="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600">Durum</button>
              <button on:click={() => openTerminal(container.Token)} class="bg-indigo-600 text-white px-3 py-1 rounded hover:bg-indigo-700">Terminal Aç</button>
            </div>
          </div>

          <!-- Dosya Yönetimi -->
          <div class="mt-4 p-3 bg-gray-50 rounded border">
            <h3 class="font-semibold mb-2">📁 Dosya Yönetimi</h3>

            <div class="flex gap-2 mb-2">
              <input
                type="file"
                on:change={(e) => handleFileChange(container.Token, e.target.files?.[0] || null)}
              />
              <button
                class="bg-indigo-600 text-white px-3 py-1 rounded hover:bg-indigo-700"
                on:click={() => uploadFile(container.Token, fileInputs[container.Token])}
              >
                Yükle
              </button>
            </div>

            {#if Array.isArray(fileStates[container.Token]) && fileStates[container.Token].length > 0}
              <ul class="space-y-1">
                {#each fileStates[container.Token] ?? [] as file}
                  <li class="flex justify-between items-center border px-2 py-1 rounded text-sm">
                    <span>{file}</span>
                    <button
                      on:click={() => deleteFile(container.Token, file)}
                      class="bg-red-500 text-white px-2 py-0.5 rounded text-xs"
                    >
                      Sil
                    </button>
                  </li>
                {/each}
              </ul>
            {:else}
              <p class="text-sm text-gray-500">Hiç dosya yok.</p>
            {/if}
          </div>
        </li>
      {/each}
    </ul>
  {/if}

  {#if showModal}
    <div class="fixed inset-0 bg-black bg-opacity-60 flex items-center justify-center z-50" on:click={closeModal}>
      <div class="bg-white p-6 rounded-lg max-w-4xl w-full" on:click|stopPropagation>
        <button class="mb-4 px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600" on:click={closeModal}>Kapat</button>
        <Terminal containerId={activeContainerToken} />
      </div>
    </div>
  {/if}
