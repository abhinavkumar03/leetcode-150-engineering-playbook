/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function (s) {
    let index = s.length - 1;

    // Skip trailing spaces
    while (index >= 0 && s[index] === ' ') {
        index--;
    }

    let length = 0;

    // Count characters of the last word
    while (index >= 0 && s[index] !== ' ') {
        length++;
        index--;
    }

    return length;
};