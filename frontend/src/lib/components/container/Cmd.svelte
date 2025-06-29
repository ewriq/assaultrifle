<script lang="ts">
  import { onMount } from 'svelte';
  import { createCmdStore } from '$lib/api/ws';

  export let id: string;

  let command = "";
  const { output, sendCommand, startWebSocket } = createCmdStore(id);

  function onSend() {
    sendCommand(command);
    command = "";
  }

  onMount(() => {
    startWebSocket();
  });
</script>

<div class="bg-black text-green-400 font-mono p-4 rounded h-[500px] overflow-y-auto text-sm">
  {#each $output as line}
    <div>{line}</div>
  {/each}
</div>

<div class="mt-2 flex gap-2">
  <input
    bind:value={command}
    on:keydown={(e) => e.key === 'Enter' && onSend()}
    class="flex-1 px-3 py-1 border rounded font-mono"
    placeholder="Komut yazın ve Enter'a basın..."
  />
  <button on:click={onSend} class="bg-blue-600 text-white px-4 rounded">Gönder</button>
</div>
