(function() {
	
	var playerLoaded = false;
	var resolutions = ['highres', 'hd2880', 'hd2160', 'hd1440', 'hd1080', 'hd720', 'large', 'medium', 'small', 'tiny'];
	
	function ytpReadyHandler()
	{
		playerLoaded = false;
		ytaq_player = ytaq_getPlayer();
	  
		if(ytaq_player)
		{
			ytaq_player.addEventListener("onStateChange",function(state) { ytpStateChangedHandler(ytaq_player,state); }, true);
			
			if (resolutions.indexOf(window.localStorage.ytaq_quality) >= resolutions.indexOf(ytaq_player.getPlaybackQuality())) {
				if(ytaq_player.setPlaybackQualityRange !== undefined)
					ytaq_setQualityRange(ytaq_player, window.localStorage.ytaq_quality);
				return;
			}
			
			var quality = resolutions.indexOf(window.localStorage.ytaq_quality);
			while(ytaq_player.getAvailableQualityLevels().indexOf(resolutions[quality]) === -1 &&
				  quality < resolutions.length) {
				quality++;
			}
			
			if (ytaq_player.getPlaybackQuality() !== resolutions[quality])
				ytaq_player.loadVideoById(ytaq_player.getVideoData().video_id, ytaq_player.getCurrentTime(), resolutions[quality]);
			
			if(ytaq_player.setPlaybackQualityRange !== undefined)
				ytaq_setQualityRange(ytaq_player, resolutions[quality]);
			
			ytaq_player.setPlaybackQuality(resolutions[quality]);
		}
	}
	
	function ytpStateChangedHandler(player, state) {
		if(window.localStorage.ytaq_autopause===true || window.localStorage.ytaq_autopause==="true") {
			if(!playerLoaded){
				player.pauseVideo();
				playerLoaded = true;
			}
		}
	}
	
	function ytaq_setQualityRange(player, res) {
		player.setPlaybackQualityRange(res, res);
	}
	
	function ytaq_getPlayer()
	{
	  var p = null;
	  if(window.videoPlayer)
	  {
		for(var i in window.videoPlayer)
		{
		  if(window.videoPlayer[i] && window.videoPlayer[i].setPlaybackQuality)
		  {
			p = window.videoPlayer[i];
			break;
		  }
		}
	  }
	  else
	  {
		p = window.document.getElementById('movie_player') ||
			window.document.getElementsByClassName("html5-video-player")[0] ||
			window.document.getElementById('movie_player-flash') ||
			window.document.getElementById('movie_player-html5') ||
			window.document.getElementById('movie_player-html5-flash');
	  }
	  return p;
	}
	
	window.addEventListener("loadstart", ytpReadyHandler, true);
	
})();