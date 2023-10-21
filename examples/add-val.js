import { Counter } from 'k6/x/atomic';

let counter1 = new Counter("id1");
let counter2 = new Counter("id2");

export default () => {
   console.log("__VU:", __VU, "__ITER:", __ITER);

   console.log(`add 2 for counter1: ${counter1.add(2)}`);
   console.log(`add 1 for counter2: ${counter2.add(1)}`);
}

export function teardown(data) {
   console.log(`teardown's value of counter1: ${counter1.val()}`);
   console.log(`teardown's value of counter2: ${counter2.val()}`);
}
