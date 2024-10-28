let currentDate = new Date();
let currentView = 'month';

// Reset currentDate to today whenever switching views
function switchView() {
    currentDate = new Date(); // Resets to today each time view changes
    
    if (currentView === 'month') {
        currentView = 'week';
        renderWeekView(getStartOfWeek(new Date()));
        document.getElementById('view-switch').innerText = 'Week';
    } else if (currentView === 'week') {
        currentView = 'day';
        renderDayView(new Date());
        document.getElementById('view-switch').innerText = 'Day';
    } else {
        currentView = 'month';
        renderMonthView(new Date());
        document.getElementById('view-switch').innerText = 'Month';
    }
    renderNavigationBar();
}

// Navigation function to update date based on view mode and direction
function navigateView(direction) {
    if (currentView === 'month') {
        currentDate.setMonth(currentDate.getMonth() + direction);
        renderMonthView(currentDate);
    } else if (currentView === 'week') {
        currentDate.setDate(currentDate.getDate() + direction * 7);
        renderWeekView(currentDate);
    } else if (currentView === 'day') {
        currentDate.setDate(currentDate.getDate() + direction);
        renderDayView(currentDate);
    }
    renderNavigationBar();
}

// Function to render the month view
function renderMonthView(date) {
    const calendarGrid = document.getElementById('calendar-grid');
    calendarGrid.innerHTML = '';
    const daysInMonth = new Date(date.getFullYear(), date.getMonth() + 1, 0).getDate();
    const firstDayOfMonth = new Date(date.getFullYear(), date.getMonth(), 1).getDay();
    const offset = (firstDayOfMonth === 0) ? 6 : firstDayOfMonth - 1; // Offset for Monday start

    for (let i = 0; i < offset; i++) {
        const emptyCell = document.createElement('div');
        calendarGrid.appendChild(emptyCell);
    }

    for (let day = 1; day <= daysInMonth; day++) {
        const dayCell = document.createElement('div');
        dayCell.innerText = day;
        const today = new Date();
        if (day === today.getDate() && date.getMonth() === today.getMonth() && date.getFullYear() === today.getFullYear()) {
            dayCell.classList.add('today');
        }
        calendarGrid.appendChild(dayCell);
    }
}

// Function to render the week view
function renderWeekView(startOfWeek) {
    const calendarGrid = document.getElementById('calendar-grid');
    calendarGrid.innerHTML = '';
    for (let i = 0; i < 7; i++) {
        const day = new Date(startOfWeek);
        day.setDate(startOfWeek.getDate() + i);

        const dayCell = document.createElement('div');
        dayCell.innerText = `${day.toDateString()}`;
        const today = new Date();
        if (day.getDate() === today.getDate() && day.getMonth() === today.getMonth() && day.getFullYear() === today.getFullYear()) {
            dayCell.classList.add('today');
        }
        calendarGrid.appendChild(dayCell);
    }
}

// Function to render the day view
function renderDayView(date) {
    const calendarGrid = document.getElementById('calendar-grid');
    calendarGrid.innerHTML = '';
    const dayCell = document.createElement('div');
    dayCell.innerText = date.toDateString();
    dayCell.classList.add('today');
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

// Helper function to format the month and year display
function getMonthYearDisplay(date) {
    return date.toLocaleDateString(undefined, { year: 'numeric', month: 'long' });
}

// Initialize on load
window.onload = () => {
    renderNavigationBar();
    renderMonthView(currentDate); // Start with month view
};
