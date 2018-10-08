# koala

### What's it?
**koala** is a CLI that join separate files generating only one file as a bundler. Nowadays, others bundlers already exist, they are more complete like [Webpack](https://github.com/webpack/webpack), but the idea of koala is to be simpler and generic"

### Get started
**koala** needs that you input an entrypoint to itself work and the entrypoint is just a file that **koala** will read like a hub to call others files, follow the sample of one entry file named as `./entry.js`.
```js
'use strict';

include lib/hello.js

hello("Koala");
```

The entry file above is used to find out the targets that is just the line of entry file prefixed with a tag, see the example below, where the tag is named as `include`.
```js
include lib/hello.js
```
Now, look the external file at `./lib/hello` where the koala named it as library, follow the sample below.

```js
function hello(name) {
  console.log(`Hello, ${name}`);
}
```
Now, look the output file at `./bin/out.js` and execute, this example execute the **koala** and [node.js](https://github.com/nodejs/node) together.
```js
'use strict';

function hello(name) {
  console.log(`Hello, ${name}`);
}

hello("Koala")
```

```sh
> koala ./entry.js ./bin/out.js include && node ./bin/out.js

2018/10/08 15:53:13 spelled successfully 104 bytes at /Users/user/remote/username/repo/bin/out.js
Hello, Koala
```
