# Monolith Project

A simple monolith project demonstrating array utility functions implemented in both Java (backend) and TypeScript (frontend).

## Project Structure

```
monolith-project/
├── backend/                 # Java backend
│   └── src/main/java/
│       └── com/example/monolith/
│           └── ArrayUtility.java
├── frontend/               # TypeScript frontend
│   ├── ArrayUtility.ts
│   ├── package.json
│   └── tsconfig.json
└── README.md
```

## Features

### Backend (Java)
- `ArrayUtility` class with `findArraySize` method
- Handles null arrays gracefully
- Simple and clean implementation

### Frontend (TypeScript)
- `ArrayUtility` class with generic type support
- Same functionality as Java version
- Full TypeScript type safety
- Handles null and undefined values

## Getting Started

### Backend (Java)
```bash
cd backend/src/main/java
javac com/example/monolith/ArrayUtility.java
java com.example.monolith.ArrayUtility
```

### Frontend (TypeScript)
```bash
cd frontend
npm install
npm run dev
```

## Array Utility Function

Both implementations provide the same core functionality:

```java
// Java
public static int findArraySize(Object[] array)

// TypeScript
public static findArraySize<T>(array: T[] | null | undefined): number
```

The function returns the size of an array, or 0 if the array is null/undefined.

## Requirements

- Java 8+ (for backend)
- Node.js 14+ (for frontend)
- TypeScript 5.0+ (for frontend)
