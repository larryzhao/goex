package goex

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSortedList(t *testing.T) {
	list := NewSortedList()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.len)
	assert.Nil(t, list.head)
}

func TestSortedListAdd(t *testing.T) {
	var list *SortedList
	// test first item add
	list = NewSortedList()
	list.Add(1000, "a")

	assert.Equal(t, 1, list.len)
	assert.NotNil(t, list.head)
	assert.Exactly(t, extractListToSlice(list), []string{"a"})

	// test if the item should be the new head
	list.Add(900, "b")
	assert.Equal(t, 2, list.Len())
	assert.Exactly(t, extractListToSlice(list), []string{"b", "a"})

	// test if the item should be at the tail
	list.Add(1100, "c")
	assert.Equal(t, 3, list.Len())
	assert.Exactly(t, extractListToSlice(list), []string{"b", "a", "c"})
	assert.Equal(t, "c", list.tail.Val.(string))

	// test if the item should be in the middle
	list.Add(1020, "d")
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, "d", list.head.Next.Next.Val.(string))
	assert.Exactly(t, extractListToSlice(list), []string{"b", "a", "d", "c"})
}

func TestPopScoreLowerThan(t *testing.T) {
	list := NewSortedList()
	list.Add(1000, "a")
	list.Add(2000, "b")
	list.Add(3000, "c")
	list.Add(4000, "d")
	list.Add(5000, "e")
	list.Add(6000, "f")
	list.Add(7000, "g")
	list.Add(8000, "h")
	list.Add(9000, "i")

	assert.Equal(t, 9, list.Len())
	assert.Exactly(t, extractListToSlice(list), []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})

	items := list.PopScoreLowerThan("3000", 2)
	assert.Equal(t, 2, len(items))
	assert.Equal(t, "a", items[0].(string))
	assert.Equal(t, "b", items[1].(string))
	assert.Equal(t, 7, list.Len())
	assert.Equal(t, "c", list.head.Val.(string))

	items = list.PopScoreLowerThan("3000", 10000)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, "c", items[0].(string))
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "d", list.head.Val.(string))

	items = list.PopScoreLowerThan("1000", 10000)
	assert.Equal(t, 0, len(items))
	assert.Equal(t, 6, list.Len())
	assert.Equal(t, "d", list.head.Val.(string))

	items = list.PopScoreLowerThan("(6000", 10000)
	assert.Equal(t, 2, len(items))
	assert.Equal(t, "d", items[0].(string))
	assert.Equal(t, "e", items[1].(string))
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, "f", list.head.Val.(string))

	items = list.PopScoreLowerThan("10000", 1000)
	assert.Equal(t, 4, len(items))
	assert.Equal(t, "f", items[0].(string))
	assert.Equal(t, "g", items[1].(string))
	assert.Equal(t, "h", items[2].(string))
	assert.Equal(t, "i", items[3].(string))
}

func BenchmarkSortedListSequenceAdd(b *testing.B) {
	list := NewSortedList()
	for i := 0; i < b.N; i++ {
		list.Add(int64(i), "val")
	}
}

func BenchmarkSortedListRandomAdd(b *testing.B) {
	count := b.N
	numbers := make([]int64, count)

	rs := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rs)
	for i := 0; i < count; i++ {
		numbers[i] = r.Int63()
	}

	list := NewSortedList()
	for i := 0; i < count; i++ {
		list.Add(numbers[i], "val")
	}
}

type concurrentAddJob struct {
	Score int64
	Val   interface{}
}

func concurrentAdd(list *SortedList, ch chan *concurrentAddJob, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range ch {
		list.Add(job.Score, job.Val)
	}
}

func BenchmarkSortedListConcurrentAdd(b *testing.B) {
	count := b.N
	numbers := make([]int64, count)

	rs := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rs)
	for i := 0; i < count; i++ {
		numbers[i] = r.Int63()
	}

	concurrency := 20
	list := NewSortedList()
	ch := make(chan *concurrentAddJob, 20)
	wg := &sync.WaitGroup{}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go concurrentAdd(list, ch, wg)
	}

	for i := 0; i < count; i++ {
		ch <- &concurrentAddJob{
			Score: numbers[i],
			Val:   "val",
		}
	}

	close(ch)
	wg.Wait()
}

func extractListToSlice(list *SortedList) []string {
	results := make([]string, list.Len())
	idx := 0
	current := list.head
	for current != nil {
		results[idx] = current.Val.(string)

		idx++
		current = current.Next
	}
	return results
}
