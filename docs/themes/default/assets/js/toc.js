(() => {
  "use strict";

  const tocs = document.querySelectorAll(".toc");

  if (!tocs) {
    return;
  }

  for (const toc of tocs) {
    const tocHeader = toc.querySelector('[data-action="toggle-toc"]');

    if (!tocHeader) {
      continue;
    }

    tocHeader.addEventListener("click", (e) => {
      e.preventDefault();

      const tocToggle = tocHeader.querySelector(".toc__toggle");

      if (!tocToggle || getComputedStyle(tocToggle).display === "none") {
        return;
      }

      toc.classList.toggle("open");
    });
  }
})();
