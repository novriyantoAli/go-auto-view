// youtube.com observer
var ytcomObserver = new MutationObserver(function(mut)
{
  for(km in mut)
  {
    for(kn in mut[km].addedNodes)
    {
      var i = mut[km].addedNodes[kn].id;
      var c = mut[km].addedNodes[kn].className;
      //if (i=="movie_player")
      if(mut[km].addedNodes[kn].nodeName=='BODY')
      {
        chrome.storage.local.get("quality", function(o)
        {
          if (o.quality!=undefined)
          {
            window.localStorage["ytaq_quality"] = o.quality;
          }
          else
          {
            window.localStorage["ytaq_quality"] = "hd1080";
          }
        });
        chrome.storage.local.get("autopause", function(o)
        {
          if (o.autopause!=undefined)
          {
            window.localStorage["ytaq_autopause"] = o.autopause;
          }
          else
          {
            window.localStorage["ytaq_autopause"] = false;
          }
        });
        chrome.storage.local.get("ytgaming", function(o)
        {
          if (o.ytgaming!=undefined)
          {
            window.localStorage["ytaq_ytgaming"] = o.ytgaming;
          }
          else
          {
            window.localStorage["ytaq_ytgaming"] = false;
          }
        });
        var yts = window.document.createElement("script");
        yts.type = "text/javascript";
        yts.src = chrome.extension.getURL("inj/ytcom.js");
        window.document.getElementsByTagName('body')[0].appendChild(yts);
        ytcomObserver.disconnect();
        
        break;
      }
    }
  }
});

// embedded video observer (experimental)
var embObserver = new MutationObserver(function(mut)
{
  for(km in mut)
  {
    for(kn in mut[km].addedNodes)
    {
      var i = mut[km].addedNodes[kn].id;
      //var c = mut[km].addedNodes[kn].className;
      //if (i=="movie_player")
      if(mut[km].addedNodes[kn].nodeName=='EMBED' || mut[km].addedNodes[kn].nodeName=='IFRAME')
      {
        if(mut[km].addedNodes[kn].src.match('^((http|https):\/\/)*(www.)*youtube\.com(.*)$'))
        {
          chrome.storage.local.get("quality", function(o)
          {
            if (o.quality!=undefined)
            {
              window.localStorage["ytaq_quality"] = o.quality;
            }
            else
            {
              window.localStorage["ytaq_quality"] = "hd1080";
            }
          });
          var yts = window.document.createElement("script");
          yts.type = "text/javascript";
          yts.src = chrome.extension.getURL("inj/emb.js");
          window.document.getElementsByTagName('body')[0].appendChild(yts);
          embObserver.disconnect();
          break;
        }
      }
    }
  }
});

if(top.location.href.match('^((http|https):\/\/)*(www.|gaming.)*youtube\.com(.*)$'))
{
	if(	top.location.href.match('^((http|https):\/\/)*(gaming.)*youtube\.com(.*)$') &&
		window.localStorage.ytaq_ytgaming &&
		( window.localStorage.ytaq_ytgaming === false | window.localStorage.ytaq_ytgaming === "false" ) ) {
	} else {
		ytcomObserver.observe(document,{"childList":true,"subtree":true,"characterData":true});
	}
	
	chrome.runtime.onMessage.addListener( function(request, sender, sendResponse) {
		if (request.type == "update_quality")
			window.localStorage["ytaq_quality"] = request.quality;
	});
	
	chrome.runtime.onMessage.addListener( function(request, sender, sendResponse) {
		if (request.type == "update_autopause")
			window.localStorage["ytaq_autopause"] = request.autopause;
	});
	
	chrome.runtime.onMessage.addListener( function(request, sender, sendResponse) {
		if (request.type == "update_ytgaming")
			window.localStorage["ytaq_ytgaming"] = request.ytgaming;
	});
  
}
else
{
	//embedded?!
  //embObserver.observe(document,{"childList":true,"subtree":true,"characterData":true});
}
