<script lang="ts">
  import Terminal from "./Cmd.svelte";
  import { createEventDispatcher } from "svelte";

  export let token: string;
  const dispatch = createEventDispatcher();

  function close() {
    dispatch("close");
  }
</script>

<div
  class="fixed inset-0 bg-black bg-opacity-60 flex items-center justify-center z-50"
  on:click={close}
>
  <div class="bg-white p-6 rounded-lg max-w-4xl w-full" on:click|stopPropagation>
    <button
      class="mb-4 px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600"
      on:click={close}
    >
      Kapat
    </button>

    <Terminal containerId={token} />
  </div>
</div>
