package intsx

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var jsonArrayReg = regexp.MustCompile(`\[\s*(\d+),\s*(\d+)\s*\]`)

// Range defines a range with integer
type Range struct {
	Start int
	End   int
}

func (r *Range) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[%d,%d]", r.Start, r.End)), nil
}

func (r *Range) UnmarshalJSON(data []byte) error {
	if r == nil {
		return errors.New("RawString: UnmarshalJSON on nil pointer")
	}

	var err error
	jsonDataStr := string(data)
	matchData := jsonArrayReg.FindAllStringSubmatch(jsonDataStr, -1)

	r.Start, err = strconv.Atoi(matchData[0][1])
	if err != nil {
		return err
	}

	r.End, err = strconv.Atoi(matchData[0][2])
	if err != nil {
		return err
	}

	return nil
}

func (r *Range) Include(val int) bool {
	return val >= r.Start && val <= r.End
}

// IncludeRange 返回当前 Range 是否包含 传入 Range
func (r *Range) IncludeRange(rg *Range) bool {
	return r.Start <= rg.Start && r.End >= rg.End
}

func (r *Range) Iter(fn func(idx int, v int)) {
	idx := 0
	i := r.Start

	for i <= r.End {
		fn(idx, i)
		i++
		idx++
	}
}
