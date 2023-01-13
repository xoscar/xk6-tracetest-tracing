import { Http, Tracetest } from "k6/x/tracetest";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "10s",
};

const testDefinition = `
type: Test
spec:
  id: VyP0gJ2VR
  name: v0.9.2test Pokeshop - Import
  description: Import a Pokemon
  trigger:
    type: http
    httpRequest:
      url: http://demo-pokemon-api.demo/pokemon/import
      method: POST
      headers:
      - key: Content-Type
        value: application/json
      body: '{"id":\${env:POKEID}}'
  outputs:
  - name: INSERTED_ID
    selector: span[tracetest.span.type="database" name="create pokeshop.pokemon" db.system="postgres"
      db.name="pokeshop" db.user="ashketchum" db.operation="create" db.sql.table="pokemon"]
    value: attr:db.result | json_path '.id'
`;

const http = new Http({ propagator: ["w3c", "b3"] });
const tracetest = new Tracetest();

export default function () {
  const response = http.get("https://test-api.k6.io");
  const run = tracetest.runFromDefinition(testDefinition, r.trace_id);

  console.log("~~~~~~~~~~");
  console.log("response.trace_id = ", response.trace_id);
  console.log("run.test.id =", run.test.id);
  console.log("run.test.id =", run.test_run.id); 
  console.log("run.test_run.trace_id =", run.test_run.trace_id);
  console.log("~~~~~~~~~~");
  console.log("");

  sleep(1);
}
