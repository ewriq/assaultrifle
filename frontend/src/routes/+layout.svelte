

<script lang="ts">
	import { page } from '$app/stores';
	import { derived } from 'svelte/store';
	import Navbar from '$lib/components/layout/Navbar.svelte';
	import Sidebar from '$lib/components/layout/Sidebar.svelte';
	let { children } = $props();
	import "../app.css";


	const layoutType = derived(page, $page => {
	  const path = $page.url.pathname;
	  
	  if (path.startsWith('/p/') || path.startsWith('/p')) return 'dashboard';
	  if (path === '/login' || path === '/register') return 'auth';
	  return 'default';
	});
  </script>
  
  {#if $layoutType === 'auth'}
	<Navbar />
  {:else if $layoutType === 'dashboard'}
	<Sidebar />
  {:else}
	<Navbar />
  {/if}
  {@render children()}

  