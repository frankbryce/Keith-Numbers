// Package collection provides data collection primitives for mining information
// about keith numbers, to hopefully find some patterns in the data
package collection

import (
	"math/big"
)

type Collection interface {
	Add(i *big.Int)
	Get() interface{}
	Reset()
	Set(target *big.Int)
}

type base struct {
	target, last *big.Int
}

func baseAdd(c *base, i *big.Int) {
	c.last = big.NewInt(0).Set(i)
}

func baseGet(c *base) {
	// no op
}

func baseReset(c *base) {
	c.last = nil
	c.target = nil
}

func (c *base) Set(target *big.Int) {
	c.target = big.NewInt(0).Set(target)
}

var Collections []Collection

func init() {
	Collections = make([]Collection, 0, 4)
	Collections = append(Collections, &last{})
	Collections = append(Collections, &underflow{})
	Collections = append(Collections, &overflow{})
	Collections = append(Collections, &count{count: 0})
}

// last collection
type last struct {
	base
}

func (c *last) Add(i *big.Int) {
	baseAdd(&c.base, i)
}

func (c *last) Get() interface{} {
	baseGet(&c.base)
	return c.last
}

func (c *last) Reset() {
	baseReset(&c.base)
}

// underflow collection
type underflow struct {
	base
	last2 *big.Int
}

func (c *underflow) Add(i *big.Int) {
    if c.last != nil {
    	c.last2 = big.NewInt(0).Set(c.last)
    }
	baseAdd(&c.base, i)
}

func (c *underflow) Get() interface{} {
	baseGet(&c.base)
	return big.NewInt(0).Sub(c.target, c.last2)
}

func (c *underflow) Reset() {
	baseReset(&c.base)
	c.last2 = nil
}

// overflow collection
type overflow struct {
	base
}

func (c *overflow) Add(i *big.Int) {
	baseAdd(&c.base, i)
}

func (c *overflow) Get() interface{} {
	baseGet(&c.base)
	return big.NewInt(0).Sub(c.last, c.target)
}

func (c *overflow) Reset() {
	baseReset(&c.base)
}

// count collection
type count struct {
	base
	count uint64
}

func (c *count) Add(i *big.Int) {
	baseAdd(&c.base, i)
	c.count++
}

func (c *count) Get() interface{} {
	baseGet(&c.base)
	return c.count
}

func (c *count) Reset() {
	baseReset(&c.base)
	c.count = 0
}
