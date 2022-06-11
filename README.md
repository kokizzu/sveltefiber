# sveltefiber


## Dependencies

- [gofiber](//gofiber.io)
- [svelte-mpa](//github.com/kokizzu/svelte-mpa)

## How to start

```shell
go mod tidy

# start golang backend
air # localhost:3001 for backend and SSR

# start auto recompile
cd svelte && npm i && npm start # localhost:5500 for client side dev mode
```

## How it works

svelte-mpa will generate `index.html` that loadded by fiber app, 
fiber app renders whatever variable in the index.html, all other API handled by normal gofiber handler.

