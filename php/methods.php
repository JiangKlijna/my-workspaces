<?php
header('Content-Type:application/json');
$arr = array(
	"server"=>$_SERVER,
	"get"=>$_GET,
	"post"=>$_POST,
	"file"=>$_FILES,
);
echo json_encode($arr);
?>