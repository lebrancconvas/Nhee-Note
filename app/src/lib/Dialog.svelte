<script lang="ts">
  import { title, description, price } from '../store';

  let titleData = '';
  let descriptionData = '';
  let priceData = 0;

  title.set(titleData);
  description.set(descriptionData);
  price.set(priceData);
</script>

<div>
  <label for="title">Title</label>
  <input type="text" name="title" id="input-title" bind:value={titleData}>

  <label for="description">Description</label>
  <input type="text" name="description" id="input-description" bind:value={descriptionData}>

  <label for="price">Price</label>
  <input type="number" name="price" id="input-price" bind:value={priceData}>
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
