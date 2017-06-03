;��������
nginx -s start
nginx -s stop
nginx -s reload
;�������
location /jars {
	proxy_buffering off;
	proxy_pass http://192.168.200.26:5050/jars;
	proxy_set_header   Host    $host;
	proxy_set_header   X-Real-IP   $remote_addr;
	proxy_set_header   X-Forwarded-For $proxy_add_x_forwarded_for;
}
;���ؾ���
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
		#�������ĵ�ַ
		proxy_pass http://backend;
	}
}