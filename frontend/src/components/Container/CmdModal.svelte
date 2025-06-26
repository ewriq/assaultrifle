<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { writable } from "svelte/store";
  
    export let containerId: string;
  
    let ws: WebSocket;
    let input = "";
    const output = writable<string>("");
  
    onMount(() => {
      ws = new WebSocket(`ws://localhost:3000/ws/container/${containerId}`);
  
      ws.onmessage = (event) => {
        output.update(o => o + event.data + "\n");
      };
  
      ws.onopen = () => {
        output.update(o => o + "Terminal bağlandı.\n");
      };
  
      ws.onerror = (e) => {
        output.update(o => o + "Hata oluştu.\n");
      };
  
      ws.onclose = () => {
        output.update(o => o + "Bağlantı kapandı.\n");
      };
    });
  
    onDestroy(() => {
      ws.close();
    });
  
    function sendCommand() {
      if (ws.readyState === WebSocket.OPEN) {
        ws.send(input);
        output.update(o => o + "> " + input + "\n");
        input = "";
      }
    }
  </script>
  
  <style>
    .terminal {
      background: black;
      color: lime;
      font-family: monospace;
      padding: 1rem;
      height: 300px;
      overflow-y: auto;
      white-space: pre-wrap;
    }
    input {
      width: 100%;
      background: black;
      color: lime;
      border: none;
      font-family: monospace;
      padding: 0.5rem;
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
  