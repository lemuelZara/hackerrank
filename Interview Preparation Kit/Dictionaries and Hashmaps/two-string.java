public class Solution {
    public static String twoStrings(String s1, String s2) {
        Map<Character, Integer> checkString = new HashMap<Character, Integer>();
        String result = "NO";
        
        for (int i = 0; i < s1.length(); i++) {
            checkString.put(s1.charAt(i), i);
        }
           
        for (int i = 0; i < s2.length(); i++) {
            if (checkString.containsKey(s2.charAt(i))) {
                result = "YES";
            }
        }
        
        return result;
    }
}