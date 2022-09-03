export namespace types {
	
	export class Token {
	    offset: number;
	    value: string;
	    classes: string[];
	
	    static createFrom(source: any = {}) {
	        return new Token(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.offset = source["offset"];
	        this.value = source["value"];
	        this.classes = source["classes"];
	    }
	}
	export class BufferData {
	    buffer: string;
	    symbol: number;
	    lines: Token[][];
	    timestamp: number;
	
	    static createFrom(source: any = {}) {
	        return new BufferData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.buffer = source["buffer"];
	        this.symbol = source["symbol"];
	        this.lines = this.convertValues(source["lines"], Token);
	        this.timestamp = source["timestamp"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetCaretResponse {
	    buffer: string;
	    start: selection.Caret;
	    end: selection.Caret;
	
	    static createFrom(source: any = {}) {
	        return new GetCaretResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.buffer = source["buffer"];
	        this.start = this.convertValues(source["start"], selection.Caret);
	        this.end = this.convertValues(source["end"], selection.Caret);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetLineLengthRequest {
	    buffer: string;
	    line: number;
	
	    static createFrom(source: any = {}) {
	        return new GetLineLengthRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.buffer = source["buffer"];
	        this.line = source["line"];
	    }
	}

}

export namespace selection {
	
	export class Caret {
	    line: number;
	    offset: number;
	
	    static createFrom(source: any = {}) {
	        return new Caret(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.line = source["line"];
	        this.offset = source["offset"];
	    }
	}

}

