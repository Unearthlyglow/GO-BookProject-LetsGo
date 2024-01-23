<script>
  import '@picocss/pico';
  import { onMount } from "svelte";

  let goPayload = [];

  onMount(async function() {
      const response = await fetch('http://localhost:8080/test');
      goPayload  = await response.json();
  });
</script>
<div>
<h1 class="header--title">Music Gear Registry</h1>

<ul>
  {#each goPayload  as t}
  <li><a href="/subdir/{t}">{t}</a></li>
  {/each}
</ul>
</div>

<style lang="scss">
	@import '$lib/styles/global.scss';


  div {
    /* background-color: blue; */
  }

  .header--title {
    color: $primary-red;
    margin-top: 10rem;
    font-size: 5rem;
  }
</style>