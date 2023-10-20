import { Counter } from 'k6/x/atomic';

let counter1 = new Counter("id1");
let counter2 = new Counter("id2");

export default () => {
   console.log("__VU:", __VU, "__ITER:", __ITER);

   console.log(`counter1 value: ${counter1.inc()}`);
   console.log(`counter2 value: ${counter2.inc()}`);
}