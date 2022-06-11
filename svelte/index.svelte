<script>
  import Form1 from './_lib/Form1.svelte'
  import Table1 from './_lib/Table1.svelte'
  import {listNames} from "./_lib/store.js";

  let rawData = [/* raw_data */]
  if(rawData && rawData.length) {
	  listNames.update(n => {
		  return rawData;
	  })
  }
  
</script>

<main>
	<Form1></Form1>
	<Table1></Table1>
</main>

<style>
</style>
