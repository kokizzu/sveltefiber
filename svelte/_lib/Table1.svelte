<script>
	import { listNames } from './store.js';

	export let list;
	$: list = $listNames
	$: console.log($listNames)
	let lastResponse = '';

	function doDelete(e, id) {
		fetch("/api/deleteName", {
			method: "POST",
			body: JSON.stringify({
				id: id,
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

<table border="1">
	<tr>
		<th>id</th>
		<th>name</th>
		<th>action</th>
	</tr>
	{#each list as row}
		<tr>
			<td>{row.id}</td>
			<td>{row.name}</td>
			<td>
				<button on:click={(e) => doDelete(e,row.id)}> Delete</button>
			</td>
		</tr>
	{/each}
</table>
<div>{lastResponse ? JSON.stringify(lastResponse) : ''}</div>
<style>
</style>
