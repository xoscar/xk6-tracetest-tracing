import { Http, Tracetest } from "k6/x/tracetest";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "10s",
};

const testId = "VyP0gJ2VR";
const http = new Http({ propagator: ["w3c", "b3"] });
const tracetest = new Tracetest();

export default function () {
  const response = http.get("https://test-api.k6.io");
  const run = tracetest.runTest(testId, response.trace_id);

  console.log("~~~~~~~~~~");
  console.log("response.trace_id = ", response.trace_id);
  console.log("run.test.id =", run.test.id);
  console.log("run.test.id =", run.test_run.id); 
  console.log("run.test_run.trace_id =", run.test_run.trace_id);
  console.log("~~~~~~~~~~");
  console.log("");

  sleep(1);
}
