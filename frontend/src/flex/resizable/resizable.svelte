<script lang="ts">
    import { onMount } from "svelte";

    export let config = {
        position: "left",
        width: {
            min: "100px",
            initial: "400px",
            collapsed: "3em",
        },
        height: {
            min: "50px",
            initial: "200px",
            collapsed: "3em",
        },
    };

    let element;
    let minimalWidth = parseInt(config.width.min, 10);
    let minimalHeight = parseInt(config.height.min, 10);

    let isToggled = true;
    let collapsed = false;

    // The current position of mouse
    let x = 0;
    let y = 0;

    // The dimension of the element
    let w = parseInt(config.width.initial, 10);
    let h = parseInt(config.height.initial, 10);

    // Handle the mousedown event
    // that's triggered when user drags the resizer
    const mouseDownHandler = function (e) {
        // Calculate the dimension of element
        const styles = window.getComputedStyle(element);

        switch (config.position) {
            case "left":
            case "right":
                x = e.clientX;
                w = parseInt(styles.width, 10);
                document.body.style.cursor = "col-resize";
                break;
            case "top":
            case "bottom":
                y = e.clientY;
                h = parseInt(styles.height, 10);
                document.body.style.cursor = "row-resize";
                break;
        }

        // Attach the listeners to `document`
        document.addEventListener("mousemove", mouseMoveHandler);
        document.addEventListener("mouseup", mouseUpHandler);
    };

    const mouseMoveHandler = function (e) {
        // How far the mouse has been moved
        const dx = e.clientX - x,
            dy = e.clientY - y;

        switch (config.position) {
            case "left":
                if (w - dx > minimalWidth) {
                    element.style.width = `${w - dx}px`;
                } else {
                    element.style.width = config.width.collapsed;
                }
                collapsed = element.style.width === config.width.collapsed;
                break;
            case "right":
                if (w + dx > minimalWidth) {
                    element.style.width = `${w + dx}px`;
                } else {
                    element.style.width = config.width.collapsed;
                }
                collapsed = element.style.width === config.width.collapsed;
                break;
            case "top":
                if (h - dy > minimalHeight) {
                    element.style.height = `${h - dy}px`;
                } else {
                    element.style.height = config.height.collapsed;
                }
                collapsed = element.style.height === config.height.collapsed;
                break;
            case "bottom":
                if (h + dy > minimalHeight) {
                    element.style.height = `${h + dy}px`;
                } else {
                    element.style.height = config.height.collapsed;
                }
                collapsed = element.style.height === config.height.collapsed;
                break;
        }
    };

    const mouseUpHandler = function () {
        // Remove the handlers of `mousemove` and `mouseup`
        document.removeEventListener("mousemove", mouseMoveHandler);
        document.removeEventListener("mouseup", mouseUpHandler);

        document.body.style.cursor = "default";
    };

    onMount(() => {
        element.style.width = config.width.initial;
        element.style.height = config.height.initial;

        switch (config.position) {
            case "left":
            case "right":
                collapsed = element.style.width === config.width.collapsed;
                break;
            case "top":
            case "bottom":
                collapsed = element.style.height === config.height.collapsed;
                break;
        }
    });

    function togglePanel() {
        isToggled = !isToggled;

        switch (config.position) {
            case "left":
            case "right":
                if (isToggled) {
                    element.style.width = `${w}px`;
                } else {
                    element.style.width = config.width.collapsed;
                }
                break;
            case "top":
            case "bottom":
                if (isToggled) {
                    element.style.heigh = `${h}px`;
                } else {
                    element.style.height = config.height.collapsed;
                }
                break;
        }
    }
</script>

<div bind:this={element} class="resizable" class:collapsed>
    <slot toggle={togglePanel} />
    <div class={`resizer ${config.position}`} on:mousedown={mouseDownHandler} />
</div>

<style>
    .resizable {
        position: relative;
        height: 100%;
    }

    .resizer.top,
    .resizer.bottom {
        position: absolute;
        cursor: row-resize;
        height: 5px;
        width: 100%;
        left: 0;
        z-index: 10;
    }

    .resizer.left,
    .resizer.right {
        position: absolute;
        cursor: col-resize;
        height: 100%;
        width: 5px;
        top: 0;
        z-index: 10;
    }

    .top {
        top: -5px;
    }
    .left {
        left: -5px;
    }
    .right {
        right: -5px;
    }
    .bottom {
        bottom: -5px;
    }

    .collapsed > .top {
        cursor: n-resize;
    }
    .collapsed > .left {
        cursor: w-resize;
    }
    .collapsed > .right {
        cursor: e-resize;
    }
    .collapsed > .bottom {
        cursor: s-resize;
    }

    :hover.resizer {
        background-color: darkslategrey;
    }

    ::-webkit-scrollbar {
        width: 0;
    }
</style>
