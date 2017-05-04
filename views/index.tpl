<!DOCTYPE html>

<html>
<head>
  <title>Robot Control</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="/static/css/style.css" >
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" >
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-grid.min.css" >
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-reboot.min.css" >
  


</head>

<body>
  <div class="container">

    <div class="row">
 

      <div class="col-md-offset-2">
               <center><button type="button" class="btn btn-success" onclick="sendUpSignal()">Move Forward</button></center><br />
               <button type="button" class="btn btn-success" onclick="sendLeftSignal()">Turn Left</button>
               <button type="button" class="btn btn-success" onclick="sendRightSignal()">Turn Right</button><br /><br />
               <center><button type="button" class="btn btn-success" onclick="sendDownSignal()">Move Backward</button></center><br/>
               <center><button type="button" class="btn btn-danger" onclick="sendStopSignal()"> Stop </button> </center>
               <br />
                <div class="alert alert-info">
                  Robot is currently : <span id="status"> <font color="red">Stopped</font></span>
                </div>
      </div>

      <div class="col-md-6 col-md-offset-8">
      <iframe class="video-frame" height="560" width="640" src="http://192.168.0.105:8081"> </iframe>
      </div>
    </div>

  </div> 

  <script
  src="http://code.jquery.com/jquery-3.2.1.min.js"
  integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4="
  crossorigin="anonymous"></script>

  <script type="text/javascript" src="/static/js/script.js"></script>

  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  </body>


</html>
