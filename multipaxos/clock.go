package multipaxos

import "fmt"

type Clock struct {
	count int
	id    string
}

func NewClock(id string) *Clock {
	return &Clock{
		count: 1,
		id:    id,
	}
}

func (c *Clock) Print() {
	fmt.Printf("[%d, %s]", c.count, c.id)
}

func (c *Clock) Incerement() {
	c.count += 1
}

func (c1 *Clock) IsGreaterThan(c2 *Clock) bool {
	if c1.count > c2.count {
		return true
	} else if c1.count == c2.count {
		return c1.id > c2.id
	}
	return false
}

func (c1 *Clock) IsEqualTo(c2 *Clock) bool {
	return c1.count == c2.count && c1.id == c2.id
}
