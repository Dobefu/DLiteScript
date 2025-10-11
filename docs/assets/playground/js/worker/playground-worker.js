(() => {
  "use strict";

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

  function init(playgroundWasmPath) {
    const go = new Go();

    WebAssembly.instantiateStreaming(
      fetch(playgroundWasmPath),
      go.importObject,
    )
      .then((result) => {
        go.run(result.instance);
      });

    return {};
  }

  function run(code) {
    return {
      method: "result",
      data: runString(code),
    };
  }
})();
