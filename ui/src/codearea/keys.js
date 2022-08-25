export default class Keys {
    constructor(file) {
        this.file = file
    }

    onKeyDown(event, cursor, elementLines) {
        var keyboardEvent = {
            buffer: this.file,
            key: event.key,
            alt: event.altKey,
            ctrl: event.ctrlKey,
            shift: event.shiftKey,
            meta: event.metaKey,
            // startLine: cursor.selection.start.line,
            // endLine: cursor.selection.end.line,
            // startOffset: cursor.selection.start.offset,
            // endOffset: cursor.selection.end.offset,
        }

        if (event.metaKey) {
            return
        }

        if (event.key.length === 1 && !event.ctrlKey && !event.altKey) {
            event.preventDefault();
            window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                if (err !== undefined) {
                    console.log(err);
                }
            });
            // cursor.selection.right(event.shiftKey, true);
            // cursor.setCurrentCursorPosition(elementLines);
            return
        }

        event.preventDefault();


        switch (event.key) {
            case "ArrowUp":
            case "ArrowDown":
            case "ArrowLeft":
            case "ArrowRight":
                window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                    if (err !== undefined) {
                        console.log(err);
                    }
                });
                // cursor.onPress(event, elementLines);
                // cursor.setCurrentCursorPosition(elementLines);
                // return;
                break;
            case "Backspace":
                window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                    if (err !== undefined) {
                        console.log(err);
                    }
                });
                // if (cursor.selection.start.line === cursor.selection.end.line
                //     && cursor.selection.start.offset === cursor.selection.end.offset) {
                //     cursor.selection.left(event.shiftKey);
                // }
                break;
            case "Enter":
                window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                    if (err !== undefined) {
                        console.log(err);
                    }
                });
                // cursor.selection.down(event.shiftKey, true);
                // cursor.offset = 0;
                break;
            case "Tab":
                window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                    if (err !== undefined) {
                        console.log(err);
                    }
                });
                // cursor.selection.right(event.shiftKey, true);
                break;
            default:
                return;
        }

        // if (cursor.selection.start.line !== cursor.selection.end.line
        //     || cursor.selection.start.offset !== cursor.selection.end.offset) {
        //     cursor.selection.collapse()
        // }

        // cursor.setCurrentCursorPosition(elementLines);
    }

    // onKeyPress(event, cursor, lines) {
    //     console.log(`keypress ${event}`)
    //     var keyboardEvent = {
    //         buffer: this.file,
    //         key: event.key,
    //         alt: event.altKey,
    //         ctrl: event.ctrlKey,
    //         shift: event.shiftKey,
    //         meta: event.metaKey,
    //         // startLine: cursor.selection.start.line,
    //         // endLine: cursor.selection.end.line,
    //         // startOffset: cursor.selection.start.offset,
    //         // endOffset: cursor.selection.end.offset,
    //     }

    //     if (!event.metaKey) {
    //         return
    //     }

    //     switch (event.key) {
    //         case 'v':
    //             window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
    //                 if (err !== undefined) {
    //                     console.log(err);
    //                 }
    //             });
    //             break;
    //         case 'c':
    //             window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
    //                 if (err !== undefined) {
    //                     console.log(err);
    //                 }
    //             });
    //             break;
    //     }
    //     event.preventDefault()
    // }
}
