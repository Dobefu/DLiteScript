(() => {
  "use strict";

  /**
   * @param {HTMLButtonElement} runBtn
   * @returns {Worker}
   */
  function createWorker(runBtn) {
    const worker = new Worker(globalThis.playgroundWorkerPath);
    worker.postMessage({ method: "init", data: globalThis.playgroundWasmPath });

    worker.onmessage = () => {
      runBtn.removeAttribute("disabled");
    };

    return worker;
  }

  const playgrounds = document.querySelectorAll(".playground");

  for (const playground of playgrounds) {
    let isRunning = false;

    /** @type {HTMLTextAreaElement | null} */
    const textarea = playground.querySelector(".playground__textarea");

    if (!textarea) {
      console.error("Textarea not found");

      continue;
    }

    /** @type {HTMLButtonElement | null} */
    const runBtn = playground.querySelector(".playground__run-btn");

    if (!runBtn) {
      console.error("Run button not found");

      continue;
    }

    /** @type {HTMLDivElement | null} */
    const runIndicator = playground.querySelector(
      ".playground__run-indicator",
    );

    if (!runIndicator) {
      console.error("Run indicator not found");

      continue;
    }

    /** @type {HTMLDivElement | null} */
    const output = playground.querySelector(".playground__output");

    if (!output) {
      console.error("Output element not found");

      continue;
    }

    let worker = createWorker(runBtn);

    globalThis.addEventListener("beforeunload", () => {
      worker.terminate();
    });

    runBtn.addEventListener("click", () => {
      if (isRunning) {
        worker.terminate();
        worker = createWorker(runBtn);

        runBtn.setAttribute("disabled", "");
        isRunning = false;
        runBtn.innerText = "Run code";
        output.innerText = "Cancelled";

        return;
      }

      isRunning = true;

      runBtn.innerText = "Cancel";
      output.innerText = "Running...";
      output.classList.remove("has-error");

      /**
       * @param {MessageEvent} e
       */
      const msgHandler = (e) => {
        if (e.data.method !== "result") {
          return;
        }

        worker.removeEventListener("message", msgHandler);

        let result;

        try {
          result = JSON.parse(e.data.data);
        } catch (error) {
          console.error(error);

          result = {
            error: "Script execution error",
          };
        }

        runIndicator.classList.add("is-animating");

        setTimeout(() => {
          runIndicator.classList.remove("is-animating");
        }, 200);

        isRunning = false;
        runBtn.innerText = "Run code";
        output.innerText = "";

        if ("error" in result) {
          output.classList.add("has-error");
          output.innerText = result.error;
        } else if ("buffer" in result) {
          output.innerText = result.buffer;
        } else {
          output.innerText = "<no output>";
        }
      };

      worker.addEventListener("message", msgHandler);
      worker.postMessage({ method: "run", data: textarea.value ?? "" });
    });
  }
})();
