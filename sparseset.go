package sparseset

import (
	"iter"
)

type denseItem[T any] struct {
	sparseIndex int
	item        T
}

type SparseSet[T any] struct {
	sparse []int
	dense  []denseItem[T]
	max    int
}

func New[T any](max int) *SparseSet[T] {
	if max < 0 {
		panic("sparse set max size cannot be less than 0")
	}

	return &SparseSet[T]{
		sparse: make([]int, max),
		dense:  make([]denseItem[T], 0, max),
		max:    max,
	}
}

func (s *SparseSet[T]) Add(idx int, item T) bool {
	if idx < 0 || idx >= s.max {
		return true
	}

	if s.Has(idx) {
		denseIdx := s.sparse[idx]
		s.dense[denseIdx].item = item

		return true
	}

	s.dense = append(s.dense, denseItem[T]{
		idx,
		item,
	})

	s.sparse[idx] = len(s.dense) - 1

	return true
}

func (s *SparseSet[T]) Has(idx int) bool {
	if idx < 0 || idx >= s.max {
		return false
	}

	denseIdx := s.sparse[idx]

	return denseIdx < len(s.dense) && idx == s.dense[denseIdx].sparseIndex
}

func (s *SparseSet[T]) Get(idx int) *T {
	if !s.Has(idx) {
		return nil
	}

	denseIdx := s.sparse[idx]

	return &s.dense[denseIdx].item
}

func (s *SparseSet[T]) Remove(idx int) {
	if !s.Has(idx) {
		return
	}

	denseIdx := s.sparse[idx]
	lastDenseIdx := len(s.dense) - 1

	s.dense[denseIdx] = s.dense[lastDenseIdx]

	movedSparseIdx := s.dense[denseIdx].sparseIndex

	s.sparse[movedSparseIdx] = denseIdx

	var zeroedDenseItem denseItem[T]
	s.dense[lastDenseIdx] = zeroedDenseItem
	s.dense = s.dense[:lastDenseIdx]
}

func (s *SparseSet[T]) Len() int {
	return len(s.dense)
}

func (s *SparseSet[T]) Capacity() int {
	return s.max
}

func (s *SparseSet[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for _, dItem := range s.dense {
			if !yield(dItem.sparseIndex, dItem.item) {
				return
			}
		}
	}
}
