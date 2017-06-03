;基本命令
nginx -s start
nginx -s stop
nginx -s reload
;反向代理
location /jars {
	proxy_buffering off;
	proxy_pass http://192.168.200.26:5050/jars;
	proxy_set_header   Host    $host;
	proxy_set_header   X-Real-IP   $remote_addr;
	proxy_set_header   X-Forwarded-For $proxy_add_x_forwarded_for;
}
;负载均衡
upstream backend {
	#ip_hash;
	server   192.168.200.26:5003;
	server   192.168.200.26:5002;
	server   192.168.200.26:5001;
}
server {
	listen       80;
	server_name  localhost;
	location / {
		#反向代理的地址
		proxy_pass http://backend;
	}
}