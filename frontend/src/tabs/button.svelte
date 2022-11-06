<script>
    import { onMount } from "svelte";
    import { CloseFile } from "../../wailsjs/go/api/Api";

    export let onclick;
    export let selected = false;
    export let buffer = "";
    export let config = {
        height: "2em",
        icon: "",
    };

    let element;

    function close() {
        console.log(`file closing ${buffer}`);
        CloseFile(buffer);
    }

    onMount(() => {
        element.scrollIntoView(true);
    });

    function selectTab(event) {
        if (event.target !== this) return;
        return onclick();
    }
</script>

<span
    bind:this={element}
    class="button"
    class:selected
    on:click={selectTab}
    style:height={config.height}
    style:line-height={config.height}
>
    <i class="material-icons-outlined md-18" on:click={selectTab}>{config.icon}</i>
    {buffer.slice(buffer.lastIndexOf("/") + 1)}
    <i class="close material-icons-outlined md-18" on:click={close}>close</i>
</span>

<style>
    .button {
        padding: 0 10px;
        width: fit-content;
        min-width: fit-content;
        cursor: pointer;
    }

    i {
        color: grey;
        vertical-align: middle;
    }

    span.selected {
        border-top: 2px white solid;
    }

    .close:hover {
        box-shadow: rgba(0, 0, 0, 0.25) 0 8px 15px;
        transform: translateY(-2px);
    }
</style>
