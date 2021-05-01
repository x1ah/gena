package generators

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/x1ah/gena"
)

// WebStackGenerator generate webstack template
type WebStackGenerator struct{}

// Run implement Generator.Run
func (ws *WebStackGenerator) Run(cfg *gena.Config, writer io.Writer) {
	tmpl := template.Must(template.New("webstack.tmpl").Funcs(map[string]interface{}{
		"icon": favicon,
		"searchEnabled": func() bool {
			return cfg.WebStack != nil && cfg.WebStack.Search != nil && cfg.WebStack.Search.Enabled && len(cfg.WebStack.Search.Engines) > 0
		},
		"renderSearchEngines": renderSearchEngines(cfg),
		"defaultSearchEngineURL": func() string {
			return defaultSearchEngineURL(cfg)
		},
	}).ParseFS(gena.Templates, "templates/webstack.tmpl"))

	if err := tmpl.Execute(writer, cfg); err != nil {
		log.Fatal("[webstack] Render template error: ", err.Error())
	}
}

type searchEngineInfo struct {
	Name     string `json:"name"`
	Img      string `json:"img"`
	Position string `json:"position"`
	URL      string `json:"url"`
}

func buildSearchIcon(input string) string {
	return fmt.Sprintf(input, "https://cdn.jsdelivr.net/gh/x1ah/webstack-assets@master/assets/images/search_icon.png")
}

var availableEngines = map[string]*searchEngineInfo{
	"google": {
		Name:     "谷歌",
		Img:      buildSearchIcon("url('%s')  -105px 0px"),
		Position: "-40px 0px",
		URL:      "https://www.google.com/search?q=",
	},
	"baidu": {
		Name:     "百度",
		Img:      buildSearchIcon("url('%s') -80px 0px"),
		Position: "0px 0px",
		URL:      "https://www.baidu.com/s?wd=",
	},
	"bing": {
		Name:     "必应",
		Img:      buildSearchIcon("url('%s')  -80px -25px"),
		Position: "0px -40px",
		URL:      "https://cn.bing.com/search?q=",
	},
	"sougou": {
		Name:     "搜狗",
		Img:      buildSearchIcon("url('%s') -80px -50px"),
		Position: "0px -80px",
		URL:      "https://www.sogou.com/web?query=",
	},
	"taobao": {
		Name:     "淘宝",
		Img:      buildSearchIcon("url('%s') -105px -50px"),
		Position: "-40px -80px",
		URL:      "https://s.taobao.com/search?q=",
	},
	"jd": {
		Name:     "京东",
		Img:      buildSearchIcon("url('%s') -80px -75px"),
		Position: "0px -120px",
		URL:      "http://search.jd.com/Search?keyword=",
	},
	"tmall": {
		Name:     "天猫",
		Img:      buildSearchIcon("url('%s') -105px -75px"),
		Position: "-40px -120px",
		URL:      "https://list.tmall.com/search_product.htm?q=",
	},
	"zhihu": {
		Name:     "知乎",
		Img:      buildSearchIcon("url('%s') -105px -100px"),
		Position: "-40px -160px",
		URL:      "https://www.zhihu.com/search?type=content&q=",
	},
	"weibo": {
		Name:     "微博",
		Img:      buildSearchIcon("url('%s') -80px -125px"),
		Position: "0px -200px",
		URL:      "https://s.weibo.com/weibo/",
	},
	"bilibili": {
		Name:     "B站",
		Img:      buildSearchIcon("url('%s') -105px -125px"),
		Position: "-40px -200px",
		URL:      "http://search.bilibili.com/all?keyword=",
	},
	"douban": {
		Name:     "豆瓣",
		Img:      buildSearchIcon("url('%s') -80px -150px"),
		Position: "0px -240px",
		URL:      "https://www.douban.com/search?source=suggest&q=",
	},
	"github": {
		Name:     "GitHub",
		Img:      buildSearchIcon("url('%s') -80px -175px"),
		Position: "0px -280px",
		URL:      "https://github.com/search?utf8=✓&q=",
	},
}

func defaultSearchEngineURL(cfg *gena.Config) string {
	if v, ok := availableEngines[cfg.WebStack.Search.Default]; ok {
		prefix := strings.Split(v.Img, " ")[0]
		url := prefix + " " + v.Position
		return fmt.Sprintf("background: %s; opacity: 1;", url)
	}
	return buildSearchIcon("background: url(%s) 0px 0px;")
}

// nolint: funlen
func renderSearchEngines(cfg *gena.Config) func() string {
	return func() string {
		var definedEngines []string
		for _, engine := range cfg.WebStack.Search.Engines {
			if v, ok := availableEngines[engine]; ok {
				buf, err := json.Marshal(v)
				if err != nil {
					if _, err2 := fmt.Fprintf(os.Stderr, "[webstack] search enging %s config error: %s", engine, err.Error()); err2 != nil {
						continue
					}
					continue
				}
				definedEngines = append(definedEngines, string(buf))
			}
		}
		var defaultEngineURL string
		v, ok := availableEngines[cfg.WebStack.Search.Default]
		if ok {
			defaultEngineURL = v.URL
		} else {
			defaultEngineURL = availableEngines["google"].URL
		}

		return fmt.Sprintf(`<script type="text/javascript">
function search() {
    $(".search-icon").css("opacity", "1");
    var listIndex = -1;
    var hotList = 0;
    var searchData = {
        "thisSearch": "%s",
        "thisSearchIcon": "%s",
        "hotStatus": true,
        "data": [%s]
    };
    var localSearchData = localStorage.getItem("searchData");
    if (localSearchData) {
        searchData = JSON.parse(localSearchData)
    }
    function filterChildren(element) {
        var thisText = $(element).contents().filter(function (index, content) {
            return content.nodeType === 3
        }).text().trim();
        return thisText
    }
    function getHotkeyword(value) {
        $.ajax({
            type: "GET",
            url: "https://sp0.baidu.com/5a1Fazu8AA54nxGko9WTAnF6hhy/su",
            async: true,
            data: {
                wd: value
            },
            dataType: "jsonp",
            jsonp: "cb",
            success: function (res) {
                $("#box ul").text("");
                hotList = res.s.length;
                if (hotList) {
                    $("#box").css("display", "block");
                    for (var i = 0; i < hotList; i++) {
                        $("#box ul").append("<li><span>" + (i + 1) + "</span> " + res.s[i] + "</li>");
                        $("#box ul li").eq(i).click(function () {
                            var thisText = filterChildren(this);
                            $("#txt").val(thisText);
                            window.open(searchData.thisSearch + thisText);
                            $("#box").css("display", "none")
                        });
                        if (i === 0) {
                            $("#box ul li").eq(i).css({
                                "border-top": "none"
                            });
                            $("#box ul span").eq(i).css({
                                "color": "#fff",
                                "background": "#f54545"
                            })
                        } else {
                            if (i === 1) {
                                $("#box ul span").eq(i).css({
                                    "color": "#fff",
                                    "background": "#ff8547"
                                })
                            } else {
                                if (i === 2) {
                                    $("#box ul span").eq(i).css({
                                        "color": "#fff",
                                        "background": "#ffac38"
                                    })
                                }
                            }
                        }
                    }
                } else {
                    $("#box").css("display", "none")
                }
            },
            error: function (res) {
                console.log(res)
            }
        })
    }
    $("#txt").keyup(function (e) {
        if ($(this).val()) {
            if (e.keyCode == 38 || e.keyCode == 40 || !searchData.hotStatus) {
                return
            }
            getHotkeyword($(this).val())
        } else {
            $(".search-clear").css("display", "none");
            $("#box").css("display", "none")
        }
    });
    $("#txt").keydown(function (e) {
        if (e.keyCode === 40) {
            listIndex === (hotList - 1) ? listIndex = 0 : listIndex++;
            $("#box ul li").eq(listIndex).addClass("current").siblings().removeClass("current");
            var hotValue = filterChildren($("#box ul li").eq(listIndex));
            $("#txt").val(hotValue)
        }
        if (e.keyCode === 38) {
            if (e.preventDefault) {
                e.preventDefault()
            }
            if (e.returnValue) {
                e.returnValue = false
            }
            listIndex === 0 || listIndex === -1 ? listIndex = (hotList - 1) : listIndex--;
            $("#box ul li").eq(listIndex).addClass("current").siblings().removeClass("current");
            var hotValue = filterChildren($("#box ul li").eq(listIndex));
            $("#txt").val(hotValue)
        }
        if (e.keyCode === 13) {
            window.open(searchData.thisSearch + $("#txt").val());
            $("#box").css("display", "none");
            $("#txt").blur();
            $("#box ul li").removeClass("current");
            listIndex = -1
        }
    });
    $("#txt").focus(function () {
        $(".search-box").css("box-show", "inset 0 1px 2px rgba(27,31,35,.075), 0 0 0 0.2em rgba(3,102,214,.3)");
        if ($(this).val() && searchData.hotStatus) {
            getHotkeyword($(this).val())
        }
    });
    $("#txt").blur(function () {
        setTimeout(function () {
            $("#box").css("display", "none")
        }, 250)
    });
    for (var i = 0; i < searchData.data.length; i++) {
        $(".search-engine-list").append('<li><span style="background:' + searchData.data[i].img + '"/></span>' +
            searchData.data[i].name + "</li>")
    }
    $(".search-icon, .search-engine").hover(function () {
        $(".search-engine").css("display", "block")
    }, function () {
        $(".search-engine").css("display", "none")
    });
    $("#hot-btn").click(function () {
        $(this).toggleClass("off");
        searchData.hotStatus = !searchData.hotStatus;
        localStorage.searchData = JSON.stringify(searchData)
    });
    searchData.hotStatus ? $("#hot-btn").removeClass("off") : $("#hot-btn").addClass("off");
    $(".search-engine-list li").click(function () {
        var index = $(this).index();
        searchData.thisSearchIcon = searchData.data[index].position;
        $(".search-icon").css("background-position", searchData.thisSearchIcon);
        searchData.thisSearch = searchData.data[index].url;
        $(".search-engine").css("display", "none");
        localStorage.searchData = JSON.stringify(searchData)
    });
    $(".search-icon").css("background-position", searchData.thisSearchIcon);
    $("#search-btn").click(function () {
        var textValue = $("#txt").val();
        if (textValue) {
            window.open(searchData.thisSearch + textValue);
            $("#box ul").html("")
        } else {
            layer.msg("请输入关键词！", {
                time: 500
            }, function () {
                $("#txt").focus()
            })
        }
    })
}
</script>`, defaultEngineURL, buildSearchIcon("%s"), strings.Join(definedEngines, ","))
	}
}
