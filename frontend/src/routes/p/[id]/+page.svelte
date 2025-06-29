<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import Cmd from '$lib/components/container/Cmd.svelte';
  import * as api from '$lib/api/container/index';

  let token = '';
  let logs = '';
  let status = '';
  let message = '';
  let error = '';

  $: token = $page.params.id;

  async function refreshStatus() {
    try {
      status = await api.getStatus(token);
    } catch {
      error = 'Durum alınamadı.';
    }
  }

  async function refreshLogs() {
    try {
      const rawLogs = await api.getLogs(token);
      logs = api.cleanLogs(rawLogs);
    } catch {
      error = 'Loglar alınamadı.';
    }
  }

  async function handleAction(action: 'start' | 'stop' | 'restart' | 'delete') {
    try {
      message = '';
      error = '';
      switch(action) {
        case 'start':
          message = await api.startContainer(token);
          await refreshStatus();
          break;
        case 'stop':
          message = await api.stopContainer(token);
          await refreshStatus();
          break;
        case 'restart':
          message = await api.restartContainer(token);
          await refreshStatus();
          break;
        case 'delete':
          message = await api.deleteContainer(token);
          break;
      }
    } catch {
      error = `${action.charAt(0).toUpperCase() + action.slice(1)} işlemi başarısız.`;
    }
  }

  onMount(() => {
    refreshStatus();
    refreshLogs();
  });
</script>

<main class="p-6 max-w-3xl mx-auto">
  <h1 class="text-2xl font-bold mb-4">Container: {token}</h1>

  {#if status}
    <p><strong>Durum:</strong> {status}</p>
  {/if}

  <div class="flex gap-2 my-4">
    <button on:click={() => handleAction('start')} class="bg-green-500 px-3 py-1 text-white rounded">Başlat</button>
    <button on:click={() => handleAction('stop')} class="bg-yellow-500 px-3 py-1 text-white rounded">Durdur</button>
    <button on:click={() => handleAction('restart')} class="bg-blue-500 px-3 py-1 text-white rounded">Yeniden Başlat</button>
    <button on:click={() => handleAction('delete')} class="bg-red-600 px-3 py-1 text-white rounded">Sil</button>
  </div>

  {#if message}
    <p class="text-green-700">{message}</p>
  {/if}

  {#if error}
    <p class="text-red-600">{error}</p>
  {/if}

  <h2 class="text-xl font-semibold mt-6">Loglar</h2>
  <pre class="bg-gray-100 p-4 rounded mt-2 overflow-x-auto" style="white-space: pre-wrap;">{logs}</pre>

  <Cmd id={token} />
</main>
