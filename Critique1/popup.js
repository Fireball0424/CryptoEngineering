document.addEventListener("DOMContentLoaded", function () {
  chrome.storage.sync.get(null, function (data) {
    const accountList = document.getElementById("account-list");

    for (let key in data) {
      const div = document.createElement("div");
      div.className = "account";
      div.textContent = key;

      div.addEventListener("click", () => {
        const password = data[key].password;
        const isHttps = data[key].isHttps || false; 
        const action = data[key].savedAction || "";
        chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
          chrome.tabs.sendMessage(tabs[0].id, {
            account: key,
            password: password,
            isHttps: isHttps, 
            savedAction: action
          });
        });
      });

      accountList.appendChild(div);
    }
  });
});
