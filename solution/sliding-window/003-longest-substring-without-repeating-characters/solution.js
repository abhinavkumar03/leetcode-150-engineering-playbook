/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
    let left = 0;
    let maxLength = 0;

    const window = new Set();

    for (let right = 0; right < s.length; right++) {
        while (window.has(s[right])) {
            window.delete(s[left]);
            left++;
        }

        window.add(s[right]);

        maxLength = Math.max(maxLength, right - left + 1);
    }

    return maxLength;
};