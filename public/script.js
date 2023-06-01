const todoForm = document.getElementById("todo-form")
const todoInput = document.getElementById("todo-input")
const todoList = document.getElementById("todo-list")

function renderTodos(todos) {
  todoList.innerHTML = ""

  todos.forEach((todo) => {
    const listItem = document.createElement("li")
    if (todo.completed) {
      listItem.classList.add(`item${todo.ID}`)
    }
    listItem.innerHTML = `
    <div class="title ${
      todo.completed ? "completed" : ""
    }" onclick="toggleCompleted(${todo})">${todo.title}</div>
    <span class="delete" onclick="deleteTodo(${todo.ID})">❌</span>
        `
    todoList.appendChild(listItem)
  })
}

async function fetchTodos() {
  try {
    const response = await fetch("/todos")
    const todos = await response.json()
    renderTodos(todos)
  } catch (error) {
    console.error(error)
  }
}

async function createTodo() {
  const title = todoInput.value.trim()
  if (title === "") return

  try {
    const response = await fetch("/todos", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title, completed: false }),
    })

    if (response.ok) {
      const todos = await response.json()
      renderTodos(todos)
      todoInput.value = ""
    }
  } catch (error) {
    console.error(error)
  }
}

async function toggleCompleted(todo) {
  try {
    const response = await fetch(`/todos/${todo.ID}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ completed: !todo.completed }),
    })

    if (response.ok) {
      const todo = await response.json()
      console.log("todo is", todo)
      const listItem = document.querySelector(`.title${todo.ID}`)
      console.log("listItem", listItem)
    }
  } catch (error) {
    console.error(error)
  }
}

async function deleteTodo(id) {
  try {
    const response = await fetch(`/todos/${id}`, {
      method: "DELETE",
    })

    if (response.ok) {
      const todos = await response.json()
      renderTodos(todos)
    }
  } catch (error) {
    console.error(error)
  }
}

todoForm.addEventListener("submit", (event) => {
  event.preventDefault()
  createTodo()
})

fetchTodos()
