# gena

Navigation website generator | [中文文档](https://github.com/x1ah/gena/blob/master/README.md)


![Test](https://github.com/x1ah/gena/workflows/Test/badge.svg) ![Lint](https://github.com/x1ah/gena/workflows/Lint/badge.svg)

---

![image](https://user-images.githubusercontent.com/14919255/115016835-3395f900-9ee8-11eb-90d7-5ed816f59872.png)


## Install

### GitHub Template (**Recommended**)

Use [gena-template](https://github.com/x1ah/gena-template) to automate generate and deploy to GitHub Pages


### From source

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

### Download binary release

Download from [Release page](https://github.com/x1ah/gena/releases)


## Generate

Define your config.yaml firstly, there are a sample config file: [config.yml](https://github.com/x1ah/gena/blob/master/config.yml)

```asciidoc
gena -c your/path/to/config.yaml 1> index.html
```

### Available templates

- [webstack](http://webstack.cc/)


## Preview

Just open `index.html` in your browser


## Use cases

- [when.run/nav](https://when.run/nav)

