chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  const accountElement = document.querySelector('input#account');
  const passwordElement = document.querySelector('input#password');
  const enterElement = document.querySelector('input[type="submit"][data-v-2b854fd1]');
  const isHttps = window.location.protocol === 'https:';

// Check if in iframe 
  if (window.top != window.self) {
    console.log("Might be in an iframe");
    return;
  }

  if (accountElement && passwordElement && enterElement) {

    // if isHttps is false and request.isHttps is true, warning user and ask whether to continue 
    if (!isHttps && request.isHttps) {
      const continueWithoutHttps = confirm("You are not using HTTPS. Are you sure you want to autofill the password?");
      if (!continueWithoutHttps) {
        return;
      }
    }

    accountElement.value = request.account;
    passwordElement.value = request.password;

    triggerAllEvents(accountElement);
    triggerAllEvents(passwordElement);

    // before send , check if form action is correct 
    const formElement = accountElement.closest('form');
    const originAction = new URL(request.savedAction);
    const currentAction = new URL(formElement.action);

    if (originAction.origin !== currentAction.origin) {
      console.log("Form action is not correct");
      return;
    }
    else{
       enterElement.click();
    }

   
  } 
  else { // error which part is not found
    if (!accountElement) {
      console.error("Account input not found.");
    }
    if (!passwordElement) {
      console.error("Password input not found.");
    }
    if (!enterElement) {
      console.error("Enter button not found.");
    }
  }
});


function triggerAllEvents(element) {
  ['input', 'change', 'blur'].forEach(eventType => {
    const event = new Event(eventType, { bubbles: true });
    element.dispatchEvent(event);
  });
}

// To ask user whether to save the password or not
const waitForEnterbuttonElement = setInterval(() => {
  const enterElement = document.querySelector('input[type="submit"][data-v-2b854fd1]');
  if (enterElement) {
    clearInterval(waitForEnterbuttonElement);

    enterElement.addEventListener('click', () => {
      const accountElement = document.querySelector('input#account');
      const passwordElement = document.querySelector('input#password');

      if (accountElement && passwordElement) {
        const isHttps = window.location.protocol === 'https:';
        const account = accountElement.value;
        const password = passwordElement.value;
        const formElement = accountElement.closest('form');
        
        // check if account has been saved 
        chrome.storage.sync.get(account, function (data) {
          if (data[account]) {
            const savedPassword = data[account].password;
            if (savedPassword !== password) {
              const updatePassword = confirm(`The password for ${account} has changed. Do you want to update it?`);
              if (updatePassword) {
                chrome.storage.sync.set({
                  [account]: {
                    password: password,
                    isHttps: isHttps, 
                    savedAction: formElement.action 
                  }
                });
              }
            }
          }
          else{
            const savePassword = confirm(`Do you want to save the password for ${account}?`);
            if (savePassword) {
              chrome.storage.sync.set({
                [account]: {
                  password: password,
                  isHttps: isHttps, 
                  savedAction: formElement.action // save the current actio
                }
               });
            }
          }
        });
      }
  });
  }
}, 1000);