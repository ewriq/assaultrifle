<script lang="ts">
    const links = [
      { href: '/dashboard', label: 'Dashboard' },
      { href: '/dashboard/containers', label: 'Containers' },
      { href: '/dashboard/settings', label: 'Ayarlar' },
      { href: '/logout', label: 'Çıkış' }
    ];
  </script>
  
  <aside class="fixed top-0 left-0 h-screen w-100 bg-gray-800 text-white p-4 shadow-lg">
    <h1 class="text-2xl font-bold mb-6">Panel</h1>
    <nav class="flex flex-col space-y-2">
      {#each links as { href, label }}
        <a href={href} class="px-4 py-2 rounded hover:bg-gray-700 transition-colors">
          {label}
        </a>
      {/each}
    </nav>
  </aside>
  