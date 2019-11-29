package pagarme

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"
)

// QueryOp is the operator of a query item
type QueryOp string

const (
	// QueryOpEquals =
	QueryOpEquals QueryOp = "="
	// QueryOpGreaterThan >
	QueryOpGreaterThan QueryOp = "=%3E"
	// QueryOpLessThan <
	QueryOpLessThan QueryOp = "=%3C"
	// QueryOpGreaterOrEqualThan >=
	QueryOpGreaterOrEqualThan QueryOp = "=%3E%3D"
	// QueryOpLessOrEqualThan <=
	QueryOpLessOrEqualThan QueryOp = "=%3C%3D"
)

type Querier interface {
	Format() string
	PName() string
}

// QueryTime is used on datetime queries
type QueryTime struct {
	Name string
	Op   QueryOp
	T    time.Time
}

// Format to pagarme query
func (qt *QueryTime) Format() string {
	return fmt.Sprintf("%v%v%v", url.PathEscape(qt.Name), string(qt.Op), qt.T.UnixNano()/1000000)
}

// PName of the query param
func (qt *QueryTime) PName() string {
	return qt.Name
}

// QueryInt is used on integer queries
type QueryInt struct {
	Name string
	Op   QueryOp
	V    int
}

// Format to pagarme query
func (qt *QueryInt) Format() string {
	return fmt.Sprintf("%v%v%v", url.PathEscape(qt.Name), string(qt.Op), qt.V)
}

// PName of the query param
func (qt *QueryInt) PName() string {
	return qt.Name
}

// QueryString is used on string queries
type QueryString struct {
	Name string
	Op   QueryOp
	V    string
}

// Format to pagarme query
func (qt *QueryString) Format() string {
	return fmt.Sprintf("%v%v%v", url.PathEscape(qt.Name), string(qt.Op), qt.V)
}

// PName of the query param
func (qt *QueryString) PName() string {
	return qt.Name
}

type QueryBuilder struct {
	items []Querier
	lock  sync.Mutex
}

// Add a querier item
func (b *QueryBuilder) Add(q Querier) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.items == nil {
		b.items = make([]Querier, 0)
	}
	b.items = append(b.items, q)
}

// Set a querier item
func (b *QueryBuilder) Set(q Querier) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.items == nil {
		b.items = make([]Querier, 0)
	}
	z := -1
	for i := range b.items {
		if b.items[i].PName() == q.PName() {
			z = i
			break
		}
	}
	if z == -1 {
		b.items = append(b.items, q)
		return
	}
	b.items[z] = q
}

// Get a querier item
func (b *QueryBuilder) Get(name string) Querier {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.items == nil {
		return nil
	}
	for i := range b.items {
		if b.items[i].PName() == name {
			return b.items[i]
		}
	}
	return nil
}

// Build the query
func (b *QueryBuilder) Build() string {
	b.lock.Lock()
	defer b.lock.Unlock()
	sb := new(strings.Builder)
	for k, v := range b.items {
		if k != 0 {
			sb.WriteRune('&')
		}
		sb.WriteString(v.Format())
	}
	return sb.String()
}
