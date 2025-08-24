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
     * Main method for testing the utility function
     */
    public static void main(String[] args) {
        // Test with different array types
        Integer[] intArray = {1, 2, 3, 4, 5};
        String[] stringArray = {"Hello", "World", "Java"};
        
        System.out.println("Integer array size: " + findArraySize(intArray));
        System.out.println("String array size: " + findArraySize(stringArray));
        System.out.println("Null array size: " + findArraySize((Object[]) null));
    }
}
