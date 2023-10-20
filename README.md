# xk6-atomic

This extension is a simple PoC/wrapper that uses underline golang's atomic package to provide in-memory counters that could be shared between VUs.

It could be used to share counters between VUs.

```javascript
import { Counter } from 'k6/x/atomic';

let counter = new Counter("some_id");

export default () => {
   console.log("__VU:", __VU, "__ITER:", __ITER);

   console.log(`increase and print the current value: ${counter.inc()}`);
}
```

See examples for more.

## Requirements

* [Golang 1.19+](https://go.dev/)
* [Git](https://git-scm.com/)
* [xk6](https://github.com/grafana/xk6) (`go install go.k6.io/xk6/cmd/xk6@lates`)



## Getting started  

1. Build the k6's binary with the module:

  ```shell
  $ make build
  ```

2. Run the example:

  ```shell
  $ ./k6 run -i 10 --vus 4 examples/script.js
  ```