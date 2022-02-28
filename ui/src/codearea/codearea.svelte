<script>
    export let file = "";

    import Line from "./line.svelte";
    import Cursor from "./cursor";
    import Keys from "./keys";
    import { afterUpdate, onMount } from "svelte";

    let lines = [];
    let element, cursor, keys;
    let elementLines = [];

    onMount(async () => {
        cursor = new Cursor(file, element);
        keys = new Keys(file, cursor);

        var bufferInit = async () => {
            let buf = await window.backend.Api.GetBuffer(file);
            console.log(`buffer loaded ${file}`);
            lines = buf.lines;
            cursor.selection.lines = buf.lines;
        };

        bufferInit();

        window.wails.Events.On("highlight_changed", (event) => {
            if (event.data.buffer === file) {
                console.log(`highlight changed ${event.data.buffer}`);
                lines = event.data.lines;
                cursor.selection.lines = event.data.lines;
            }
        });
    });

    function onKeyDown(event) {
        keys.onKeyDown(event, cursor, elementLines);
    }

    function onMouse() {
        var offset = cursor.getCurrentCursorPosition();
        cursor.onMouse(offset);
    }

    afterUpdate(() => {
        if (cursor !== undefined) {
            console.log(`after update`);
            cursor.setCurrentCursorPosition(elementLines);
        }
    });
</script>

<div class={"codearea"}>
    <div class={"gutter"}>
        {#each lines as _, index (index)}
            <code class={"gutter-line"}>{index + 1}</code>
        {/each}
    </div>
    <div
        bind:this={element}
        class="container"
        contenteditable="true"
        on:keydown={onKeyDown}
        on:mouseup={onMouse}
    >
        {#each lines as line, index (line)}
            <Line
                bind:node={elementLines[index]}
                buffer={file}
                line={index}
                tokens={line}
            />
        {/each}
    </div>
</div>

<style>
    .codearea {
        display: flex;
        height: 100%;
        overflow: hidden;
        user-select: none;
        -webkit-user-select: none;
    }

    .container {
        padding: 0 5px;
        text-align: left;
        display: inline;
        outline: none;
        user-select: none;
        -webkit-user-select: none;
    }

    .gutter {
        padding: 0 5px;
        box-sizing: border-box;
        border-right: 1px black solid;
    }

    .gutter-line {
        white-space: pre;
        display: block;
        text-align: right;
        outline: none;
    }

    :global(code, div.line) {
        display: block;
        font: 15px/24px "Fira Code", sans-serif;
        letter-spacing: 1px;
        user-select: text;
    }
</style>
