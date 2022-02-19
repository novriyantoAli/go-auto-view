var ytaq_player = null;
var ytaq_initPassed = false;
var ytaq_stateListenerSet = false;

originReadyHandler = window.onYouTubePlayerReady || function(){};

function ytpStateChangedHandler(state)
{
  if(ytaq_player==null)
  {
    ytaq_player = ytaq_getPlayer();
  }
  if (ytaq_player!=null) {
    // tiny
    // small
    // medium
    // large
    // hd720
    // hd1080
    // hd1440
    // highres
    ytaq_player.setPlaybackQuality(window.localStorage["ytaq_quality"]);
    
    if(state === 1)
    {
      if (ytaq_initPassed)
      {
        if(ytaq_stateListenerSet) {
          ytaq_player.removeEventListener('onStateChange', 'ytpStateChangedHandler');
          ytaq_stateListenerSet = false;
        }
      }
      else
      {
        ytaq_initPassed=true;
        ytaq_player.pauseVideo();
      }
    }
    else if(state === 2 || state === 5)
    {
      ytaq_player.playVideo();
    }
  }
}

function ytaq_getPlayer(playerId)
{
  var p = null;
  for(i in window.document.getElementsByTagName('embed'))
  {
    var el = window.document.getElementsByTagName('embed')[i];
    //if(el.src.match('^((http|https):\/\/)*(www.)*youtube\.com(.*)\&playerapiid\='+playerId+'(.*)$'))
    if(el.getAttribute("flashvars").match('^(.*)playerapiid\='+playerId+'(.*)$'))
    {
      p = el;
      break;
    }
  }
  if(p==null)
  {
    for(i in window.document.getElementsByTagName('iframe'))
    {
      var el = window.document.getElementsByTagName('embed')[i];
      //if(el.src.match('^((http|https):\/\/)*(www.)*youtube\.com(.*)\&playerapiid\='+playerId+'(.*)$'))
      if(el.getAttribute("flashvars").match('^(.*)playerapiid\='+playerId+'(.*)$'))
      {
        p = el;
        break;
      }
    }
  }
  
  return p;
}

function ytpReadyHandler(playerId)
{
  if(playerId)
  {
    ytaq_player = ytaq_getPlayer(playerId);
    if(!ytaq_stateListenerSet)
    {
      ytaq_initPassed = false;
      ytaq_stateListenerSet = true;
      ytaq_player.addEventListener('onStateChange','ytpStateChangedHandler');
    }
    
    if(ytaq_player.getPlayerState()!==-1)
    {
      //ytaq_initPassed = true;
      ytaq_player.pauseVideo();
    }
  }
  originReadyHandler(playerId);
}

window.onYouTubePlayerReady = window.ytpReadyHandler;
