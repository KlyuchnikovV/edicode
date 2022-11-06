<script lang="ts">
	import "@fortawesome/fontawesome-free/css/all.css";
	import Tabs from "./tabs/tabs.svelte";
	import Flex from "./flex/flex.svelte";
	import Explorer from "./explorer/explorer.svelte";
	import Sidepanel from "./sidepanel/sidepanel.svelte";
	import LazyLoader from "./lazy-loader/lazy-loader.svelte";
	import { EventsOn } from "../wailsjs/runtime/runtime";
	import { GetRightPanels } from "../wailsjs/go/api/Api.js";

	let hasLeft = true;
	let hasRight = true;
	let hasFooter = true;

	let plugins = [];
	let items = [
		{
			component: Explorer,
			argument: {
				name: "explorer",
				icon: "manage_search",
			},
		},
	];

	EventsOn(`plugins_loaded_all`, async () => {
		console.log("plugins all loaded");
		plugins = await GetRightPanels();

		console.log(plugins);

		plugins.forEach((element) => {
			items.push({
				component: LazyLoader,
				argument: element,
			});
		});

		items = items;

		console.log(items);
	});
</script>

<main>
	<Flex bind:hasLeft bind:hasRight bind:hasFooter>
		<div class="left" slot="left">
			<h2 class="left-panel">Left</h2>
		</div>

		<div class="center" slot="center">
			<Tabs />
		</div>

		<div class="right-panel" slot="right" let:toggle>
			<Sidepanel {items} {toggle} />
		</div>
		<h2 class="bottom-panel" slot="footer">Bottom-bar</h2>
	</Flex>
	<footer class="footer">Footer</footer>
</main>

<style>
	main {
		text-align: center;
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		overflow: hidden;
		height: 100%;
		display: flex;
		flex-direction: column;
		user-select: none;
		-webkit-user-select: none;
	}

	h2 {
		text-align: center;
		justify-content: center;

		font-size: 20px;
		margin: 0;
	}

	.left {
		width: 100%;
		height: 100%;
		user-select: none;
	}

	.center {
		width: 100%;
		height: 100%;
	}

	.right-panel {
		width: 100%;
		overflow: hidden;
		user-select: none;
		height: 100%;
		border-left: solid 1px black;
	}

	.left-panel {
		width: 100%;
		overflow: hidden;
		user-select: none;
		height: 100%;
		border-right: solid 1px black;
	}

	.bottom-panel {
		bottom: 0;
		width: 100%;
		border-top: solid 1px black;
	}

	.footer {
		bottom: 0;
		width: 100%;
	}

	:global(::-webkit-scrollbar) {
		width: 12px; /* ширина всей полосы прокрутки */
	}

	:global(::-webkit-scrollbar-track) {
		background: inherit; /*#282c34;*/ /* цвет зоны отслеживания */
	}

	:global(::-webkit-scrollbar-thumb) {
		background-color: #504e56; /* цвет бегунка */
		border-radius: 3px; /* округлось бегунка */
	}

	:global(::-webkit-scrollbar-corner) {
		background: inherit; /*#282c34;*/
	}

	:global(.material-icons-outlined.md-18) {
		font-size: 18px;
	}
	:global(.material-icons-outlined.md-24) {
		font-size: 24px;
	}
	:global(.material-icons-outlined.md-36) {
		font-size: 36px;
	}
	:global(.material-icons-outlined.md-48) {
		font-size: 48px;
	}

	/* Overscroll of html element */
	:global(html) {
		overflow: hidden;
		height: 100%;
	}
	:global(body) {
		height: 100%;
	}
</style>
