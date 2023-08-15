<script>
  // your script goes here
  let debtList = [
    {
      title: '',
      description: '',
      price: 0,
      createdAt: ''
    }
  ];
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
    <section id="action-section">
      <button id="add-new-note-button">Add New Note</button>
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

  #add-new-note-button {
    background-color: rgb(255, 164, 37);
    font-size: 14px;
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
