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

function DELETE() {
	
}

function POST() {
	if(count($_FILES) > $GLOBALS['FILE_SIZE']) exit('[]');
	
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
	case "POST":
		POST();
		exit;
	case "DELETE":
		DELETE();
		exit;
	case "GET":
		break;
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
	</style>
	<body>
		<div id="app">
		</div>
	</body>
	<script>
		// native ajax
		var Ajax = function (method, url, callback) {
			var xhr = new XMLHttpRequest();
			xhr.onreadystatechange = function() {
				if (xhr.readyState === 4) {
					callback({
						status: xhr.status,
						response: xhr.response
					});
				}
			}
			xhr.open(method, url, true);
			xhr.send(null);
		}
		var app = document.getElementById('app');
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
	</script>
</html>
