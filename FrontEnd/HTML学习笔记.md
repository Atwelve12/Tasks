# HTML学习笔记

[实践网页]()

1. 用vscode新建htlm文件

2. ！+tap补全[^网页]
   
   ```html
   <!DOCTYPE html>
   <html lang="en">
   <head>
       <meta charset="UTF-8">
       <meta name="viewport" content="width=device-width, initial-scale=1.0">
       <title>Document</title>
   </head>
   <body>
   
   </body>
   </html>
   ```

> "en"表示英文
> 
> title为标签页名称
> 
> body为正文

3. 在body添加标签

```html
<h1>标题</h1>
<p>这是一个段落。</p>
<p>这是另外一个段落。</p>
<a href="http://www.xxx.com>网址</a>
<div>盒子 </div>
```

[标签手册](https://www.runoob.com/tags/html-reference.html)

4. 用CSS美化

```css
<style>
 div {

}   
</style>
```

可以用class给每一个元素起一个名字

```css
<div class="name1"> </div>


<style>
 .name1{

}   
</style>
```



****





[^网页]:
    [类型] <!DOCTYPE html>
    [外信封开始] <html>
                                          [信头开始] <head>
                                            [编码说明] <meta charset>
                                            [手机适配] <meta viewport>
                                            [信件主题] <title>
                                          [信头结束] </head>
                                          [正文开始] <body>
                                            [这里放实际内容]
                                          [正文结束] </body>
    
    [外信封结束] </html>
