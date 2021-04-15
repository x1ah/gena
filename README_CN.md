## gena

导航网站生成器 | [英文文档](https://github.com/x1ah/gena/blob/master/README.md)


![Test](https://github.com/x1ah/gena/workflows/Test/badge.svg) ![Lint](https://github.com/x1ah/gena/workflows/Lint/badge.svg)

---

### 安装

**go1.16 required**

```asciidoc
go get github.com/x1ah/gena/cmd/gena
```

```asciidoc
> gena --help
Usage of ./gena:
  -c string
    	Config file (default "config.yml")
  -t string
    	Template name, available: webstack (default "webstack")
```

### 生成网站

参照 [config.yml](https://github.com/x1ah/gena/blob/master/config.yml) 填入自己的网站列表

```asciidoc
gena -c your/path/to/config.yaml > index.html
```

### 效果预览

浏览器打开生成的 `index.html`

![image](https://user-images.githubusercontent.com/14919255/114753545-27942500-9d8a-11eb-984a-a9cbb9c7e3b7.png)

