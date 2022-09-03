import { SetCursor, MouseDown, MouseUp } from "../../wailsjs/go/api/Api.js";

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
