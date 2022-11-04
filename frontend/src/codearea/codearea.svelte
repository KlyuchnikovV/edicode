<script lang="ts">
    import Line from "./line.svelte";
    import Cursor from "./cursor";
    import { onMount } from "svelte";
    import { EventsOn } from "../../wailsjs/runtime/runtime";
    import {
        GetBuffer,
        GetCursor,
        HandleKeyboardEvent,
    } from "../../wailsjs/go/api/Api";

    export let buffer = "";
    export let shouldScrollPastEnd = true;

    let lines = [];
    let cursor;
    let elementLines = [];
    let codearea;

    $: {
        async()=> {
            lines = (await GetBuffer(buffer)).lines;
        }
    }

    onMount(async () => {
        console.log(`codearea ${buffer}`);
        cursor = new Cursor(buffer, codearea);

        var bufferInit = async () => {
            lines = (await GetBuffer(buffer)).lines;
        };

        bufferInit();

        EventsOn(`highlight_changed_${buffer}`, (event) => {
            lines = event.data.lines;

            GetCursor(buffer).then((selection: any, err: Error) => {
                if (err) {
                    console.error(err);
                } else {
                    cursor.setSelection(selection, elementLines);
                }
            });
        });

        EventsOn(`cursor_moved_${buffer}`, (event) => {
            cursor.setSelection(event.data, elementLines);
        });
    });

    function onKeyDown(event) {
        HandleKeyboardEvent({
            buffer: buffer,
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

<div id={`codearea-${buffer}`} class="codearea">
    <div class="gutter">
        {#each lines as _, index (index)}
            <code class={"gutter-line"}>{index + 1}</code>
        {/each}
    </div>
    <div
        bind:this={codearea}
        class="text-container"
        class:scrollpastend={shouldScrollPastEnd}
        contenteditable="true"
        on:keydown|self|capture|stopPropagation|preventDefault={onKeyDown}
    >
        {#each lines as line, index (line)}
            <Line
                bind:node={elementLines[index]}
                {buffer}
                line={index}
                tokens={line}
            />
        {/each}
    </div>
</div>

<style>
    .codearea {
        /* visibility: hidden; */
        display: flex;
        /* height: 100%; */
        /* overflow: scroll; */
        user-select: none;
        -webkit-user-select: none;
    }

    .text-container {
        padding: 0 5px;
        text-align: left;
        display: inline;
        outline: none;
        user-select: none;
        -webkit-user-select: none;
        /* resize: horizontal; */
        margin-left: 10px;

        overflow: scroll;
    }

    .scrollpastend {
        padding-bottom: 87vh;
    }

    .gutter {
        padding: 0 5px;
        box-sizing: border-box;
        border-right: 1px black solid;
        /* position: absolute; */
        /* background-color: rgba(27, 38, 54, 1); */
    }

    .gutter-line {
        white-space: pre;
        display: block;
        text-align: right;
        outline: none;
        color: gray;
        overflow: hidden;
        user-select: none;
        -webkit-user-select: none;
        /* border-bottom: 1px white solid; */
    }

    :global(code, div.line) {
        display: block;
        font: 15px/24px "Fira Code", sans-serif;
        letter-spacing: 1px;
        user-select: text;
        padding: 0;
    }
</style>
