(() => {
  "use strict";

  /**
   * @param {MessageEvent} e
   */
  self.addEventListener("message", (e) => {
    if (typeof e.data !== "object" || e.data === null) {
      console.error("Invalid data", e.data);

      return;
    }

    /**
     * @type {object}
     * @property {string|undefined} method
     * @property {string|undefined} data
     */
    const data = e.data;

    switch (data.method) {
      case "init":
        postMessage(init(data.data));

        break;

      case "run":
        postMessage(run(data.data));

        break;

      default:
        console.error("Unknown method", data.method);
    }
  });

  /**
   * @param {string} playgroundWasmPath
   * @returns {object}
   */
  function init(playgroundWasmPath) {
    const go = new globalThis.Go();

    WebAssembly.instantiateStreaming(
      fetch(playgroundWasmPath),
      go.importObject,
    )
      .then((result) => {
        go.run(result.instance);
      });

    return {
      method: "init",
    };
  }

  /**
   * @param {string} code
   * @returns {object}
   */
  function run(code) {
    /** @type {object} */
    let result;

    try {
      result = globalThis.runString(code);
    } catch (error) {
      result = error.message
    }

    return {
      method: "result",
      data: result,
    };
  }
})();
