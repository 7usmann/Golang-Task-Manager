package handlers

import (
	"net/http"
)

// HomePageHandler handles the root page.
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
        <html>
            <head>
                <title>Golang Task Manager</title>
            </head>
            <body>
                <h1>Welcome to the Golang Task Manager</h1>
                <ul>
                    <li><a href="/tasks">View All Tasks</a></li>
                    <li><a href="/tasks/new">Create a New Task</a></li>
                    <li><a href="/tasks/{id}">View a Specific Task</a></li>
                </ul>
            </body>
        </html>
    `))
}
