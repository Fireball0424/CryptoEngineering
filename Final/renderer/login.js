function goToRegister() {
    window.location.href = "register.html";
}

document.getElementById("loginForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
  
    console.log("Username:", username);
    console.log("Password:", password);

    const result = await window.electronAPI.login(username, password);
    
    const msg = document.getElementById("message");
  
    if (result.success) {
      msg.innerHTML = "Login successful! Redirecting...";
      setTimeout(() => {
        window.location.href = "main.html";
      }, 2000);
    } else {
      msg.innerHTML = result.error || "Login failed!";
      msg.style.color = "red";
    }
  });