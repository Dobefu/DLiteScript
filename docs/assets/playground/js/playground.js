(() => {
  "use strict";

  /**
   * @param {HTMLButtonElement} runBtn
   * @returns {Worker}
   */
  function createWorker(runBtn) {
    const worker = new Worker(playgroundWorkerPath);
    worker.postMessage({ method: "init", data: playgroundWasmPath });

    worker.onmessage = () => {
      runBtn.removeAttribute("disabled", "");
    };

    return worker;
  }

  const playgrounds = document.querySelectorAll(".playground");

  for (const playground of playgrounds) {
    let isRunning = false;

    const textarea = playground.querySelector(".playground__textarea");
    const runBtn = playground.querySelector(".playground__run-btn");
    const runIndicator = playground.querySelector(
      ".playground__run-indicator",
    );
    const output = playground.querySelector(".playground__output");

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
