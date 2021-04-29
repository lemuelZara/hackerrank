public class Solution {
    static long getMaxValue(long[] arr) {
        long max = 0;
        long sum = 0;

        for (int i = 0; i < arr.length; i++) {
            sum = sum + arr[i];

            if (sum > max) {
                max = sum;
            }
        }
        
        return max;
    }

    static long arrayManipulation(int n, int[][] queries) {
        long[] arr = new long[n + 2];

        for (int i = 0; i < queries.length; i++) {
            int a = queries[i][0];
            int b = queries[i][1];
            int k = queries[i][2];

            arr[a] = arr[a] + k;
            arr[b + 1] = arr[b + 1] - k;
        }

        return getMaxValue(arr);
    }
}
