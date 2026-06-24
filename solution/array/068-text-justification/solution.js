var fullJustify = function (words, maxWidth) {
    const result = [];
    let index = 0;

    while (index < words.length) {
        const lineStart = index;
        let lineLength = words[index].length;

        index++;

        while (
            index < words.length &&
            lineLength + 1 + words[index].length <= maxWidth
        ) {
            lineLength += 1 + words[index].length;
            index++;
        }

        const lineWords = words.slice(lineStart, index);
        const isLastLine = index === words.length;

        result.push(
            buildLine(lineWords, maxWidth, isLastLine)
        );
    }

    return result;
};

function buildLine(words, maxWidth, isLastLine) {
    const wordCount = words.length;

    let totalWordLength = 0;

    for (const word of words) {
        totalWordLength += word.length;
    }

    if (isLastLine || wordCount === 1) {
        let line = words.join(" ");
        line += " ".repeat(maxWidth - line.length);
        return line;
    }

    const totalSpaces = maxWidth - totalWordLength;
    const gaps = wordCount - 1;

    const baseSpaces = Math.floor(totalSpaces / gaps);
    const extraSpaces = totalSpaces % gaps;

    let line = "";

    for (let i = 0; i < wordCount; i++) {
        line += words[i];

        if (i === wordCount - 1) {
            continue;
        }

        let spaces = baseSpaces;

        if (i < extraSpaces) {
            spaces++;
        }

        line += " ".repeat(spaces);
    }

    return line;
}