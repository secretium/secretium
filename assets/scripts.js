import htmx from 'htmx.org';

// Set HTMX to the window object.
window.htmx = require('htmx.org');

// HTMX config.
htmx.config.selfRequestsOnly = true;
htmx.config.globalViewTransitions = true;
htmx.config.historyEnabled = false;