#set su password
sudo passwd
#home/dir zh_CN->en_US
export LANG=en_US
xdg-user-dirs-gtk-update
export LANG=zh_US
#remove amazon
sudo apt-get remove unity-webapps-common
#mirror for dl.google.com, use CDN
gmirror.org
#install *.deb
dpkg -i <xxx.deb>
#extract *.tar.gz
tar -xzf *.tar.gz
#python27->python35
sudo apt-get install python3-pip
pip install --upgrade pip
http://www.lfd.uci.edu/~gohlke/pythonlibs/

sudo update-alternatives --install /usr/bin/python python /usr/bin/python2 100
sudo update-alternatives --install /usr/bin/python python /usr/bin/python3 150
#install java
sudo add-apt-repository ppa:webupd8team/java
sudo apt-get update
sudo apt-get install oracle-java8-installer
sudo apt-get install oracle-java8-set-default

sudo apt-get install openjdk-8-jdk

sudo update-alternatives --config java
#install maven
wget http://apache.fayea.com/maven/maven-3/3.3.9/binaries/apache-maven-3.3.9-bin.tar.gz

sudo apt-gets install maven
#install tomcat
wget http://mirrors.cnnic.cn/apache/tomcat/tomcat-8/v8.0.37/bin/apache-tomcat-8.0.37.tar.gz
#install go
#install git
sudo apt-get install git
#install atom
https://github.com/atom/atom/releases
#install eclipse
http://www.eclipse.org/downloads/
#install dbeaver
https://github.com/serge-rider/dbeaver/releases
#install qt
sudo apt-get install cmake qt-default qtcreator
#install android studio
#install charles
https://www.charlesproxy.com/download/
#install 163music
http://music.163.com/#/download

