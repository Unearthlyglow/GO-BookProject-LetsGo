<script>
  import { onMount } from "svelte";
  let data = [];

  onMount(async function() {
  try {
    const response = await fetch('http://localhost:8080/test');
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const responseData = await response.text();
    console.log('Response Data:', responseData);  // Log the response data
    console.log('Looking over the Payload:', responseData);
    data = JSON.parse(responseData);
    
  } catch (error) {
    console.error('Error fetching data:', error.message); 
  }
 
});
</script>

<div>
  <h1 class="header--title">Music Gear Registry</h1>

 <ul>
    {#each data as item (`${item.ID}-${item.AnotherProperty}`)}
      <li>{item.ID}: {item.Content}</li>
      <!-- Replace 'ColumnName' with the actual column name you want to display -->
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