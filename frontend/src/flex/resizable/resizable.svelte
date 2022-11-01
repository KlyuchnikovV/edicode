<script lang="ts">
    import { onMount } from "svelte";

    export let position = "left";
    export let width = "400px";
    export let minWidth = "100px";
    export let minHeight = "50px";
    export let collapsedWidth = "3em";
    export let collapsedHeight = "3em";

    let element;
    let minimalWidth = parseInt(minWidth, 10);
    let minimalHeight = parseInt(minHeight, 10);

    let isToggled = true;

    // The current position of mouse
    let x = 0;
    let y = 0;

    // The dimension of the element
    let w = parseInt(width, 10);
    let h = 0;

    // Handle the mousedown event
    // that's triggered when user drags the resizer
    const mouseDownHandler = function (e) {
        // Calculate the dimension of element
        const styles = window.getComputedStyle(element);

        switch (position) {
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

        switch (position) {
            case "left":
                if (w - dx > minimalWidth) {
                    element.style.width = `${w - dx}px`;
                } else {
                    element.style.width = collapsedWidth;
                }
                break;
            case "right":
                if (w - dx > minimalWidth) {
                    element.style.width = `${w + dx}px`;
                } else {
                    element.style.width = collapsedWidth;
                }
                break;
            case "top":
                if (h - dy > minimalHeight) {
                    element.style.height = `${h - dy}px`;
                } else {
                    element.style.height = collapsedHeight;
                }
                break;
            case "bottom":
                if (h - dy > minimalHeight) {
                    element.style.heigh = `${h + dy}px`;
                } else {
                    element.style.heigh = collapsedHeight;
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
        element.style.width = `${w}px`;
    });

    function togglePanel() {
        isToggled = !isToggled;

        switch (position) {
            case "left":
            case "right":
                if (isToggled) {
                    element.style.width = `${w}px`;
                } else {
                    element.style.width = collapsedWidth;
                }
                break;
            case "top":
            case "bottom":
                if (isToggled) {
                    element.style.heigh = `${h}px`;
                } else {
                    element.style.height = collapsedHeight;
                }
                break;
        }
    }
</script>

<div bind:this={element} class="resizable">
    <slot toggle={togglePanel} />
    <div
        class={`resizer resizer-r  ${position}`}
        on:mousedown={mouseDownHandler}
    />
</div>

<style>
    ::-webkit-scrollbar {
        width: 0;
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
    }

    /* Placed at the right side */
    .resizer-r {
        width: 5px;
        z-index: 10;
    }

    .left {
        cursor: col-resize;
        height: 100%;
        left: -1px;
        top: 0;
    }

    .right {
        cursor: col-resize;
        height: 100%;
        right: 0;
        top: 0;
    }

    .top {
        cursor: row-resize;
        width: 100%;
        left: 0;
        top: 0;
    }

    .bottom {
        cursor: row-resize;
        width: 100%;
        left: 0;
        bottom: 0;
    }

    :hover.resizer-r {
        background-color: darkslategrey;
    }
</style>
