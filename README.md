# SvelteFiber example

minimal example how to integrate svelte-mpa with [Golang](//go.dev) (in this case Fiber framework)

## Dependencies

- [gofiber](//gofiber.io)
- [svelte-mpa](//github.com/kokizzu/svelte-mpa)

## How to start

```shell
go mod tidy

# start golang backend
air # localhost:3001 for backend and SSR

# start auto recompile
cd svelte 
npm i && npm start # localhost:5500 for client side dev mode
# or using bun
bun install && bun start 
```

## How it works

- svelte-mpa will generate `index.html` that loaded by fiber app
- fiber app renders/replace any special [template](//github.com/kokizzu/gotro/tree/master/Z) keywords in the `index.html`
- all other API handled by normal gofiber handler

