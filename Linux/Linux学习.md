# Linux 常用命令学习笔记

****

## 目录与文件操作

* ls

> 列出文件



> ls -l 列出文件属性
> ls -a列出所有文件包括隐藏文件
> ls -la列出所有文件带属性

> **缩写 ll** 列出文件带属性

* cd files

> 进入文件files

> cd ..进入上一层目录
> cd ../..上上层目录
> cd -上一个目录

* pwd

> 打印当前路径

* cat files

> 读取文件

* head files

> 读取文件前10行

> head --line=3 files 看前3行

* tail files

> 读取文件后10行

> tail --line=3 files 看后3行

* less files

> 看文件全文用上下翻看

* mkdir files

> 创建目录

* rm file

> 删除文件

> rm -r dirs
> 
> 删除目录与其中子目录
> 
> rm -f 
> 
> 强制删除（不带提示）
> 
> rm -fr *
> 
> 删除当前目录里的所有文件
> 
> rm -i files.txt
> 
> 删除前逐一询问确认

+ touch file.txt

> 1.改变已有文件的时间属性
> 
> 2.创建新的空文件

> touch -a file.txt
> 
> 改变access时间(读取时间)
> 
> touch -m files.txt
> 
> 改变modify时间(修改时间)
> 
> touch -c
> 
> touch --no-create
> 
> 不创建新的文件
> 
> touch file1.txt -r file2.txt
> 
> 将file1的时间戳修改为file2相同的时间戳
> 
> touch -t 2602011553.55 file.txt
> 
> 将文件时间戳改为26年2月1日15时53分55秒

## 权限与执行

![c13f7944-9829-480e-86cd-5f86445d5308](file:///D:/Users/ppwqHU/Pictures/Typedown/c13f7944-9829-480e-86cd-5f86445d5308.png)

+ chmod

> 改变用户权限

> chmod -R a+r *
> 
> 将当前目录所有文件及递归目录文件设置为所有人可读取
> 
> chmod ug+w,o-w file.txt
> 
> 
> 
> chmod a+r,a+w,a+x file.txt
> 
> = chmod 777 file.txt

* sudo

> 以管理员的身份运行文件

> sudo visudo
> 
> 看查配置文件
> 
> sudo -l
> 
> 看查用户权限
> 
> sudo -u name whoami
> 
> 指定用户执行命令
> 
> sudo  !!
> 
> 使用root权限执行上条命令

## 网络与端口（了解即可）

+ ping 目标主机

> 测试主机间网络连通性

> ping -c 4 www.baidu.com
> 
> ping4次

+ netstat

> 显示网络状态

> netstat -a
> 
> 显示所有连接信息 

## 进程与服务（了解即可）

+ ps 参数

> 显示进程

+ systemctl

> 控制软件的启动和关闭

> systemctl start name
> 
> systemctl stop name
> 
> systemctl restrat name

> systermctl enable name
> 
> 设置开机自启
> 
> systemctl disable name

## 调试与日志查看

+ curl 参数 网站源码

> 文件传输
> 
>  curl -o baidu.html www.baidu.com

+ tail -f

> 跟踪打印文件末尾变化


