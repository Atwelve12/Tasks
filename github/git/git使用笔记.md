# git个人仓库的使用

1. 新建文件夹

2. 创建本地仓库

```git
$ git init
```

```git
$ git add abc.txt
 或
$ git add .
```



```git
$ git commit -m"note"
```



****

# git个人远程仓库使用

1. 添加远程仓库

```git
$ git clone git@github.com:Atwelve12/Task.git
```

*或者*



新建文件夹

```git
$ git init
```

> 创造本地文件

```git
$ git remote add origin git@github.com:Atwelve12/Task.git
```

> 创建远程仓库并与github仓库关联
> 
> origin是常用的远程仓库名

2. 修改你的代码,文件.
   
   
   
   

3. 上传文件

```git
$ git add *
```

> 上传暂存区

```git
$ git commit -m"备注"
```

> 备注

```git
 $ git push origin main
```

> 上传[^上传]

4. 下载代码

```git
$ git pull origin main
```

或

```git
$ git pull
```

> 拉取最新代码

****

# git使用还是太麻烦了

**有更好简便的办法吗**

有的兄弟有的

* 用git Desktop可视化操作

* 用vs code配置git

> 就不用在敲击代码,直接可视化操作



****















[^上传]: 需要提前绑定自己的SSHkey
