# gena

导航网站生成器 | [English Document](https://github.com/x1ah/gena/blob/master/README_EN.md)


![Test](https://github.com/x1ah/gena/workflows/Test/badge.svg) ![Lint](https://github.com/x1ah/gena/workflows/Lint/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/x1ah/gena)](https://goreportcard.com/report/github.com/x1ah/gena)

---

![image](https://user-images.githubusercontent.com/14919255/115016835-3395f900-9ee8-11eb-90d7-5ed816f59872.png)


## 安装

### 一键生成（**推荐**）

从 [gena-template](https://github.com/x1ah/gena-template) 自动生成并自动部署到 GitHub Pages

### 源码安装

**go1.16 required**

```asciidoc
go get -u github.com/x1ah/gena/cmd/gena
```

```asciidoc
> gena --help
Usage of ./gena:
  -c string
    	Config file (default "config.yml")
```

### 可执行文件下载

从 [Release page](https://github.com/x1ah/gena/releases) 下载对应平台文件

## 生成网站

参照 [config.yml](https://github.com/x1ah/gena/blob/master/config.yml) 填入自己的网站列表

```asciidoc
gena -c your/path/to/config.yaml 1> index.html
```

### 可选的模板

- [webstack](http://webstack.cc/)


## 效果预览

浏览器打开生成的 `index.html`

## 使用案例

- [when.run/nav](https://when.run/nav/)
