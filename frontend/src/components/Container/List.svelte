<script lang="ts">
  import { onMount } from "svelte";
  import axios from "axios";
  import Cookies from "js-cookie";

  const user = Cookies.get("token");

  type Container = {
    name: string;
    password: string;
    port: number;
    token: string;
    type: string;
  };

  let containers: Container[] = [];
  let loading = true;
  let error: string | null = null;

  const API = "http://localhost:3000";

  async function fetchContainers() {
    loading = true;
    try {
      const res = await axios.post(`${API}/api/container/list`, { user });
      containers = res.data.data; // doğru kullanım
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

  onMount(fetchContainers);
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
            <h2 class="font-bold text-lg">{container.name}</h2>
            <p class="text-sm text-gray-600">Port: {container.port}</p>
            <p class="text-sm text-gray-600">Token: {container.token}</p>
            <p class="text-sm text-gray-600">Tür: {container.type}</p>
          </div>
          <div class="space-x-2">
            <button
              on:click={() => startContainer(container.token)}
              class="bg-green-500 text-white px-3 py-1 rounded hover:bg-green-600"
              >Başlat</button
            >
            <button
              on:click={() => stopContainer(container.token)}
              class="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600"
              >Durdur</button
            >
            <button
              on:click={() => deleteContainer(container.token)}
              class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600"
              >Sil</button
            >
          </div>
        </div>
      </li>
    {/each}
  </ul>
{/if}
