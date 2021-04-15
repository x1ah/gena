# gena

Navigation website generator | [中文文档](https://github.com/x1ah/gena/blob/master/README_CN.md)


![Test](https://github.com/x1ah/gena/workflows/Test/badge.svg) ![Lint](https://github.com/x1ah/gena/workflows/Lint/badge.svg)

---

![image](https://user-images.githubusercontent.com/14919255/114878771-29fe8980-9e33-11eb-834c-515379882664.png)


## Install

### GitHub Template (**Recommended**)

Use [gena-template](https://github.com/x1ah/gena-template) to automate generate


### From source

**go1.16 required**

```asciidoc
go install github.com/x1ah/gena/cmd/gena
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


## Use cases

- [x1ah's nav](https://when.run/nav)
