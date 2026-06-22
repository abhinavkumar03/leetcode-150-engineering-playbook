/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
function convert(s, numRows) {
    if (numRows === 1 || numRows >= s.length) {
        return s;
    }

    const rows = Array.from(
        { length: numRows },
        () => []
    );

    let currentRow = 0;
    let direction = 1;

    for (const char of s) {
        rows[currentRow].push(char);

        if (currentRow === 0) {
            direction = 1;
        } else if (currentRow === numRows - 1) {
            direction = -1;
        }

        currentRow += direction;
    }

    return rows
        .map(row => row.join(""))
        .join("");
}