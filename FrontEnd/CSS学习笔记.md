# CSS学习笔记

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


