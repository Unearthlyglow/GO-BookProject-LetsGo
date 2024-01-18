<script>
  import '@picocss/pico';
  import { onMount } from "svelte";

  let test = [];

  onMount(async function() {
      const response = await fetch('http://localhost:8080/test');
      test = await response.json();
  });
</script>
<div>
<h1>Combining SvelteKit and Golang server</h1>

<ul>
  {#each test as t}
  <li><a href="/subdir/{t}">{t}</a></li>
  {/each}
</ul>
</div>

<style>

  div {
    /* background-color: blue; */
  }
</style>