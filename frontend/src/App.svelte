<script>
	import Codearea from "./codearea/codearea.svelte";
	import { Tabs, Tab, TabList, TabPanel } from "svelte-tabs";
	import { GetBufferNames } from "../wailsjs/go/api/Api.js";

	let getBufferNamesPromise = getBufferNames();

	async function getBufferNames() {
		return await GetBufferNames();
	}
</script>

<main>
	<Tabs>
		<TabList>
			{#await getBufferNamesPromise then bufferNames}
				{#each bufferNames as buffer}
					<Tab>{buffer}</Tab>
				{/each}
			{:catch err}
				{console.log(err)}
			{/await}
		</TabList>
		{#await getBufferNamesPromise then bufferNames}
			{#each bufferNames as buffer (buffer)}
				<TabPanel>
					<Codearea file={buffer} />
				</TabPanel>
			{/each}
		{:catch err}
			{console.log(err)}
		{/await}
	</Tabs>
</main>

<style>
	main {
		text-align: center;
		max-width: 240px;
		margin: 0 auto;
		max-height: 97vh;
		overflow: hidden;
		display: flex;
	}

	:global(.svelte-tabs) {
		max-height: 100%;
		width: 100%;
	}

	:global(.svelte-tabs ul) {
		text-align: left;
		max-height: 100%;
	}
	:global(.svelte-tabs > div) {
		overflow: auto;
		max-height: 93%;
	}

	:global(.svelte-tabs li.svelte-tabs__tab) {
		color: white;
		font: 15px/24px "Fira Code", sans-serif;
	}

	:global(.svelte-tabs li.svelte-tabs__selected) {
		color: #4f81e5;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

	:global(body) {
		/* TODO: to file */
		background-color: #282c34;
		color: #abb2bf;
	}
</style>
