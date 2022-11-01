<script lang="ts">
    import { onMount } from "svelte";
    import { EventsOn } from "../../wailsjs/runtime/runtime";
    import Button from "./button.svelte";

    export let items = [];
    export let toggle;

    let selected = items[0].argument.name;
    let visible = true;

    function onclick(id) {
        return function () {
            if (!visible) {
                visible = true;
                selected = id;
                toggle();
            } else {
                if (selected === id) {
                    visible = false;
                    toggle();
                } else {
                    selected = id;
                }
            }
        };
    }

    EventsOn("sidepanel_open", (name) => {
        console.log(`selected is ${selected}, name is ${name.data}`);
        selected = name.data;
    });
</script>

<div class="container">
    <div class="debug-panel sticky center-element">Debug panel</div>
    <div class="panel-container center-element">
        {#each items as item, id (id)}
            <div
                class="child {selected === item.argument.name
                    ? 'visible'
                    : 'not-visible'}"
            >
                <svelte:component
                    this={item.component}
                    plugin={item.argument.name}
                    file={item.argument.file}
                />
            </div>
        {/each}
    </div>
    <span class="sidepanel sticky">
        {#each items as item, id (id)}
            <Button
                selected={selected === item.argument.name}
                icon={item.argument.icon}
                onclick={onclick(item.argument.name)}
            />
        {/each}
    </span>
</div>

<style>
    .visible {
        display: block;
    }

    .not-visible {
        display: none;
    }

    .child {
        overflow-y: auto;
        height: 100%;
    }

    .container {
        display: inline-flex;
        position: relative;
        height: 100%;
    }

    .panel-container {
        margin-top: 3em;
    }

    .debug-panel {
        height: 3em;
        width: 100%;
        border-bottom: 1px solid black;
        padding-top: 0.75em;
    }

    .center-element {
        display: flex;
        flex-direction: column;
        overflow: hidden;
        width: inherit;
        min-width: 0px;
        margin-right: 3em;
        /* Firefox */
        width: -moz-calc(100% - 3em);
        /* Opera */
        width: -o-calc(100% - 3em);
        /* WebKit */
        width: -webkit-calc(100% - 3em);
        /* Standard */
        width: calc(100% - 3em);
    }

    .sidepanel {
        z-index: 3;
        height: 100%; /* Full-height: remove this if you want "auto" height */
        width: 3em; /* Set the width of the sidebar */
        right: 0;
        background-color: #504e56;
        /* padding-top: 0.25em; */
    }

    .sticky {
        top: 0; /* Stay at the top */
        z-index: 1; /* Stay on top */
        overflow-x: hidden; /* Disable horizontal scroll */
        min-width: 2em;
        position: absolute;
    }
</style>
