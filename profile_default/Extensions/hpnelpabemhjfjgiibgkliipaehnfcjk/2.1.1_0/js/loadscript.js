"chrome"==(-1!=navigator.userAgent.indexOf("Chrome")?"chrome":"firefox")&&chrome.windows.getAll(function(e){for(var o in e)chrome.tabs.getAllInWindow(e[o].id,function(e){for(var o in e)-1!==e[o].url.indexOf("youtube.com")&&chrome.tabs.executeScript(e[o].id,{file:"js/myapp.js"})})});