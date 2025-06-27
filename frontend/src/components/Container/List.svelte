<script lang="ts">
  import { onMount } from "svelte";
  import TerminalModal from "../Cmd/Cmd.svelte";
  import {
    fetchContainers,
    containerActions,
    fetchFiles,
    fileStates,
    fileInputs,
    handleFileChange,
    uploadFile,
    deleteFile
  } from "$lib/container";
  import Cookies from "js-cookie";

  let containers = [];
  let loading = true;
  let error: string | null = null;
  let showModal = false;
  let activeContainerToken = "";

  const loadContainers = async () => {
    loading = true;
    try {
      const user = Cookies.get("token");
      if (!user) throw new Error("Kullanıcı tokenı yok!");

      containers = await fetchContainers();
      console.log("Yüklenen containerlar:", containers);

      await Promise.all(containers.map(c => fetchFiles(c.Token)));
    } catch (e) {
      console.error("Hata:", e);
      error = e.message || "Bir hata oluştu";
    } finally {
      loading = false;
    }
  };

  const openTerminal = (token: string) => {
    activeContainerToken = token;
    showModal = true;
  };

  const closeModal = () => {
    activeContainerToken = "";
    showModal = false;
  };

  onMount(loadContainers);
</script>

{#if loading}
  <p class="text-gray-600 text-sm">Yükleniyor...</p>
{:else if error}
  <p class="text-red-500">Hata: {error}</p>
{:else if containers.length === 0}
  <p class="text-gray-500">Hiç konteyner bulunamadı.</p>
{:else}
  <ul class="space-y-4">
    {#each containers as c}
      <li class="p-4 border rounded bg-white shadow">
        <div class="flex justify-between items-center">
          <div>
            <h2 class="font-bold text-lg">{c.Name}</h2>
            <p class="text-sm text-gray-600">Port: {c.Port} | Token: {c.Token} | Tür: {c.Types}</p>
          </div>

          <div class="flex flex-wrap gap-2">
            <button class="bg-green-500 hover:bg-green-600 text-white px-3 py-1 rounded" on:click={() => containerActions.start(c.Token)}>Başlat</button>
            <button class="bg-yellow-400 hover:bg-yellow-500 text-white px-3 py-1 rounded" on:click={() => containerActions.stop(c.Token)}>Durdur</button>
            <button class="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded" on:click={() => containerActions.delete(c.Token)}>Sil</button>
            <button class="bg-blue-500 hover:bg-blue-600 text-white px-3 py-1 rounded" on:click={() => containerActions.status(c.Token)}>Durum</button>
            <button class="bg-purple-500 hover:bg-purple-600 text-white px-3 py-1 rounded" on:click={() => openTerminal(c.Token)}>Terminal</button>
          </div>
        </div>

        <div class="mt-4 p-3 bg-gray-50 rounded border">
          <h3 class="font-semibold mb-2">📁 Dosya Yönetimi</h3>
          <div class="flex flex-wrap gap-2 mb-2 items-center">
            <input
              type="file"
              class="border px-2 py-1 rounded"
              on:change={(e) => handleFileChange(c.Token, e.target.files?.[0] || null)} />
            <button class="bg-indigo-500 hover:bg-indigo-600 text-white px-3 py-1 rounded" on:click={() => uploadFile(c.Token, fileInputs[c.Token])}>Yükle</button>
          </div>

          {#if fileStates[c.Token]?.length > 0}
            <ul class="space-y-1 text-sm">
              {#each fileStates[c.Token] as file}
                <li class="flex justify-between items-center border px-2 py-1 rounded">
                  <span>{file}</span>
                  <button class="bg-red-500 hover:bg-red-600 text-white px-2 py-1 text-xs rounded" on:click={() => deleteFile(c.Token, file)}>Sil</button>
                </li>
              {/each}
            </ul>
          {:else}
            <p class="text-gray-500 text-sm">Hiç dosya yok.</p>
          {/if}
        </div>
      </li>
    {/each}
  </ul>
{/if}

{#if showModal}
  <TerminalModal token={activeContainerToken} on:close={closeModal} />
{/if}
