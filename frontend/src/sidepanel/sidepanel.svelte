<script>
    import { GetActionsList, MakeAction } from "../../wailsjs/go/api/Api.js";
    let value = "";
    let actions = [];

    function oninput(event) {
        console.log(event);
        console.log(value);

        GetActionsList(value).then((result) => {
            actions = result;
        });
    }

    function onclick(event, action) {
        var selectedTab = document.getElementsByClassName(
            "svelte-tabs__selected"
        )[0];
        console.log(selectedTab);

        MakeAction({
            action: action,
            param: selectedTab.textContent,
        });
    }
</script>

<div class="sidepanel">
    <input
        bind:value
        type="text"
        class="palette"
        placeholder="Search.."
        on:input={oninput}
    />
    <div>
        {#each actions as action, id (action)}
            <div
                id={`action-${id}`}
                class="action"
                on:click={(event) => {
                    onclick(event, action);
                }}
            >
                {action}
            </div>
        {/each}
    </div>
</div>

<style>
    .sidepanel {
        -webkit-flex: 1;
        -ms-flex: 1;
        flex: 1;
        resize: horizontal;
        background-color: #1e2025;
        border-left: 1px #cccccc solid;
    }

    .palette {
        width: 90%;
        margin-top: 1em;

        /* display: block; */
        font: 15px/24px "Fira Code", sans-serif;
        letter-spacing: 1px;
        user-select: text;

        background-color: #282c34;
        color: #abb2bf;
    }
    /*
    .action {
        font: 15px/24px "Fira Code", sans-serif;
        letter-spacing: 1px;
        user-select: text;

        width: 90%;
        margin-top: 1em;
        background-color: #1e2025;
    } */

    .action {
        appearance: none;
        background-color: #000000;
        border: 2px solid #1a1a1a;
        border-radius: 15px;
        box-sizing: border-box;
        color: #ffffff;
        cursor: pointer;
        display: inline-block;
        font: 15px/24px "Fira Code", sans-serif;
        line-height: normal;
        margin: 0;
        min-height: 60px;
        min-width: 0;
        outline: none;
        padding: 16px 24px;
        text-align: center;
        text-decoration: none;
        transition: all 300ms cubic-bezier(0.23, 1, 0.32, 1);
        user-select: none;
        -webkit-user-select: none;
        touch-action: manipulation;
        width: 90%;
        will-change: transform;
        margin-top: 1em;
    }

    .action:disabled {
        pointer-events: none;
    }

    .action:hover {
        box-shadow: rgba(0, 0, 0, 0.25) 0 8px 15px;
        transform: translateY(-2px);
    }

    .action:active {
        box-shadow: none;
        transform: translateY(0);
    }
</style>
