public class Solution {
    static void swap(int[] arr, int start, int end) {
        int temp = arr[start];
        arr[start] = arr[end];
        arr[end] = temp;
    }
    
    public static void minimumBribes(List<Integer> q) {
        int bribes = 0;
        
        int[] arr = q.stream().mapToInt(i -> i).toArray();
        
        for (int index = arr.length - 1; index >= 0; index--) {            
            if (arr[index] != index + 1) {
                if (((index - 1) >= 0) && arr[index - 1] == (index + 1)){
                    bribes++;
                    
                    swap(arr, index, index - 1);
                } else if (((index - 2) >= 0) && arr[index - 2] == (index + 1)) {
                    bribes += 2;
                    
                    swap(arr, index - 2, index - 1);
                    swap(arr, index - 1, index);
                } else {
                    System.out.println("Too chaotic");
                    return;
                }
            }
        }
        
        System.out.println(bribes);
    }
}
    