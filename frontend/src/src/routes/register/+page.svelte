<script lang="ts">
  import { registerUser } from '$lib/api/auth/register';

  let email = '';
  let password = '';
  let username = '';
  let error = '';

  const register = async () => {
    await registerUser(
      email,
      password,
      username,
      () => {
        error = 'Kayıt başarılı! Ana sayfaya yönlendiriliyorsunuz...';
        setTimeout(() => location.href = '/', 3000); // goto('/') de kullanılabilir
      },
      (errMsg) => {
        error = errMsg;
      }
    );
  };
</script>

<main>
  <div class="flex justify-center items-center">
    <div class="mb-1.5 w-full max-w-sm p-4 border-gray-200 rounded-lg shadow sm:p-6 md:p-8">
      <form class="space-y-6" on:submit|preventDefault={register}>
        <h5 class="text-xl font-medium text-gray-900">Register</h5>

        {#if error}
          <div class="bg-green-100 text-green-700 p-4 rounded" role="alert">
            <p class="font-bold">{error}</p>
          </div>
        {/if}

        <div>
          <label for="username" class="block mb-2 text-sm font-medium text-gray-900">Username</label>
          <input
            type="text"
            bind:value={username}
            id="username"
            class="input"
            placeholder="company"
            required
          />
        </div>

        <div>
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900">Email</label>
          <input
            type="email"
            bind:value={email}
            id="email"
            class="input"
            placeholder="name@company.com"
            required
          />
        </div>

        <div>
          <label for="password" class="block mb-2 text-sm font-medium text-gray-900">Password</label>
          <input
            type="password"
            bind:value={password}
            id="password"
            class="input"
            placeholder="••••••••"
            required
          />
        </div>

        <button
          type="submit"
          class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
        >
          Register
        </button>
      </form>
    </div>
  </div>
</main>

<style>
  
</style>