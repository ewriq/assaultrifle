<script lang="ts">
  import Cmd from "./index.svelte";
  import { createEventDispatcher } from "svelte";

  export let token: string;
  const dispatch = createEventDispatcher();

  const close = () => dispatch("close");
</script>

<div
  class="fixed inset-0 bg-black bg-opacity-60 flex items-center justify-center z-50"
  on:click={close}
>
  <div class="bg-white p-6 rounded-lg max-w-4xl w-full" on:click|stopPropagation>
    <div class="flex justify-between items-center mb-4">
      <h2 class="text-lg font-bold">Terminal</h2>
      <button
        class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600"
        on:click={close}
      >
        Kapat
      </button>
    </div>
    <Cmd containerId={token} />
  </div>
</div>
