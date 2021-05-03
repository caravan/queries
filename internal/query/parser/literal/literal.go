package literal

// String parses a string literal
var String = WS(quotedParser("'"))
