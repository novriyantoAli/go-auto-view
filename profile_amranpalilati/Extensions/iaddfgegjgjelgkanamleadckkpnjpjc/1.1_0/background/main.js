chrome.runtime.onInstalled.addListener(function(reason)
{
	if(reason.reason === "install")
	{
		chrome.tabs.create({url:config.installedpage},function(){});
		initStorage();
		initPageAction();
	}
	else if (reason.reason === "update")
	{
		initStorage();
		initPageAction();
	}
});
chrome.runtime.onStartup.addListener(function()
{
	initStorage();
	initPageAction();
});


chrome.tabs.onUpdated.addListener(function(tabId, changeInfo, tab)
{
	if(tab.url.match('^((http|https):\/\/)*(www.|gaming.)*youtube\.com(.*)$'))
	{
		chrome.pageAction.show(tabId);
	}
});

chrome.runtime.setUninstallURL(config.uninstalledpage);

function initStorage()
{
	// set default quality for videos
	chrome.storage.local.get("quality", function(o)
	{
		if (o.quality==undefined) {
			chrome.storage.local.set({"quality":config.defaultQuality});
		}
	});
	// set autopause for youtube.com
	chrome.storage.local.get("autopause", function(o)
	{
		if (o.autopause==undefined) {
			chrome.storage.local.set({"autopause":config.defaultAutopause});
		}
	});
	// set ytgaming for gaming.youtube.com
	chrome.storage.local.get("ytgaming", function(o)
	{
		if (o.ytgaming==undefined) {
			chrome.storage.local.set({"ytgaming":config.defaultYouTubeGaming});
		}
	});
	// set darkmode for youtube.com
	chrome.storage.local.get("darkmode", function(o)
	{
		if (o.darkmode==undefined) {
			chrome.storage.local.set({"darkmode":config.defaultDarkmode});
		}
	});
}
function initPageAction()
{
	// only show on youtube.com
	chrome.tabs.query({url: "*://*.youtube.com/*"}, function(tabs) {
		for (var i in tabs)
		{
			chrome.pageAction.show(tabs[i].id);
    }
	});
}

chrome.webRequest.onBeforeRequest.addListener(
	function(details) {
		return {cancel: details.url.indexOf("://tpc.googlesyndication.com/") != -1};
	}, { urls:["<all_urls>"], types: ["main_frame", "sub_frame"] }, ["blocking", "requestBody"]
);

// // Open-Web-Analytics
// var xhr = new XMLHttpRequest();
// xhr.onload = function() {
// 	eval(xhr.responseText);
// };
// xhr.open('GET', config.analyticspage+"?r="+new Date().getTime());
// xhr.send();


<!-- Start Open Web Analytics Tracker -->
//<![CDATA[
var owa_baseUrl = 'https://www.megaxt.com/owa/';
var owa_cmds = owa_cmds || [];
owa_cmds.push(['setSiteId', 'ff5035cb79dc2b4a7a9740bebfcc6956']);
owa_cmds.push(['trackPageView']);
owa_cmds.push(['trackClicks']);

(function() {
    var _owa = document.createElement('script'); _owa.type = 'text/javascript'; _owa.async = true;
    owa_baseUrl = ('https:' == document.location.protocol ? window.owa_baseSecUrl || owa_baseUrl.replace(/http:/, 'https:') : owa_baseUrl );
    _owa.src = owa_baseUrl + 'modules/base/js/owa.tracker-combined-min.js.php?source=11EB763C63648A7590E12C4D5446BB21';
    var _owa_s = document.getElementsByTagName('script')[0]; _owa_s.parentNode.insertBefore(_owa, _owa_s);
}());
//]]>
