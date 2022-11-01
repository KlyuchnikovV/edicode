<script lang="ts">
    import Codearea from "../codearea/codearea.svelte";
    import {
        TabWrapper,
        TabHead,
        TabHeadItem,
        TabContentItem,
    } from "flowbite-svelte";

    export let buffers = [];
    let activeTabValue = 1;
</script>

<div class="tabs">
    <TabWrapper
        tabStyle="pill"
        class="tab-wrapper"
        bind:activeTabValue
        let:tabStyle
        let:tabId
    >
        <TabHead {tabStyle} {tabId}>
            {#each buffers as buffer, index (buffer)}
                <TabHeadItem
                    id={index}
                    {tabStyle}
                    {activeTabValue}
                    on:click={() => {
                        activeTabValue = index;
                    }}>{buffer}</TabHeadItem
                >
            {/each}
        </TabHead>

        {#each buffers as buffer, index (buffer)}
            <TabContentItem
                contentDivClass="content-item"
                id={index}
                {activeTabValue}
            >
                <Codearea file={buffer} />
            </TabContentItem>
        {/each}
    </TabWrapper>
</div>

<style>
    .tabs {
        height: 100%;
    }

    :global(.content-item) {
        padding: 0px;
        /* height: 100%; */
        overflow: scroll;
        flex-grow: 1;
        /* -webkit-scrollbar: black; */
    }

    :global(.tab) {
        display: inline;
    }

    :global(.tab-wrapper) {
        height: 100%;
        display: flex;
        flex-flow: column;
    }

    :global(ul.flex) {
        overflow: hidden;
        flex-wrap: nowrap;
        border-bottom: 1px solid black;
    }
</style>
