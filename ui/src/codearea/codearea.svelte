<script>
    export let file = "";

    import Line from "./line.svelte";
    import Cursor from "./cursor";
    import Keys from "./keys";
    import { afterUpdate, beforeUpdate, onMount } from "svelte";

    let lines = [];
    let element, cursor, keys;
    let elementLines = [];

    onMount(async () => {
        keys = new Keys(file, cursor);
        cursor = new Cursor(file, element);

        var bufferInit = async () => {
            let buf = await window.backend.Api.GetBuffer(file);
            console.log(`buffer loaded ${file}`);
            lines = buf.lines;
        };

        bufferInit();

        window.wails.Events.On(`highlight_changed_${file}`, (event) => {
            console.log(`highlight changed ${file}`);

            // console.log(event.data.lines);
            lines = event.data.lines;

            window.backend.Api.GetCursor(file).then((offset, err) => {
                if (err !== undefined) {
                    console.log(err);
                    return;
                }

                console.log(`cursor is ${offset.cursor}`);
                cursor.setCurrentCursorPosition(elementLines, offset);
            });
        });

        window.wails.Events.On(`cursor_moved_${file}`, (event) => {
            cursor.setCurrentCursorPosition(elementLines, event.data);
        });
    });

    function onKeyDown(event) {
        keys.onKeyDown(event, cursor, elementLines);
    }

    function onKeyPress(event) {
        // keys.onKeyPress(event, cursor, elementLines);
    }
</script>

<div class="codearea">
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
        on:keypress={onKeyPress}
    >
        {#each lines as line, index (line)}
            <Line
                bind:node={elementLines[index]}
                buffer={file}
                line={index}
                tokens={line}
                lineCursor={function (line) {
                    cursor.onMouseByLine(line, elementLines);
                }}
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
