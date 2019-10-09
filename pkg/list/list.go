package list

type List []Item

type Item struct {
	Text string
}

var MyTODOList = List{
	Item{"write a talk"},
	Item{"go to the meetup"},
	Item{"speak at the meetup"},
	Item{"drink beer"},
	Item{"sleep"},
}
