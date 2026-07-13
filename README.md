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
```

## How it works

- svelte-mpa will generate `index.html` that loaded by fiber app
- fiber app renders/replace any special [template](//github.com/kokizzu/gotro/tree/master/Z) keywords in the `index.html`
- all other API handled by normal gofiber handler

## Maintenance checklist

- [x] Go runtime updated to 1.26.5.
- [x] Go dependencies refreshed and module files tidied.
- [x] Frontend runtime documented for Node 24 and npm 11.
- [x] Svelte MPA build tooling updated, vulnerable dev-server/notification dependencies removed, and `npm audit` is clean.
- [x] Svelte table markup fixed for Svelte 5 validation.
- [x] `make test` runs Go tests and the Svelte production build.
- [x] `make verify-dependency-security` and `make vulncheck` check dependency security.
