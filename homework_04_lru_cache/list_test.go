package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.Remove(l.Front())

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		l.PushFront(500) // [500, 10, 30]
		require.Equal(t, 500, l.Front().Value)
		l.Remove(l.Front()) // [10, 30]
		require.Equal(t, 10, l.Front().Value)

		l.PushBack(500) // [10, 30, 500]
		require.Equal(t, 500, l.Back().Value)
		l.Remove(l.Back()) // [10, 30]
		require.Equal(t, 30, l.Back().Value)

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		require.Equal(t, 7, l.Len())
		l.MoveToFront(l.Back()) // [70, 80, 60, 40, 10, 30, 50]
		require.Equal(t, 7, l.Len())
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)

		l.MoveToFront(l.Front().Next)
		require.Equal(t, 80, l.Front().Value) // 80, 70, 60, 40, 10, 30, 50

		l.MoveToFront(l.Front().Next.Next)
		require.Equal(t, 60, l.Front().Value)                // 60, 80, 70, 40, 10, 30, 50
		require.Equal(t, 40, l.Front().Next.Next.Next.Value) // 60, 80, 70, 40, 10, 30, 50

		l.MoveToFront(l.Back().Prev.Prev)
		require.Equal(t, 10, l.Front().Value)      // 10, 60, 80, 70, 40, 30, 50
		require.Equal(t, 60, l.Front().Next.Value) // 10, 60, 80, 70, 40, 30, 50
		l.MoveToFront(l.Back())
		require.Equal(t, 30, l.Back().Value) // 50, 10, 60, 80, 70, 40, 30
	})
}
