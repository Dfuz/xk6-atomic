import { Counter } from 'k6/x/atomic';

let counter = new Counter("some_id");

export default () => {
   // increase and store the current value
   let current = counter.inc();
   console.log("__VU:", __VU, "__ITER:", __ITER, ` current value is: ${current}`);
}
