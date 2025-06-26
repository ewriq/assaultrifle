<script lang="ts">
  import { goto } from "$app/navigation";
  import Cookies from "js-cookie";
  import { loginUser } from "$lib/login";

  let email = "";
  let password = "";
  let error: string = "";
  let success: string = "";

  async function login() {
    error = "";
    success = "";

    try {
      const token = await loginUser(email, password);
      Cookies.set("token", token);
      success = "Başarılı! Yönlendiriliyorsunuz...";
      setTimeout(() => goto("/p"), 2000);
    } catch (err: any) {
      error = err.message || "Bilinmeyen bir hata oluştu.";
    }
  }
</script>

<main class="min-h-screen flex justify-center items-center bg-gray-100 p-6">
  <div class="w-full max-w-sm bg-white p-6 rounded-lg shadow-md">
    <h2 class="text-2xl font-bold mb-4 text-center">Giriş Yap</h2>

    {#if error}
      <div class="bg-red-100 text-red-700 p-3 mb-3 rounded">{error}</div>
    {/if}

    {#if success}
      <div class="bg-green-100 text-green-700 p-3 mb-3 rounded">{success}</div>
    {/if}

    <form class="space-y-4" on:submit|preventDefault={login}>
      <div>
        <label for="email" class="block text-sm font-medium text-gray-700">E-posta</label>
        <input
          id="email"
          type="email"
          bind:value={email}
          required
          class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm shadow-sm focus:outline-none focus:ring focus:ring-blue-200"
          placeholder="ornek@mail.com"
        />
      </div>

      <div>
        <label for="password" class="block text-sm font-medium text-gray-700">Şifre</label>
        <input
          id="password"
          type="password"
          bind:value={password}
          required
          class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 text-sm shadow-sm focus:outline-none focus:ring focus:ring-blue-200"
          placeholder="••••••••"
        />
      </div>

      <button
        type="submit"
        class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded"
      >
        Giriş Yap
      </button>
    </form>
  </div>
</main>
