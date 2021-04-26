public class Solution {
    static int[] rotLeft(int[] a, int d) {
        int len = a.length;
        int[] rotateArr = new int[len];

        for (int i = 0; i < len; i++) {
            int j = (i + len - d) % len;

            rotateArr[j] = a[i];
        }

        return rotateArr;
    }
}
