package problem1480

import (
    "reflect"
    "testing"
)

func TestRunningSumFunc(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want []int
    }{
        {
            name: "Example 1",
            nums: []int{1, 2, 3, 4},
            want: []int{1, 3, 6, 10},
        },
        {
            name: "Example 2",
            nums: []int{1, 1, 1, 1, 1},
            want: []int{1, 2, 3, 4, 5},
        },
        {
            name: "Example 3",
            nums: []int{3, 1, 2, 10, 1},
            want: []int{3, 4, 6, 16, 17},
        },
        {
            name: "Single element",
            nums: []int{5},
            want: []int{5},
        },
        {
            name: "Empty array",
            nums: []int{},
            want: []int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := RunningSum(tt.nums); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("RunningSum() = %v, want %v", got, tt.want)
            }
        })
    }
}
