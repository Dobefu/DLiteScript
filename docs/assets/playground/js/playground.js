if (globalThis.playgroundWasmPath === undefined) {
  throw new TypeError("playgroundWasmPath is not defined");
}

if (globalThis.playgroundWorkerPath === undefined) {
  throw new TypeError("playgroundWorkerPath is not defined");
}

(() => {
  "use strict";

  /**
   * @param {HTMLButtonElement} runBtn
   * @returns {{worker: Worker, abortController: AbortController}}
   */
  function createWorker(runBtn) {
    const worker = new Worker(globalThis.playgroundWorkerPath);
    const abortController = new AbortController();

    /**
     * @param {MessageEvent} e
     */
    worker.addEventListener("message", (e) => {
      if (e.data.method !== "init") {
        return;
      }

      runBtn.removeAttribute("disabled");
    }, { signal: abortController.signal });

    worker.postMessage({ method: "init", data: globalThis.playgroundWasmPath });

    return { worker, abortController };
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

    let { worker, abortController } = createWorker(runBtn);

    globalThis.addEventListener("beforeunload", () => {
      abortController.abort();
      worker.terminate();
    });

    runBtn.addEventListener("click", () => {
      if (isRunning) {
        abortController.abort();
        worker.terminate();
        ({ worker, abortController } = createWorker(runBtn));

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

      worker.addEventListener("message", msgHandler, {
        signal: abortController.signal,
        once: true,
      });

      worker.postMessage({ method: "run", data: textarea.value ?? "" });
    });
  }
})();
