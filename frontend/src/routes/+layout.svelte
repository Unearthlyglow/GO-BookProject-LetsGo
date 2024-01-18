<script>
	import '@picocss/pico';


	// import metropolis_building from '$lib/images/metropolis2.svg';


</script>

<svelte:head>
	<meta name="google" content="notranslate" />
	<meta title="Muse DB" />
</svelte:head>

<div>
		<slot />
</div>

<style lang="scss">
	$primary-color: #ec3232;

	div {
		background-color: $primary-color;
	}

</style>

