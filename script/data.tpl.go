package gomoji

// ListSize are the number of available emoji
const ListSize uint = $EMOJIS_NB

/* List is the full emoji array of Emoji elements */
var List = [...]Emoji{
	%EMOJIS
}

// ListGroupSize are the number of emoji group
const ListCategorySize uint = $CAT_NB

// ListGroup is an array of Group elements
// That usefull to allocate capacity in memory
var ListCategory = Groups{
	$CAT
}

const ListKeywordSize uint = $KEY_NB

var ListKeyword = Groups{
	$KEY
}
