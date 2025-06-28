<script lang="ts">
  import { onMount, onDestroy } from "svelte";
  import { writable } from "svelte/store";

  export let containerId: string;

  let ws: WebSocket;
  let input = "";
  let terminalDiv: HTMLDivElement;
  const output = writable("Bağlanıyor...\n");

  onMount(() => {
    const protocol = location.protocol === "https:" ? "wss" : "ws";
    ws = new WebSocket(`${protocol}://${location.hostname}:3000/ws/container/${containerId}`);

    ws.onopen = () => output.update(o => o + "✅ Bağlantı kuruldu.\n");
    ws.onmessage = e => output.update(o => o + e.data + "\n");
    ws.onerror = () => output.update(o => o + "❌ Bağlantı hatası.\n");
    ws.onclose = () => output.update(o => o + "🔌 Bağlantı kapandı.\n");
  });

  onDestroy(() => {
    ws?.close();
  });

  function sendCommand() {
    if (input.trim() && ws.readyState === WebSocket.OPEN) {
      ws.send(input);
      output.update(o => o + "> " + input + "\n");
      input = "";
      scrollToBottom();
    }
  }

  function scrollToBottom() {
    setTimeout(() => {
      terminalDiv?.scrollTo({ top: terminalDiv.scrollHeight, behavior: "smooth" });
    }, 10);
  }
</script>

<style>
  .terminal {
    background: #111;
    color: #0f0;
    font-family: monospace;
    padding: 1rem;
    height: 300px;
    overflow-y: auto;
    white-space: pre-wrap;
    border-radius: 0.5rem;
  }

  input {
    width: 100%;
    background: #111;
    color: #0f0;
    font-family: monospace;
    padding: 0.5rem;
    border: 1px solid #444;
    border-radius: 0.25rem;
    outline: none;
  }
</style>

<div class="terminal" bind:this={terminalDiv}>
  {$output}
</div>

<input
  type="text"
  bind:value={input}
  on:keydown={(e) => e.key === "Enter" && sendCommand()}
  placeholder="Komut girin..."
  autofocus
/>
