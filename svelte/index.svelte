<script>
  import Form1 from './_lib/Form1.svelte'
  import Table1 from './_lib/Table1.svelte'
  
  let rawData = [/* raw_data */];
</script>

<main>
	<Form1></Form1>
	<Table1></Table1>

	{JSON.stringify(rawData)}
</main>

<style>
</style>
