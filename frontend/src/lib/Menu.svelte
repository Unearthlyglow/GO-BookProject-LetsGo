<script>
	import { fly, scale } from 'svelte/transition';
	import { quadOut } from 'svelte/easing';

	/**
	 * @type {any}
	 */
	export let open;
</script>

{#if open}
	<div class="menu">
		<div class="menu--item"><a href="/#body">Home</a></div>
		<div class="menu--item"><a href="/writingmap/#body">Writing</a></div>
		<div class="menu--item"><a href="/projectmap#body">Projects</a></div>
		<div class="menu--item"><a href="/about/#body">About</a></div>
	</div>

	<hr transition:scale={{ duration: 750, easing: quadOut, opacity: 1 }} />
{/if}

<style lang="scss">
	@import '$lib/styles/global.scss';
	.menu {
		display: block;
		margin-top: 2rem;
		text-align: center;
		font-size: 1.5em;
		letter-spacing: 0.15em;
		padding: 1em;
		padding-top: 0;
		height: 15rem;
	}
	div > a {
		color: $primary-white;
	}

	a:hover {
		color: $primary-yellow;
	}

	.menu--item {
		margin: 2rem;
	}

	.menu--item:hover {
		text-decoration: underline;
		color: $primary-red;
	}
</style>
