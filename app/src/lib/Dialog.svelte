<script lang="ts">
  import { setContext } from 'svelte';

  export let title: string = '';
  export let description: string = '';
  export let price: number = 0;

</script>

<div>
  <label for="title">Title</label>
  <input type="text" name="title" id="input-title" bind:value={title}>

  <label for="description">Description</label>
  <input type="text" name="description" id="input-description" bind:value={description}>

  <label for="price">Price</label>
  <input type="number" name="price" id="input-price" bind:value={price}>
</div>

<style>
  * {
    /* outline: 2px solid black; */
    margin: 0 auto;
    z-index: 99;
  }

  input {
    outline: none;
    border: none;
    font-size: 16px;
  }
</style>
