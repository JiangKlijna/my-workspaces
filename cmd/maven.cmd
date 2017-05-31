mvn archetype:generate -DgroupId=com.jiangKlijna.web -DartifactId=springboot -DarchetypeArtifactId=maven-archetype-webapp -DinteractiveMode=false
mvn package

mvn archetype:generate -DgroupId=com.jiangKlijna.maven -DartifactId=MavenTest -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
cd MavenTest
mvn compile
mvn exec:java -Dexec.mainClass="com.jiangKlijna.maven.App"
mvn test
mvn clean
mvn site
mvn validate
mvn package//validate , compile , test

tomcat:deploy	;部署一个web war包
tomcat:reload	;重新加载web war包
tomcat:start	;启动tomcat
tomcat:stop		;停止tomcat
tomcat:undeploy	;停止一个war包
tomcat:run		;启动嵌入式tomcat ，并运行当前项目


;1.直接从浏览器上下载该文件；http://repo.maven.apache.org/maven2/archetype-catalog.xml
;2.然后复制到C:\Users\Administrator\.m2\repository\org\apache\maven\archetype\archetype-catalog\2.4 下面；
;3.然后在执行的命令后面加上增加参数-DarchetypeCatalog=local，变成读取本地文件即可。