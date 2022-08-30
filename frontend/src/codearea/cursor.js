import { SetCursor } from "../../wailsjs/go/api/Api.js";

export default class Cursor {
    constructor(file, codearea) {
        this.file = file;
        this.codearea = codearea;
    }

    setSelection(selection, lines) {
        var start = this.getNode(lines[selection.start.line].children, selection.start.offset)
        var end = this.getNode(lines[selection.end.line].children, selection.end.offset)

        var range = document.createRange()
        var sel = window.getSelection()

        range.setStart(start.node.firstChild, start.offset)
        range.setEnd(end.node.firstChild, end.offset)

        sel.removeAllRanges()
        sel.addRange(range)

        start.node.scrollIntoView({ block: "nearest", behavior: "smooth" });
    }

    sendSelection(l, t, lines) {
        var selection = window.getSelection();
        var range = selection.getRangeAt(0);

        var endContainer = range.endContainer.parentNode;
        var line = lines.indexOf(endContainer.parentNode)
        var token = Array.from(lines[line].children).indexOf(endContainer);

        console.log(`line is ${line} token is ${token}`)
        console.log(lines[line])

        var offset = this.getOffset(lines[line].children, lines[line].children[token], range.endOffset);

        console.log(selection)
        console.log(range)

        var offsets = {
            start: {
                line: line,
                offset: offset
            },
            end: {
                line: line,
                offset: offset
            }
        }

        switch (selection.type) {
            case "Caret":
            case "Range":
                var startContainer = range.startContainer.parentNode;
                var lineIndex = lines.indexOf(startContainer.parentNode)
                line = lines[lineIndex].children;
                token = Array.from(line).indexOf(startContainer);

                offsets.start = {
                    line: lineIndex,
                    offset: this.getOffset(
                        line,
                        line[token],
                        range.startOffset,
                    ),
                }

                break;
            case "None":
                break;
        }

        if (offsets.start.line === offsets.end.line && offsets.start.offset === offsets.end.offset) {
            if (offsets.start.line !== l || token !== t) {
                console.log(`rewriting ${l}-${t}`)
                offsets.start = {
                    line: l,
                    offset: this.getOffset(
                        l,
                        lines[l][t],
                        range.startOffset,
                    ),
                }
                offsets.end = offsets.start
            }
        }

        console.log(offsets)

        SetCursor({
            buffer: this.file,
            start: offsets.start,
            end: offsets.end,
        }).then((_, err) => {
            if (err) {
                console.log(err);
                return;
            }
        })
    }

    getOffset(line, node, baseOffset) {
        console.log(line)

        if (line.length === 1 && line[0].textContent === "\u200B") {
            return 0
        }


        var offset = baseOffset

        for (var i = 0; i < line.length; i++) {
            if (line[i].textContent === "\u200B") {
                continue
            }

            if (line[i].isEqualNode(node)) {
                break;
            }

            offset += line[i].textContent.length
        }

        console.log(`offset is ${offset}`)

        return offset
    }

    getNode(line, offset) {
        var i = 0;

        for (; i < line.length; i++) {
            if (line[i].textContent === "\u200B") {
                continue
            }

            if (line[i].textContent.length > offset) {
                break;
            }

            offset -= line[i].textContent.length
        }

        if (i === line.length) {
            i = line.length - 1
            offset = line[i].textContent.length
        }

        return {
            node: line[i],
            offset: offset,
        }
    }
}
