# XML 语法提示

## 安装 DTD 文件

```bash
sh -c "$(curl -fsSL https://gobatis.co/dtd/dtd.sh)"
```

## Goland 配置

**第一步：**
打开一个 mapper.xml 文件，将光标放在指定位置：

![](/assets/dtd/1.png)

**第二步：**

* Mac 下按 `Alt + Enter`；

弹出如下选项：
![](/assets/dtd/2.png)

**第三步：**

在弹窗中配置 DTD 映射路径：
::: tip 提示 
注意：需要使用用户名绝对路径，Goland 有路径提示。
:::

![](/assets/dtd/3.png)

完成配置后，在使用 `gobatis.co/dtd/mapper.dtd` 的 xml 文件中便会出现标签和属性的提示。

