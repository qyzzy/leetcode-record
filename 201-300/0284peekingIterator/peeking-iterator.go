package peekingIterator

type Iterator struct{}

func (this *Iterator) hasNext() bool { return false }

func (this *Iterator) next() int { return 0 }

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter     *Iterator
	_hasNext bool
	_next    int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, iter.hasNext(), iter.next()}
}

func (it *PeekingIterator) hasNext() bool {
	return it._hasNext
}

func (it *PeekingIterator) next() int {
	ret := it._next
	it._hasNext = it.iter.hasNext()
	if it._hasNext {
		it._next = it.iter.next()
	}
	return ret
}

func (it *PeekingIterator) peek() int {
	return it._next
}
