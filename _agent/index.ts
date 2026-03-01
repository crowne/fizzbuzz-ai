import * as dotspec from "./specs/dotspec";
import * as hello from "./specs/hello";

async function main() {

  console.log("Agent starting");

  // await hello.generate();

  await dotspec.generate("./_agent/specs/calc.dot");

}

main().catch(console.error);