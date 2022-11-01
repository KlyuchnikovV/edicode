<script>
    export let onclick;
    export let icon;
    export let selected = false;
</script>

<div on:click={onclick} class:selected>
    <i class:selected class="material-icons-outlined md-36">{icon}</i>
</div>

<style>
    div {
        height: 3em;
        padding: 6px;
    }

    i {
        /* width: 2em;
        height: 2em; */
        cursor: pointer;
        color: grey;
    }

    div.selected {
        border-right: 2px white solid;
    }

    i.selected {
        color: white;
    }
</style>
