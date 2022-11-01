<script lang="ts">
	import Resizable from "./resizable/resizable.svelte";

	export let hasLeft = true;
	export let hasRight = true;
	export let hasFooter = true;
</script>

<div class="flexbox">
	<div class="container">
		<div class="left" style="display:{hasLeft ? 'flex' : 'none'}">
			<Resizable position="right">
				<slot name="left" />
			</Resizable>
		</div>

		<div class="center">
			<slot name="center" />
		</div>

		<div class="right" style="display:{hasRight ? 'flex' : 'none'}">
			<Resizable
				position="left"
				width="350px"
				minWidth="150px"
				let:toggle
			>
				<slot name="right" {toggle} />
			</Resizable>
		</div>
	</div>

	<footer style="display:{hasFooter ? 'flex' : 'none'}">
		<slot name="footer" />
	</footer>
</div>

<style>
	.flexbox {
		height: 100%;
		width: 100%;
		/* border: solid 1px black; */
		display: flex;
		flex-direction: column;
		align-items: stretch;
	}

	.container {
		display: flex;
		flex-direction: row;
		height: 100%;
		width: 100%;
		/* user-select: none;
		-webkit-user-select: none; */
		position: relative;
		overflow: hidden;
		max-width: none;
	}

	div > div > :not(.center) {
		border-top: none;
		border: solid 1px black;
	}

	.right {
		right: 0;
		user-select: none;
		-webkit-user-select: none;
		height: 100%;
	}

	.left {
		left: 0;
		user-select: none;
		-webkit-user-select: none;
	}

	.center {
		width: 100%;
		display: block;
		/* user-select: none;
		-webkit-user-select: none; */
		overflow: hidden;
		border-bottom: 1px solid black;
	}

	footer {
		flex: 3;
		bottom: 0;
		height: 25px;
	}
</style>
