<script lang="ts">
	import '../app.css';
	import Navbar from '../../src/components/Navbar.svelte' 
	let { children } = $props();
</script>
<Navbar/>
{@render children()} 
