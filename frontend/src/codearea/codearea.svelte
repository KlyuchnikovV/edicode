<script>
    export let file = "";

    import Line from "./line.svelte";
    import Cursor from "./cursor";
    import Keys from "./keys";
    import { onMount } from "svelte";
    import { GetBuffer, GetCursor } from "../../wailsjs/go/api/Api";
    import { EventsOn } from "../../wailsjs/runtime/runtime";

    let lines = [];
    let cursor, keys;
    let elementLines = [];
    let codearea;

    onMount(async () => {
        keys = new Keys(file);
        cursor = new Cursor(file, codearea);

        var bufferInit = async () => {
            let buf = await GetBuffer(file);
            console.log(`buffer loaded ${file}`);
            lines = buf.lines;
        };

        bufferInit();

        EventsOn(`highlight_changed_${file}`, (event) => {
            console.log(`highlight changed ${file}`);

            lines = event.data.lines;

            GetCursor(file).then((selection, err) => {
                if (err) {
                    console.log(err);
                    return;
                }

                console.log(selection);
                cursor.setSelection(selection, elementLines);
                console.log("set");
            });
        });

        EventsOn(`cursor_moved_${file}`, (event) => {
            console.log(event);
            console.log(elementLines);
            cursor.setSelection(event.data, elementLines);
        });
    });

    function onKeyDown(event) {
        console.log(event);
        keys.onKeyDown(event, cursor, elementLines);
    }
</script>

<div id={`codearea-${file}`} class="codearea">
    <div class={"gutter"}>
        {#each lines as _, index (index)}
            <code class={"gutter-line"}>{index + 1}</code>
        {/each}
    </div>
    <div bind:this={codearea} class="container" contenteditable="true">
        {#each lines as line, index (line)}
            <Line
                bind:node={elementLines[index]}
                buffer={file}
                line={index}
                tokens={line}
                sendSelection={(line) => {
                    cursor.sendSelection(
                        line,
                        elementLines[line].children.length - 1,
                        elementLines
                    );
                }}
                tokenSendSelection={(line, token) => {
                    cursor.sendSelection(line, token, elementLines);
                }}
            />
        {/each}
    </div>
</div>

<svelte:window on:keydown={onKeyDown} />

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
        width: 100%;
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
