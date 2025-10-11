(() => {
  "use strict";

  const playgrounds = document.querySelectorAll(".playground");

  for (const playground of playgrounds) {
    let worker = new Worker(playgroundWorkerPath);
    worker.postMessage({ method: "init", data: playgroundWasmPath });

    let isRunning = false;

    const textarea = playground.querySelector(".playground__textarea");
    const runBtn = playground.querySelector(".playground__run-btn");
    const runIndicator = playground.querySelector(
      ".playground__run-indicator",
    );
    const output = playground.querySelector(".playground__output");

    worker.onmessage = () => {
      runBtn.removeAttribute("disabled", "");
    };

    runBtn.addEventListener("click", () => {
      if (isRunning) {
        worker.terminate();
        runBtn.setAttribute("disabled", "");
        isRunning = false;
        runBtn.innerText = "Run code";
        output.innerText = "Cancelled";

        worker = new Worker(playgroundWorkerPath);
        worker.postMessage({ method: "init", data: playgroundWasmPath });
        worker.onmessage = () => {
          runBtn.removeAttribute("disabled", "");
        };

        return;
      }

      isRunning = true;

      runBtn.innerText = "Cancel";
      output.innerText = "Running...";
      output.classList.remove("has-error");

      const msgHandler = (e) => {
        if (e.data.method !== "result") {
          return;
        }

        worker.removeEventListener("message", msgHandler);

        const result = JSON.parse(e.data.data);

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
        }
      };

      worker.addEventListener("message", msgHandler);
      worker.postMessage({ method: "run", data: textarea.value ?? "" });
    });
  }
})();
