<script>
    import { MakeAction } from "../../wailsjs/go/plugins/Manager";
    import { EventsOn } from "../../wailsjs/runtime/runtime";
    import { GetActionsList } from "../../wailsjs/go/api/Api";

    export let config = {
        fontSize: "15px",
        fontFamily: '"Fira Code", sans-serif',
        lineHeight: "24px",
    };

    let value = "";
    let actions = [];

    function oninput(event) {
        console.log(event);
        console.log(value);

        GetActionsList(value).then((result) => {
            console.log("actions");
            console.log(result);
            actions = result;
        });
    }

    function onclick(event, action) {
        MakeAction(action);
    }

    EventsOn(`plugins_loaded_all`, async () => {
        GetActionsList(value).then((result) => {
            console.log("actions");
            console.log(result);
            actions = result;
        });
    });
</script>

<div class="explorer">
    <input
        bind:value
        type="text"
        class="palette"
        placeholder="Search..."
        on:input={oninput}
        style:font-size={config.fontSize}
        style:font-family={config.fontFamily}
        style:line-height={config.lineHeight}
    />
    <div>
        {#each actions as action, id (id)}
            <div
                id={`action-${id}`}
                class="action"
                on:click={(event) => {
                    onclick(event, action.action);
                }}
            >
                {action.name}
            </div>
        {/each}
    </div>
</div>

<style>
    .explorer {
        /* margin: 1em; */
        -webkit-flex: 1;
        -ms-flex: 1;
        flex: 1;
        resize: horizontal;
        /* background-color: #1e2025; */
        /* border-left: 1px #cccccc solid; */
        display: block;
        margin: 1em;
    }

    .palette {
        /* width: 90%; */
        /* margin-top: 1em; */

        /* font: 15px/24px "Fira Code", sans-serif; */
        letter-spacing: 1px;
        user-select: text;
        color: black;
        /* margin-right: 2em; */
        flex: 1;
        width: 100%;

        /* min-width: 0px; */
        /* width: calc(100%-2em); */

        /* background-color: #282c34;
        color: #abb2bf; */
    }

    .action {
        /* font: 15px/24px "Fira Code", sans-serif; */
        letter-spacing: 1px;
        user-select: text;

        width: 90%;
        margin-top: 1em;
        cursor: pointer;
        overflow: hidden;
        /* background-color: #1e2025; */
    }

    /* .action:hover {
        background-color: gray;
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
        /* font: 15px/24px "Fira Code", sans-serif; */
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
