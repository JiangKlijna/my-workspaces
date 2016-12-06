
import pip
from os import system
#from subprocess import call

#pip list	#列出所有安装的库
#pip list --outdated	#列出所有过期的库
#python -m pip install -U pip	#升级pip
#pip install --upgrade [库名]	#更新库

system('pip list --format=columns')
system('python -m pip install -U pip')
system('pip list --outdated --format=columns')

for dist in pip.get_installed_distributions():
    system("pip install --upgrade " + dist.project_name)
