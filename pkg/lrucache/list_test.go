package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPushFront(t *testing.T) {
	// when the list is empty, PushFront should set the root value as the passed in value
	val := 5
	t.Run("empty list", func(t *testing.T) {
		l := NewList()
		node := l.Unshift("key1", 5)

		require.Equal(t, val, node.Value.(int))
		require.Equal(t, val, l.Head().Value.(int))
	})

	// when the list is not empty, PushFront creates a new node and sets it as the head value
	t.Run("non-empty list", func(t *testing.T) {
		foobar := "foobar"
		l := NewList()

		l.Unshift("key1", 5)
		node := l.Unshift("key2", foobar)

		assert.Equal(t, foobar, node.Value.(string))
		assert.Equal(t, foobar, l.Head().Value.(string))
		assert.Equal(t, 5, l.Tail().Value.(int))
	})
}

func TestRemove(t *testing.T) {
	t.Run("remove node", func(t *testing.T) {
		l := NewList()
		node := l.Unshift("key2", "foo")

		l.Remove(node)

		assert.Equal(t, 0, l.Length())
		assert.Nil(t, l.Head().Value)
	})

	t.Run("remove tail", func(t *testing.T) {
		l := NewList()
		node1 := l.Unshift("key2", "foo")
		node2 := l.Unshift("key1", 5)

		assert.Equal(t, node1, l.Tail())
		assert.Equal(t, 2, l.Length())

		l.Remove(l.Tail())

		assert.Equal(t, node2, l.Tail())
		assert.Equal(t, 1, l.Length())
	})
}

func TestMoveFront(t *testing.T) {
	t.Run("move node to the front", func(t *testing.T) {
		l := NewList()
		node1 := l.Unshift("key1", 5)
		node2 := l.Unshift("key2", "foobar")

		assert.Equal(t, node2, l.Head())

		l.MoveFront(node1)

		assert.Equal(t, node1, l.Head())
	})
}
