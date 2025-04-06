chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  const accountElement = document.querySelector('input#account');
  const passwordElement = document.querySelector('input#password');
  const enterElement = document.querySelector('input[type="submit"][data-v-2b854fd1]');

  if (accountElement && passwordElement && enterElement) {
    accountElement.value = request.account;
    passwordElement.value = request.password;

    triggerAllEvents(accountElement);
    triggerAllEvents(passwordElement);
    enterElement.click();
  } else {
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