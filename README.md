# Monolith Project

A simple monolith project demonstrating array utility functions implemented in Java (backend), TypeScript (frontend), and Go (backend-go).

## Project Structure

```
monolith-project/
├── backend/                 # Java backend
│   └── src/main/java/
│       └── com/example/monolith/
│           └── ArrayUtility.java
├── backend-go/             # Go backend
│   ├── src/
│   │   └── array_utility.go
│   ├── go.mod
│   └── README.md
├── frontend/               # TypeScript frontend
│   ├── ArrayUtility.ts
│   ├── package.json
│   └── tsconfig.json
└── README.md
```

## Features

### Backend (Java)
- `ArrayUtility` class with `findArraySize` method
- `findSumOfThree` method for calculating sum of three numbers
- Handles null arrays gracefully
- Simple and clean implementation

### Backend (Go)
- `ArrayUtility` struct with `FindArraySize` method
- `FindSumOfThree` method for calculating sum of three numbers
- Package-level functions equivalent to Java static methods
- Uses Go's reflection for generic slice handling
- Handles nil slices gracefully

### Frontend (TypeScript)
- `ArrayUtility` class with generic type support
- Same functionality as Java and Go versions
- Full TypeScript type safety
- Handles null and undefined values

## Getting Started

### Backend (Java)
```bash
cd backend/src/main/java
javac com/example/monolith/ArrayUtility.java
java com.example.monolith.ArrayUtility
```

### Backend (Go)
```bash
cd backend-go
go run src/array_utility.go
```

### Frontend (TypeScript)
```bash
cd frontend
npm install
npm run dev
```

## Array Utility Functions

All three implementations provide the same core functionality:

```java
// Java
public static int findArraySize(Object[] array)
public static int findSumOfThree(int a, int b, int c)

// Go
func (au *ArrayUtility) FindArraySize(arr interface{}) int
func (au *ArrayUtility) FindSumOfThree(a, b, c int) int

// TypeScript
public static findArraySize<T>(array: T[] | null | undefined): number
public static findSumOfThree(a: number, b: number, c: number): number
```

The functions return the size of an array/slice, or 0 if null/nil/undefined, and calculate the sum of three numbers.

## Requirements

- Java 8+ (for Java backend)
- Go 1.21+ (for Go backend)
- Node.js 14+ (for frontend)
- TypeScript 5.0+ (for frontend)

## Migration Notes

The Go backend was migrated from Java with these key differences:
- Uses slices instead of arrays
- Package-level functions instead of static methods
- Reflection for generic type handling
- Go naming conventions (PascalCase for exported functions)
