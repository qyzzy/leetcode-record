package heap

import "sort"

type hp struct {
	sort.IntSlice
}

func (h *hp) Pop() interface{} {
	val := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return val
}

func (h *hp) Push(val interface{}) {
	h.IntSlice = append(h.IntSlice, val.(int))
}
