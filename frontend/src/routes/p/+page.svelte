<script lang="ts">
  import { onMount } from "svelte";
  import Cookies from "js-cookie";
  import List from "$lib/components/container/List.svelte";
  import Add from "$lib/components/container/Add.svelte";
  import { getUser } from "$lib/api/auth/user";

  let user: any = null;
  const token = Cookies.get("token");

  onMount(async () => {
    if (!token) {
      window.location.href = "/login";
      return;
    }

    user = await getUser(token);
  });
</script>

<main class="p-6 max-w-6xl mx-auto space-y-10">
  {#if user}
    <h1 class="text-2xl font-bold">Welcome, {user[0].Username}!</h1>
    <Add />
    <List />
  {:else}
    <p class="text-gray-600">Kullanıcı verisi yükleniyor...</p>
  {/if}
</main>

<style>
  main {
    font-family: system-ui, sans-serif;
  }
</style>
