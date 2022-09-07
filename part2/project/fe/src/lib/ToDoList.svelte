<script>
  import { onMount } from "svelte";
  import { getTodos, createTodo, updateTodo, deleteTodo } from './services'

  let newTodoContent = "";

  let todoList = [];

  onMount(async () => {
    const response = await getTodos()
    todoList = response
  });

  async function addToList () {
    const newTodo = { Id: "", Content: newTodoContent, IsDone: false }
    const createdTodo = await createTodo(newTodo)
    console.log("created todo:", createdTodo)
    todoList = [...todoList, createdTodo]
    newTodoContent = "";
  }

  async function _updateTodo(index) {
    const todo = todoList[index]
    const updatedTodo = await updateTodo(todo.Id, {Content: todo.Content, IsDone: todo.IsDone})
    todoList[index] = updatedTodo
  }

  async function _deleteTodo(index) {
    const todo = todoList[index]
    await deleteTodo(todo.Id)
    
    todoList = todoList.filter(_todo => _todo.Id !== todo.Id)
  }
</script>

<div class="list">
  <input bind:value={newTodoContent} type="text" placeholder="new todo item.." />
  <button on:click={addToList}>Add</button>

  <br />
  {#each todoList as item, index}
    <input
      bind:checked={item.IsDone}
      type="checkbox"
      on:change={()=>{_updateTodo(index)}} />
    <span class:checked={item.IsDone}>{item.Content}</span> 
    <span on:click={() => _deleteTodo(index)}>❌</span>
    <br />
  {/each}

  <!-- {JSON.stringify(data, null, 2)} -->
</div>

<style>
  .checked {
    text-decoration: line-through;
  }
  .list {
    text-align: left;
  }
</style>
