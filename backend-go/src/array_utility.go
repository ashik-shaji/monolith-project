package main

import (
	"fmt"
	"reflect"
	"strings"
)

// ArrayUtility provides utility functions for array operations
type ArrayUtility struct{}

// FindArraySize finds the size of an array (slice in Go)
// This function works with any type of slice
func (au *ArrayUtility) FindArraySize(arr interface{}) int {
	if arr == nil {
		return 0
	}

	// Use reflection to get the length of the slice
	val := reflect.ValueOf(arr)
	if val.Kind() == reflect.Slice {
		return val.Len()
	}

	return 0
}

// FindSumOfThree finds the sum of three numbers
func (au *ArrayUtility) FindSumOfThree(a, b, c int) int {
	return a + b + c
}

// ConcatStrings concatenates multiple strings with a separator
func (au *ArrayUtility) ConcatStrings(strSlice []string, separator string) string {
	if len(strSlice) == 0 {
		return ""
	}

	if separator == "" {
		separator = " "
	}

	return strings.Join(strSlice, separator)
}

// Static functions (Go doesn't have static methods, but we can use package-level functions)
// FindArraySizeStatic is a package-level function equivalent to Java's static method
func FindArraySizeStatic(arr interface{}) int {
	if arr == nil {
		return 0
	}

	val := reflect.ValueOf(arr)
	if val.Kind() == reflect.Slice {
		return val.Len()
	}

	return 0
}

// FindSumOfThreeStatic is a package-level function equivalent to Java's static method
func FindSumOfThreeStatic(a, b, c int) int {
	return a + b + c
}

// ConcatStringsStatic is a package-level function equivalent to Java's static method
func ConcatStringsStatic(strSlice []string, separator string) string {
	if len(strSlice) == 0 {
		return ""
	}

	if separator == "" {
		separator = " "
	}

	return strings.Join(strSlice, separator)
}

func main() {
	// Create an instance of ArrayUtility
	au := &ArrayUtility{}

	// Test with different array types
	intArray := []int{1, 2, 3, 4, 5}
	stringArray := []string{"Hello", "World", "Go"}

	fmt.Printf("Integer array size: %d\n", au.FindArraySize(intArray))
	fmt.Printf("String array size: %d\n", au.FindArraySize(stringArray))
	fmt.Printf("Null array size: %d\n", au.FindArraySize(nil))

	// Test the sum function
	fmt.Printf("Sum of 10, 20, 30: %d\n", au.FindSumOfThree(10, 20, 30))
	fmt.Printf("Sum of 5, 15, 25: %d\n", au.FindSumOfThree(5, 15, 25))

	// Test the string concatenation function
	testStrings := []string{"Hello", "World", "Go", "Programming"}
	fmt.Printf("Concatenated with space: '%s'\n", au.ConcatStrings(testStrings, " "))
	fmt.Printf("Concatenated with dash: '%s'\n", au.ConcatStrings(testStrings, "-"))
	fmt.Printf("Concatenated with comma: '%s'\n", au.ConcatStrings(testStrings, ", "))
	fmt.Printf("Concatenated with empty separator: '%s'\n", au.ConcatStrings(testStrings, ""))

	// Test static functions
	fmt.Printf("Static - Integer array size: %d\n", FindArraySizeStatic(intArray))
	fmt.Printf("Static - Sum of 7, 14, 21: %d\n", FindSumOfThreeStatic(7, 14, 21))
	fmt.Printf("Static - Concatenated: '%s'\n", ConcatStringsStatic(testStrings, " | "))
}
