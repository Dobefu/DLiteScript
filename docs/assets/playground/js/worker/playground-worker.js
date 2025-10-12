(() => {
  "use strict";

  /**
   * @param {MessageEvent} e
   */
  onmessage = (e) => {
    switch (e.data.method) {
      case "init":
        postMessage(init(e.data.data));

        break;

      case "run":
        postMessage(run(e.data.data));

        break;

      default:
        console.log(e.data);
    }
  };

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

    return {};
  }

  /**
   * @param {string} code
   * @returns {object}
   */
  function run(code) {
    return {
      method: "result",
      data: globalThis.runString(code),
    };
  }
})();
