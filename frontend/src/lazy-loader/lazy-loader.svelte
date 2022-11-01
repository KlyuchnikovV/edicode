<script lang="ts">
    export let plugin = "";
    export let file = "";
</script>

{#await import(`../../../plugins/${plugin}/${file}.svelte`) then Module}
    <Module.default />
{/await}
