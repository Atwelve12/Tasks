# git进阶笔记

****



## 远程创建分支

```git
$ git branch -r develop
```



```git
$ git push origin --delete develop
//远程删除分支
$ git branch -r
//看查远程分支
$ git branch -a
//看查所有分支
```

****



## 创建分支

```git
$ git branch 
```

> 查看分支

```git
$ git checkout -b develop
```

> 创建分支develop
> 
> 可用switch代替checkout

```git
$ git add *
```

```git
$ git commit -m-"note"
```

```git
$ git  checkout main
```

> 回到main分支

```git
$ git merge develop
```

> 合并代码

```git
$ git push origin main
```

> 上传完成

```git
$ git branch -d develop
```

> 删除分支
> 
> 可用 -D 强制删除

****

## 团队合作工作流

![一种工作流](https://github.com/Atwelve12/Task/blob/4e1efab1d66e9cb9ddfebf6b1fdacd8f690205b9/photo/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202025-11-10%20174700.png)

****

## 回溯代码

```git
 $ git log
```

> 看查上传记录



```git
$ git log --stat
```

> 看查改动文件



```git
git diff 9f7720326d0ef955348ce7693957f1f5b519f513
```

> 通过commit id具体看查改动



```git
git reset --hard 88cd16b249c8bbc6640072531ef77bc09a0e32fd
//或
git checkout 88cd16b249c8bbc6640072531ef77bc09a0e32fd
```

> 回溯代码




