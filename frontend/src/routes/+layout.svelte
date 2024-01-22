<script>
	import '@picocss/pico';
	import Footer from '$lib/Footer.svelte';
	import Body from '$lib/Body.svelte';
	import Nav from '$lib/Nav.svelte';
	// import metropolis_building from '$lib/images/metropolis2.svg';


</script>

<svelte:head>
	<meta name="google" content="notranslate" />
	<meta title="Muse DB" />
</svelte:head>

<div class="container-fluid">
	<!-- <Nav /> -->
	<Body>
		<Nav />
		<slot />
	</Body>
		<Footer />
</div>

<style lang="scss">
	@import '$lib/styles/global.scss';

	.container-fluid {
		min-height: 100%;
		min-width: 100%;
		display: grid;
		grid-template-columns: repeat(auto-fit, 1fr);
		padding-right: 0px;
		padding-left: 0px;
		background-color: $primary-white;
	}
	

	
	
	$primary-red: #8C2525;

	div {
		background-color: $primary-red;
	}

</style>

