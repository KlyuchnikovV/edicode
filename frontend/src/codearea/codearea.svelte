<script>
    import Line from "./line.svelte";
    import Cursor from "./cursor";
    import { onMount } from "svelte";
    import { EventsOn } from "../../wailsjs/runtime/runtime";
    import {
        GetBuffer,
        GetCursor,
        HandleKeyboardEvent,
    } from "../../wailsjs/go/api/Api";

    export let file = "";
    
    let lines = [];
    let cursor;
    let elementLines = [];
    let codearea;

    onMount(async () => {
        cursor = new Cursor(file, codearea);

        var bufferInit = async () => {
            lines = (await GetBuffer(file)).lines;
        };

        bufferInit();

        EventsOn(`highlight_changed_${file}`, (event) => {
            lines = event.data.lines;

            GetCursor(file).then((selection, err) => {
                if (err) {
                    console.error(err);
                } else {
                    cursor.setSelection(selection, elementLines);
                }
            });
        });

        EventsOn(`cursor_moved_${file}`, (event) => {
            cursor.setSelection(event.data, elementLines);
        });
    });

    function onKeyDown(event) {
        HandleKeyboardEvent({
            buffer: file,
            key: event.key,
            alt: event.altKey,
            ctrl: event.ctrlKey,
            shift: event.shiftKey,
            meta: event.metaKey,
        }).then((err) => {
            if (err) {
                console.error(err);
            }
        });
    }
</script>

<div id={`codearea-${file}`} class="codearea">
    <div class={"gutter"}>
        {#each lines as _, index (index)}
            <code class={"gutter-line"}>{index + 1}</code>
        {/each}
    </div>
    <div
        bind:this={codearea}
        class="container"
        contenteditable="true"
        on:keydown|self|capture|stopPropagation|preventDefault={onKeyDown}
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
        overflow: scroll;
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
        resize: horizontal;
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
