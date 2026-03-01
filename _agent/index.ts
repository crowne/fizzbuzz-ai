import * as dotspec from "./specs/dotspec";
import * as hello from "./specs/hello";

async function main() {

  console.log("Agent starting");

  await dotspec.generate("./_agent/specs/calc.dot");

  // await hello.generate();
  
}

main().catch(console.error);