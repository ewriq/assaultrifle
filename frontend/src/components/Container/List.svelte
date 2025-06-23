<script lang="ts">
  import { onMount } from "svelte";
  import axios from "axios";
  import Cookies from "js-cookie";

  // Terminal component
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
            <h2 class="font-bold text-lg">{container.Name}</h2>
            <p class="text-sm text-gray-600">Port: {container.Port}</p>
            <p class="text-sm text-gray-600">Token: {container.Token}</p>
            <p class="text-sm text-gray-600">Tür: {container.Types}</p>
          </div>
          <div class="space-x-2 flex items-center">
            <button
              on:click={() => startContainer(container.Token)}
              class="bg-green-500 text-white px-3 py-1 rounded hover:bg-green-600"
            >Başlat</button>

            <button
              on:click={() => stopContainer(container.Token)}
              class="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600"
            >Durdur</button>

            <button
              on:click={() => deleteContainer(container.Token)}
              class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600"
            >Sil</button>

            <button
              on:click={() => statusContainer(container.Token)}
              class="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600"
            >Durum</button>

            <button
              on:click={() => openTerminal(container.Token)}
              class="bg-indigo-600 text-white px-3 py-1 rounded hover:bg-indigo-700"
            >Terminal Aç</button>
          </div>
        </div>
      </li>
    {/each}
  </ul>
{/if}

{#if showModal}
  <div
    class="fixed inset-0 bg-black bg-opacity-60 flex items-center justify-center z-50"
    on:click={closeModal}
  >
    <div
      class="bg-white p-6 rounded-lg max-w-4xl w-full"
      on:click|stopPropagation
    >
      <button
        class="mb-4 px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600"
        on:click={closeModal}
      >
        Kapat
      </button>

      <Terminal containerId={activeContainerToken} />
    </div>
  </div>
{/if}
