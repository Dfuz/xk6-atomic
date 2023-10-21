import { Counter } from 'k6/x/atomic';

let counter = new Counter("another_id");

export default () => {
   // decrease and store the current value
   let current = counter.dec();
   console.log("__VU:", __VU, "__ITER:", __ITER, ` current value is: ${current}`);
}