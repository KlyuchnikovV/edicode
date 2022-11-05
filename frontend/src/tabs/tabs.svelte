<script lang="ts">
    import Button from "./button.svelte";
    import Codearea from "../codearea/codearea.svelte";
    import { onMount } from "svelte";
    import { GetBufferNames } from "../../wailsjs/go/api/Api.js";
    import { EventsOn, EventsEmit } from "../../wailsjs/runtime/runtime";

    export let config = {
        header: {
            height: "2em",
        },
    };

    let buffers = [];
    let selected = "";

    function setSelected(name) {
        selected = name;
        EventsEmit("tab_changed", selected);
    }

    function onclick(id) {
        return function () {
            setSelected(id);
        };
    }

    onMount(async () => {
        buffers = await GetBufferNames();
        setSelected(buffers[buffers.length - 1]);
    });

    EventsOn(`buffer_opened`, async (name) => {
        buffers = await GetBufferNames();
        setSelected(name.data);
    });

    EventsOn(`buffer_closed`, async (name) => {
        buffers = await GetBufferNames();
        if (selected === name.data) {
            setSelected(name.data);
        }
    });
</script>

<div class="tabs-container">
    <div
        class="header"
        style:height={config.header.height}
        style:min-width={config.header.height}
    >
        {#each buffers as buffer (buffer)}
            <Button
                {buffer}
                selected={selected === buffer}
                onclick={onclick(buffer)}
                config={{
                    height: config.header.height,
                    icon: "description",
                }}
            />
        {/each}
    </div>
    <div class="panel-container" style:margin-top={config.header.height}>
        {#each buffers as buffer (buffer)}
            <div
                class="child {selected === buffer ? 'visible' : 'not-visible'}"
            >
                <Codearea {buffer} />
            </div>
        {/each}
    </div>
</div>

<style>
    .visible {
        display: block;
    }

    .not-visible {
        display: none;
    }

    .child {
        height: 100%;
    }

    .tabs-container {
        display: flex;
        position: relative;
        height: 100%;
    }

    .panel-container {
        display: flex;
        flex-direction: column;
        overflow: hidden;
        width: 100%;
        min-width: 0px;
    }

    .header {
        display: inline-flex;
        width: 100%;
        border-bottom: 1px solid black;
        user-select: none;
        top: 0; /* Stay at the top */
        z-index: 1; /* Stay on top */
        overflow: hidden;
        position: absolute;
    }

    .header::-webkit-scrollbar {
        height: 2px;
        outline: 1px solid slategrey;
    }

    .header:hover {
        overflow-x: auto;
    }
</style>
