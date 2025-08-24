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
    
    /**
     * Finds the sum of three numbers
     * @param a First number
     * @param b Second number
     * @param c Third number
     * @returns The sum of the three numbers
     */
    public static findSumOfThree(a: number, b: number, c: number): number {
        return a + b + c;
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
    
    // Test the new sum function
    console.log("Sum of 10, 20, 30:", ArrayUtility.findSumOfThree(10, 20, 30));
    console.log("Sum of 5, 15, 25:", ArrayUtility.findSumOfThree(5, 15, 25));
}
