const apiUrl = import.meta.env.VITE_API_URL;

export const getTodos = async (): Promise<object[]> => {
    const response = await fetch(`${apiUrl}/api/todos`)
    const json = await response.json()
    console.log(typeof json)
    return json
}

export const createTodo = async (todo: object): Promise<object> => {
    const response = await fetch(`${apiUrl}/api/todos`, {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        body: JSON.stringify(todo) // body data type must match "Content-Type" header
    })

    return await response.json()
}

export const updateTodo = async (id: string, todo: object): Promise<object> => {
    const response = await fetch(`${apiUrl}/api/todos/${id}`, {
        method: 'PATCH', // *GET, POST, PUT, DELETE, etc.
        body: JSON.stringify(todo) // body data type must match "Content-Type" header
    })

    return await response.json()
}

export const deleteTodo = async (id: string): Promise<object> => {
    const response = await fetch(`${apiUrl}/api/todos/${id}`, {
        method: 'DELETE', // *GET, POST, PUT, DELETE, etc.
    })

    return await response.json()
}
