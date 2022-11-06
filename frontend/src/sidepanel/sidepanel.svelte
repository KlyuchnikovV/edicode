<script lang="ts">
    import { EventsOn } from "../../wailsjs/runtime/runtime";
    import Button from "./button.svelte";

    export let items = [];
    export let toggle;
    export let config = {
        debugPanel: {
            height: "2em",
        },
        sidePanel: {
            width: "3em",
        },
    };

    let selected = items[0].argument.name;
    let visible = true;

    function onclick(id: string) {
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
    
    EventsOn("sidepanel_open", (name: any) => {
        console.log(`selected is ${selected}, name is ${name.data}`);
        selected = name.data;
    });
</script>

<div class="panel">
    <div
        class="center-element debug-panel sticky"
        style:height={config.debugPanel.height}
        style:line-height={config.debugPanel.height}
        style:margin-right={config.sidePanel.width}
    >
        Debug panel
    </div>
    <div
        class="center-element"
        style:margin-top={config.debugPanel.height}
        style:margin-right={config.sidePanel.width}
    >
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
    <span
        class="sidepanel sticky"
        style:width={config.sidePanel.width}
        style:min-width={config.sidePanel.width}
        style:line-height={config.sidePanel.width}
    >
        {#each items as item, id (id)}
            <Button
                selected={selected === item.argument.name}
                onclick={onclick(item.argument.name)}
                config={{
                    icon: item.argument.icon,
                    width: config.sidePanel.width,
                }}
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

    .panel {
        /* display: inline-flex;
        position: relative; */
        height: 100%;
    }

    .debug-panel {
        width: 100%;
        border-bottom: 1px solid black;
        vertical-align: middle;
    }

    .center-element {
        display: flex;
        flex-direction: column;
        height: 100%;
        width: inherit;
        min-width: 0px;
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
        right: 0;
        background-color: #504e56;
    }

    .sticky {
        top: 0; /* Stay at the top */
        z-index: 1; /* Stay on top */
        overflow: hidden; /* Disable horizontal scroll */
        position: absolute;
    }
</style>
