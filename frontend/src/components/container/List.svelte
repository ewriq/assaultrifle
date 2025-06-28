<script lang="ts">
    import { onMount } from "svelte";
    import { fetchAllContainers, type Container } from "$lib/container/list";
  
    let containers: Container[] = [];
    let loading = true;
    let error: string | null = null;
  
    onMount(async () => {
      loading = true;
      const result = await fetchAllContainers();
      containers = result.data;
      error = result.error;
      loading = false;
    });
  </script>
  
  {#if loading}
    <p class="text-gray-600 text-sm">Yükleniyor...</p>
  {:else if error}
    <p class="text-red-500">Hata: {error}</p>
  {:else if containers.length === 0}
    <p class="text-gray-500">Hiç konteyner bulunamadı.</p>
  {:else}
    <ul class="space-y-2">
      {#each containers as c}
        <li class="border p-3 rounded bg-white shadow">
          <strong>{c.Name}</strong> — Port: {c.Port} — Tür: {c.Types}
        </li>
      {/each}
    </ul>
  {/if}
  