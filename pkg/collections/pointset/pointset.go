// Package PointSet is a naive but simple map to store small point values
// on a grid, represented as two uint32 values. We are aiming for speed
package pointset

import (
	"fmt"
	"strings"
)

// this will represent our combined int using x and y
// we shift the x value and combine it with y to make
// a unique value that is fast to calculate
type Point = uint32

type PointSet struct {
	len        int
	buckets    [8192][100]Point
	bucketLens [8192]int
}

func pack(x, y uint16) Point {
	return (uint32(x)<<16 | uint32(y))
}

func hash(x, y uint16) uint16 {
	return (uint16(x)<<8 | y)
}

func unpack(point Point) (x, y uint16) {
	x = uint16(point >> 16)
	y = uint16(point)
	return
}

func New(buckets int) *PointSet {
	return &PointSet{}
}

func (ps *PointSet) Insert(x uint16, y uint16) {
	point := pack(x, y)
	hash := hash(x, y)
	bucketIdx := int(hash) % len(ps.buckets)
	// check if point already exists so we don't duplicates
	for i := 0; i < ps.bucketLens[bucketIdx]; i++ {
		if ps.buckets[bucketIdx][i] == point {
			return
		}
	}
	ps.buckets[bucketIdx][ps.bucketLens[bucketIdx]] = point
	ps.len++
	ps.bucketLens[bucketIdx]++
}

func (ps *PointSet) Contains(x uint16, y uint16) bool {
	point := pack(x, y)
	hash := hash(x, y)
	bucketIdx := int(hash) % len(ps.buckets)
	for i := 0; i < ps.bucketLens[bucketIdx]; i++ {
		if ps.buckets[bucketIdx][i] == point {
			return true
		}
	}
	return false
}

func (ps *PointSet) Length() int {
	return ps.len
}

func (ps *PointSet) PrintBucketHistogram() {
	for _, b := range ps.buckets {
		fmt.Println(strings.Repeat("#", len(b)))
	}
}
