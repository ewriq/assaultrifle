<script lang="ts">
  import { goto } from '$app/navigation';
  import Cookies from 'js-cookie';
  import { registerUser } from '$lib/auht/register';

  let email = '';
  let password = '';
  let username = '';
  let error = '';
  let success = '';

  async function handleRegister() {
    error = '';
    success = '';

    try {
      const token = await registerUser({ email, password, username });
      Cookies.set('token', token);
      success = 'Kayıt başarılı! Ana sayfaya yönlendiriliyorsunuz...';
      setTimeout(() => goto('/'), 2000);
    } catch (err: any) {
      error = err.message || 'Bilinmeyen bir hata oluştu.';
    }
  }
</script>

<main class="min-h-screen flex items-center justify-center bg-gray-100 px-4">
  <div class="w-full max-w-sm bg-white p-6 rounded-lg shadow-md">
    <h2 class="text-2xl font-bold mb-4 text-center">Kayıt Ol</h2>

    {#if error}
      <div class="bg-red-100 text-red-700 p-3 mb-3 rounded">{error}</div>
    {/if}

    {#if success}
      <div class="bg-green-100 text-green-700 p-3 mb-3 rounded">{success}</div>
    {/if}

    <form on:submit|preventDefault={handleRegister} class="space-y-4">
      <div>
        <label for="username" class="block text-sm font-medium text-gray-700">Kullanıcı Adı</label>
        <input
          id="username"
          type="text"
          bind:value={username}
          required
          placeholder="örnekKullanici"
          class="mt-1 w-full border border-gray-300 rounded px-3 py-2 text-sm shadow-sm focus:outline-none focus:ring focus:ring-blue-200"
        />
      </div>

      <div>
        <label for="email" class="block text-sm font-medium text-gray-700">E-posta</label>
        <input
          id="email"
          type="email"
          bind:value={email}
          required
          placeholder="ornek@mail.com"
          class="mt-1 w-full border border-gray-300 rounded px-3 py-2 text-sm shadow-sm focus:outline-none focus:ring focus:ring-blue-200"
        />
      </div>

      <div>
        <label for="password" class="block text-sm font-medium text-gray-700">Şifre</label>
        <input
          id="password"
          type="password"
          bind:value={password}
          required
          placeholder="••••••••"
          class="mt-1 w-full border border-gray-300 rounded px-3 py-2 text-sm shadow-sm focus:outline-none focus:ring focus:ring-blue-200"
        />
      </div>

      <button
        type="submit"
        class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded"
      >
        Kayıt Ol
      </button>
    </form>
  </div>
</main>
