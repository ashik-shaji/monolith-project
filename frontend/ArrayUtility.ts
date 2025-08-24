/**
 * Simple utility class for array operations in TypeScript
 */
export class ArrayUtility {
    
    /**
     * Finds the size of an array
     * @param array The array to find the size of
     * @returns The size of the array, or 0 if null/undefined
     */
    public static findArraySize<T>(array: T[] | null | undefined): number {
        if (array == null) {
            return 0;
        }
        return array.length;
    }
}

// Example usage and testing
if (require.main === module) {
    // Test with different array types
    const intArray: number[] = [1, 2, 3, 4, 5];
    const stringArray: string[] = ["Hello", "World", "TypeScript"];
    
    console.log("Number array size:", ArrayUtility.findArraySize(intArray));
    console.log("String array size:", ArrayUtility.findArraySize(stringArray));
    console.log("Null array size:", ArrayUtility.findArraySize(null));
    console.log("Undefined array size:", ArrayUtility.findArraySize(undefined));
}
