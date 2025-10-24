(() => {
  "use strict";

  /** @type {NodeListOf<HTMLTextAreaElement>} */
  const ides = document.querySelectorAll(".ide");

  for (const ide of ides) {
    let wasEscapePressed = false;

    /**
     * @param {KeyboardEvent} e
     */
    ide.addEventListener("keydown", (e) => {
      switch (e.key) {
        case "Tab":
          if (!wasEscapePressed) {
            e.preventDefault();
            const lines = ide.value.split("\n");
            let selectionStart = ide.selectionStart;
            let selectionEnd = ide.selectionEnd;

            const selectedLines = getLinesFromSelection(
              lines,
              selectionStart,
              selectionEnd,
            );

            if (e.shiftKey) {
              for (const lineIdx of selectedLines) {
                const lineLen = lines[lineIdx].length;
                lines[lineIdx] = lines[lineIdx].replace(/^\s{1,2}/, "");

                selectionEnd -= lineLen - lines[lineIdx].length;
              }
            } else {
              for (const lineIdx of selectedLines) {
                lines[lineIdx] = `  ${lines[lineIdx]}`;

                selectionEnd += 2;
              }
            }

            ide.value = lines.join("\n");
            ide.dispatchEvent(new Event("input"));
            ide.dispatchEvent(new Event("change"));

            ide.selectionStart = selectionStart;
            ide.selectionEnd = selectionEnd;
          }

          wasEscapePressed = false;

          break;

        case "Escape":
          wasEscapePressed = true;

          break;

        case "Control":
        case "Alt":
        case "Meta":
        case "Shift":
          break;

        default:
          wasEscapePressed = false;

          break;
      }
    });
  }

  /**
   * @param {string[]} lines
   * @param {number} start
   * @param {number} end
   */
  function getLinesFromSelection(lines, start, end) {
    /** @type {number[]} */
    const selectedLines = [];

    let lineIdx = 0;
    let charIdx = 0;

    do {
      charIdx += lines[lineIdx].length;

      if (charIdx >= start) {
        selectedLines.push(lineIdx);
      }

      charIdx++;
      lineIdx++;
    } while (charIdx <= end);

    return selectedLines;
  }
})();
