package skip_list

type SkipNode struct {
	points []*SkipNode
	parent *SkipNode
	key    string
	val    int
}

type SkipList struct {
	keyword string
	front   *SkipNode
	tail    *SkipNode
}

func InitSkipList(keyword string) *SkipList {
	return &SkipList{
		keyword: keyword,
		front:   nil,
		tail:    nil,
	}
}

func (list *SkipList) insert() {

}

func (list *SkipList) find() {

}

func (list *SkipList) delete() {}
