import { Attractor, createVerboseLogger } from "attractor";
import { readFile } from "fs/promises";

async function generate(filePath: string) {

  const dotSource = await readFile(filePath, "utf8");

  const attractor = await Attractor.create({
    // read calc.dot contents into dotSource
    dotSource: dotSource,
    provider: "ollama",
    model: "qwen3-coder:30b",   // must match a model you pulled in Step 2
    onEvent: createVerboseLogger(),
  });

  const result = await attractor.run();

  console.log("Pipeline finished:", result.state);
  console.log("Nodes executed:", result.results.length);

  // for (const r of result.results) {
  //   console.log(`  [${r.node}] ${r.status}`);
  // }
}

export { generate };