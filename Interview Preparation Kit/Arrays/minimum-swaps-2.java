public class Solution {    
    static int minimumSwaps(int[] arr) {
        int swaps = 0;
        
        for (int index = 0; index < arr.length; index++) {
            while (arr[index] != index + 1) {
                int swapIndex = arr[index] - 1;
                int valueAtIndex = arr[index];
                int valueAtSwapIndex = arr[swapIndex];
                
                arr[index] = valueAtSwapIndex;
                arr[swapIndex] = valueAtIndex;
                
                swaps++;
            }
        }
        
        return swaps;
    }
}