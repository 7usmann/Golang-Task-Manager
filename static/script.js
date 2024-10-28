let currentDate = new Date();
let currentView = 'month';


// Reset currentDate to today whenever switching views

// Reset currentDate to today whenever switching views
function switchView() {
    currentDate = new Date(); // Resets to today each time view changes
    console.log("Switch view triggered");  // Add this line to confirm it's called

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
        renderMonthView(currentDate);
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

function renderMonthView(date) {
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

    // Fill the calendar with days of the month
    for (let day = 1; day <= daysInMonth; day++) {
        const dayCell = document.createElement('div');
        dayCell.innerText = day;

        // Make each cell clickable
        dayCell.onclick = function () {
            const formattedDate = `${year}-${String(month + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
            console.log(`Redirecting to /date/${formattedDate}`);
            window.location.href = `/date/${formattedDate}`;
        };

        // Highlight today's date
        if (day === new Date().getDate() && month === new Date().getMonth() && year === new Date().getFullYear()) {
            dayCell.classList.add('today');
        }
        calendarGrid.appendChild(dayCell);
    }
}


// Function to render the week view
function renderWeekView(startOfWeek) {
    const calendarGrid = document.getElementById('calendar-grid');
    calendarGrid.innerHTML = '';
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth();
    for (let i = 0; i < 7; i++) {
        const day = new Date(startOfWeek);
        day.setDate(startOfWeek.getDate() + i);

        const dayCell = document.createElement('div');
        dayCell.innerText = `${day.toDateString()}`;
        dayCell.onclick = function () {
            const formattedDate = `${year}-${String(month + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
            console.log(`Redirecting to /date/${formattedDate}`);
            window.location.href = `/date/${formattedDate}`;
        };
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
    const year = date.getFullYear();
    const month = date.getMonth();
    calendarGrid.innerHTML = '';
    const dayCell = document.createElement('div');
    dayCell.innerText = date.toDateString();
    dayCell.classList.add('today');
    dayCell.onclick = function () {
        const formattedDate = `${year}-${String(month + 1).padStart(2, '0')}-${String(currentDate).padStart(2, '0')}`;
        console.log(`Redirecting to /date/${formattedDate}`);
        window.location.href = `/date/${formattedDate}`;
    };
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

// Helper function to format the month and year display
function getMonthYearDisplay(date) {
    return date.toLocaleDateString(undefined, { year: 'numeric', month: 'long' });
}

// Initialize on load
window.onload = () => {
    renderNavigationBar();
    renderMonthView(currentDate); // Start with month view
};
