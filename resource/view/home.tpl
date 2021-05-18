<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <title>萤火虫之心</title>
    <link rel="stylesheet" type="text/css" href="public/css/main.css"/>
    <meta name="apple-mobile-web-app-capable" content="YES"/>
    <meta name="apple-mobile-web-app-status-bar-style" content="black"/>
    <link rel="apple-touch-icon" href="/public/images/apple-touch-icon.png"/>
    <meta name="title" content="firefly.love"/>
    <meta name="description" content="233"/>
    <link rel="image_src" href="/public/images/fb-image.jpg"/>
</head>
<body>
<canvas id="canvas"></canvas>

<div id="ui">
    <div id="fps"></div>
    <input id="chat" type="text"/>
    <div id="chatText"></div>
    <aside id="frogMode">
        <h3>Frog Mode</h3>
        <section id="tadpoles">
            <h4>Tadpoles</h4>
            <ul id="tadpoleList">
            </ul>
        </section>
        <section id="console">
            <h4>Console</h4>
        </section>
    </aside>

    <div id="cant-connect">
        与服务器断开连接了。您可以重新刷新页面。
    </div>
    <div id="unsupported-browser">
        <p>
            您的浏览器不支持 <a rel="external" href="http://en.wikipedia.org/wiki/WebSocket">WebSockets</a>.
            推荐您使用以下浏览器
        </p>
        <ul>
            <li><a rel="external" href="http://www.google.com/chrome">Google Chrome</a></li>
            <li><a rel="external" href="http://apple.com/safari">Safari 4</a></li>
            <li><a rel="external" href="http://www.mozilla.com/firefox/">Firefox 4</a></li>
            <li><a rel="external" href="http://www.opera.com/">Opera 11</a></li>
        </ul>
        <p>
            <a href="#" id="force-init-button">仍然浏览!</a>
        </p>
    </div>

</div>

<script src="/public/js/lib/parseUri.js"></script>
<script src="/public/js/lib/modernizr-1.5.min.js"></script>
<script src="/public/js/jquery.min.js"></script>
<script src="/public/js/lib/Stats.js"></script>

<script src="/public/js/App.js"></script>
<script src="/public/js/Model.js"></script>
<script src="/public/js/Settings.js"></script>
<script src="/public/js/Keys.js"></script>
<script src="/public/js/WebSocketService.js"></script>
<script src="/public/js/Camera.js"></script>

<script src="/public/js/Tadpole.js"></script>
<script src="/public/js/TadpoleTail.js"></script>

<script src="/public/js/Message.js"></script>
<script src="/public/js/WaterParticle.js"></script>
<script src="/public/js/Arrow.js"></script>
<script src="/public/js/formControls.js"></script>

<script src="/public/js/Cookie.js"></script>
<script src="/public/js/main.js"></script>

</body>
</html>