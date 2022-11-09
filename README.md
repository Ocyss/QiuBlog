<p align="center">
  <a href="https://github.com/qiu-lzsnmb/QiuBlog">
    <img src="https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/android-icon-192x192.png" alt="Logo" width="180" height="180">
  </a>

  <h1 align="center">QiuBlog</h1>
  <p align="center">
    无限趋近于完美的自用博客，长期维护 <b>自用，学习</b>
    <br />
     <br />
    <a href="http://xn--xe4a.cf/">查看Demo</a>
    <a href="https://github.com/qiu-lzsnmb/QiuBlog/issues">报告Bug</a>
    <a href="https://github.com/qiu-lzsnmb/QiuBlog/issues">提出新特性</a>
  </p>

## 技术栈

#### 后端 Golang 1.19

- Gin 1.8.1             [(Web框架)](https://gin-gonic.com/zh-cn/)
- GORM 1.23.8     [(ORM)](https://gorm.io/zh_CN/)
- MySQL 8             [(数据库)](https://www.mysql.com/)

#### 前端 Vue.js 3

- Vite 3.1.0                      [(构建工具)](https://cn.vitejs.dev/)

- naive-ui / ionicons5    [(UI库 / 图标库)](https://www.naiveui.com/zh-CN/os-theme)
- sass                               [(css预处理语言)](https://www.sass.hk/)
- wangeditor                  [(富文本编辑器)](https://www.wangeditor.com/)
- vuedraggable              [(拖拽库)](https://github.com/SortableJS/Vue.Draggable)
- axios                             (网络请求库)

## 预览

<img src="https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/2022-10-30/2.png" style="zoom: 33%;" />

<img src="https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/2022-10-30/3.png" style="zoom:25%;" />

<img src="https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/2022-10-30/4.png" style="zoom:25%;" />

<img src="https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/2022-10-30/1.png" style="zoom: 25%;" />

### 文件目录说明
后端

```
│─api // 接口
│  └─v1
├─config // 配置文件
├─database // 数据库备份文件（初始化）
├─log // 项目日志
├─middleware // 中间件
├─model // 数据模型层
├─routes
│   router.go // 路由入口   
├─utils // 项目公用工具库
│  ├─ask // 预返回状态码
│  ├─errmsg // 接口错误返回处理
│  └─tool // 工具包
└─web // 前端源码
```

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 联系方式

​	qiudie@88.com

​	<a href="http://www.coolapk.com/u/3117607">酷安:熬鈛愿庸 </a>

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/qiu-lzsnmb/QiuBlog/blob/master/LICENSE)

### 鸣谢

[wejectchen GinBlog](https://github.com/wejectchen/Ginblog) 非常感谢wejectchen的开源项目，和录制的教程

[naive-ui-admin](https://github.com/jekip/naive-ui-admin) 部分代码 借鉴 了次项目

[shaojintian README模板](https://github.com/shaojintian/Best_README_template)  一下子就正式起来了

