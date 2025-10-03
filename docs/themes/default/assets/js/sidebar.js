(() => {
  "use strict";

  const sidebar = document.querySelector(".sidebar");

  if (!sidebar) {
    return;
  }

  const sidebarToggles = document.querySelectorAll(
    '[data-action="toggle-sidebar"]',
  );

  for (const toggle of sidebarToggles) {
    toggle.addEventListener("click", (e) => {
      e.preventDefault();

      sidebar.classList.add("is-transitioning");
      sidebar.classList.toggle("open");

      setTimeout(() => {
        sidebar.classList.remove("is-transitioning");
      }, 200);
    });
  }
})();
