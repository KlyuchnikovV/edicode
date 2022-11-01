<script lang="ts">
    import { SetCursor, MouseDown, MouseUp } from "../../wailsjs/go/api/Api.js";

    export let text = "";
    export let classes = [];
    export let buffer = "";
    export let line = 0;
    export let token = 0;
    export let offset = 0;

    async function syncSelection() {
        await new Promise((r) => setTimeout(r, 1));
    }

    async function mousedown(event) {
        await syncSelection();

        var selection = window.getSelection();
        var lineOffset = offset;

        console.log(selection);

        if (selection.anchorNode.parentNode.isEqualNode(event.target)) {
            lineOffset += selection.focusOffset;
        } else if (selection.focusNode.parentNode.isEqualNode(event.target)) {
            lineOffset += selection.anchorOffset;
        }

        MouseDown({
            buffer: buffer,
            line: line,
            offset: lineOffset,
        }).then((err) => {
            if (err) {
                console.error(err);
                return;
            }
        });
    }

    async function mouseup(event) {
        await syncSelection();

        var selection = window.getSelection();
        var lineOffset = offset;

        console.log(selection);

        if (selection.focusNode.parentNode.isEqualNode(event.target)) {
            lineOffset += selection.focusOffset;
        } else if (selection.anchorNode.parentNode.isEqualNode(event.target)) {
            lineOffset += selection.anchorOffset;
        }

        MouseUp({
            buffer: buffer,
            line: line,
            offset: lineOffset,
        }).then((err) => {
            if (err) {
                console.error(err);
                return;
            }
        });
    }

    function dblclick(event) {
        SetCursor({
            buffer: buffer,
            start: {
                line: line,
                offset: offset,
            },
            end: {
                line: line,
                offset: offset + text.length,
            },
        }).then((err) => {
            if (err) {
                console.error(err);
                return;
            }
        });
    }
</script>

<code
    id={`${buffer}-${line + 1}:${token}`}
    class={classes.join(" ")}
    contenteditable="true"
    on:mousedown|self|stopPropagation={mousedown}
    on:mouseup|self|stopPropagation={mouseup}
    on:dblclick|self|capture|stopPropagation={dblclick}
>
    {text}
</code>

<!-- TODO: on mouse up in every place in document -->
<style>
    code {
        white-space: pre;
        tab-size: 4;
        outline: none;
        display: inline;
        min-width: 3px;
        height: 100%;
        caret-color: whitesmoke;

        user-select: text;
        -webkit-user-select: text;
    }

    code::selection,
    code::-webkit-selection {
        color: aqua;
    }

    .delimiter {
        color: blueviolet;
    }

    .error {
        background-color: indianred;
        border-radius: 3px;
    }

    code:hover {
        border: thin whitesmoke dashed;
        border-radius: 3px;
    }
</style>
