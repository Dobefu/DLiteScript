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
          const output = playground.querySelector(".playground__output");

          runBtn.addEventListener("click", () => {
            output.innerHTML = "";
            output.classList.remove("has-error");

            result = JSON.parse(runString(textarea.value ?? ""));

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
