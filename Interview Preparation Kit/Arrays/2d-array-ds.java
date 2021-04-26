public class Solution {
    static int hourglassSum(int[][] arr) {
        int result = 0;
        int max = -63;

        for (int i = 0; i <= arr.length / 2; i++) {
            for (int j = 0; j <= arr[i].length / 2; j++) {
                int top =
                    arr[i][j] +
                    arr[i][1 + j] +
                    arr[i][2 + j];
                int mid = arr[1 + i][1 + j];
                int bot =
                    arr[2 + i][j] +
                    arr[2 + i][1 + j] +
                    arr[2 + i][2 + j];

                result = top + mid + bot;

                if (result > max) {
                    max = result;
                }
            }
        }

        return max;
    }
}
