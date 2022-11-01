<script lang="ts">
    import Token from "./token.svelte";
    import { MouseDown } from "../../wailsjs/go/api/Api.js";

    export let node = null;
    export let buffer = "";
    export let line = 0;
    export let tokens = [];

    function mousedown(event) {
        event.stopPropagation();

        var offset = 0;
        if (tokens.length > 0) {
            offset =
                tokens[tokens.length - 1].offset +
                tokens[tokens.length - 1].value.length;
        }

        MouseDown({
            buffer: buffer,
            line: line,
            offset: offset,
        }).then((err) => {
            if (err) {
                console.error(err);
                return;
            }
        });
    }
</script>

<div
    bind:this={node}
    class="line"
    id={`${buffer}-${line + 1}`}
    contenteditable={true}
    on:mousedown|self={mousedown}
>
    {#if tokens.length > 0}
        {#each tokens as token, id (token)}
            <Token
                {buffer}
                {line}
                token={id}
                text={token.value}
                classes={token.classes}
                offset={token.offset}
            />
        {/each}
    {:else}
        <Token
            {buffer}
            {line}
            token={0}
            text={"\u200B"}
            classes={[]}
            offset={0}
        />
    {/if}
</div>

<style>
    .line {
        /* width: inherit; */
        user-select: none;
        -webkit-user-select: none;
    }

    /*
    div:empty {
        white-space: pre;
        tab-size: 4;
        outline: none;
        height: 100%;
    } */
</style>
