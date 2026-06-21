/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function (s) {
    const words = s.trim().split(/\s+/);

    let left = 0;
    let right = words.length - 1;

    while (left < right) {
        [words[left], words[right]] = [words[right], words[left]];
        left++;
        right--;
    }

    return words.join(" ");
};