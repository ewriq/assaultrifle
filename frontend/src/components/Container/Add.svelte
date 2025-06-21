<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import axios from "axios";
    import Cookies from "js-cookie";
  
    const dispatch = createEventDispatcher();
    const user = Cookies.get("token") || "";
  
    let open = false;
    let name = "";
    let password = "";
    let port = "";
    let type = "MySql";
    let message = "";
    let loading = false;
  
    async function submit() {
      loading = true;
      message = "";
  
      try {
        const res = await axios.post("http://localhost:3000/api/container/add", {
          name,
          password,
          port,
          user,
          type
        });
  
        if (res.data.status === "OK") {
          dispatch("added", { token: res.data.token });
          message = "Container başarıyla oluşturuldu.";
          open = false;
          resetForm();
        } else {
          message = res.data.error || "Bilinmeyen hata.";
        }
      } catch (e: any) {
        message = "Sunucu hatası: " + e.message;
      } finally {
        loading = false;
      }
    }
  
    function resetForm() {
      name = "";
      password = "";
      port = "";
      type = "";
    }
  </script>
  
  <!-- Modal Açma Butonu -->
  <button class="bg-blue-600 text-white px-4 py-2 rounded" on:click={() => (open = true)}>
    Yeni Container
  </button>
  
  <!-- Modal -->
  {#if open}
    <div class="fixed inset-0 bg-black/60 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg w-full max-w-md p-6 relative">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">Yeni Container</h2>
  
        {#if message}
          <p class="mb-3 text-sm {message.startsWith('Container') ? 'text-green-600' : 'text-red-500'}">
            {message}
          </p>
        {/if}
  
        <div class="space-y-3">
          <div>
            <label class="text-sm font-medium block mb-1">İsim</label>
            <input class="w-full p-2 border rounded" bind:value={name} placeholder="örnek_db" />
          </div>
  
          <div>
            <label class="text-sm font-medium block mb-1">Şifre</label>
            <input class="w-full p-2 border rounded" type="password" bind:value={password} placeholder="root" />
          </div>
  
          <div>
            <label class="text-sm font-medium block mb-1">Port</label>
            <input class="w-full p-2 border rounded" bind:value={port} placeholder="3306" />
          </div>
  
          <div>
            <label class="text-sm font-medium block mb-1">Veritabanı Türü</label>
            <input class="w-full p-2 border rounded" bind:value={type} placeholder="MySql / PostgreSql" />
          </div>
        </div>
  
        <div class="flex justify-end gap-2 mt-6">
          <button
            class="bg-gray-200 hover:bg-gray-300 text-gray-800 px-4 py-2 rounded"
            on:click={() => (open = false)}
          >
            İptal
          </button>
          <button
            class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded disabled:opacity-50"
            on:click={submit}
            disabled={loading}
          >
            {loading ? "Oluşturuluyor..." : "Kaydet"}
          </button>
        </div>
  
        <button class="absolute top-2 right-3 text-gray-400 hover:text-gray-600" on:click={() => (open = false)}>
          ✕
        </button>
      </div>
    </div>
  {/if}
  