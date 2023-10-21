# xk6-atomic

This extension is a simple PoC/wrapper that uses underline golang's atomic package to provide in-memory counters that could be shared between VUs.

It could be used to share counters between VUs.

```javascript
import { Counter } from 'k6/x/atomic';

let counter = new Counter("some_id");

export default () => {
   // increase and store the current value
   let current = counter.inc();
   console.log("__VU:", __VU, "__ITER:", __ITER, ` current value is: ${current}`);
}
```

The `Counter` class has the following methods:

* `inc()` - increase the counter by 1 and return the current value
* `dec()` - decrease the counter by 1 and return the current value
* `add(n)` - increase (or decrease) the counter by `n` and return the current value
* `value()` - return the current value

See examples for more.

## Requirements

* [Golang 1.19+](https://go.dev/)
* [Git](https://git-scm.com/)
* [xk6](https://github.com/grafana/xk6) (`go install go.k6.io/xk6/cmd/xk6@lates`)


## Getting started  

1. Build the k6's binary with the module:

  ```shell
  $ xk6 build v0.47.0 --with github.com/olegbespalov/xk6-atomic
  ```

2. Run the example:

  ```shell
  $ ./k6 run -i 10 --vus 4 examples/script.js
  ```