<script>
    import Token from "./token.svelte";

    export let node;
    export let buffer = "";
    export let line = 0;
    export let tokens = [];
    export let sendSelection;
    export let tokenSendSelection;
</script>

<div
    bind:this={node}
    class="line"
    id={`${buffer}-${line + 1}`}
    on:mouseup={(event) => {
        console.log(`triggered by line ${line+1}`)
        event.stopPropagation();
        sendSelection(line);
    }}
>
    {#if tokens.length > 0}
        {#each tokens as value, token (value)}
            <Token
                {buffer}
                {line}
                {token}
                text={value.value}
                classes={value.classes}
                sendSelection={tokenSendSelection}
            />
        {/each}
    {:else}
        <Token
            {buffer}
            {line}
            token={0}
            text={"\u200B"}
            classes={[]}
            sendSelection={tokenSendSelection}
        />
    {/if}
</div>

<style>
    .line {
        width: inherit;
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
