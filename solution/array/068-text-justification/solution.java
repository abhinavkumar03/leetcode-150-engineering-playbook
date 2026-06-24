import java.util.ArrayList;
import java.util.List;

class Solution {

    public List<String> fullJustify(String[] words, int maxWidth) {
        List<String> result = new ArrayList<>();

        int index = 0;

        while (index < words.length) {
            int lineStart = index;
            int lineLength = words[index].length();

            index++;

            while (index < words.length &&
                   lineLength + 1 + words[index].length() <= maxWidth) {
                lineLength += 1 + words[index].length();
                index++;
            }

            List<String> lineWords = new ArrayList<>();

            for (int i = lineStart; i < index; i++) {
                lineWords.add(words[i]);
            }

            boolean isLastLine = index == words.length;

            result.add(
                buildLine(lineWords, maxWidth, isLastLine)
            );
        }

        return result;
    }

    private String buildLine(
        List<String> words,
        int maxWidth,
        boolean isLastLine
    ) {
        int wordCount = words.size();

        int totalWordLength = 0;
        for (String word : words) {
            totalWordLength += word.length();
        }

        if (isLastLine || wordCount == 1) {
            String line = String.join(" ", words);

            StringBuilder builder = new StringBuilder(line);

            while (builder.length() < maxWidth) {
                builder.append(' ');
            }

            return builder.toString();
        }

        int totalSpaces = maxWidth - totalWordLength;
        int gaps = wordCount - 1;

        int baseSpaces = totalSpaces / gaps;
        int extraSpaces = totalSpaces % gaps;

        StringBuilder builder = new StringBuilder();

        for (int i = 0; i < wordCount; i++) {
            builder.append(words.get(i));

            if (i == wordCount - 1) {
                continue;
            }

            int spaces = baseSpaces;

            if (i < extraSpaces) {
                spaces++;
            }

            for (int j = 0; j < spaces; j++) {
                builder.append(' ');
            }
        }

        return builder.toString();
    }
}