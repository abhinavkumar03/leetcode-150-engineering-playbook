/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function (haystack, needle) {
    const n = haystack.length;
    const m = needle.length;

    if (m === 0) {
        return 0;
    }

    for (let start = 0; start <= n - m; start++) {
        let match = true;

        for (let offset = 0; offset < m; offset++) {
            if (haystack[start + offset] !== needle[offset]) {
                match = false;
                break;
            }
        }

        if (match) {
            return start;
        }
    }

    return -1;
};