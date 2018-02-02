<?php
/*
 * $_FILES["file"]["name"] - 被上传文件的名称
 * $_FILES["file"]["type"] - 被上传文件的类型
 * $_FILES["file"]["size"] - 被上传文件的大小，以字节计
 * $_FILES["file"]["tmp_name"] - 存储在服务器的文件的临时副本的名称
 * $_FILES["file"]["error"] - 由文件上传导致的错误代码
 */

$FILE_SIZE = 3;
$MAX_FILE_SIZE = 1024*1024;
$DIR = date("Y/md/");

class UploadFile {
	var $key;
	var $file;
 
    public function __construct($key, $file) {
		$this->key = $key;
        $this->file = $file;
    }
 
    public function handler() {
		if($this->file["error"] != 0) {
			return $this->status(-1, "Error:".$this->file["error"]);
		}
		if($this->file["size"] > $GLOBALS['MAX_FILE_SIZE']) {
			return $this->status(-2, "file size > MAX_FILE_SIZE");
		}
		try{
			$fn = $this->newfilename();
			$re = move_uploaded_file($this->file["tmp_name"], $fn);
			if ($re == 1)
				return $this->status(0, $fn);
			else
				return $this->status(-4, 'move file error');
		}catch(Exception $e){
			return $this->status(-5, $e->getMessage());
		}
    }
	
	private function newfilename() {
		$arr = explode('.', $this->file["name"]);
		return $GLOBALS['DIR'].time().rand().'.'.end($arr);
	}
	
	private function status($code, $msg) {
		return array(
			"code"=>$code,
			"msg"=>$msg,
			"name"=>$this->key
		);
	}
}

function PUT() {
	$_PUT = array();
	parse_str(file_get_contents('php://input'), $_PUT);
	
}

function POST() {
	if (empty($_FILES)) exit('[]');
	if (count($_FILES) > $GLOBALS['FILE_SIZE']) exit('[]');
	
	if (!file_exists($DIR)) mkdir($GLOBALS['DIR'], 0777, true);
	
	$arr = array();
	for ($key= key($_FILES); $key = key($_FILES); next($_FILES)) {
		$uf = new UploadFile($key, $_FILES[$key]);
		$re = $uf->handler();
		array_push($arr, $re);
	}
	echo json_encode($arr);
}

switch ($_SERVER['REQUEST_METHOD']){
	case "GET":
		break;
	case "POST":
		POST();
		exit;
	case "PUT":
		PUT();
		exit;
	default:
		header('HTTP/1.1 404 Not Found'); 
		echo '<title>404 Not Found</title><h1>Not Found</h1><p>The requested URL upload.php was not found on this server.</p>';
		exit;
}
?>
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>php文件上传</title>
	</head>
	<style>
		* {
			margin: 0;
			padding: 0;
			font-family: consolas;
		}
		html, body {
			width: 100%;
			height: 100%;
			text-align: center;
			background-color: #f1f1f1;
		}
		#app {
			text-align: left;
			border-radius: 6px;
			background-color: white;
			box-shadow: 0 0 8px rgba(0,0,0,.3);
			vertical-align: middle;
			overflow-y: auto;
			overflow-x: hidden;
			display: inline-block;
		}
		#app header {
			padding: 10px;
			background: none;
			color: #50c87e;
			font-size: 23px;
			font-weight: 600;
			border-bottom: 1px rgba(0,0,0,.3) solid;
			cursor: default;
		}
		#app header:hover {
			color: white;
			background: #50c87e;
			border-bottom: 1px #eee solid;
		}
		#app article {
			padding: 10px;
		}
	</style>
	<body>
		<div id="app">
			<header>文件列表</header>
			<article>
				<ul id="files"></ul>
			</article>
		</div>
	</body>
	<script>
		// 先获得app dom对象
		var app = document.getElementById('app');
		// native ajax 闭包
		(function(w, n){
			var objToStr = function(obj) {
				var sb = [];
				for (var i in obj) {
					sb.push('&', i, '=', obj[i]);
				}
				sb.shift();
				return sb.join('');
			}
			var Ajax = function (method, url, params, data, callback) {
				var paramsUrl = params === null ? url : url + '?' + objToStr(params);
				var requestBody = data === null ? data : objToStr(data);
				var xhr = new XMLHttpRequest();
				xhr.onreadystatechange = function() {
					if (xhr.readyState === 4) {
						callback({
							url: paramsUrl,
							status: xhr.status,
							response: xhr.response,
							requestBody: requestBody
						});
					}
				}
				xhr.open(method, paramsUrl, true);
				if (data !== null)
				xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
				xhr.send(requestBody);
			};
			w[n] = Ajax;
		})(window, 'Ajax');
		// 调整app大小的闭包
		(function(w, a){
			var style = a.style;
			w.onresize = function() {
				if (w.innerWidth < 800) { // 调整为移动端样式
					style.marginTop = 0;
					style.height = '100%';
					style.width = '100%';
				} else { // 调整为pc端样式
					var marginTop = parseInt((w.innerWidth-800) / (w.screen.width-800) * 60);
					style.marginTop = marginTop + 'px';
					style.height = (w.innerHeight - marginTop * 2) + 'px';
					style.width = '800px';
				}
			}
			w.onresize();
		})(window, app);
		// 绘制app的闭包
		(function(w, a, n){
			var files = a.querySelector('#files');
			var Draw = function(){}
			var draw = w[n] = new Draw();
			Draw.prototype.generate = function(list) {
				files.innerHTML = list.toString();
			}
		})(window, app, 'Draw');
		// 监听锚点的闭包
		(function(w, n){
			var Router = function(){}
			var router = w[n] = new Router();
			Router.prototype.refresh = function(hash) {
				if (hash !== null) {
					w.location.hash = hash;
				}
				
			}
			router.refresh(w.location.hash === '' ? '/' : null);
		})(window, 'Router');
	</script>
</html>
