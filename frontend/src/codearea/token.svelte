<script>
    export let text = "";
    export let classes = [];
    export let buffer = "";
    export let line = 0;
    export let token = 0;
    export let sendSelection;
</script>

<code
    id={`${buffer}-${line+1}:${token}`}
    class={classes.join(" ")}
    contenteditable="true"
    on:mouseup={(event) => {
        console.log(`triggered by token (${line+1}-${token})`)
        event.stopPropagation();
        sendSelection(line, token);
    }}
>
    {text}
</code>

<style>
    code {
        white-space: pre;
        tab-size: 4;
        outline: none;
        display: inline-block;
        min-width: 3px;
        height: 100%;
        caret-color: whitesmoke;

        user-select: text;
        -webkit-user-select: text;
    }

    code::selection,
    code::-webkit-selection {
        color: yellowgreen;
        outline: blue 2px;
    }

    .symbols {
        color: whitesmoke;
    }

    .delimiter {
        color: blueviolet;
    }
</style>
