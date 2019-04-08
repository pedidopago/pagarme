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
}

// QueryTime is used on datetime queries
type QueryTime struct {
	Name string
	Op   QueryOp
	T    time.Time
}

// Format to pagarme query
func (qt *QueryTime) Format() string {
	return fmt.Sprintf("%v%v%v", url.PathEscape(qt.Name), string(qt.Op), url.PathEscape(qt.T.Format(time.RFC3339Nano)))
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
