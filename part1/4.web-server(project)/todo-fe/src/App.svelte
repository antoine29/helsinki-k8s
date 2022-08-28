<script lang="ts">
  import svelteLogo from './assets/svelte.svg'
  import Counter from './lib/Counter.svelte'
  import ToDoList from './lib/ToDoList.svelte';

  const getDailyPicNumber = () => {
    const date = new Date();
    const number = `${date.getDate()}${date.getMonth()+1}`
    return number
  }

  const apiUrl = import.meta.env.VITE_API_URL

  const click = (event) => {console.log(apiUrl)}
</script>


<main class="container">
  <img
    src={`${apiUrl}/image/${getDailyPicNumber()}`}
    alt="Vite Logo"
    on:click={click} />
 
  <h1>Go + Vite + Svelte + TS => ToDo list</h1>
  <!-- <div >
    <a href="https://vitejs.dev" target="_blank"> 
      <img src="/vite.svg" class="logo" alt="Vite Logo" />
    </a>
    <a href="https://svelte.dev" target="_blank"> 
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
  </div> -->

  <!-- <div class="card">
    <Counter />
  </div>
 -->
  <ToDoList />

  <!-- <p>
    Check out <a href="https://github.com/sveltejs/kit#readme" target="_blank">SvelteKit</a>, the official Svelte app framework powered by Vite!
  </p>

  <p class="read-the-docs">
    Click on the Vite and Svelte logos to learn more
  </p> -->
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  /* .read-the-docs {
    color: #888;
  } */
  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;  }
</style>