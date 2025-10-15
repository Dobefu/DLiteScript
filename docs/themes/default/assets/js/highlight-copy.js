(() => {
  "use strict";

  const highlights = document.querySelectorAll(".highlight");

  if (!highlights) {
    return;
  }

  for (const highlight of highlights) {
    const copyBtn = document.createElement("button");
    copyBtn.classList.add("highlight__copy");

    highlight.addEventListener("click", () => {
      /** @type {HTMLPreElement | null} */
      const codeElement = highlight.querySelector("[data-lang]");
      const code = (codeElement?.innerText ?? "").replaceAll("\n\n", "\n");

      navigator.clipboard.writeText(code);
      copyBtn.classList.add("copied");

      setTimeout(() => {
        copyBtn.classList.remove("copied");
      }, 1000);
    });

    highlight.insertBefore(copyBtn, highlight.firstChild);
  }
})();
