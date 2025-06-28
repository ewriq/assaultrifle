<script lang="ts">
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import { getUser } from "$lib/api/auth/user";
    import Add from "$lib/components/container/Add.svelte";
    import List from "$lib/components/container/List.svelte";
  
    let user: any = null;
  
    onMount(() => {
      getUser(
        (data) => {
          user = data;
        },
        () => {
          goto("/login");
        }
      );
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
  