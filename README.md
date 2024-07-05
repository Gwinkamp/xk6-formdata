# xk6-formdata

Extension for [k6](https://k6.io). Allows to create a request body in multipart form-data format.  
In the original `k6/jslib` library, I noticed a problem with Cyrillic encodings when creating multipart form-data request payload. This extension is designed to solve this problem.

## Requirements

* [Golang](https://go.dev/)
* [xk6](https://k6.io/blog/extending-k6-with-xk6/)

```shell
go install go.k6.io/xk6/cmd/xk6@latest
```

## Build

From local repository:

```shell
xk6 build --with xk6-formdata=.
```

From remote repository:

```shell
xk6 build --with github.com/Gwinkamp/xk6-formdata
```

## Usage

In load testing scenarios

```javascript
import { check } from "k6";
import http from "k6/http";
import formdata from 'k6/x/formdata';

export const options = {
  target: 1,
  duration: "10s",
};

const file1 = open("path/to/file1", "b");
const file2 = open("path/to/file2", "b");

export default function () {
  const builder = new formdata.Builder();

  builder.add("param1", "value");  // add simple string field
  builder.addBytes("param2", file1);  // add field with bytes (ArrayBuffer)
  builder.addFile("param3", "test.txt", file2);  // add file field

  const payload = builder.build();  // create form-data payload

  const response = http.post(
    "http://localhost:8080/test",
    payload,
    {
      headers: {
        "Content-Type": builder.getContentType(),  // get ContentType header with boundary
      },
    }
  );
  check(response, { "status is 200": (r) => r.status === 200 });
}
```

To run this script, you need to run k6 executable file, which was previously built with `xk6 build` command

```shell
./k6 run scripts/example.js
```

## Contribution

Freely. I am always glad to have suggestions