
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

function post() {
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
		post();
		break;
	default:
		header('HTTP/1.1 404 Not Found');
		echo '<title>404 Not Found</title><h1>Not Found</h1><p>The requested URL upload.php was not found on this server.</p>';
		break;
}
?>
