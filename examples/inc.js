import { Counter } from 'k6/x/atomic';

const counter = new Counter("some_id");

export default () => {
   // increase and store the current value
   const current = counter.inc();
   console.log("__VU:", __VU, "__ITER:", __ITER, ` current value is: ${current}`);
}
