<nav class="container-tabs">
	<li>
		<a href="/#body" class="single--tab"> Home</a>
	</li>
	<li class="class-divide">||</li>
	<li role="list" dir="ltr">
		<p class="tab-intact" aria-haspopup="listbox">Writing</p>
		<ul role="listbox">
			<li class="manifesto--test"><a href="/jazz-manifesto/#body">Jazz Manifesto</a></li>

			<li>
				<a href="https://unearthlyglow.github.io/RustMDBooks-CCMBlog/" target="_blank"
					>CCM General Blog(MD Book)</a
				>
			</li>
		</ul>
	</li>
	<li class="class-divide">||</li>
	<li role="list" dir="rtl">
		<p class="tab-intact" aria-haspopup="listbox">Projects</p>
		<ul role="listbox">
			<li><a href="/softwaremap/#body">Software Portfolio</a></li>
			<li><a href="/visual/#body">Visual Media</a></li>
			<!-- <li><a href="/metro/#body">Metropolis Design</a></li> -->
			<li><a href="/art/#body">Art I've Enjoyed</a></li>
			<li><a href="/embedded/#body">The Road To Embedded Programming</a></li>
		</ul>
	</li>
	<li class="class-divide">||</li>
	<li role="list" dir="rtl">
		<p class="tab-intact" aria-haspopup="listbox">About</p>
		<ul role="listbox">
			<li><a href="/about/#body">About Me</a></li>
			<!-- <li><a href="/about/resume">Resume</a></li> -->
			<li><a href="/todos/#body">Website To-Do List!</a></li>
		</ul>
	</li>
</nav>

<style lang="scss">
	@import '$lib/styles/global.scss';

	nav {
		background-color: $primary-white;
		color: white;
	}

	.single--tab {
		font-weight: 900;
	}

	.container-tabs {
		justify-content: center;

		& a {
			color: rgb(13, 81, 90);
			font-size: 1.5rem;
		}
	}

	.tab-intact {
		color: rgb(13, 81, 90);
		font-weight: 900;
		margin: 0;
		font-size: 1.5rem;
	}

	.tab-intact:hover {
		color: #942bef;
	}

	.class-divide {
		color: #8c2525;
	}

	ul[role='listbox'] {
		background-color: $primary-white;
		box-shadow: 10px 5px 5px rgba(0, 0, 0, 0.98);
		border: none;
		text-align: center;

		& a {
			font-size: 0.87rem;
			color: $primary-green;
		}
	}

	ul[role='listbox']:hover {
		background-color: $primary-white;
		transition: background-color 0.5s;

		& a {
			color: $primary-green;
		}

		& a:hover {
			text-decoration: underline;
			color: $primary-yellow;
		}
	}

	a:hover {
		color: #d60675;
	}

	@media only screen and (max-width: 548px) {
		.container-tabs {
			display: none;
		}
	}
</style>
