# Go Backend - Array Utility

This is the Go (Golang) implementation of the ArrayUtility functionality, migrated from the Java backend.

## Features

### ArrayUtility Struct
- **FindArraySize**: Finds the size of any slice (Go's equivalent of arrays)
- **FindSumOfThree**: Calculates the sum of three integers
- **ConcatStrings**: Concatenates multiple strings with a custom separator

### Package-Level Functions (Static Equivalent)
- **FindArraySizeStatic**: Package-level function for array size
- **FindSumOfThreeStatic**: Package-level function for sum calculation
- **ConcatStringsStatic**: Package-level function for string concatenation

## Key Differences from Java

1. **No Static Methods**: Go doesn't have static methods, so we use package-level functions
2. **Interface{} Type**: Uses `interface{}` (similar to Java's Object) for generic slice handling
3. **Reflection**: Uses Go's reflection package to determine slice types and lengths
4. **Slices vs Arrays**: Go uses slices (dynamic arrays) instead of fixed arrays
5. **String Operations**: Uses Go's `strings.Join` for efficient string concatenation

## How to Run

### Prerequisites
- Go 1.21 or later installed

### Run the Program
```bash
cd backend-go
go run src/array_utility.go
```

### Build and Run
```bash
cd backend-go
go build -o array_utility src/array_utility.go
./array_utility
```

## Expected Output
```
Integer array size: 5
String array size: 3
Null array size: 0
Sum of 10, 20, 30: 60
Sum of 5, 15, 25: 45
Concatenated with space: 'Hello World Go Programming'
Concatenated with dash: 'Hello-World-Go-Programming'
Concatenated with comma: 'Hello, World, Go, Programming'
Concatenated with empty separator: 'Hello World Go Programming'
Static - Integer array size: 5
Static - Sum of 7, 14, 21: 42
Static - Concatenated: 'Hello | World | Go | Programming'
```

## Code Structure

```go
type ArrayUtility struct{}

// Instance methods
func (au *ArrayUtility) FindArraySize(arr interface{}) int
func (au *ArrayUtility) FindSumOfThree(a, b, c int) int
func (au *ArrayUtility) ConcatStrings(strSlice []string, separator string) string

// Package-level functions (static equivalent)
func FindArraySizeStatic(arr interface{}) int
func FindSumOfThreeStatic(a, b, c int) int
func ConcatStringsStatic(strSlice []string, separator string) string
```

## String Concatenation Function

The `ConcatStrings` function provides flexible string concatenation:

- **Parameters**: 
  - `strSlice []string`: Slice of strings to concatenate
  - `separator string`: Separator between strings (defaults to space if empty)
- **Returns**: Concatenated string with separators
- **Features**:
  - Handles empty slices gracefully
  - Customizable separators
  - Uses Go's efficient `strings.Join` function

## Testing

The main function includes comprehensive tests for all functionality:
- Array size calculation for different types
- Sum calculation with various numbers
- String concatenation with different separators
- Null/empty slice handling
- Both instance methods and package-level functions
