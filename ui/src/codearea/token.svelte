<script>
    export let text = "";
    export let classes = [];
    export let buffer = "";
    export let line = 0;
    export let token = 0;
</script>

<code id={`${buffer}-${line}:${token}`} class={classes} contenteditable="true">
    {text}
</code>

<style>
    code {
        white-space: pre;
        tab-size: 4;
        outline: none;
        display: inline-block;
        user-select: auto;
        -webkit-user-select: auto;
        min-width: 3px;
        height: 100%;
    }

    code::selection {
        /* background: rgba(81, 92, 126, 0.25);
        color: rgba(81, 92, 126, 0.27); */
        background: cyan;
    }
</style>
