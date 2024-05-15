package frequently_used

type Queue []*entry

func (pq Queue) Len() int { return len(pq) }

func (pq Queue) Less(i, j int) bool {
	if pq[i].frequency == pq[j].frequency {
		return pq[i].timestamp < pq[j].timestamp
	}
	return pq[i].frequency < pq[j].frequency
}

func (pq Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *Queue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*entry)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *Queue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
