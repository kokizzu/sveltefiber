<script>

	import { listNames, endpointUrl } from './store.js';
	let name = '';
	let lastResponse = '';

	function doSubmit(e) {
		e.preventDefault()
		fetch(endpointUrl+"/api/createName", {
			method: "POST",
			body: JSON.stringify({
				name: name,
			}),
			headers: {
				"Content-type": "application/json; charset=UTF-8"
			}
		})
			.then(response => response.json())
			.then(data => {
				console.log(data)
				lastResponse = data
				listNames.update(n => {
					return data.records || [];
				})
			})
	}
</script>

<!--<img src="/public/svelte.png"/>-->
<form on:submit={doSubmit}>
	Name: <input type="text" bind:value={name}/> <br/>
	<button on:click={doSubmit}>Send</button>
</form>
<div>{lastResponse ? JSON.stringify(lastResponse) : ''}</div>
<style>
</style>
