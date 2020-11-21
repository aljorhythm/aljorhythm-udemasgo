package algos

import "testing"

type testcase struct {
	target   int
	expected int
}

func TestBinarySearchRecursive(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	cases := []testcase{
		testcase{0, -1},
		testcase{1, 0},
		testcase{2, 1},
		testcase{3, 2},
		testcase{4, 3},
		testcase{5, 4},
		testcase{6, -1},
	}

	for _, testcase := range cases {
		actual := BinarySearchRecursive(arr, testcase.target)
		if actual != testcase.expected {
			t.Errorf("Finding %d, Index should be %d, got %d instead", testcase.target, testcase.expected, actual)
		}
	}
}

func TestBinarySearchIterative(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	cases := []testcase{
		testcase{0, -1},
		testcase{1, 0},
		testcase{2, 1},
		testcase{3, 2},
		testcase{4, 3},
		testcase{5, 4},
		testcase{6, -1},
	}

	for _, testcase := range cases {
		actual := BinarySearchIterative(arr, testcase.target)
		if actual != testcase.expected {
			t.Errorf("Finding %d, Index should be %d, got %d instead", testcase.target, testcase.expected, actual)
		}
	}
}
