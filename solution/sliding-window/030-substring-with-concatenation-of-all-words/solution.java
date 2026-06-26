import java.util.*;

class Solution {

    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> result = new ArrayList<>();

        if (s == null || s.isEmpty() || words.length == 0) {
            return result;
        }

        int wordLength = words[0].length();
        int wordCount = words.length;
        int windowLength = wordLength * wordCount;

        if (s.length() < windowLength) {
            return result;
        }

        Map<String, Integer> targetFrequency = new HashMap<>();

        for (String word : words) {
            targetFrequency.put(word, targetFrequency.getOrDefault(word, 0) + 1);
        }

        for (int offset = 0; offset < wordLength; offset++) {

            int left = offset;
            int wordsInWindow = 0;

            Map<String, Integer> currentFrequency = new HashMap<>();

            for (int right = offset; right + wordLength <= s.length(); right += wordLength) {

                String currentWord = s.substring(right, right + wordLength);

                if (!targetFrequency.containsKey(currentWord)) {
                    currentFrequency.clear();
                    wordsInWindow = 0;
                    left = right + wordLength;
                    continue;
                }

                currentFrequency.put(
                    currentWord,
                    currentFrequency.getOrDefault(currentWord, 0) + 1
                );

                wordsInWindow++;

                while (currentFrequency.get(currentWord) > targetFrequency.get(currentWord)) {
                    String leftWord = s.substring(left, left + wordLength);

                    currentFrequency.put(
                        leftWord,
                        currentFrequency.get(leftWord) - 1
                    );

                    wordsInWindow--;
                    left += wordLength;
                }

                if (wordsInWindow == wordCount) {
                    result.add(left);

                    String leftWord = s.substring(left, left + wordLength);

                    currentFrequency.put(
                        leftWord,
                        currentFrequency.get(leftWord) - 1
                    );

                    wordsInWindow--;
                    left += wordLength;
                }
            }
        }

        return result;
    }
}