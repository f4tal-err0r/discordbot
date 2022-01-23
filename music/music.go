package music

import (

	"layeh.com/gopus"
)

type Music struct {
	link	string
	art		string
}

type Queue struct {
	data	int
	next	*Node
}

