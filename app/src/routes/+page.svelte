<script lang="ts">
  // your script goes here
  import type { IDebt } from '../@types';
  let debtList: IDebt[] = [];

  let titleData = '';
  let descriptionData = '';
  let priceData = 0;

  let addDialog = false;

  function addNewNote() {
    addDialog = true;
  }

  function validateDialog(title: string, price: number): boolean {
    if(title === '') {
      alert('Please fill all required fields');
      return false
    }

    if(price <= 0) {
      alert('Price cannot be 0 or less than 0');
      return false
    }

    return true;
  }

  function submit() {
    if(!validateDialog(titleData, priceData)) {
      return;
    }

    const newDebt: IDebt = {
      title: titleData,
      description: descriptionData,
      price: priceData,
      createdAt: new Date().toLocaleDateString(),
    };

    debtList = [...debtList, newDebt];
    closeDialog();
  }

  function closeDialog() {
    addDialog = false;
  }

</script>

<svelte:head>
  <title>Nhee Note | บันทึกหนี้</title>
</svelte:head>

<div>
  <header>
    <h1>Nhee Note</h1>
  </header>
  <main>
    <section id="debt-section">
      <!-- Create Table -->
      <table>
        <tr>
          <th>ID</th>
          <th>Title</th>
          <th>Description</th>
          <th>Price</th>
          <th>Created At</th>
        </tr>
        {#each debtList as debt, index}
        <tr>
          <td>{index + 1}</td>
          <td>{debt.title}</td>
          <td>{debt.description}</td>
          <td>{debt.price}</td>
          <td>{debt.createdAt}</td>
        </tr>
        {/each}
      </table>
    </section>
    {#if addDialog}
    <section id="dialog-section">
      <div id="dialog-container">
        <div>
          <label for="title">Title</label>
          <input type="text" name="title" id="input-title" bind:value={titleData}>

          <label for="description">Description</label>
          <input type="text" name="description" id="input-description" bind:value={descriptionData}>

          <label for="price">Price</label>
          <input type="number" name="price" id="input-price" bind:value={priceData}>
        </div>
        <div id="dialog-action">
          <button on:click={submit}>Submit</button>
          <button on:click={closeDialog}>Close</button>
        </div>
      </div>
    </section>
    {/if}
    <section id="action-section">
      <button id="add-new-note-button" on:click={addNewNote}>Add New Note</button>
    </section>
  </main>
  <footer>
    <p>Copyright 2023, Made by <a href="https://github.com/lebrancconvas">Poom Yimyuean (@lebrancconvas)</a></p>
  </footer>
</div>

<style>
  /* your styles go here */
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Open Sans', sans-serif;
  }

  header {
    text-align: center;
    margin-bottom: 20px;
  }

  main {
    text-align: center;
  }

  table {
    margin: 0 auto;
    border: 1px solid black;
    margin-bottom: 10px;
    padding: 7px;
  }

  button {
    outline: 0;
    border: 0;
    border-radius: 10px;
    padding: 8px;
    cursor: pointer;
  }

  button:active {
    transform: scale(0.98);
  }

  input {
    outline: none;
    border: none;
    font-size: 16px;
  }

  #add-new-note-button {
    background-color: rgb(255, 164, 37);
    font-size: 14px;
  }

  #dialog-container {
    outline: 2px solid black;
    z-index: 99;
    margin: 0 auto;
    width: 50%;
    height: 100px;
    box-shadow: 0 0 0 9999px rgba(0, 0, 0, 0.5);
  }

  footer {
    text-align: center;
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 50px;
    background-color: white;
  }
</style>

<!-- markup (zero or more items) goes here -->
