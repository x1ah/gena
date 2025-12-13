# gena

导航网站生成器 | [English Document](https://github.com/x1ah/gena/blob/master/README_EN.md)

![Test](https://github.com/x1ah/gena/workflows/Test/badge.svg) ![Lint](https://github.com/x1ah/gena/workflows/Lint/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/x1ah/gena)](https://goreportcard.com/report/github.com/x1ah/gena) [![goproxy.cn](https://goproxy.cn/stats/github.com/x1ah/gena/badges/download-count.svg)](https://github.com/x1ah/gena)

---

![image](https://user-images.githubusercontent.com/14919255/116836115-ec9e4800-abf7-11eb-9a83-82eb379c130d.png)

## 简介

`gena` 是一个用 Go 语言编写的导航网站生成器，可以通过简单的 YAML 配置文件快速生成美观的导航网站。支持多种模板主题，适合个人、团队或组织使用。

### 特性

- 🚀 **快速生成**：通过 YAML 配置文件快速生成静态 HTML 网站
- 🎨 **多种模板**：支持 webstack、tilde-enhanced 等多种精美模板
- 📱 **响应式设计**：所有模板均支持移动端访问
- 🔍 **搜索功能**：支持多种搜索引擎（webstack 模板）
- ⌨️ **键盘导航**：支持快捷键快速访问网站（tilde 模板）
- 🎯 **简单易用**：只需一个配置文件即可生成完整网站
- 📦 **单文件部署**：生成的 HTML 文件可独立部署到任何静态托管服务

## 安装

### 一键生成（**推荐**）

从 [gena-template](https://github.com/x1ah/gena-template) 自动生成并自动部署到 GitHub Pages

### 源码安装

**go1.21+ required**

```bash
go get -u github.com/x1ah/gena/cmd/gena
```

或者克隆仓库后编译：

```bash
git clone https://github.com/x1ah/gena.git
cd gena
go build -o gena ./cmd/gena
```

查看帮助信息：

```bash
gena --help
```

输出：
```
Usage of ./gena:
  -c string
    	Config file (default "config.yml")
```

### 可执行文件下载

从 [Release page](https://github.com/x1ah/gena/releases) 下载对应平台的可执行文件

## 快速开始

### 1. 创建配置文件

参照 [config.yml](https://github.com/x1ah/gena/blob/master/config.yml) 创建你自己的配置文件：

```yaml
# 网站标题
title: 我的导航站

# 网站描述
description: 我的个人导航网站

# 网站图标
favicon: https://example.com/favicon.ico

# 网站 URL
url: https://example.com

# GitHub 链接（可选）
github: https://github.com/yourname

# 页脚信息
footer: © 2024 My Navigation

# 选择模板：webstack 或 tilde
template: webstack

# 网站内容
content:
  categories:
    - name: 常用网站
      sites:
        - name: GitHub
          description: 代码托管平台
          url: https://github.com/
          icon: https://github.com/favicon.ico
        - name: Google
          description: 搜索引擎
          url: https://www.google.com/
```

### 2. 生成网站

```bash
gena -c config.yml > index.html
```

### 3. 预览

直接在浏览器中打开生成的 `index.html` 文件即可预览效果。

## 模板说明

### webstack 模板

基于 [WebStack](http://webstack.cc/) 设计的精美导航模板，适合展示大量网站链接。

**特性：**
- 侧边栏分类导航
- 支持多种搜索引擎
- 响应式设计
- 夜间模式支持
- 图标懒加载

**配置示例：**

```yaml
template: webstack

webstack:
  search:
    enabled: true
    default: google
    engines:
      - google
      - github
      - baidu
      - bing
```

### tilde 模板

基于 [tilde-enhanced](https://github.com/Ozencb/tilde-enhanced) 的极简起始页模板，支持键盘快速导航。

**特性：**
- 极简设计，专注内容
- 键盘快捷键导航（输入首字母快速跳转）
- 支持搜索功能
- 支持 URL 直接访问
- 支持特殊命令（如 `invert!` 切换主题）
- 深色/浅色主题切换

**配置示例：**

```yaml
template: tilde

tilde:
  search:
    url: https://duckduckgo.com/?q=
    placeholder: "Search or enter URL"
  theme: dark  # dark 或 light
  show_keys: false  # 是否显示键盘快捷键
```

**使用技巧：**
- 输入网站名称的首字母，按回车快速访问
- 输入 `?` 查看帮助
- 输入 `invert!` 切换主题
- 输入 `[数字]!` 启动对应分类的所有网站
- 输入 `[key]:[query]` 在指定网站搜索
- 输入 `[key]/[path]` 访问网站特定路径
- 直接输入 URL 或域名访问

## 配置说明

### 基础配置

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `title` | string | 是 | 网站标题 |
| `description` | string | 是 | 网站描述 |
| `template` | string | 是 | 模板名称：`webstack` 或 `tilde` |
| `url` | string | 否 | 网站 URL |
| `favicon` | string | 否 | 网站图标 URL |
| `github` | string | 否 | GitHub 链接 |
| `footer` | string | 否 | 页脚信息 |
| `google_analytics` | string | 否 | Google Analytics ID |

### 内容配置

```yaml
content:
  categories:
    - name: 分类名称
      sites:
        - name: 网站名称
          description: 网站描述
          url: https://example.com
          icon: https://example.com/icon.png  # 可选，不填则自动获取
```

### webstack 模板配置

```yaml
webstack:
  search:
    enabled: true              # 是否启用搜索
    default: google            # 默认搜索引擎
    engines:                   # 可用的搜索引擎列表
      - google
      - github
      - baidu
      - bing
      # ... 更多搜索引擎
```

支持的搜索引擎：`google`, `baidu`, `bing`, `sougou`, `taobao`, `jd`, `tmall`, `zhihu`, `weibo`, `bilibili`, `douban`, `github`

### tilde 模板配置

```yaml
tilde:
  search:
    url: https://duckduckgo.com/?q=  # 搜索 URL
    placeholder: "Search or enter URL"  # 搜索框占位符
  theme: dark  # 主题：dark 或 light
  show_keys: false  # 是否显示键盘快捷键（true 显示按键，false 显示图标）
```

## 部署

### GitHub Pages

1. 将生成的 `index.html` 推送到 GitHub 仓库
2. 在仓库设置中启用 GitHub Pages
3. 选择主分支的根目录作为源

### 其他静态托管服务

生成的 `index.html` 是纯静态文件，可以部署到任何静态托管服务：

- [Netlify](https://www.netlify.com/)
- [Vercel](https://vercel.com/)
- [Cloudflare Pages](https://pages.cloudflare.com/)
- 自己的服务器

## 使用案例

- [when.run/nav](https://when.run/nav/) - 使用 webstack 模板的导航站

## 开发

### 运行测试

```bash
go test ./...
```

### 代码检查

```bash
golangci-lint run
```

### 贡献

欢迎提交 Issue 和 Pull Request！

## 交流群

|QQ 群|
|:--:|
|群号：100916933<br><img src="https://user-images.githubusercontent.com/14919255/118518920-372cd200-b76b-11eb-8271-0ee04e8df04c.jpeg" height="350px" />|

## 许可证

本项目采用 MIT 许可证，详见 [LICENSE](LICENSE) 文件。

## 致谢

- [WebStack](http://webstack.cc/) - webstack 模板设计
- [tilde-enhanced](https://github.com/Ozencb/tilde-enhanced) - tilde 模板设计
