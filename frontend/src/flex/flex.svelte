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
						initial: "0px",
						collapsed: "0px",
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
</div>

<style>
	.flexbox {
		height: 100%;
		width: 100%;
		border-bottom: solid 1px black;
	}

	.flex-container {
		display: flex;
		flex-direction: row;
		height: 100%;
		width: 100%;
		position: relative;
		overflow: hidden;
		max-width: none;
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
		height: 100%;
		border-top: none;
	}

	.center {
		height: 100%;
		overflow: hidden;
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
		bottom: 0px;
		user-select: none;
		-webkit-user-select: none;
	}
</style>
