function goToMainCalendar() {
    currentView = 'month';
    renderMonthView(new Date());
    document.getElementById('view-switch').innerText = 'Month';
    renderNavigationBar();
}




let currentDate = new Date();
let currentView = 'month';

// Reset currentDate to today whenever switching views

// Reset currentDate to today whenever switching views

// Navigation function to update date based on view mode and direction
function switchView() {
    currentDate = new Date(); // Reset to today each time view changes

    if (currentView === 'month') {
        currentView = 'week';
        renderWeekView(getStartOfWeek(currentDate));
        document.getElementById('view-switch').innerText = 'Week';
    } else if (currentView === 'week') {
        currentView = 'day';
        renderDayView(currentDate);
        document.getElementById('view-switch').innerText = 'Day';
    } else {
        currentView = 'month';
        refreshCalendarGrid(); // Fetch tasks when switching back to month view
        document.getElementById('view-switch').innerText = 'Month';
    }
    renderNavigationBar();
}



document.addEventListener("DOMContentLoaded", async () => {
    const currentDate = new Date();
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth() + 1;

    try {
        currentTasks = await fetchTasksForMonth(year, month);
        console.log("Fetched tasks for month:", currentTasks);
        renderMonthView(currentDate, currentTasks);
    } catch (error) {
        console.error("Error fetching tasks for month:", error);
    }
});

async function fetchTasksForMonth(year, month) {
    console.log(`Fetching tasks from URL: /api/tasks/month/${year}/${month}`);
    try {
        const response = await fetch(`/api/tasks/month/${year}/${month}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error("Error fetching tasks:", error);
        return [];
    }
}

async function refreshCalendarGrid() {1
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth() + 1;
    const timestamp = new Date().getTime(); // Cache-busting timestamp

    try {
        const tasks = await fetchTasksForMonth(year, month);
        renderMonthView(currentDate, tasks);
    } catch (error) {
        console.error("Error refreshing calendar grid:", error);
    }
}



// Modified `renderMonthView` to accept tasks for easier rendering
function renderMonthView(date, tasks = []) {
    // Ensure tasks is an array, even if null or undefined
    tasks = tasks || [];

    const calendarGrid = document.getElementById('calendar-grid');
    calendarGrid.innerHTML = ''; // Clear existing content
    calendarGrid.className = 'calendar-grid'; // Ensure the grid layout is applied

    const year = date.getFullYear();
    const month = date.getMonth();
    const daysInMonth = new Date(year, month + 1, 0).getDate();

    // Get the day of the week for the 1st of the current month
    const firstDayOfMonth = new Date(year, month, 1).getDay();
    const offset = (firstDayOfMonth === 0) ? 6 : firstDayOfMonth - 1; // Offset for Monday start

    // Fill empty cells before the first day of the month
    for (let i = 0; i < offset; i++) {
        const emptyCell = document.createElement('div');
        calendarGrid.appendChild(emptyCell);
    }

    // Fill the calendar with days of the month and add tasks to their relevant cells
    for (let day = 1; day <= daysInMonth; day++) {
        const dayCell = document.createElement('div');
        dayCell.classList.add('day-cell');
        dayCell.innerHTML = `<span class="day-number">${day}</span>`;

        // Filter tasks for this specific day
        const dayTasks = tasks.filter(task => {
            const taskDate = new Date(task.task_date);
            return (
                taskDate.getDate() === day &&
                taskDate.getMonth() === month &&
                taskDate.getFullYear() === year
            );
        });

        // Create a div to hold task titles
        const tasksDiv = document.createElement('div');
        tasksDiv.className = 'tasks-container';

        // Display each task title in a smaller div within the tasks container
        dayTasks.forEach(task => {
            const taskDiv = document.createElement('div');
            taskDiv.className = 'task-title';
            taskDiv.textContent = task.title;
            tasksDiv.appendChild(taskDiv);
        });

        dayCell.appendChild(tasksDiv);

        // Make each cell clickable to navigate to the specific date
        dayCell.onclick = function () {
            const formattedDate = `${String(day).padStart(2, '0')}/${String(month + 1).padStart(2, '0')}/${year}`;
            window.location.href = `/date/${formattedDate}`;
        };

        // Highlight today's date
        if (
            day === new Date().getDate() &&
            month === new Date().getMonth() &&
            year === new Date().getFullYear()
        ) {
            dayCell.classList.add('today');
        }
        calendarGrid.appendChild(dayCell);
    }
}

// Function to render the week view
async function renderWeekView(startOfWeek) {
    const tasks = await fetchTasksForWeek(startOfWeek);
    const calendarGrid = document.getElementById('calendar-grid');
    calendarGrid.innerHTML = ''; // Clear grid

    for (let i = 0; i < 7; i++) {
        const dayDate = new Date(startOfWeek);
        dayDate.setDate(startOfWeek.getDate() + i);

        const dayFormatted = `${String(dayDate.getDate()).padStart(2, '0')}/${String(dayDate.getMonth() + 1).padStart(2, '0')}/${dayDate.getFullYear()}`;

        const dayCell = document.createElement('div');
        dayCell.classList.add('day-cell');
        dayCell.innerHTML = `<span class="day-number">${dayFormatted}</span>`;

        // Event listener for redirection
        dayCell.addEventListener('click', function () {
            const formattedDate = `${String(dayDate.getDate()).padStart(2, '0')}/${String(dayDate.getMonth() + 1).padStart(2, '0')}/${dayDate.getFullYear()}`;
            console.log(`Redirecting to: /date/${formattedDate}`);
            window.location.href = `/date/${formattedDate}`;
        });

        // Filter tasks for this specific day
        const dayTasks = tasks.filter(task => {
            const taskDate = new Date(task.task_date);
            return (
                taskDate.getDate() === dayDate.getDate() &&
                taskDate.getMonth() === dayDate.getMonth() &&
                taskDate.getFullYear() === dayDate.getFullYear()
            );
        });

        const tasksDiv = document.createElement('div');
        tasksDiv.className = 'tasks-container';

        // Display each task title in a smaller div within the tasks container
        dayTasks.forEach(task => {
            const taskDiv = document.createElement('div');
            taskDiv.className = 'task-title';
            taskDiv.textContent = task.title;
            tasksDiv.appendChild(taskDiv);
        });

        

        dayCell.appendChild(tasksDiv);
        calendarGrid.appendChild(dayCell);
    }
}




// Function to render the day view
async function renderDayView(date) {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();

    let tasks = await fetchTasksForDay(year, month, day);
    tasks = tasks || [];  // Default to an empty array if tasks is null or undefined

    const calendarGrid = document.getElementById('calendar-grid');
    calendarGrid.innerHTML = ''; // Clear existing content

    const dayFormatted = `${String(day).padStart(2, '0')}/${String(month).padStart(2, '0')}/${year}`;

    const dayCell = document.createElement('div');
    dayCell.classList.add('day-cell');
    dayCell.innerHTML = `<span class="day-number">${dayFormatted}</span>`;

    const tasksDiv = document.createElement('div');
    tasksDiv.className = 'tasks-container';

    // Add click event to the day cell for redirection
    dayCell.onclick = function () {
        console.log(`Redirecting to /date/${dayFormatted}`);
        window.location.href = `/date/${dayFormatted}`;
    };

    // Populate tasks within the day cell
    tasks.forEach(task => {
        const taskDiv = document.createElement('div');
        taskDiv.className = 'task-title';
        taskDiv.textContent = task.title;
        tasksDiv.appendChild(taskDiv);
    });

    // Highlight today's date if it matches
    const today = new Date();
    if (date.getDate() === today.getDate() && date.getMonth() === today.getMonth() && date.getFullYear() === today.getFullYear()) {
        dayCell.classList.add('today');
    }

    dayCell.appendChild(tasksDiv);
    calendarGrid.appendChild(dayCell);
}



// Helper to get the start of the week (Monday)
function getStartOfWeek(date) {
    const startOfWeek = new Date(date);
    const day = startOfWeek.getDay();
    const diff = startOfWeek.getDate() - day + (day === 0 ? -6 : 1); // Adjust to Monday start
    startOfWeek.setDate(diff);
    return startOfWeek;
}

// Render navigation bar with month/year and navigation buttons
function renderNavigationBar() {
    const calendarHeader = document.getElementById('calendar-header');
    console.log('Calendar Header:', calendarHeader); // This should log the element or null
    if (!calendarHeader) {
        console.error("calendar-header element not found.");
        return; // Early return if element is not found
    }
    calendarHeader.innerHTML = '';
    
    const monthYearDisplay = document.createElement('span');
    monthYearDisplay.id = 'calendar-month';
    monthYearDisplay.innerText = getMonthYearDisplay(currentDate);
    calendarHeader.appendChild(monthYearDisplay);

    const prevButton = document.createElement('button');
    prevButton.innerText = '❮';
    prevButton.className = 'nav-button';
    prevButton.onclick = () => navigateView(-1);
    calendarHeader.insertBefore(prevButton, monthYearDisplay);

    const nextButton = document.createElement('button');
    nextButton.innerText = '❯';
    nextButton.className = 'nav-button';
    nextButton.onclick = () => navigateView(1);
    calendarHeader.appendChild(nextButton);
}

function navigateView(direction) {
    if (currentView === 'month') {
        currentDate.setMonth(currentDate.getMonth() + direction);
        refreshCalendarGrid();
    } else if (currentView === 'week') {
        currentDate.setDate(currentDate.getDate() + direction * 7);
        renderWeekView(getStartOfWeek(currentDate));
    } else if (currentView === 'day') {
        currentDate.setDate(currentDate.getDate() + direction);
        renderDayView(currentDate);
    }
    renderNavigationBar();
}

// Helper function to format the month and year display
function getMonthYearDisplay(date) {
    return date.toLocaleDateString(undefined, { year: 'numeric', month: 'long' });
}

async function fetchTasksForWeek(startDate) {
    const endDate = new Date(startDate);
    endDate.setDate(startDate.getDate() + 6);

    const start = `${startDate.getFullYear()}-${String(startDate.getMonth() + 1).padStart(2, '0')}-${String(startDate.getDate()).padStart(2, '0')}`;
    const end = `${endDate.getFullYear()}-${String(endDate.getMonth() + 1).padStart(2, '0')}-${String(endDate.getDate()).padStart(2, '0')}`;

    try {
        const response = await fetch(`/api/tasks/week?start=${start}&end=${end}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data || []; // Ensure an empty array if data is null or undefined
    } catch (error) {
        console.error("Error fetching tasks for week:", error);
        return []; // Return an empty array on error
    }
}


async function fetchTasksForDay(year, month, day) {
    try {
        const response = await fetch(`/api/tasks/day/${year}/${month}/${day}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error("Error fetching tasks for day:", error);
        return [];
    }
}




// Initialize on load
window.onload = () => {
    renderNavigationBar();
    renderMonthView(currentDate); // Start with month view
    refreshCalendarGrid()
};
window.addEventListener("pageshow", () => {
    refreshCalendarGrid();
});


