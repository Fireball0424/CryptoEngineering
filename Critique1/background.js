// search data in storage 
chrome.runtime.onMessage.addListener((message, sender, sendResponse) =>{
    if(message.type === "get_data"){
        const account = message.payload.account;
        chrome.storage.sync.get(account, function(data) {
            sendResponse({data:data[account]});
        });
    }
    return true;
});