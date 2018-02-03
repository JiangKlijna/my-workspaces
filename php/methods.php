<?php
header('Content-Type:application/json');
if (empty($_POST)) {
	$_DATA = array();
	parse_str(file_get_contents('php://input'), $_DATA);
}
$arr = array(
	"server"=>$_SERVER,
	"get"=>$_GET,
	"post"=>$_POST,
	"file"=>$_FILES,
	"data"=>$_DATA,
);
echo json_encode($arr);
?>