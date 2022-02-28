// // let code = "function add(a, b) {\n\treeturn a + b\n}";
// // let codeElement;

// // let selections = [];
// let currentSelection = null;

// export function updateSelection(event) {
//     currentSelection = {
//         start: event.currentTarget.selectionStart,
//         end: event.currentTarget.selectionEnd,
//     };
// }

// // function recordSelection() {
// //     if (!currentSelection) return;

// //     selections = [...selections, currentSelection];
// // }

// export function playbackSelection(codeElement) {
//     // selections.forEach((selection) => {
//     // setTimeout(() => {
//     // currentSelection = selection;
//     selection = currentSelection
//     codeElement.focus();
//     codeElement.setSelectionRange(selection.start, selection.end);
//     //}, index * 1000);
//     // });
//     // clear at the end
//     // setTimeout(() => {
//     //     window.getSelection().empty();
//     //     currentSelection = null;
//     // }, selections.length * 1000);
// }

export default class Selection {
    constructor() {
        this.start = {
            line: 1,
            offset: 0,
        }
        this.end = {
            line: 1,
            offset: 0,
        }
        this.lines = []
        this.goingDown = null
    }

    up(shifted) {
        if (this.start.line > 1) {
            this.start.line--
        } else {
            this.start.offset = 0;
        }

        if (!shifted) {
            this.end = {
                line: this.start.line,
                offset: this.start.offset,
            }
        }

        this.goingDown = false

        console.log(`up - start { line: ${this.start.line}, offset: ${this.start.offset} }, end { line: ${this.end.line}, offset: ${this.end.offset} }`)

    }

    down(shifted, grown) {
        let lineLen = this.lenOfLine(this.end.line - 1);
        if (this.end.line < this.lines.length || (this.end.line == this.lines.length && grown)) {
            this.end.line++;
        } else {
            this.end.offset = lineLen;
        }

        if (!shifted) {
            this.start = {
                line: this.end.line,
                offset: this.end.offset,
            }
        }

        this.goingDown = true

        console.log(`down - start { line: ${this.start.line}, offset: ${this.start.offset} }, end { line: ${this.end.line}, offset: ${this.end.offset} }`)
    }

    left(shifted) {
        let lineLen = this.lenOfLine(this.start.line - 1);
        if (this.start.offset > lineLen) {
            this.start.offset = lineLen;
        }
        if (this.start.offset > 0) {
            this.start.offset--;
        } else if (this.start.line <= 1) {
            this.start.offset = 0;
        } else {
            this.start.line--;
            lineLen = 0;
            this.lines[this.start.line - 1].forEach((element) => {
                lineLen += element.value.length;
            });
            this.start.offset = lineLen;
        }

        if (!shifted) {
            this.end = {
                line: this.start.line,
                offset: this.start.offset,
            }
        }

        this.goingDown = false
    }

    right(shifted, grown) {
        let lineLen = this.lenOfLine(this.end.line - 1);
        if (this.end.offset > lineLen) {
            this.end.offset = lineLen;
        }


        if (this.end.offset > lineLen || (this.end.offset == lineLen && grown)) {
            if (grown) {
                this.end.offset = lineLen + 1
            } else {
                this.end.offset = lineLen
            }
        } else if (this.end.offset < lineLen) {
            this.end.offset++
        } else if (this.end.line < this.lines.length) {
            this.end.line++
            this.end.offset = 0
        }

        if (!shifted) {
            this.start = {
                line: this.end.line,
                offset: this.end.offset,
            }
        }
        this.goingDown = true
    }

    lenOfLine(line) {
        var tokens = this.lines[line]
        if (tokens === undefined) {
            return 0
        }

        let lineLen = 0;
        tokens.forEach((element) => {
            lineLen += element.value.length;
        });
        return lineLen
    }
}
