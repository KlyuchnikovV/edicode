export default class Keys {
    constructor(file) {
        this.file = file
    }

    onKeyDown(event, cursor, elementLines) {
        var offsetLocal = cursor.offset;
        var lineLen = cursor.selection.lenOfLine(cursor.line - 1);

        if (offsetLocal > lineLen) {
            offsetLocal = lineLen;
        }

        var keyboardEvent = {
            buffer: this.file,
            key: event.key,
            alt: event.altKey,
            ctrl: event.ctrlKey,
            shift: event.shiftKey,
            line: cursor.line,
            offset: offsetLocal,
        }

        if (event.key.length === 1) {
            window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                if (err !== undefined) {
                    console.log(err);
                }
            });
            cursor.right(true);
            // cursor.setCurrentCursorPosition(elementLines);
            return
        }

        event.preventDefault();
        switch (event.key) {
            case "ArrowUp":
            case "ArrowDown":
            case "ArrowLeft":
            case "ArrowRight":
                cursor.onPress(event);
                break;
            case "Backspace":
                window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                    if (err !== undefined) {
                        console.log(err);
                    }
                });
                cursor.left();
                break;
            case "Enter":
                window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                    if (err !== undefined) {
                        console.log(err);
                    }
                });
                cursor.down(event.shiftKey, true);
                cursor.offset = 0;
                break;
            case "Tab":
                window.backend.Api.HandleKeyboardEvent(keyboardEvent).then((err) => {
                    if (err !== undefined) {
                        console.log(err);
                    }
                });
                cursor.right(true);
                break;
            default:
                return;
        }
        cursor.setCurrentCursorPosition(elementLines);
    }
}
