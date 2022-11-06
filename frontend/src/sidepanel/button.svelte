<script>
    export let onclick;
    export let selected = false;

    export let config = {
        width: "2em",
        icon: "",
    };
</script>

<div on:click={onclick} class:selected style:height={config.width}>
    <i class:selected class="material-icons-outlined md-36">{config.icon}</i>
</div>

<style>
    i {
        cursor: pointer;
        color: grey;
        vertical-align: middle;
    }

    div.selected {
        border-right: 2px white solid;
    }

    i.selected {
        color: white;
    }
</style>
