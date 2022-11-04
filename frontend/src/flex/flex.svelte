<script lang="ts">
	import Resizable from "./resizable/resizable.svelte";

	export let hasLeft = true;
	export let hasRight = true;
	export let hasFooter = true;
</script>

<div class="flexbox">
	<div class="flex-container">
		<div class="left" style:display={hasLeft ? "flex" : "none"}>
			<Resizable
				config={{
					position: "right",
					height: {
						min: "100%",
						initial: "100%",
						collapsed: "100%",
					},
					width: {
						min: "50px",
						initial: "0",
						collapsed: "0",
					},
				}}
			>
				<slot name="left" />
			</Resizable>
		</div>

		<div class="flex-container center">
			<div class="center">
				<slot name="center" />
			</div>
			<div class="bottom" style:display={hasFooter ? "flex" : "none"}>
				<Resizable
					config={{
						position: "top",
						height: {
							min: "50px",
							initial: "150px",
							collapsed: "0px",
						},
						width: {
							min: "100%",
							initial: "100%",
							collapsed: "100%",
						},
					}}
					let:toggle
				>
					<slot name="footer" />
				</Resizable>
			</div>
		</div>

		<div
			class="right"
			style:display={hasRight ? "flex" : "none"}
			style="border-top: none"
		>
			<Resizable
				config={{
					position: "left",
					width: {
						min: "150px",
						initial: "350px",
						collapsed: "3em",
					},
					height: {
						min: "100%",
						initial: "100%",
						collapsed: "3em",
					},
				}}
				let:toggle
			>
				<slot name="right" {toggle} />
			</Resizable>
		</div>
	</div>

	<!-- <footer style="display:{hasFooter ? 'flex' : 'none'}"> -->
	<!-- <Resizable position="top" width="350px" minWidth="150px" let:toggle>
		<slot name="footer" />
	</Resizable> -->
	<!-- <slot name="footer" /> -->
	<!-- </footer> -->
</div>

<style>
	.flexbox {
		height: 100%;
		width: 100%;
		/* border: solid 1px black; */
		/* display: flex; */
		/* flex-direction: column;
		align-items: stretch; */

		border-bottom: solid 1px black;
	}

	.flex-container {
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
		/* border-top: none; */
		/* border: solid 1px black; */
	}

	.right {
		right: 0;
		user-select: none;
		-webkit-user-select: none;
		height: 100%;
		/* border-top: none; */
		/* border-left: solid 1px black; */
	}

	.left {
		left: 0;
		user-select: none;
		-webkit-user-select: none;
		height: 100%;
		border-top: none;
		/* border-right: solid 1px black; */
	}

	.center {
		height: 100%;
		/* width: 100%;
		display: flex; */
		/* user-select: none;
		-webkit-user-select: none; */
		overflow: hidden;
		/* border-bottom: 1px solid black; */
		flex-direction: column;
		align-items: stretch;
	}

	.flex-container > ::slotted {
		height: 100%;
		width: 100%;
		overflow: hidden;
		border-bottom: 1px solid black;
		flex: 1;
	}

	.bottom {
		/* footer { */
		/* flex: 1; */
		bottom: 0px;
		user-select: none;
		-webkit-user-select: none;
		/* border-top: solid 1px black; */
		/* height: 25px; */
	}
</style>
