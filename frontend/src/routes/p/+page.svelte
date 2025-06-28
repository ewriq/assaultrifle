<script lang="ts">
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import Cookies from "js-cookie";
  import { fetchUser } from "$lib/auht/user";

  import List from "../../components/Container/List.svelte";
  import Add from "../../components/Container/Add.svelte";

  let user: { Username: string } | null = null;
  const token = Cookies.get("token");

  onMount(async () => {
    if (!token) {
      goto("/login");
      return;
    }

    try {
      user = await fetchUser(token);
    } catch (e) {
      console.error("Kullanıcı verisi alınamadı:", e);
      goto("/login");
    }
  });
</script>

<main class="p-6 max-w-6xl mx-auto space-y-10">
  {#if user}
    <h1 class="text-2xl font-bold">Welcome, {user.Username}!</h1>
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
