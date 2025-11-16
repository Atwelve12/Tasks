# CSS学习笔记

****



## 渐变背景的一种实现方法

* **背景设置**
  
  ```css
  body {
            margin:  0;
            padding: 0;
            height: 100vh;
            background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1, #96ceb4);
            background-size: 400% 400%;
            animation: gradientBG 15s ease infinite;
        }
  ```

> * `margin: 0; padding: 0;` - 清除默认边距，让背景充满整个屏幕
> 
> * `height: 100vh;` - 设置高度为100%视口高度（viewport height）
> 
> * `background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1, #96ceb4);` - 创建45度角的多色线性渐变

> * `background-size: 400% 400%;` - **关键技巧**：将背景放大4倍，为动画创造移动空间

> * `animation: gradientBG 15s ease infinite;` - 应用动画：

> * `gradientBG` - 动画名称

> * `15s` - 动画时长15秒

> * `ease` - 缓动函数，让动画更自然

> * `infinite` - 无限循环

**效果**

![效果](https://github.com/Atwelve12/Task/blob/e48946905ef712a1454ae6ea075e2c319bb2803c/photo/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202025-11-11%20195505.png)

* **@keyframes动画**

```css
@keyframes gradientBG { 
    0% {
     background-position: 0% 50%;
     } 
    50% {     
    background-position: 100% 50%;
     }
     100% {
     background-position: 0% 50%; 
    }
}
```

> `background-position: 0% 50%;`-背景位置,水平位置,垂直位置[^位置]

****

* 





```css
.divname {
    padding: 50px;
    color: white;
    text-align: center;
    font-family: Arial, sans-serif;
}
```

> * `padding: 50px;` - 内边距，让内容有呼吸空间
> 
> * `color: white;` - 白色文字，在彩色背景上保持可读性
> 
> * `text-align: center;` - 文字居中对齐
> 
> * `font-family: Arial, sans-serif;` - 设置字体族



****





[^位置]: 水平位置：0% 表示背景图像的左边缘与容器的左边缘对齐，100% 表示背景图像的右边缘与容器的右边缘对齐。  
垂直位置：50% 表示背景图像的垂直中心与容器的垂直中心对齐。
