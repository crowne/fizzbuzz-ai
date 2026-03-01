import { Attractor, createVerboseLogger } from "attractor";

async function generate() {

  const attractor = await Attractor.create({
    dotSource: "digraph { start [shape=ellipse] }",
    provider: "ollama",
    model: "qwen3-coder:30b",   // must match a model you pulled in Step 2
    onEvent: createVerboseLogger(),
  });

  const response = await attractor.runAgent(
    "Create a file called cmd/hello.go that prints 'Hello, World!' when executed.",
  );

  console.log("Agent response:", response);
}

export { generate };