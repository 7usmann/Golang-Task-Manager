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

async function loadTasks(day, month, year) {
    try {
        const response = await fetch(`/api/tasks/${day}/${month}/${year}`);
        const tasks = await response.json();

        const taskListDiv = document.getElementById("tasks-list");
        taskListDiv.innerHTML = ''; // Clear existing tasks

        tasks.forEach(task => {
            const taskItem = document.createElement("li");
            taskItem.className = "task-item";
            taskItem.innerHTML = `
                <h4>${task.title}</h4>
                <p>${task.description}</p>
                <p>Type: ${task.task_type}</p>
                <button onclick="deleteTask(${task.id})">Delete</button>
                <button onclick="showUpdateForm(${task.id}, '${task.title}', '${task.description}', '${task.task_type}')">Update</button>
            `;
            taskListDiv.appendChild(taskItem);
        });
    } catch (error) {
        console.error("Error fetching tasks:", error);
    }
}

document.getElementById("new-task-form").addEventListener("submit", function(event) {
    event.preventDefault();
    console.log("Form submitted");  // Debugging line
    console.log("Form submitted");  // Debugging line

    const urlParts = window.location.pathname.split("/");
    const day = urlParts[2];
    const month = urlParts[3];
    const year = urlParts[4];
    const taskDate = `${year}-${month}-${day}`; // Format date as YYYY-MM-DD

    const taskData = {
        title: document.getElementById("title").value,
        description: document.getElementById("description").value,
        completed: false,
        task_type: document.getElementById("task_type").value,
        task_date: taskDate,
    };

    fetch(`/api/tasks/${day}/${month}/${year}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(taskData),
    })
    .then(response => {
        if (response.ok) {
            alert("Task created successfully");
            loadTasks(day, month, year); // Reload tasks for the current date
            document.getElementById("new-task-form").reset(); // Clear the form
        } else {
            alert("Failed to create task");
        }
        return response.json();
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Error creating task");
    });
});

async function deleteTask(taskId) {
    const confirmed = confirm("Are you sure you want to delete this task?");
    if (!confirmed) return;

    try {
        const response = await fetch(`/api/tasks/${taskId}`, { method: "DELETE" });
        if (response.ok) {
            alert("Task successfully deleted");
            const urlParts = window.location.pathname.split("/");
            loadTasks(urlParts[2], urlParts[3], urlParts[4]); // Reload tasks
        } else {
            alert("Failed to delete task");
        }
    } catch (error) {
        console.error("Error deleting task:", error);
        alert("Error deleting task");
    }
}

function showUpdateForm(taskId, title, description, taskType) {
    const updateFormContainer = document.getElementById("update-form-container");
    updateFormContainer.innerHTML = `
        <h3>Update Your Task</h3>
        <form id="update-task-form">
            <label for="update-title">Title:</label>
            <input type="text" id="update-title" name="title" value="${title}" required><br>
            
            <label for="update-description">Description:</label>
            <input type="text" id="update-description" name="description" value="${description}"><br>
            
            <label for="update-task_type">Task Type:</label>
            <select id="update-task_type" name="task_type" required>
                <option value="academia" ${taskType === "academia" ? "selected" : ""}>Academia</option>
                <option value="work" ${taskType === "work" ? "selected" : ""}>Work</option>
                <option value="extracurricular" ${taskType === "extracurricular" ? "selected" : ""}>Extracurricular</option>
                <option value="other" ${taskType === "other" ? "selected" : ""}>Other</option>
            </select><br>
            
            <button type="button" onclick="updateTask(${taskId})">Confirm</button>
        </form>
    `;
}

async function updateTask(taskId) {
    const updatedTaskData = {
        title: document.getElementById("update-title").value,
        description: document.getElementById("update-description").value,
        task_type: document.getElementById("update-task_type").value,
    };

    try {
        const response = await fetch(`/api/tasks/${taskId}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(updatedTaskData),
        });

        if (response.ok) {
            alert("Task successfully updated");
            const urlParts = window.location.pathname.split("/");
            loadTasks(urlParts[2], urlParts[3], urlParts[4]); // Reload tasks
            document.getElementById("update-form-container").innerHTML = ""; // Clear the update form
        } else {
            alert("Failed to update task");
        }
    } catch (error) {
        console.error("Error updating task:", error);
        alert("Error updating task");
    }
}


