function goToRegister() {
    window.location.href = "register.html";
}

const {ipcRenderer} = require('electron');

document.getElementById("loginForm").addEventListener("submit", async (e) => {
    e.preventDefault(); // Prevent the default form submission behavior

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const result = await ipcRenderer.invoke("login", { username, password });
    
    const msg = document.getElementById("message");

    if(result.success){
        msg.innerHTML = "Login successful! Redirecting...";
        setTimeout(() => {
            window.location.href = "main.html"; // Redirect to main page after 2 seconds
        }, 2000);
    }else{
        msg.innerHTML = result.error || "Login failed!";
        msg.style.color = "red";
    }
});