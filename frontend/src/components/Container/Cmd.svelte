<script lang="ts">
    import { onMount } from "svelte";
    import Cookies from "js-cookie";
  
    export let containerId: string;
    let socket: WebSocket;
    let input = "";
    let output = "";
  
    const token = Cookies.get("token");
  
    function sendCommand() {
      if (input.trim() && socket.readyState === WebSocket.OPEN) {
        socket.send(input);
        output += `> ${input}\n`;
        input = "";
      }
    }
  
    onMount(() => {
      if (!token) {
        output = "❌ Kullanıcı tokeni bulunamadı.";
        return;
      }
  
      const url = `ws://localhost:3000/ws/${containerId}?token=${token}`;
      socket = new WebSocket(url);
  
      socket.onmessage = (event) => {
        output += event.data + "\n";
      };
  
      socket.onerror = () => {
        output += "❌ WebSocket bağlantı hatası\n";
      };
  
      socket.onclose = () => {
        output += "🔌 Bağlantı kapatıldı\n";
      };
    });
  </script>
  
  <div class="bg-black text-green-400 p-4 rounded-lg font-mono text-sm h-96 overflow-y-auto whitespace-pre-wrap">
    {output}
  </div>
  
  <input
    bind:value={input}
    on:keydown={(e) => e.key === 'Enter' && sendCommand()}
    placeholder="Komut girin..."
    class="mt-2 w-full px-3 py-2 border rounded"
  />
  