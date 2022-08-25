<script>
    export let text = "";
    export let classes = [];
    export let buffer = "";
    export let line = 0;
    export let token = 0;
    // export let cursor;

    function onMouse(event) {
        // cursor(event, line - 1, token);
    }
</script>

<code
    id={`${buffer}-${line}:${token}`}
    class={classes}
    contenteditable="true"
    on:mouseup={onMouse}
>
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
        caret-color: whitesmoke;
    }

    code::selection {
        /* background: rgba(81, 92, 126, 0.25);
        color: rgba(81, 92, 126, 0.27); */
        background: cyan;
    }

    .symbols {
        color: whitesmoke;
    }

    .whitespace {
        /* background: red; */
    }

    .delimiter {
        color: blueviolet;
    }
</style>
