!function(){"use strict";const e=!1;v("YSAT");var n="on",t="",o="on";const i=chrome||browser,r=-1!=navigator.userAgent.indexOf("Chrome")?"chrome":"firefox",a=r.charAt(0).toUpperCase()+r.slice(1);let d={},c=null,u="chrome"===r?"chrome-extension":"firefox-add-on";function s(){v("New Page");var e=document.getElementsByClassName("video-stream")[0]||!1;v(e),e&&(l(),f(),e.onprogress=e.ontimeupdate=function(){l(),f()})}function l(){var e=[".videoAdUiSkipButton.videoAdUiAction.videoAdUiFixedPaddingSkipButton",".ytp-ad-skip-button",".ytp-ad-overlay-close-button"];for(var n in e)p(e[n])}function f(){var e=["ytp-ad-overlay-close-button"];for(var n in v("function-initCloseAdBanner"),e){var t=document.getElementsByClassName(e[n]);t.length>0&&(v("function-initCloseAdBanner-hasBtn"),g(t))}}function g(e){if(v("function-triggerCloseAdBanner"),"off"===n)return!1;if("off"===o)return!1;for(var t=0;t<e.length;t++)v("function-triggerCloseAdBanner-click"),y(e[t],"click")}function p(e){v("Its "+n),!1!==_(e)&&(v("Skipping Ad in 1 seconds for: "+e),setTimeout(function(){if(!1!==_(e)){var n="skip_after_30_secs"===t?3e4:1;v("Skipping in: "+n),setTimeout(function(){document.querySelector(e).click()},n),function(){if(0==document.querySelectorAll(".ytp-ad-button-text").length)return;var e=document.querySelector(".ytp-ad-button-text").innerText;m({t:"event",ec:a+" "+(d.installType||"development"),ea:"Ad button text",el:e||""})}(),function(e){m({t:"event",ec:a+" "+(d.installType||"development"),ea:"Skip button clicked",el:"Skip button class: "+(e||"")})}(e)}},100))}function _(e){if("off"===n)return!1;if(0==document.querySelectorAll(e).length)return!1;if("skip_never"===t)return!1;if("skip_immediately"===t)return!0;if("skip_after_countdown"===t){var o=".ytp-ad-skip-button-slot";return 0!=document.querySelectorAll(o).length&&"none"!==window.getComputedStyle(document.querySelector(o)).display}if("skip_after_30_secs"===t){o=".ytp-ad-skip-button-slot";return 0!=document.querySelectorAll(o).length}}function v(){e&&console.log.apply(null,arguments)}function m(e){v("GA",e.t);let n=c||" ",t=u||"",o="Background Page: "+(u||""),i=d.name+" - "+u,r=Math.random().toString().replace("0.",""),a=window.location.hostname,s=window.location.host,l={v:"1",tid:"UA-57562361-4",cid:n,ds:t,z:r,t:"pageview",aip:"1",npa:"1",dl:window.location.href,dh:s,dp:a,dt:o,an:i,aid:a,av:d.version},f=Object.assign(l,e||{}),g=[];f.ev&&f.ev!==parseInt(f.ev)&&delete f.ev;for(let[e,n]of Object.entries(f))g.push(e+"="+n);let p=g.join("&");try{let e=new XMLHttpRequest;e.open("POST","https://www.google-analytics.com/collect",!0),e.send(p),v("Sent:",f),v("Sent:",p)}catch(e){v("Error sending report to Google Analytics.\n"+e)}}function y(e,n){if(e)if(v("triggering click"),e.fireEvent)e.fireEvent("on"+n);else{var t=document.createEvent("Events");t.initEvent(n,!0,!1),e.dispatchEvent(t)}}!function(){function e(){v("Hash changed"),s()}s(),document.addEventListener("spfdone",function(){v("spfdone"),s()}),document.addEventListener("transitionend",function(e){v("transitionend"),s()}),document.addEventListener("DOMContentLoaded",function(){v("DOMContentLoaded"),s()}),window.addEventListener("popstate",function(){v("Postate changed"),s()}),"onhashchange"in window&&(window.onhashchange=e);window.document.onload=e,window.onload=e,setTimeout(function(){s()},1e3)}(),i.storage.sync.get(["youtube_ad_skip_trigger_config","uuid","extensionDetails"],function(e){n=!1===e.youtube_ad_skip_trigger_config.status?"off":"on",t=e.youtube_ad_skip_trigger_config.skip_video_ad||"skip_after_countdown",o=!1===e.youtube_ad_skip_trigger_config.close_ad_banner?"off":"on",c=e&&e.uuid?e.uuid:" ",e&&e.extensionDetails&&(d=e.extensionDetails)}),i.storage.onChanged.addListener(function(e,i){n=!1===e.youtube_ad_skip_trigger_config.newValue.status?"off":"on",t=e.youtube_ad_skip_trigger_config.newValue.skip_video_ad||"skip_after_countdown",o=!1===e.youtube_ad_skip_trigger_config.newValue.close_ad_banner?"off":"on"}),v("Loading Youtube Video Skip Ad")}();