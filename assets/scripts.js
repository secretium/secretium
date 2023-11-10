// HTMX config.
htmx.config.selfRequestsOnly = true;
htmx.config.globalViewTransitions = true;
htmx.config.historyEnabled = false;

// Get secret sharer JWT from local storage.
function getSecretSharerJWT() {
    return localStorage.getItem('secretium_jwt');
}

// HTMX request interceptor for setting secret sharer JWT to the request headers.
document.body.addEventListener('htmx:configRequest', function (evt) {
    // If secret sharer JWT exists, add it to request headers.
    if (getSecretSharerJWT() !== null) {
        // Add secret sharer JWT to request headers.
        evt.detail.headers['Authorization'] = `Bearer ${getSecretSharerJWT()}`;
    }
});

// HTMX event listener for saving secret sharer JWT to local storage.
document.body.addEventListener("jwtSaveToLocalStorage", function (evt) {
    // Save secret sharer JWT to local storage.
    localStorage.setItem('secretium_jwt', evt.detail.value);
})

// HTMX event listener for removing secret sharer JWT from local storage.
document.body.addEventListener("jwtRemoveFromLocalStorage", function () {
    // Remove secret sharer JWT from local storage.
    localStorage.removeItem('secretium_jwt');
})