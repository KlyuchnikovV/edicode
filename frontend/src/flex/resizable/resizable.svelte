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
                break;
            case "top":
            case "bottom":
                y = e.clientY;
                h = parseInt(styles.height, 10);
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
                break;
            case "right":
                if (w + dx > minimalWidth) {
                    element.style.width = `${w + dx}px`;
                } else {
                    element.style.width = config.width.collapsed;
                }
                break;
            case "top":
                if (h - dy > minimalHeight) {
                    element.style.height = `${h - dy}px`;
                } else {
                    element.style.height = config.height.collapsed;
                }
                break;
            case "bottom":
                if (h + dy > minimalHeight) {
                    element.style.height = `${h + dy}px`;
                } else {
                    element.style.height = config.height.collapsed;
                }
                break;
        }
    };

    const mouseUpHandler = function () {
        // Remove the handlers of `mousemove` and `mouseup`
        document.removeEventListener("mousemove", mouseMoveHandler);
        document.removeEventListener("mouseup", mouseUpHandler);
    };

    onMount(() => {
        element.style.width = config.width.initial;
        element.style.height = config.height.initial;
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

<div bind:this={element} class="resizable">
    <slot toggle={togglePanel} />
    <div
        class={`resizer resizer-${config.position}  ${config.position}`}
        on:mousedown={mouseDownHandler}
    />
</div>

<style>
    ::-webkit-scrollbar {
        width: 0;
    }

    .collapsed {
        border: none;
    }

    .visible {
        display: block;
        /* visibility: visible; */
    }

    .not-visible {
        display: none;
        /* visibility: hidden;
        max-height: 0px;
        max-width: 0px; */
    }

    .resizable {
        position: relative;
        height: 100%;
        /* overflow: scroll; */
    }

    .resizer {
        /* All resizers are positioned absolutely inside the element */
        position: absolute;
        z-index: 999999;
    }

    /* Placed at the right side */
    .resizer-left,
    .resizer-right {
        width: 5px;
        z-index: 10;
    }

    .resizer-top,
    .resizer-bottom {
        height: 5px;
        z-index: 10;
    }

    .left {
        cursor: col-resize;
        height: 100%;
        left: -5px;
        top: 0;
    }

    .right {
        cursor: col-resize;
        height: 100%;
        right: -5px;
        top: 0;
    }

    .top {
        cursor: row-resize;
        width: 100%;
        left: 0;
        top: -5px;
    }

    .bottom {
        cursor: row-resize;
        width: 100%;
        left: 0;
        bottom: -5px;
    }

    :hover.resizer {
        background-color: darkslategrey;
    }
</style>
