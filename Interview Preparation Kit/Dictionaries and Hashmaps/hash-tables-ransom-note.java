public class Solution {
    static void checkMagazine(String[] magazine, String[] note) {
        List<String> magazineWords = Arrays
            .stream(magazine)
            .sorted()
            .collect(Collectors.toList());
        
        boolean notRemovedWord = Arrays
            .stream(note)
            .anyMatch(word -> !magazineWords.remove(word));

        System.out.println(notRemovedWord ? "No": "Yes");
    }
}