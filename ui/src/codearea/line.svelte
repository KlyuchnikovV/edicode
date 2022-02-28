<script>
    import Token from "./token.svelte";
    import { onMount } from "svelte";

    export let tokens = [];
    export let buffer = "";
    export let line = 0;
    export let node;

    // onMount(async () => {
    //     window.wails.Events.On("line_changed", (event) => {
    //         console.log("highlight upd");
    //         if (event.data === file) {
    //             lines = event.Data.Lines;
    //         }
    //     });

    //     cursor = new Cursor(file, element);
    // });
</script>

<div bind:this={node} class="line" id={`${buffer}-${line + 1}`}>
    {#if tokens.length > 0}
        {#each tokens as token, index (token)}
            <Token
                {buffer}
                line={line + 1}
                token={index}
                text={token.value}
                classes={token.classes}
            />
        {/each}
    {:else}
        <Token
            {buffer}
            line={line + 1}
            token={0}
            text={"\u200B"}
            classes={[]}
        />
    {/if}
</div>

<style>
    .line {
        width: fit-content;
        user-select: none;
        -webkit-user-select: none;
    }

    div:empty {
        white-space: pre;
        tab-size: 4;
        outline: none;
        height: 100%;
    }
</style>
