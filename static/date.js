// Extract date from URL and format it for display
document.addEventListener("DOMContentLoaded", () => {
    const urlParts = window.location.pathname.split("/");
    const day = urlParts[2];
    const month = urlParts[3];
    const year = urlParts[4];
    const formattedDate = `${day}/${month}/${year}`;

    // Display the formatted date
    document.getElementById("task-date").innerText = formattedDate;

    // Load tasks for the specific date
    loadTasks(day, month, year);
});

// Function to load tasks from the API
async function loadTasks(day, month, year) {
    try {
        const response = await fetch(`/api/tasks/${day}/${month}/${year}`);
        const tasks = await response.json();

        const taskListDiv = document.getElementById("tasks-list");
        taskListDiv.innerHTML = ''; // Clear existing tasks

        tasks.forEach(task => {
            const taskItem = document.createElement("li");
            taskItem.className = "task-item";
            taskItem.innerHTML = `<h4>${task.title}</h4><p>${task.description}</p><p>Type: ${task.task_type}</p>`;
            taskListDiv.appendChild(taskItem);
        });
    } catch (error) {
        console.error("Error fetching tasks:", error);
    }
}

// Handle form submission to add a new task
document.getElementById("new-task-form").addEventListener("submit", function(event) {
    event.preventDefault();

    const urlParts = window.location.pathname.split("/");
    const taskDate = `${urlParts[4]}-${urlParts[3]}-${urlParts[2]}`; // Extract date from URL as YYYY-MM-DD

    const taskData = {
        title: document.getElementById("title").value,
        description: document.getElementById("description").value,
        completed: false,
        task_type: document.getElementById("task_type").value,
        task_date: taskDate,
    };

    fetch("/api/tasks", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(taskData),
    })
    .then(response => response.json())
    .then(data => {
        console.log("Task created:", data);
        loadTasks(taskDate.split("-")[2], taskDate.split("-")[1], taskDate.split("-")[0]); // Reload tasks for the current date
    })
    .catch(error => console.error("Error:", error));
});
