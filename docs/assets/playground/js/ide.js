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
})();
