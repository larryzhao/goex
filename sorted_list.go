package goex

import (
	"math"
	"strconv"
	"sync"
)

var maxScore int64 = math.MaxInt64
var minScore int64 = math.MinInt64

type sortedListItem struct {
	Val   interface{}
	Score int64
	Next  *sortedListItem
	Prev  *sortedListItem
}

type SortedList struct {
	head *sortedListItem
	tail *sortedListItem
	mu   *sync.Mutex
	len  int
}

// NewSortedList create a sorted list
func NewSortedList() *SortedList {
	return &SortedList{
		head: nil,
		tail: nil,
		mu:   &sync.Mutex{},
		len:  0,
	}
}

// Len length of the current list
func (list *SortedList) Len() int {
	return list.len
}

func (list *SortedList) Add(score int64, val interface{}) {
	defer list.mu.Unlock()
	list.mu.Lock()

	newItem := &sortedListItem{
		Val:   val,
		Score: score,
		Next:  nil,
		Prev:  nil,
	}

	// when list is empty, the item is the first
	if list.Empty() {
		list.head = newItem
		list.tail = newItem
		list.len++
		return
	}

	// when it should be the new head
	if score <= list.head.Score {
		newItem.Next = list.head
		list.head.Prev = newItem

		list.head = newItem
		list.len++
		return
	}

	prev := list.head
	current := list.head.Next

	// when there's only the head, and the item should go after it
	if current == nil {
		prev.Next = newItem
		newItem.Prev = prev
		list.tail = newItem
		list.len++
		return
	}

	for {
		// found the place to be
		if score <= current.Score {
			prev.Next = newItem
			newItem.Prev = prev

			newItem.Next = current
			current.Prev = newItem

			list.len++
			return
		}

		// when we meets the tail
		if current.Next == nil {
			current.Next = newItem
			newItem.Prev = current
			list.tail = newItem
			list.len++
			return
		}

		// move next
		prev = current
		current = prev.Next
	}
}

func (list *SortedList) Empty() bool {
	return list.head == nil
}

func (list *SortedList) PopScoreLowerThan(limit string, count uint32) []interface{} {
	var current *sortedListItem
	var idx uint32

	up := upperComparer(limit)
	var results []interface{}

	list.mu.Lock()
	defer func() {
		if list.head != current {
			list.head = current
			list.len = list.len - int(idx)
		}
		list.mu.Unlock()
	}()

	if count == 0 || limit == "-inf" || list.Empty() {
		return []interface{}{}
	}

	current = list.head
	for idx < count {
		if current == nil || !up(current.Score) {
			return results
		}

		results = append(results, current.Val)
		current = current.Next
		idx++
	}

	return results
}

func lowComparer(limit string) func(score int64) bool {
	if limit == "-inf" {
		return func(score int64) bool {
			return score >= minScore
		}
	}

	if limit == "+inf" {
		return func(score int64) bool {
			return score >= maxScore
		}
	}

	rr := []rune(limit)

	if rr[0] == '(' {
		// TODO: handle error
		limitScore, _ := strconv.ParseInt(string(rr[1:]), 10, 64)
		return func(score int64) bool {
			return score > limitScore
		}
	}

	limitScore, _ := strconv.ParseInt(limit, 10, 64)
	return func(score int64) bool {
		return score >= limitScore
	}
}

func upperComparer(limit string) func(score int64) bool {
	if limit == "-inf" {
		return func(score int64) bool {
			return score <= minScore
		}
	}

	if limit == "+inf" {
		return func(score int64) bool {
			return score <= maxScore
		}
	}

	rr := []rune(limit)

	if rr[0] == '(' {
		// TODO: handle error
		limitScore, _ := strconv.ParseInt(string(rr[1:]), 10, 64)
		return func(score int64) bool {
			return score < limitScore
		}
	}

	limitScore, _ := strconv.ParseInt(limit, 10, 64)
	return func(score int64) bool {
		return score <= limitScore
	}
}
