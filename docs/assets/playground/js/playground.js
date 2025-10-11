(() => {
  "use strict";

  const go = new Go();

  WebAssembly.instantiateStreaming(fetch(playgroundWasmPath), go.importObject)
    .then(
      (result) => {
        go.run(result.instance);

        const playgrounds = document.querySelectorAll(".playground");

        for (const playground of playgrounds) {
          const textarea = playground.querySelector(".playground__textarea");
          const runBtn = playground.querySelector(".playground__run-btn");
          const runIndicator = playground.querySelector(
            ".playground__run-indicator",
          );
          const output = playground.querySelector(".playground__output");

          runBtn.addEventListener("click", () => {
            runBtn.setAttribute("disabled", "");
            output.innerHTML = "Running...";
            output.classList.remove("has-error");

            result = JSON.parse(runString(textarea.value ?? ""));
            runIndicator.classList.add("is-animating");

            setTimeout(() => {
              runIndicator.classList.remove("is-animating");
              runBtn.removeAttribute("disabled");
            }, 200);

            output.innerText = "";

            if ("error" in result) {
              output.classList.add("has-error");
              output.innerText = result.error;
            } else if ("buffer" in result) {
              output.innerText = result.buffer;
            }
          });
        }
      },
    );
})();
