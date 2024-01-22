<script>
	// import '@fontsource/lato';
	import yellowBuilding from '$lib/images/yellowBuilding.svg';
	import plant from '$lib/images/planth.svg';
	import chanda from '$lib/images/chanda.svg';
	import table from '$lib/images/table.png';
	import couch from '$lib/images/couch.png';
	import { page } from '$app/stores';
</script>

<svelte:head>
	<!-- <title>CCM | {$page.data.page.title}</title> -->
	<!-- <meta name="description" content={$page.data.page.description} /> -->
</svelte:head>

<div id="body">
	<div class="body--inner">
		<div class="building-logos">
			<img class="building-logo" alt="The project logo" src={yellowBuilding} />
			<img class="building-logo2" alt="The project logo" src={yellowBuilding} />
		</div>

		<hr class="hr--top" />
		<div class="chanda">
			<img alt="chandelier" src={chanda} />
		</div>
		<div class="content">
			<slot />
		</div>

		<p class="fin">フィン</p>

		<hr class="hr--bottom" />

		<hr class="hr--bottom2" />

		<!-- <p class="body--back-to-top"><a href="/{$page.data.page.title}#body">Back to Top</a></p> -->
		<div class="body--home">
			<div><img class="plant" alt="plant" src={plant} /></div>
			<div><img class="couch" alt="table" src={couch} /></div>
			<div><img class="table" alt="table" src={table} /></div>
		</div>
	</div>
</div>


<style lang="scss">
	
	@import '$lib/styles/global.scss';

	.chanda {
		text-align: center;
		transform: translateY(-3rem);
	}

	#body {
		background-color: rgb(46, 46, 51);

		border-bottom: 60px solid rgb(32, 26, 56);
		border-top: 14px solid #fddf92;
		border-left: 2.5px solid #fddf92;
		border-right: 2.5px solid #fddf92;
		min-height: 40rem;
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
	}

	p,
	li,
	a,
	div {
		font-family: 'Lato', sans-serif;
		color: white;
	}

	.content {
		margin-left: 10rem;
		margin-right: 10rem;
	}

	a:hover {
		color: #fac53d;
	}

	button {
		background: none;
	}

	.body--inner {
		margin-top: 1rem;
		margin-left: 1rem;
		margin-right: 1rem;

		/* border-right: 5px #d3a632 solid;
		border-left: 5px #d3a632 solid; */
	}

	.building-logos {
		display: flex;
		justify-content: space-between;
	}
	.building-logo {
		width: 3rem;
		height: 3rem;
	}

	.body--home {
		display: flex;

		justify-content: space-between;
		align-items: flex-end;
	}

	.body--home > img {
		height: 10rem;
		width: 10rem;
	}

	.plant {
		height: 10rem;
		width: 10rem;
		transform: translateX(-2rem);
		transform: translateY(1rem);
	}

	.couch {
		width: 18rem;
		height: 7rem;
		transform: translateY(2.2rem);
	}

	.table {
		width: 10rem;
		height: 10rem;
		transform: translateX(1rem);
		transform: translateY(1rem);
	}

	.building-logo2 {
		width: 3rem;
		height: 3rem;
		-moz-transform: scaleX(-1);
		-o-transform: scaleX(-1);
		-webkit-transform: scaleX(-1);
		transform: scaleX(-1);
		filter: FlipH;
		-ms-filter: 'FlipH';
	}

	.hr--top {
		margin-bottom: 3rem;
		margin-top: 0rem;
		clear: both;
		display: block;
		width: 100%;
		background-color: $primary-white;
		height: 7px;
	}

	.hr--bottom {
		margin-top: 10rem;
		width: 40%;
		background-color: $primary-white;
		height: 7px;
		margin-bottom: 0rem;
	}

	.hr--bottom2 {
		width: 30%;
		background-color: #712828;
		height: 15px;
		margin-top: 0rem;
	}

	.fin {
		transform: translateY(9rem);
		font-weight: 900;
		font-size: 0.9rem;
		color: $primary-white;
		text-align: end;
	}

	.body--back-to-top {
		text-align: center;
		font-size: 1.2rem;
	}

	@media only screen and (min-width: 1800px) {
		.body--inner {
			margin-left: 2rem;
			margin-right: 2rem;
		}
	}

	@media only screen and (max-width: 900px) {
		.body--inner {
			margin-left: 1rem;
			margin-right: 1rem;
		}

		.content {
			margin-left: 1rem;
			margin-right: 1rem;
		}

		.chanda > img {
			width: 5rem;
		}

		#body {
			// margin-top: 2rem;
		}

		.body--home {
			height: 10rem;
		}

		.plant {
			width: 5rem;
			height: 6rem;
		}

		.couch {
			width: 10rem;
			height: 6rem;
		}

		.table {
			width: 5rem;
			height: 5rem;
		}
	}

	@media only screen and (max-width: 500px) {
		#body {
			margin-top: 0rem;
		}
	}
</style>
