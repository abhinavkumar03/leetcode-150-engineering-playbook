/**
 * @param {string} s
 * @param {string[]} words
 * @return {number[]}
 */
function findSubstring(s, words) {
    if (s.length === 0 || words.length === 0) {
        return [];
    }

    const wordLength = words[0].length;
    const wordCount = words.length;
    const windowLength = wordLength * wordCount;

    if (s.length < windowLength) {
        return [];
    }

    const targetFrequency = new Map();

    for (const word of words) {
        targetFrequency.set(
            word,
            (targetFrequency.get(word) || 0) + 1
        );
    }

    const result = [];

    for (let offset = 0; offset < wordLength; offset++) {

        let left = offset;
        let wordsInWindow = 0;
        const currentFrequency = new Map();

        for (
            let right = offset;
            right + wordLength <= s.length;
            right += wordLength
        ) {
            const currentWord = s.substring(right, right + wordLength);

            if (!targetFrequency.has(currentWord)) {
                currentFrequency.clear();
                wordsInWindow = 0;
                left = right + wordLength;
                continue;
            }

            currentFrequency.set(
                currentWord,
                (currentFrequency.get(currentWord) || 0) + 1
            );

            wordsInWindow++;

            while (
                currentFrequency.get(currentWord) >
                targetFrequency.get(currentWord)
            ) {
                const leftWord = s.substring(left, left + wordLength);

                currentFrequency.set(
                    leftWord,
                    currentFrequency.get(leftWord) - 1
                );

                wordsInWindow--;
                left += wordLength;
            }

            if (wordsInWindow === wordCount) {
                result.push(left);

                const leftWord = s.substring(left, left + wordLength);

                currentFrequency.set(
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