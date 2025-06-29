<script lang="ts">
    import { page } from "$app/stores";
    import { onMount } from "svelte";
    import { get } from "svelte/store";
    import { useFileManager } from "$lib/api/container/file";

    const id = get(page).params.id;
    const {
        path,
        files,
        loading,
        error,
        message,
        selectedFile,
        fetchFiles,
        deleteFile,
        uploadFile,
        fileNameFromLine,
    } = useFileManager(id);

    $: if ($path) fetchFiles($path);

    onMount(() => {
        fetchFiles($path);
    });

    function handleDelete(file: string) {
        const name = fileNameFromLine(file);
        if (confirm(`${name} dosyasını silmek istediğinize emin misiniz?`)) {
            deleteFile($path, name);
        }
    }

    function handleUpload() {
        if ($selectedFile) {
            uploadFile($path, $selectedFile);
        } else {
            alert("Lütfen bir dosya seçin.");
        }
    }
</script>

<div class="p-4 max-w-xl mx-auto border rounded shadow">
    <h1 class="text-xl font-semibold mb-3">Container Dosyaları - {id}</h1>

    <label for="path-input" class="block font-medium mb-2">Path Seçin:</label>
    <input
        id="path-input"
        type="text"
        bind:value={$path}
        placeholder="/"
        class="w-full mb-4 p-2 border rounded"
    />

    {#if $loading}
        <p>Yükleniyor...</p>
    {:else}
        {#if $error}
            <p class="text-red-600 mb-2">{$error}</p>
        {/if}
        {#if $message}
            <p class="text-green-600 mb-2">{$message}</p>
        {/if}
        <ul class="mb-4">
            {#each $files as file}
                <li class="flex justify-between items-center border-b py-1">
                    <span>{fileNameFromLine(file)}</span>
                    <button
                        class="text-red-600 hover:text-red-800"
                        on:click={() => handleDelete(file)}
                    >
                        Sil
                    </button>
                </li>
            {/each}
            {#if $files.length === 0}
                <li>Dosya bulunamadı.</li>
            {/if}
        </ul>
    {/if}

    <div class="mb-2">
        <label for="file-upload" class="block mb-1 font-medium"
            >Dosya Yükle</label
        >
        <input
            id="file-upload"
            type="file"
            on:change={(e) => selectedFile.set(e.target.files?.[0] ?? null)}
        />
    </div>

    <button
        class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700"
        on:click={handleUpload}
        disabled={!$selectedFile}
    >
        Yükle
    </button>
</div>
