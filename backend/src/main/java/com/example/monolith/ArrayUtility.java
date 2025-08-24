package com.example.monolith;

/**
 * Simple utility class for array operations
 */
public class ArrayUtility {
    
    /**
     * Finds the size of an array
     * @param array The array to find the size of
     * @return The size of the array, or 0 if null
     */
    public static int findArraySize(Object[] array) {
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
     * @return The sum of the three numbers
     */
    public static int findSumOfThree(int a, int b, int c) {
        return a + b + c;
    }
    
    /**
     * Main method for testing the utility functions
     */
    public static void main(String[] args) {
        // Test with different array types
        Integer[] intArray = {1, 2, 3, 4, 5};
        String[] stringArray = {"Hello", "World", "Java"};
        
        System.out.println("Integer array size: " + findArraySize(intArray));
        System.out.println("String array size: " + findArraySize(stringArray));
        System.out.println("Null array size: " + findArraySize((Object[]) null));
        
        // Test the new sum function
        System.out.println("Sum of 10, 20, 30: " + findSumOfThree(10, 20, 30));
        System.out.println("Sum of 5, 15, 25: " + findSumOfThree(5, 15, 25));
    }
}
