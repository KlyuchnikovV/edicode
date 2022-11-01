<script lang="ts">
	import "@fortawesome/fontawesome-free/css/all.css";
	import { onMount } from "svelte";
	import { EventsOn } from "../wailsjs/runtime/runtime";
	import { GetBufferNames, GetRightPanels } from "../wailsjs/go/api/Api.js";
	import Tabs from "./tabs/tabs.svelte";
	import Flex from "./flex/flex.svelte";
	import LazyLoader from "./lazy-loader/lazy-loader.svelte";
	import Sidepanel from "./sidepanel/sidepanel.svelte";
	import Explorer from "./explorer/explorer.svelte";

	let hasLeft = false;
	let hasRight = true;
	let hasFooter = true;

	let buffers = [];
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

	onMount(async () => {
		buffers = await GetBufferNames();
		console.log(plugins);
	});

	EventsOn(`buffer_opened`, async () => {
		buffers = await GetBufferNames();
	});
	EventsOn(`buffer_closed`, async () => {
		buffers = await GetBufferNames();
	});
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
			<h2>Left</h2>
		</div>

		<div class="center" slot="center">
			<Tabs {buffers} />
		</div>

		<div class="right-panel" slot="right" let:toggle>
			<Sidepanel {items} {toggle} />
		</div>

		<h2 class="footer" slot="footer">Footer</h2>
	</Flex>
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
	}

	h2 {
		text-align: center;
		justify-content: center;

		font-size: 20px;
		margin: 0;
		/* user-select: none;
		-webkit-user-select: none; */
	}

	/* div {
		user-select: none;
		-webkit-user-select: none;
	} */
	/*
	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	} */

	.left {
		/* background-color: blue; */
		width: 100%;
		user-select: none;
	}

	.center {
		/* background-color: green; */
		width: 100%;
		height: 100%;
		/* user-select: none; */
	}

	.right-panel {
		/* background-color: yellow; */
		width: 100%;
		overflow: hidden;
		user-select: none;
		height: 100%;
	}

	.footer {
		/* background-color: black; */
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
		/* border: 3px solid; отступ вокруг бегунка */
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
</style>
