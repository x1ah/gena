# gena

Navigation website generator | [中文文档](https://github.com/x1ah/gena/blob/master/README_CN.md)


![Test](https://github.com/x1ah/gena/workflows/Test/badge.svg) ![Lint](https://github.com/x1ah/gena/workflows/Lint/badge.svg)

---

## Install

### From source

**go1.16 required**

```asciidoc
go get github.com/x1ah/gena/cmd/gena
```

```asciidoc
> gena --help
Usage of ./gena:
  -c string
    	Config file (default "config.yml")
```

### Download binary release

Download from [Release page](https://github.com/x1ah/gena/releases)


## Generate

Define your config.yaml firstly, there are a sample config file: [config.yml](https://github.com/x1ah/gena/blob/master/config.yml)

```asciidoc
gena -c your/path/to/config.yaml > index.html
```

### Available templates

- [webstack](http://webstack.cc/)


## Preview

Just open `index.html` in your browser

![image](https://user-images.githubusercontent.com/14919255/114753545-27942500-9d8a-11eb-984a-a9cbb9c7e3b7.png)

