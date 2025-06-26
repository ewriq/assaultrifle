<script lang="ts">
  import { onMount } from "svelte";
  import { fetchContainers, startContainer, stopContainer, deleteContainer, statusContainer, 
           fileStates, fileInputs, handleFileChange, uploadFile, fetchFiles, deleteFile } from "$lib/container;
  import TerminalModal from "./Cmd.svelte";

  let containers = [];
  let loading = true;
  let error: string | null = null;
  let showModal = false;
  let activeContainerToken = "";

  async function loadContainers() {
    loading = true;
    try {
      containers = await fetchContainers();
      // Her konteyner için dosya listesini getir
      await Promise.all(containers.map(c => fetchFiles(c.Token)));
      error = null;
    } catch (e) {
      error = e.message || "Bir hata oluştu";
    } finally {
      loading = false;
    }
  }

  function openTerminal(token: string) {
    activeContainerToken = token;
    showModal = true;
  }
  function closeModal() {
    showModal = false;
    activeContainerToken = "";
  }

  onMount(() => {
    loadContainers();
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

        <!-- Dosya yönetim alanı -->
        <div class="mt-4 p-3 bg-gray-50 rounded border">
          <h3 class="font-semibold mb-2">📁 Dosya Yönetimi</h3>
          <div class="flex gap-2 mb-2">
            <input type="file" on:change={(e) => handleFileChange(container.Token, e.target.files?.[0] || null)} />
            <button class="bg-indigo-600 text-white px-3 py-1 rounded hover:bg-indigo-700" on:click={() => uploadFile(container.Token, fileInputs[container.Token])}>Yükle</button>
          </div>

          {#if fileStates[container.Token]?.length > 0}
            <ul class="space-y-1">
              {#each fileStates[container.Token] as file}
                <li class="flex justify-between items-center border px-2 py-1 rounded text-sm">
                  <span>{file}</span>
                  <button on:click={() => deleteFile(container.Token, file)} class="bg-red-500 text-white px-2 py-0.5 rounded text-xs">Sil</button>
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
  <TerminalModal token={activeContainerToken} on:close={closeModal} />
{/if}
