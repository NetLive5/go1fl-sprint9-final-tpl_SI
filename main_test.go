package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantNil bool
	}{
		{name: "negative size", size: -1, wantNil: true},
		{name: "zero size", size: 0, wantNil: true},
		{name: "one element", size: 1},
		{name: "multiple elements", size: 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := generateRandomElements(test.size)
			if test.wantNil {
				require.Nil(t, got)
				return
			}

			require.Len(t, got, test.size)
			for i, value := range got {
				assert.GreaterOrEqual(t, value, 0, "generated value at index %d should be non-negative", i)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: nil, want: 0},
		{name: "single element", data: []int{7}, want: 7},
		{name: "maximum at beginning", data: []int{9, 2, 3}, want: 9},
		{name: "maximum in middle", data: []int{1, 9, 3}, want: 9},
		{name: "maximum at end", data: []int{1, 2, 9}, want: 9},
		{name: "negative values", data: []int{-8, -2, -5}, want: -2},
		{name: "duplicate maximum", data: []int{4, 9, 9, 2}, want: 9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, maximum(test.data))
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: nil, want: 0},
		{name: "single element", data: []int{7}, want: 7},
		{name: "fewer elements than chunks", data: []int{2, 8, 3}, want: 8},
		{name: "evenly divided", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "remainder and maximum in final part", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 25, 10}, want: 25},
		{name: "negative values", data: []int{-8, -2, -5}, want: -2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, maxChunks(test.data))
		})
	}
}
