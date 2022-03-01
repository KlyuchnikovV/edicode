import Selection from "./selection";

export default class Cursor {
    constructor(file, container) {
        this.file = file
        // this.line = 1; // based with 1
        // this.offset = 0;
        this.container = container
        // this.lines = []

        this.selection = new Selection();
    }

    onPress(event) {
        switch (event.key) {
            case "ArrowUp":
                this.selection.up(event.shiftKey);
                break;
            case "ArrowDown":
                this.selection.down(event.shiftKey, false);
                break;
            case "ArrowLeft":
                this.selection.left(event.shiftKey);
                break;
            case "ArrowRight":
                this.selection.right(event.shiftKey, false);
                break;
        }
    }

    onMouse(event, lines) {
        var selection = window.getSelection();
        var range = selection.getRangeAt(0);
        var startLine = 0, startNode = 0, endLine = 0, endNode = 0;
        var startOffset = 0, endOffset = 0;
        console.log(range)
        for (var i = 0; i < lines.length; i++) {
            var len = 0
            for (var j = 0; j < lines[i].children.length; j++) {
                if (range.startContainer.parentNode.isEqualNode(lines[i].children[j])) {
                    startLine = i+1;
                    startNode = j;
                    startOffset = len + range.startOffset;
                }
                if (range.endContainer.parentNode.isEqualNode(lines[i].children[j])) {
                    endLine = i+1;
                    endNode = j;
                    endOffset = len + range.endOffset;
                }
                if (startLine !== 0 && endLine !== 0) {
                    break;
                }
                len += lines[i].children[j].textContent.length
            }
        }

        this.selection.start = {
            line: startLine,
            offset: startOffset,
        }
        this.selection.end = {
            line: endLine,
            offset: endOffset,
        }

        console.log(this.selection)
    }

    // TODO: selection left-right expansion
    // TODO: select empty line

    setCurrentCursorPosition(elements) {
        var startElement = elements[this.selection.start.line - 1]
        var endElement = elements[this.selection.end.line - 1]
        var selection = window.getSelection();

        console.log(`start ${this.selection.start.line - 1}, end ${this.selection.end.line - 1}`)
        console.log(startElement)

        let range = this.rangeTill(startElement, this.selection.start.offset, endElement, this.selection.end.offset)

        selection.removeAllRanges();
        selection.addRange(range);

        if (this.selection.goingDown === null) {
            return
        }

        if (this.selection.goingDown) {
            endElement.scrollIntoView({ block: "nearest", behavior: "smooth" });
        } else {
            startElement.scrollIntoView({ block: "nearest", behavior: "smooth" });
        }
    }

    getCurrentCursorPosition() {
        var selection = window.getSelection(),
            charCount = 0,
            node;

        if (this._isChildOf(selection.focusNode, this.container)) {
            node = selection.focusNode;
            charCount = selection.focusOffset;

            while (node !== null) {
                if (node === this.container) {
                    break;
                }

                if (node !== undefined && node.classList !== undefined && node.classList.contains("line")) {
                    charCount++
                }

                if (node.previousSibling) {
                    node = node.previousSibling;
                    charCount += node.textContent.length;
                } else {
                    node = node.parentNode;
                }
            }
        }

        return charCount - 1;
    }

    _isChildOf(node, parentElement) {
        while (node !== null) {
            if (node === parentElement) {
                return true;
            }
            node = node.parentNode;
        }

        return false;
    }

    rangeTill(start, startOffset, end, endOffset) {
        let range = document.createRange()

        // console.log("1")
        // console.log(`node is '${start.children[0].textContent}'`)

        // console.log("2")
        var node = start.children[0]
        for (var i = 1; i < start.children.length && startOffset > node.textContent.length; i++) {
            // console.log(`node is '${node.textContent}'`)
            startOffset -= node.textContent.length
            node = start.children[i]
        }

        if (startOffset > node.textContent.length) {
            startOffset = node.textContent.length
        }
        if (node.firstChild !== null) {
            node = node.firstChild
        }

        // console.log("3")
        // console.log(node)
        // console.log(`offset ${startOffset}`)
        range.setStart(node, startOffset);

        // console.log("4")

        // console.log("5")
        var node = end.children[0]
        for (var i = 1; i < end.children.length && endOffset > node.textContent.length; i++) {
            // console.log(node)
            endOffset -= node.textContent.length
            node = end.children[i]
        }

        if (endOffset > node.textContent.length) {
            endOffset = node.textContent.length
        }
        if (node.firstChild !== null) {
            node = node.firstChild
        }

        // console.log("6")
        // console.log(node)
        // console.log(`offset ${endOffset}`)
        range.setEnd(node, endOffset);
        // console.log("7")
        return range;
    }
}
