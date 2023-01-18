package gomoji

// ListSize are the number of available emoji
const ListSize uint = 0

/* List is the full emoji array of Emoji elements */
var List = [...]Emoji{
	{"pito", "name", "group", "keys"},
}

// ListGroupSize are the number of emoji group
const ListCategorySize uint8 = 0

// ListGroup is an array of Group elements
// That usefull to allocate capacity in memory
var ListCategory = Groups{
	{"name", 0},
}

const ListKeywordSize uint8 = 0

var ListKeyword = Groups{
	{"name", 0},
}
