package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

// func TestContent(t *testing.T) {
// 	reader := strings.NewReader(raw)
// 	fmt.Println(Content(reader))
// }

func TestText(t *testing.T) {
	reader := strings.NewReader(raw_text)
	doc, _ := html.Parse(reader)
	for _, f := range COMPUTE_FUNC {
		Tranverse(doc, f)
	}
	fmt.Println(largest_score, Attr(largest_node, "class"))
	res := getNodeText(largest_node)
	fmt.Println(res)
	fmt.Println("---------")
	fmt.Println(optimizationEnter(res))
}

var raw_text = `


<!DOCTYPE html>
<html lang="en" xmlns:wb="http://open.weibo.com/wb">
  <head>
    <meta charset="utf-8">
    <title>Cyeam</title>
    
    <meta name="author" content="Bryce">

    <meta name="google-translate-customization" content="a4136e955b3e09f2-45a74b56dc13e741-gf616ffda6e6360e0-11"></meta>

    <!-- Enable responsive viewport -->
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <!-- Le HTML5 shim, for IE6-8 support of HTML elements -->
    <!--[if lt IE 9]>
      <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->

    <!-- Le styles -->
    <link href="/assets/themes/twitter/bootstrap/css/bootstrap.2.2.2.min.css" rel="stylesheet">
    <link href="/assets/themes/twitter/css/style.css?body=1" rel="stylesheet" type="text/css" media="all">

    <!-- Le fav and touch icons -->
  <!-- Update these with your own images -->
    <link rel="shortcut icon" href="/assets/images/c32.ico">
  <!--  <link rel="apple-touch-icon" href="images/apple-touch-icon.png">
    <link rel="apple-touch-icon" sizes="72x72" href="images/apple-touch-icon-72x72.png">
    <link rel="apple-touch-icon" sizes="114x114" href="images/apple-touch-icon-114x114.png">
  -->

    <!-- atom & rss feed -->
    <link href="/atom.xml" type="application/atom+xml" rel="alternate" title="Sitewide ATOM Feed">
    <link href="/rss.xml" type="application/rss+xml" rel="alternate" title="Sitewide RSS Feed">

    <!-- jQuery -->
  <!--<script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>-->
    <script src="http://libs.baidu.com/jquery/1.7.2/jquery.min.js"></script>
    <script>!window.jQuery && document.write('<script src="/assets/js/jquery-1.7.2.min.js"><\/script>');</script>
  <!--<script type="text/javascript" src="http://app.stacktack.com/jquery.stacktack.min.js"></script>-->
  <script type="text/javascript" src="/assets/js/cyeam.js"></script>
  <script type="text/javascript">
    $().ready(function(){
      //wallpaper();
      // $(document).stacktack();
    })
  </script>
  </head>

  <body>
    <div class="navbar">
      <div class="navbar-inner">
        <div class="container-narrow">
          <div class="wrapperLogo" id="wrapperLogo" style="float: left; width: 80px;">
            <a href="/"><img src="/assets/images/logo.jpg" style="width:50px"></a>
          </div>
          <ul class="nav">
            <li><a href="/"><span class="icon-home"></span> Home</a></li>
            <li><a href="/categories.html"><span class="icon-th-list"></span> Categories</a></li>
            <li><a href="/archive.html"><span class="icon-ok"></span> Archive</a></li>
            <li><a href="/tags.html"><span class="icon-tags"></span> Tags</a></li>
            <li><a href="/pages.html"><span class="icon-th"></span> Pages</a></li>
            <!--<li><a href="/cyeam.html"><span class="icon-bookmark"></span> Cyeam</a></li>-->
            <!--
            
            


  
    
      
        
        <li><a href="/archive.html">Archive</a></li>
        
      
    
  
    
      
    
  
    
      
        
        <li><a href="/categories.html">Categories</a></li>
        
      
    
  
    
      
    
  
    
      
        
        <li><a href="/pages.html">Pages</a></li>
        
      
    
  
    
      
    
  
    
      
    
  
    
  
    
      
        
        <li><a href="/tags.html">Tags</a></li>
        
      
    
  


-->
          </ul>
        </div>
      </div>
    </div>

    <div id="container-narrow" class="container-narrow opacity">
      
      <div class="content" id="content">
        
<div class="page-header">
  <h1>Cyeam  <small>你不要用战术的勤奋掩盖战略的懒惰。</small></h1>
</div>

<div class="row-fluid">
  <div class="span12">
    
<h2>
    <a id="垂直搜索爬虫——maodou" href="/se/2015/04/05/maodou" target="_blank">
        垂直搜索爬虫——maodou
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-04-05
</p>
<p class="description">用Golang实现的搜索引擎爬虫maodou。</p>

<p class="read-all">
    <a href="/se/2015/04/05/maodou" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br"><br /></h2>

<h2>
    <a id="探测局域网里面的设备" href="/network/2015/03/16/fing" target="_blank">
        探测局域网里面的设备
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-03-16
</p>
<p class="description">用Golang实现Fing的功能。</p>

<p class="read-all">
    <a href="/network/2015/03/16/fing" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/network/2015/03/16/fing" target="_blank">
        <img src="http://cyeam.qiniudn.com/fing.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-1"><br /></h2>

<h2>
    <a id="网络协议的端口号" href="/network/2015/03/14/port" target="_blank">
        网络协议的端口号
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-03-14
</p>
<p class="description">关于端口号的解读。</p>

<p class="read-all">
    <a href="/network/2015/03/14/port" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-2"><br /></h2>

<h2>
    <a id="DNS协议Golang实现" href="/network/2015/02/03/dns" target="_blank">
        DNS协议Golang实现
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-02-03
</p>
<p class="description">这几天分析了一下DNS协议的内容，用Go语言。</p>

<p class="read-all">
    <a href="/network/2015/02/03/dns" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-3"><br /></h2>

<h2>
    <a id="DNS协议分析" href="/network/2015/01/29/dns" target="_blank">
        DNS协议分析
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-01-29
</p>
<p class="description">这几天分析了一下DNS协议的内容，用Go语言。</p>

<p class="read-all">
    <a href="/network/2015/01/29/dns" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-4"><br /></h2>

<h2>
    <a id="安全URL的Base64编码" href="/web/2015/01/20/urlsafebase64" target="_blank">
        安全URL的Base64编码
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-01-20
</p>
<p class="description">今天又见识了一种新的Base64编码方式。</p>

<p class="read-all">
    <a href="/web/2015/01/20/urlsafebase64" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-5"><br /></h2>

<h2>
    <a id="使用Git工具统计代码" href="/kaleidoscope/2015/01/17/gitstats" target="_blank">
        使用Git工具统计代码
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-01-17
</p>
<p class="description">GitStats源码分析。</p>

<p class="read-all">
    <a href="/kaleidoscope/2015/01/17/gitstats" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-6"><br /></h2>

<h2>
    <a id="字符串查找算法（二）" href="/golang/2015/01/15/go_index" target="_blank">
        字符串查找算法（二）
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-01-15
</p>
<p class="description">RK算法和FNV算法。</p>

<p class="read-all">
    <a href="/golang/2015/01/15/go_index" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-7"><br /></h2>

<h2>
    <a id="北京地铁视频播放原理猜想" href="/kaleidoscope/2015/01/14/bjsubway" target="_blank">
        北京地铁视频播放原理猜想
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2015-01-14
</p>
<p class="description">今天回家路上坐地铁的新发现。</p>

<p class="read-all">
    <a href="/kaleidoscope/2015/01/14/bjsubway" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-8"><br /></h2>

<h2>
    <a id="如何开发搜索引擎——爬虫（一）" href="/se/2014/12/26/search_engine" target="_blank">
        如何开发搜索引擎——爬虫（一）
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-12-26
</p>
<p class="description">想自学一下搜索引擎，记录一下目前了解到的东西，也和大家分享。</p>

<p class="read-all">
    <a href="/se/2014/12/26/search_engine" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-9"><br /></h2>

<h2>
    <a id="Golang通过邻接表实现有向图" href="/golang/2014/12/20/golang_dg" target="_blank">
        Golang通过邻接表实现有向图
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-12-20
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/golang/2014/12/20/golang_dg" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-10"><br /></h2>

<h2>
    <a id="Golang通过pup实现HTML解析" href="/web/2014/12/01/golang_pup" target="_blank">
        Golang通过pup实现HTML解析
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-12-01
</p>
<p class="description">在GitHub上面找了一个能解析HTML的包。</p>

<p class="read-all">
    <a href="/web/2014/12/01/golang_pup" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-11"><br /></h2>

<h2>
    <a id="Golang实现HTTP发送gzip请求" href="/golang/2014/11/29/golang_gzip" target="_blank">
        Golang实现HTTP发送gzip请求
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-11-29
</p>
<p class="description">beego的httplib不支持发送gzip请求，自己研究了一下。</p>

<p class="read-all">
    <a href="/golang/2014/11/29/golang_gzip" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-12"><br /></h2>

<h2>
    <a id="Golang字符串切割函数Split" href="/golang/2014/11/22/golang_split" target="_blank">
        Golang字符串切割函数Split
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-11-22
</p>
<p class="description">升级beego 1.4.2之后，根据发现的不兼容的地方顺藤摸瓜。</p>

<p class="read-all">
    <a href="/golang/2014/11/22/golang_split" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-13"><br /></h2>

<h2>
    <a id="beego/config包源码分析" href="/beego/2014/11/12/beego_config" target="_blank">
        beego/config包源码分析
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-11-12
</p>
<p class="description">基于beego 1.4.2版本。</p>

<p class="read-all">
    <a href="/beego/2014/11/12/beego_config" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-14"><br /></h2>

<h2>
    <a id="Apache Bench 压力测试" href="/test/2014/11/11/ab" target="_blank">
        Apache Bench 压力测试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-11-11
</p>
<p class="description">ab压力测试，说一下我的经验</p>

<p class="read-all">
    <a href="/test/2014/11/11/ab" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-15"><br /></h2>

<h2>
    <a id="MySQL的field_in_set函数" href="/database/2014/10/13/field_in_set" target="_blank">
        MySQL的field_in_set函数
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-10-13
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/database/2014/10/13/field_in_set" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-16"><br /></h2>

<h2>
    <a id="用Golang写了一个更新壁纸的小程序" href="/kaleidoscope/2014/10/09/go_wallpaper" target="_blank">
        用Golang写了一个更新壁纸的小程序
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-10-09
</p>
<p class="description">把程序放到启动目录里面，这样就能每次开机更新啦。</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/10/09/go_wallpaper" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-17"><br /></h2>

<h2>
    <a id="基于TCP套接字，通过Golang模拟HTTP请求（续）" href="/network/2014/09/29/go_http2" target="_blank">
        基于TCP套接字，通过Golang模拟HTTP请求（续）
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-09-29
</p>
<p class="description">使用Golang语言，了解HTTP协议通信过程。</p>

<p class="read-all">
    <a href="/network/2014/09/29/go_http2" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-18"><br /></h2>

<h2>
    <a id="基于TCP套接字，通过Golang模拟HTTP请求" href="/network/2014/09/28/go_http" target="_blank">
        基于TCP套接字，通过Golang模拟HTTP请求
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-09-28
</p>
<p class="description">使用Golang语言，了解HTTP协议通信过程。</p>

<p class="read-all">
    <a href="/network/2014/09/28/go_http" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-19"><br /></h2>

<h2>
    <a id="Go语言的测试" href="/golang/2014/09/25/go_test" target="_blank">
        Go语言的测试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-09-25
</p>
<p class="description">Golang源码testing包。</p>

<p class="read-all">
    <a href="/golang/2014/09/25/go_test" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-20"><br /></h2>

<h2>
    <a id="Ubuntu通过Bing壁纸自动更新" href="/kaleidoscope/2014/09/17/ubuntu_bing_bg" target="_blank">
        Ubuntu通过Bing壁纸自动更新
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-09-17
</p>
<p class="description">用代码给生活家点小情调。</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/09/17/ubuntu_bing_bg" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-21"><br /></h2>

<h2>
    <a id="Go语言的插入排序实现" href="/golang/2014/09/01/go_insertsort" target="_blank">
        Go语言的插入排序实现
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-09-01
</p>
<p class="description">Golang源码sort包。</p>

<p class="read-all">
    <a href="/golang/2014/09/01/go_insertsort" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-22"><br /></h2>

<h2>
    <a id="Go语言的堆排序实现" href="/golang/2014/08/29/go_heapsort" target="_blank">
        Go语言的堆排序实现
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-29
</p>
<p class="description">Golang源码sort包。</p>

<p class="read-all">
    <a href="/golang/2014/08/29/go_heapsort" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-23"><br /></h2>

<h2>
    <a id="龙哥与刘翔的约战" href="/ctalk/2014/08/28/zealer_t" target="_blank">
        龙哥与刘翔的约战
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-28
</p>
<p class="description">你说手机不好，我说你人不好，你说手机不好，我说别人手机也这样，你说手机不好，我说你给我找个除了苹果三星以外好的。你说手机不好，我说你从事了犯法的事情。你说手机不好，我说你被包养了没资格所说话。两个人说了半天，手机到底好不好？</p>

<p class="read-all">
    <a href="/ctalk/2014/08/28/zealer_t" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/ctalk/2014/08/28/zealer_t" target="_blank">
        <img src="http://cyeam.qiniudn.com/luoyulong.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-24"><br /></h2>

<h2>
    <a id="Solr里的整形字段" href="/solr/2014/08/27/solr_trieint" target="_blank">
        Solr里的整形字段
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-27
</p>
<p class="description">东西不多，关于Solr查库遇到的小问题。</p>

<p class="read-all">
    <a href="/solr/2014/08/27/solr_trieint" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-25"><br /></h2>

<h2>
    <a id="Golang的map迭代" href="/golang/2014/08/25/go_map" target="_blank">
        Golang的map迭代
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-25
</p>
<p class="description">上个礼拜比较忙，没有来得及更新。其中忙的一件事，就是改关于map迭代的bug。问题很简单，主要讲一下我的低级失误和Golang的map迭代策略。</p>

<p class="read-all">
    <a href="/golang/2014/08/25/go_map" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-26"><br /></h2>

<h2>
    <a id="接口安全机制" href="/framework/2014/08/18/go_hmacsha1" target="_blank">
        接口安全机制
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-18
</p>
<p class="description">说一下关于接口安全的设计机制。</p>

<p class="read-all">
    <a href="/framework/2014/08/18/go_hmacsha1" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-27"><br /></h2>

<h2>
    <a id="大数乘法" href="/golang/2014/08/15/go_largenumberx" target="_blank">
        大数乘法
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-15
</p>
<p class="description">用大数乘法实现长度不限的数字乘法。</p>

<p class="read-all">
    <a href="/golang/2014/08/15/go_largenumberx" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-28"><br /></h2>

<h2>
    <a id="字符串反转" href="/golang/2014/08/14/go_reverse" target="_blank">
        字符串反转
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-14
</p>
<p class="description">Golang的实现。</p>

<p class="read-all">
    <a href="/golang/2014/08/14/go_reverse" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-29"><br /></h2>

<h2>
    <a id="把一个字符串中的字符从小写转为大写" href="/golang/2014/08/12/go_upper" target="_blank">
        把一个字符串中的字符从小写转为大写
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-12
</p>
<p class="description">东西很简单，主要围绕Golang包的源码和Golang语言特性进行介绍。</p>

<p class="read-all">
    <a href="/golang/2014/08/12/go_upper" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-30"><br /></h2>

<h2>
    <a id="Golang通过反射实现结构体转成JSON数据" href="/golang/2014/08/11/go_json" target="_blank">
        Golang通过反射实现结构体转成JSON数据
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-11
</p>
<p class="description">Golang的'reflect'包是一个简易的反射包，能够通过此包实现将结构体编码成JSON数据流。</p>

<p class="read-all">
    <a href="/golang/2014/08/11/go_json" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-31"><br /></h2>

<h2>
    <a id="字符串查找算法" href="/golang/2014/08/08/go_index" target="_blank">
        字符串查找算法
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-08
</p>
<p class="description">发现好多笔试题，都问的是库函数。</p>

<p class="read-all">
    <a href="/golang/2014/08/08/go_index" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-32"><br /></h2>

<h2>
    <a id="Golang——append的可变参数" href="/golang/2014/08/07/go_append" target="_blank">
        Golang——append的可变参数
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-07
</p>
<p class="description">合并slice，被copy虐了一通，后来才发现，append能虐copy几条街。</p>

<p class="read-all">
    <a href="/golang/2014/08/07/go_append" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-33"><br /></h2>

<h2>
    <a id="Golang——for循环的两种用法" href="/golang/2014/08/06/go_for" target="_blank">
        Golang——for循环的两种用法
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-06
</p>
<p class="description">一般语言都支持两种遍历方式，这里做下介绍。</p>

<p class="read-all">
    <a href="/golang/2014/08/06/go_for" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-34"><br /></h2>

<h2>
    <a id="Golang——json数据处理" href="/json/2014/08/04/go_json" target="_blank">
        Golang——json数据处理
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-08-04
</p>
<p class="description">关于Unicode的介绍和Golang的处理方法。</p>

<p class="read-all">
    <a href="/json/2014/08/04/go_json" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-35"><br /></h2>

<h2>
    <a id="布隆过滤器" href="/hash/2014/07/30/bloomfilter" target="_blank">
        布隆过滤器
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-30
</p>
<p class="description">曾经听一位大神提起过布隆过滤器，一般用来过滤黑名单。之前去搜狗面试，最后一个题也是这个。当时果断没想起来。</p>

<p class="read-all">
    <a href="/hash/2014/07/30/bloomfilter" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-36"><br /></h2>

<h2>
    <a id="Golang binary包——byte数组如何转int？" href="/hash/2014/07/29/go_bytearraytoint" target="_blank">
        Golang binary包——byte数组如何转int？
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-29
</p>
<p class="description">看布隆过滤器源码https://github.com/willf/bloom，里面用了binary包，在这里做记录。</p>

<p class="read-all">
    <a href="/hash/2014/07/29/go_bytearraytoint" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-37"><br /></h2>

<h2>
    <a id="常见哈希函数FNV和MD5" href="/hash/2014/07/28/fnv_md5" target="_blank">
        常见哈希函数FNV和MD5
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-28
</p>
<p class="description">关于Golang语言的哈希函数FNV和MD5的简单介绍。</p>

<p class="read-all">
    <a href="/hash/2014/07/28/fnv_md5" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-38"><br /></h2>

<h2>
    <a id="网址压缩的调研分析（续）" href="/web/2014/07/25/short_url2" target="_blank">
        网址压缩的调研分析（续）
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-25
</p>
<p class="description">将Python实现的short_url.py改写成了Golang实现。</p>

<p class="read-all">
    <a href="/web/2014/07/25/short_url2" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-39"><br /></h2>

<h2>
    <a id="网址压缩的调研分析" href="/web/2014/07/24/short_url" target="_blank">
        网址压缩的调研分析
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-24
</p>
<p class="description">关于网址压缩的调研和自己的理解。</p>

<p class="read-all">
    <a href="/web/2014/07/24/short_url" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-40"><br /></h2>

<h2>
    <a id="Golang开发Thrift接口" href="/golang/2014/07/22/go_thrift" target="_blank">
        Golang开发Thrift接口
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-22
</p>
<p class="description">用Go语言开发Thrift接口的hello, world。</p>

<p class="read-all">
    <a href="/golang/2014/07/22/go_thrift" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-41"><br /></h2>

<h2>
    <a id="Golang 接口实现" href="/golang/2014/07/20/go_inte" target="_blank">
        Golang 接口实现
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-20
</p>
<p class="description">Golang 接口语法学习。</p>

<p class="read-all">
    <a href="/golang/2014/07/20/go_inte" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-42"><br /></h2>

<h2>
    <a id="Golang 处理命令行启动参数" href="/golang/2014/07/20/go_flag" target="_blank">
        Golang 处理命令行启动参数
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-20
</p>
<p class="description">Golang flag 包的学习。</p>

<p class="read-all">
    <a href="/golang/2014/07/20/go_flag" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-43"><br /></h2>

<h2>
    <a id="Linux Mint 下ShadowSocks安装试用" href="/kaleidoscope/2014/07/18/shadowsocks" target="_blank">
        Linux Mint 下ShadowSocks安装试用
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-18
</p>
<p class="description">自从6月份开始，上网越来越难。goagent基本废了，只能另寻它法。</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/07/18/shadowsocks" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-44"><br /></h2>

<h2>
    <a id="Android虚拟机模拟摄像头" href="/android/2014/07/17/avd_camera" target="_blank">
        Android虚拟机模拟摄像头
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-17
</p>
<p class="description">Linux Mint下为Android虚拟机开启摄像头。网上有很多教程，感觉都老了点，对不上。我这个是试了好久试出来的，再次记录一下，和大家分享。</p>

<p class="read-all">
    <a href="/android/2014/07/17/avd_camera" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-45"><br /></h2>

<h2>
    <a id="Linux下定时执行——crontab" href="/linux/2014/07/16/crontab" target="_blank">
        Linux下定时执行——crontab
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-16
</p>
<p class="description">想做一个服务，能够每天定时调用执行。比如，每天定时推送微信，或者发送邮件，又或者为Android端推送消息。形式多种多样，内容主要是能有一个地方能够提醒展示今天要做的事情。像结婚纪念日、好朋友生日这种，还是挺重要的。</p>

<p class="read-all">
    <a href="/linux/2014/07/16/crontab" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-46"><br /></h2>

<h2>
    <a id="Go 语言简单实现HashSet" href="/golang/2014/07/15/go_hashset" target="_blank">
        Go 语言简单实现HashSet
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-15
</p>
<p class="description">公司有个需求，就是能够对列表去重。本屌原本想直接用for循环实现，后来去查了查Java的实现方式，大开眼界。</p>

<p class="read-all">
    <a href="/golang/2014/07/15/go_hashset" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-47"><br /></h2>

<h2>
    <a id="Android L 预览版试用" href="/android/2014/07/14/android_l_preview" target="_blank">
        Android L 预览版试用
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-14
</p>
<p class="description">在微博上面，各种所谓关注科技的官微都在发各种Android L的更新截图，各种吵。我也就没忍住，去刷了一把。然后今天又刷了回来。</p>

<p class="read-all">
    <a href="/android/2014/07/14/android_l_preview" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/android/2014/07/14/android_l_preview" target="_blank">
        <img src="http://cyeam.qiniudn.com/7d3f0ccdjw1eic56ksdlnj20fp07fmxp.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-48"><br /></h2>

<h2>
    <a id="Nginx根据域名转发" href="/nginx/2014/07/12/nginx_router" target="_blank">
        Nginx根据域名转发
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-12
</p>
<p class="description">虽然一直没有直接配置过公司的Nginx服务器，但是还是耳濡目染了解到了一些相关内容，知道Nginx能够根据域名进行转发请求。这样，一台服务器就能够配置多个域名和多个应用程序。</p>

<p class="read-all">
    <a href="/nginx/2014/07/12/nginx_router" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/nginx/2014/07/12/nginx_router" target="_blank">
        <img src="http://cyeam.qiniudn.com/nginx.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-49"><br /></h2>

<h2>
    <a id="数据库访问的缓存与最大连接数" href="/golang/2014/07/11/db_cache_macconn" target="_blank">
        数据库访问的缓存与最大连接数
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-11
</p>
<p class="description">之前有人提点，现在自己写，就忘记加了。弱爆了。。。</p>

<p class="read-all">
    <a href="/golang/2014/07/11/db_cache_macconn" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-50"><br /></h2>

<h2>
    <a id="数据库访问时区问题" href="/golang/2014/07/09/sql_time_zone" target="_blank">
        数据库访问时区问题
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-09
</p>
<p class="description">今天第二篇，都一次一天发多篇干货</p>

<p class="read-all">
    <a href="/golang/2014/07/09/sql_time_zone" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-51"><br /></h2>

<h2>
    <a id="Go语言接口开发——不确定JSON数据结构的解析" href="/golang/2014/07/09/api_result_parse_go" target="_blank">
        Go语言接口开发——不确定JSON数据结构的解析
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-09
</p>
<p class="description">今天开发的小总结</p>

<p class="read-all">
    <a href="/golang/2014/07/09/api_result_parse_go" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-52"><br /></h2>

<h2>
    <a id="百度云推送——Go语言实现类库" href="/golang/2014/07/08/baiduyunpush" target="_blank">
        百度云推送——Go语言实现类库
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-08
</p>
<p class="description">把之前写好的百度云推送封装成了Go包。只是实现了最简单的消息推送。。</p>

<p class="read-all">
    <a href="/golang/2014/07/08/baiduyunpush" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/golang/2014/07/08/baiduyunpush" target="_blank">
        <img src="http://blog.cyeam.com/assets/images/logo.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-53"><br /></h2>

<h2>
    <a id="如何离线完成go get——安装Apache Thrift有感" href="/kaleidoscope/2014/07/03/fuckgfwforgo" target="_blank">
        如何离线完成go get——安装Apache Thrift有感
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-03
</p>
<p class="description">由于一些原因（你懂的），如果golang代码存放在Google Code上面，想通过go get下载编译就是在骗自己。今天就通过一些方法解决了。做天朝的程序员不易，且行且珍惜吧。</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/07/03/fuckgfwforgo" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/kaleidoscope/2014/07/03/fuckgfwforgo" target="_blank">
        <img src="http://thrift.apache.org/static/images/favicon.ico" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-54"><br /></h2>

<h2>
    <a id="正则表达式" href="/regex/2014/07/02/regex" target="_blank">
        正则表达式
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-02
</p>
<p class="description">正则表达式是传说中黑客的必备技能，也在知道创宇技能表里，之前看过一些，每次都记不住，这次记到这里，忘记了回来看好了。</p>

<p class="read-all">
    <a href="/regex/2014/07/02/regex" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/regex/2014/07/02/regex" target="_blank">
        <img src="http://img.wdjimg.com/mms/icon/v1/7/63/419428a17d65f4d604f6c3e00f404637_256_256.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-55"><br /></h2>

<h2>
    <a id="grep与日志开发" href="/linux/2014/07/01/grep" target="_blank">
        grep与日志开发
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-07-01
</p>
<p class="description">后台开发离不开日志，日志能帮助检查bug。而大量的日志并不能通过人工阅读进行检查，一般都是借助grep工具。这里将通过学习grep命令来对日志的打印方式进行分析。</p>

<p class="read-all">
    <a href="/linux/2014/07/01/grep" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/linux/2014/07/01/grep" target="_blank">
        <img src="http://cyeam.qiniudn.com/shell.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-56"><br /></h2>

<h2>
    <a id="Nutch和Solr的安装和简单测试" href="/nutch/2014/06/21/nutchsolr" target="_blank">
        Nutch和Solr的安装和简单测试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-06-21
</p>
<p class="description">之前看过《Lucene in Action》一部分，里面介绍了Nutch这个Java实现的网络爬虫，把测试的结果放在这里</p>

<p class="read-all">
    <a href="/nutch/2014/06/21/nutchsolr" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/nutch/2014/06/21/nutchsolr" target="_blank">
        <img src="http://cyeam.qiniudn.com/nutch.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-57"><br /></h2>

<h2>
    <a id="流程图打脸" href="/kaleidoscope/2014/06/19/fuckprocesschart" target="_blank">
        流程图打脸
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-06-19
</p>
<p class="description">记我的答辩趣闻</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/06/19/fuckprocesschart" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/kaleidoscope/2014/06/19/fuckprocesschart" target="_blank">
        <img src="http://cyeam.qiniudn.com/1788_m.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-58"><br /></h2>

<h2>
    <a id="AndroidHTTP库——Asynchronous Http Client for Android" href="/android/2014/06/17/asyncandroidhttp" target="_blank">
        AndroidHTTP库——Asynchronous Http Client for Android
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-06-17
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/android/2014/06/17/asyncandroidhttp" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-59"><br /></h2>

<h2>
    <a id="Android端RTSP解决方案——libstreaming" href="/2014/06/15/libstreaming" target="_blank">
        Android端RTSP解决方案——libstreaming
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-06-15
</p>
<p class="description">Postgraduate design</p>

<p class="read-all">
    <a href="/2014/06/15/libstreaming" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/2014/06/15/libstreaming" target="_blank">
        <img src="http://cyeam.qiniudn.com/libstreaming_icon.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-60"><br /></h2>

<h2>
    <a id="百度云推送" href="/golang/2014/06/11/baiduyunpush" target="_blank">
        百度云推送
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-06-11
</p>
<p class="description">百度云推送的Go语言实现。Github 地址： https://github.com/mnhkahn/BaiduYunPush。</p>

<p class="read-all">
    <a href="/golang/2014/06/11/baiduyunpush" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/golang/2014/06/11/baiduyunpush" target="_blank">
        <img src="http://cyeam.qiniudn.com/baiduyunpush.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-61"><br /></h2>

<h2>
    <a id="『hello, world』 如何运行" href="/computer/2014/05/15/gcc" target="_blank">
        『hello, world』 如何运行
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-05-15
</p>
<p class="description">《深入理解计算机系统》读书笔记</p>

<p class="read-all">
    <a href="/computer/2014/05/15/gcc" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/computer/2014/05/15/gcc" target="_blank">
        <img src="http://img3.douban.com/lpic/s4510534.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-62"><br /></h2>

<h2>
    <a id="笔试题总结" href="/cyeam/2014/05/14/exam" target="_blank">
        笔试题总结
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-05-14
</p>
<p class="description">基本每一次笔试，我都会在结束后记录下来答过的题，回来好好研究研究。马上就要毕业了，论文也差不多了，现在来总结一下。</p>

<p class="read-all">
    <a href="/cyeam/2014/05/14/exam" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/cyeam/2014/05/14/exam" target="_blank">
        <img src="http://cyeam.qiniudn.com/offer_temple.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-63"><br /></h2>

<h2>
    <a id="浪潮之巅——腾讯帝国？" href="/ctalk/2014/05/07/tencent_empire" target="_blank">
        浪潮之巅——腾讯帝国？
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-05-07
</p>
<p class="description">向吴军博士的《浪潮之巅》致敬。</p>

<p class="read-all">
    <a href="/ctalk/2014/05/07/tencent_empire" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/ctalk/2014/05/07/tencent_empire" target="_blank">
        <img src="http://weixin.qq.com/zh_CN/htmledition/images/weixin/qr_url_download0ec594.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-64"><br /></h2>

<h2>
    <a id="Android 配置" href="/postgraduate/2014/04/23/pager_settings" target="_blank">
        Android 配置
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-23
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/postgraduate/2014/04/23/pager_settings" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/23/pager_settings" target="_blank">
        <img src="http://cyeam.qiniudn.com/android.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-65"><br /></h2>

<h2>
    <a id="查看自己的top10 Linux命令" href="/linux/2014/04/19/linux_top10_command" target="_blank">
        查看自己的top10 Linux命令
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-19
</p>
<p class="description">在V2EX发现的有趣的查看自己输入历史的top10 命令。</p>

<p class="read-all">
    <a href="/linux/2014/04/19/linux_top10_command" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/linux/2014/04/19/linux_top10_command" target="_blank">
        <img src="http://cyeam.qiniudn.com/shell.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-66"><br /></h2>

<h2>
    <a id="Android 自定义形状" href="/postgraduate/2014/04/18/pager_ui_design" target="_blank">
        Android 自定义形状
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-18
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/postgraduate/2014/04/18/pager_ui_design" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/18/pager_ui_design" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-67"><br /></h2>

<h2>
    <a id="基于流媒体的对讲机系统——数据库模块" href="/postgraduate/2014/04/18/pager_sqlite" target="_blank">
        基于流媒体的对讲机系统——数据库模块
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-18
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/postgraduate/2014/04/18/pager_sqlite" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/18/pager_sqlite" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-68"><br /></h2>

<h2>
    <a id="Android 通知" href="/postgraduate/2014/04/18/pager_notification" target="_blank">
        Android 通知
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-18
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/postgraduate/2014/04/18/pager_notification" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/18/pager_notification" target="_blank">
        <img src="http://cyeam.qiniudn.com/android.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-69"><br /></h2>

<h2>
    <a id="视频编解码技术H264" href="/postgraduate/2014/04/17/pager_video" target="_blank">
        视频编解码技术H264
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-17
</p>
<p class="description">关于H264视频编码格式的介绍和分析。</p>

<p class="read-all">
    <a href="/postgraduate/2014/04/17/pager_video" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/17/pager_video" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-70"><br /></h2>

<h2>
    <a id="SDP协议及开发设计" href="/postgraduate/2014/04/17/pager_sdp" target="_blank">
        SDP协议及开发设计
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-17
</p>
<p class="description">SDP协议分析及开发设计</p>

<p class="read-all">
    <a href="/postgraduate/2014/04/17/pager_sdp" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/17/pager_sdp" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-71"><br /></h2>

<h2>
    <a id="实时流媒体协议RTSP" href="/postgraduate/2014/04/17/pager_rtsp" target="_blank">
        实时流媒体协议RTSP
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-17
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/postgraduate/2014/04/17/pager_rtsp" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/17/pager_rtsp" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-72"><br /></h2>

<h2>
    <a id="实时传输协议RTP" href="/postgraduate/2014/04/17/pager_rtp" target="_blank">
        实时传输协议RTP
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-17
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/postgraduate/2014/04/17/pager_rtp" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/17/pager_rtp" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-73"><br /></h2>

<h2>
    <a id="基于流媒体的对讲机系统——系统开发准备" href="/postgraduate/2014/04/17/pager_prepare" target="_blank">
        基于流媒体的对讲机系统——系统开发准备
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-17
</p>
<p class="description">系统开发需要的预备知识——工欲善其事，必先利其器</p>

<p class="read-all">
    <a href="/postgraduate/2014/04/17/pager_prepare" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/17/pager_prepare" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-74"><br /></h2>

<h2>
    <a id="基于流媒体的对讲机系统——Android的开发准备" href="/postgraduate/2014/04/17/pager_android_framework" target="_blank">
        基于流媒体的对讲机系统——Android的开发准备
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-17
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/postgraduate/2014/04/17/pager_android_framework" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/17/pager_android_framework" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-75"><br /></h2>

<h2>
    <a id="基于流媒体的对讲机系统的设计与实现" href="/postgraduate/2014/04/13/pager" target="_blank">
        基于流媒体的对讲机系统的设计与实现
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-04-13
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/postgraduate/2014/04/13/pager" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/04/13/pager" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-76"><br /></h2>

<h2>
    <a id="SIP协议的分析以及opensips注册和通话的研究" href="/postgraduate/2014/03/05/sip" target="_blank">
        SIP协议的分析以及opensips注册和通话的研究
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-03-05
</p>
<p class="description">毕设题目《基于流媒体的语音视频通话系统》，基于Android实现。关于SIP协议的分析以及opensips注册和通话的研究。</p>

<p class="read-all">
    <a href="/postgraduate/2014/03/05/sip" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/03/05/sip" target="_blank">
        <img src="http://mnhkahn.github.io/assets/images/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-77"><br /></h2>

<h2>
    <a id="iPhone的CSS显示" href="/kaleidoscope/2014/02/07/iphone_css" target="_blank">
        iPhone的CSS显示
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-02-07
</p>
<p class="description">iPhone的CSS显示，费了我stackoverflow 6个声望问到的。。。</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/02/07/iphone_css" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-78"><br /></h2>

<h2>
    <a id="Android Quick Start" href="/postgraduate/2014/02/05/android_quickstart" target="_blank">
        Android Quick Start
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-02-05
</p>
<p class="description">毕设题目《基于流媒体的语音视频通话系统》，基于Android实现。Android基础。</p>

<p class="read-all">
    <a href="/postgraduate/2014/02/05/android_quickstart" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/02/05/android_quickstart" target="_blank">
        <img src="http://cyeam.qiniudn.com/android.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-79"><br /></h2>

<h2>
    <a id="安卓对讲机开发评估" href="/postgraduate/2014/02/04/postgraduate_design_evaluate" target="_blank">
        安卓对讲机开发评估
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-02-04
</p>
<p class="description">毕设题目《基于流媒体的语音视频通话系统》，基于Android实现。先在这里做一下技术评估。</p>

<p class="read-all">
    <a href="/postgraduate/2014/02/04/postgraduate_design_evaluate" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/postgraduate/2014/02/04/postgraduate_design_evaluate" target="_blank">
        <img src="http://cyeam.qiniudn.com/c168.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-80"><br /></h2>

<h2>
    <a id="Linux Mint下安装Nexus 7 驱动" href="/kaleidoscope/2014/02/01/linux_android_drivers" target="_blank">
        Linux Mint下安装Nexus 7 驱动
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-02-01
</p>
<p class="description">为了能够让我的Nexus 7翻墙，我决定root。是在Linux Mint下root Android。看着也不难，结果搞了好几天。越到一半机子驱动问题，没办法继续了，看着一块砖头放在那，真叫个急啊。root成功后，twitter还是上不去，还得再写一篇文章来总结一下Android翻墙。</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/02/01/linux_android_drivers" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/kaleidoscope/2014/02/01/linux_android_drivers" target="_blank">
        <img src="http://cyeam.qiniudn.com/writing_udev_rules.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-81"><br /></h2>

<h2>
    <a id="对联" href="/kaleidoscope/2014/01/28/couplet" target="_blank">
        对联
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-01-28
</p>
<p class="description">一年一度春节又要到了， 我也要为家里写对联了，作为一年一次的练习。</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/01/28/couplet" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-82"><br /></h2>

<h2>
    <a id="Linux命令整理" href="/linux/2014/01/22/linux" target="_blank">
        Linux命令整理
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-01-22
</p>
<p class="description">用到的一些常见的Linux命令，在这里记录一下。</p>

<p class="read-all">
    <a href="/linux/2014/01/22/linux" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/linux/2014/01/22/linux" target="_blank">
        <img src="http://cyeam.qiniudn.com/shell.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-83"><br /></h2>

<h2>
    <a id="beego介绍" href="/golang/2014/01/21/beego" target="_blank">
        beego介绍
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-01-21
</p>
<p class="description">关于golang语言和beego框架的使用的介绍。golang所有的设计都是围绕减少行数展开。</p>

<p class="read-all">
    <a href="/golang/2014/01/21/beego" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/golang/2014/01/21/beego" target="_blank">
        <img src="http://cyeam.qiniudn.com/beego_logo.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-84"><br /></h2>

<h2>
    <a id="一些网站架构中常见的术语和技术" href="/framework/2014/01/14/meeting" target="_blank">
        一些网站架构中常见的术语和技术
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-01-14
</p>
<p class="description">1月10日参加公司后台架构的技术总结</p>

<p class="read-all">
    <a href="/framework/2014/01/14/meeting" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-85"><br /></h2>

<h2>
    <a id="2013读书总结" href="/notebook/2014/01/13/2013_reading" target="_blank">
        2013读书总结
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-01-13
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/notebook/2014/01/13/2013_reading" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/notebook/2014/01/13/2013_reading" target="_blank">
        <img src="http://cyeam.qiniudn.com/%E6%98%8E%E6%9C%9D%E9%82%A3%E4%BA%9B%E4%BA%8B%E5%84%BF.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-86"><br /></h2>

<h2>
    <a id="Markdown" href="/kaleidoscope/2014/01/12/markdown" target="_blank">
        Markdown
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-01-12
</p>
<p class="description">Markdown备忘。</p>

<p class="read-all">
    <a href="/kaleidoscope/2014/01/12/markdown" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/kaleidoscope/2014/01/12/markdown" target="_blank">
        <img src="http://cyeam.qiniudn.com/md.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-87"><br /></h2>

<h2>
    <a id="Linux Mint 64bit下安装Dota 2" href="/toss/2014/01/11/dota2" target="_blank">
        Linux Mint 64bit下安装Dota 2
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-01-11
</p>
<p class="description">现在的新电脑大都用的双显卡，一张Intel的集成显卡，一张Nvidia的独立显卡。默认运行集成显卡，在玩游戏这些需要大量图形计算的时候运行独立显卡。这个自动切换的过程在Windows和Mac环境下，都是由Nvidia的显卡驱动自动完成的，而在Linux下，伟大的Nvidia却不提供这样的切换功能了。所以Linus问候了它。</p>

<p class="read-all">
    <a href="/toss/2014/01/11/dota2" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/toss/2014/01/11/dota2" target="_blank">
        <img src="http://cyeam.qiniudn.com/Linus-Torvalds-Fuck-You-Nvidia.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-88"><br /></h2>

<h2>
    <a id="Java面试宝典" href="/java/2014/01/02/javacollection" target="_blank">
        Java面试宝典
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2014-01-02
</p>
<p class="description">从2013年9月开始找工作，在几个月的Java程序员求职过程中，总结了一些被问到的笔试题和面试题。Java语言博大精深，是整个程序界的上乘语言，应该得到重视。</p>

<p class="read-all">
    <a href="/java/2014/01/02/javacollection" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/java/2014/01/02/javacollection" target="_blank">
        <img src="http://cyeam.qiniudn.com/java_logo.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-89"><br /></h2>

<h2>
    <a id="电影《色|戒》最后王佳芝为什么没有吃下那颗为特工准备的自杀胶囊？" href="/ctalk/2013/11/25/lust_caution" target="_blank">
        电影《色|戒》最后王佳芝为什么没有吃下那颗为特工准备的自杀胶囊？
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-11-25
</p>
<p class="description">明知道说了那句“快走”，自己无论如何都是死了，为什么最后摸了摸胶囊，还是没吃呢。</p>

<p class="read-all">
    <a href="/ctalk/2013/11/25/lust_caution" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/ctalk/2013/11/25/lust_caution" target="_blank">
        <img src="http://cyeam.qiniudn.com/lust_caution.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-90"><br /></h2>

<h2>
    <a id="搜狗2014校园招聘笔试题" href="/collection/2013/11/23/sogou" target="_blank">
        搜狗2014校园招聘笔试题
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-11-23
</p>
<p class="description">11月23号去清华参加的笔试，那时候Java准备的还算充分了，笔试答的还可以，面试的时候看到Java和数据结构方面分别错了2个，很不错了。搜狗是一家非常好的公司，据说待遇是13.5x15，还带期权，马上要上市的公司。我也很想去。一面很顺利，一共面了20min，面试官说知道我的能力，让我直接进复试。当时我特别高兴，因为过了一面其实已经一只脚踩进了搜狗。二面的时候，面试官很强势，再加上我很期待进入这家公司，很是紧张。他直奔主题，问我看过哪些Java源代码，Java虚拟机垃圾回收具体用的是什么方法，后台服务器解决并发请求时常见的工具和解决方案是什么。这些都答的马马虎虎。最后一个算法，有10G的文件，里面保存着网站的黑名单，怎么样在处理用户请求时查找到该用户时候存在于黑名单中。此算法我写了半个多小时，勉强写完。我和搜狗的情节也就此画上了句号。一方面自己能力还是有限，另一方面，其实我还是有机会进入这种牛逼互联网公司的，加油。最后附上当时拍的搜狗的照片。</p>

<p class="read-all">
    <a href="/collection/2013/11/23/sogou" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/11/23/sogou" target="_blank">
        <img src="http://cyeam.qiniudn.com/sogou_logo.gif" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-91"><br /></h2>

<h2>
    <a id="中科曙光2014校园招聘笔试" href="/collection/2013/11/05/sugon" target="_blank">
        中科曙光2014校园招聘笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-11-05
</p>
<p class="description">算是国企里面非常厉害的了，而且效益也不错。但是今年没怎么招人。</p>

<p class="read-all">
    <a href="/collection/2013/11/05/sugon" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/11/05/sugon" target="_blank">
        <img src="http://cyeam.qiniudn.com/sugou.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-92"><br /></h2>

<h2>
    <a id="方正国际2014校园招聘笔试" href="/collection/2013/10/30/founder" target="_blank">
        方正国际2014校园招聘笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-30
</p>
<p class="description">这套笔试题当时做的时候觉得很难，还涉及到了一些架构的东西，当时都不会。现在看来，也就那样，熟悉Java后台开发的应该问题不大。后来还去面过方正科技，就在北大旁边那个，号称是方正的电脑和一些高科技产品都是这个子公司生产的。还搞了群面，一群人在那胡扯。笔试很水，全是C#题目，难度很低。这个公司感觉不怎么样了。</p>

<p class="read-all">
    <a href="/collection/2013/10/30/founder" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/30/founder" target="_blank">
        <img src="http://cyeam.qiniudn.com/founder.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-93"><br /></h2>

<h2>
    <a id="神州航天软件2014校园招聘笔试" href="/collection/2013/10/25/bsast" target="_blank">
        神州航天软件2014校园招聘笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-25
</p>
<p class="description">神州航天软件2014校园招聘笔试。该公司在航天城，工资每月6000，带1000补贴，解决北京户口。待遇还是不错的。一次签5年。</p>

<p class="read-all">
    <a href="/collection/2013/10/25/bsast" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/25/bsast" target="_blank">
        <img src="http://cyeam.qiniudn.com/bsast.gif" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-94"><br /></h2>

<h2>
    <a id="乐视2014校园招聘笔试" href="/collection/2013/10/24/letv" target="_blank">
        乐视2014校园招聘笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-24
</p>
<p class="description">有一个工作人员穿着youtobe的衣服，应该是从Google online store买到的，看着挺牛逼的，^_^。公司目前在发展乐视盒子和乐视电视，进军客厅。</p>

<p class="read-all">
    <a href="/collection/2013/10/24/letv" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/24/letv" target="_blank">
        <img src="http://cyeam.qiniudn.com/letv_logo.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-95"><br /></h2>

<h2>
    <a id="趣游2014校园招聘笔试" href="/collection/2013/10/24/gamewave" target="_blank">
        趣游2014校园招聘笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-24
</p>
<p class="description">好像是做手游的公司。一个创业公司。</p>

<p class="read-all">
    <a href="/collection/2013/10/24/gamewave" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/24/gamewave" target="_blank">
        <img src="http://cyeam.qiniudn.com/gamewave.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-96"><br /></h2>

<h2>
    <a id="巨人网络2014校园招聘笔试" href="/collection/2013/10/21/giant" target="_blank">
        巨人网络2014校园招聘笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-21
</p>
<p class="description">巨人网络2014校园招聘笔试。据说笔试通过的会去上海参加面试。我果断没去成。</p>

<p class="read-all">
    <a href="/collection/2013/10/21/giant" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/21/giant" target="_blank">
        <img src="http://cyeam.qiniudn.com/giant_logo.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-97"><br /></h2>

<h2>
    <a id="去哪网2014校园招聘Java开发面经" href="/collection/2013/10/19/qunarinterview" target="_blank">
        去哪网2014校园招聘Java开发面经
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-19
</p>
<p class="description">去哪网待遇非常给力，一个FE都能给到14x16。而且刚刚上市，发展很不错。我的面试官是重庆人，看到我的简历上写着本科是重庆大学的，也很照顾我，可惜我很不给力。一直摊开了问，从基础到Java源码、C++源码都问了一圈，还有C语言小技巧。面试难度符合他们公司的工资水平。</p>

<p class="read-all">
    <a href="/collection/2013/10/19/qunarinterview" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/19/qunarinterview" target="_blank">
        <img src="http://cyeam.qiniudn.com/qunar_logo.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-98"><br /></h2>

<h2>
    <a id="完美世界2014校园招聘笔试" href="/collection/2013/10/19/perfectworld" target="_blank">
        完美世界2014校园招聘笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-19
</p>
<p class="description">只是简单了参加了一下笔试，地点是在清华。没有面试机会。这种游戏公司招人要求都有点高。</p>

<p class="read-all">
    <a href="/collection/2013/10/19/perfectworld" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/19/perfectworld" target="_blank">
        <img src="http://cyeam.qiniudn.com/perfectworld_logo.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-99"><br /></h2>

<h2>
    <a id="广联达2014校园招聘笔试" href="/collection/2013/10/17/glodon" target="_blank">
        广联达2014校园招聘笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-17
</p>
<p class="description">公司是做建筑方面的建模工具产品。没有面试机会。</p>

<p class="read-all">
    <a href="/collection/2013/10/17/glodon" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/17/glodon" target="_blank">
        <img src="http://cyeam.qiniudn.com/glodon_logo.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-100"><br /></h2>

<h2>
    <a id="触控科技2014校园招聘笔试题" href="/collection/2013/10/15/chukong" target="_blank">
        触控科技2014校园招聘笔试题
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-15
</p>
<p class="description">触控科技笔试题，也参加了面试，跟我讲说11月份会通知我，结果也没消息了。公司不大，老板爱喝茶，办公室里全是专业的喝茶器具。之前收购了一个开源游戏引擎Cocos2d。貌似最近又搞来了小鳄鱼爱洗澡，传闻也快上市了。发展不错，不过前面还有可怕的腾讯。。。</p>

<p class="read-all">
    <a href="/collection/2013/10/15/chukong" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/15/chukong" target="_blank">
        <img src="http://cyeam.qiniudn.com/chukong_logo.png" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-101"><br /></h2>

<h2>
    <a id="百度2014校园招聘后台开发笔试" href="/collection/2013/10/13/baidu" target="_blank">
        百度2014校园招聘后台开发笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-10-13
</p>
<p class="description">在北大参加的笔试，我觉得答得非常好，除了最后一题都会做。没有面试机会。。。</p>

<p class="read-all">
    <a href="/collection/2013/10/13/baidu" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/10/13/baidu" target="_blank">
        <img src="http://cyeam.qiniudn.com/bdlogo.gif" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-102"><br /></h2>

<h2>
    <a id="阿里巴巴2014校园招聘面经" href="/collection/2013/09/16/alibabainterview" target="_blank">
        阿里巴巴2014校园招聘面经
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-09-16
</p>
<p class="description">本来准备先面一些小公司，积累点经验，再去面自己想去的公司的。结果才试了一家，阿里就来面试了。而且我觉得我的水平确实不够阿里的层次。虽然我开发经验比较多一些，但是他们更关注你对底层实现的理解，说来惭愧，一直在做东西，没有时间关注这些。这次面试，就我本人而言，没有任何亮点，没有说出来一个让人觉得还可以的东西，所以没戏。</p>

<p class="read-all">
    <a href="/collection/2013/09/16/alibabainterview" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/09/16/alibabainterview" target="_blank">
        <img src="http://cyeam.qiniudn.com/alibaba.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-103"><br /></h2>

<h2>
    <a id="神州航天软件发展规划部实习笔试" href="/collection/2013/09/09/bsastintern" target="_blank">
        神州航天软件发展规划部实习笔试
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-09-09
</p>
<p class="description">9月初的实习笔试，后面还有面试，通过了。果断拒了。</p>

<p class="read-all">
    <a href="/collection/2013/09/09/bsastintern" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-104"><br /></h2>

<h2>
    <a id="阿里巴巴2013年实习生笔试题" href="/collection/2013/05/23/alibaba" target="_blank">
        阿里巴巴2013年实习生笔试题
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-05-23
</p>
<p class="description">答题说明：1.答题时间90分钟，请注意把握时间；2.试题分为四个部分：单项选择题（10题，20分）、不定向选择题（4题，20分）、填空问答（5题，40分）、综合体（1题，20分）。</p>

<p class="read-all">
    <a href="/collection/2013/05/23/alibaba" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-105"><br /></h2>

<h2>
    <a id="微软2012年实习生笔试题" href="/collection/2013/04/20/microsoft_intern_2012" target="_blank">
        微软2012年实习生笔试题
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-04-20
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/collection/2013/04/20/microsoft_intern_2012" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/collection/2013/04/20/microsoft_intern_2012" target="_blank">
        <img src="http://cyeam.qiniudn.com/microsoft_bldg.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-106"><br /></h2>

<h2>
    <a id="快速排序(QUICK SORT)" href="/computer/2013/04/17/quicksort" target="_blank">
        快速排序(QUICK SORT)
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-04-17
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/computer/2013/04/17/quicksort" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-107"><br /></h2>

<h2>
    <a id="冒泡排序(BUBBLE SORT)" href="/computer/2013/04/16/bubblesort" target="_blank">
        冒泡排序(BUBBLE SORT)
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-04-16
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/computer/2013/04/16/bubblesort" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-108"><br /></h2>

<h2>
    <a id="微软2013年夏季实习生笔试题" href="/collection/2013/04/07/microsoft_intern_2013" target="_blank">
        微软2013年夏季实习生笔试题
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-04-07
</p>
<p class="description">2013 Microsoft Campus Intern Hiring Written Test</p>

<p class="read-all">
    <a href="/collection/2013/04/07/microsoft_intern_2013" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-109"><br /></h2>

<h2>
    <a id="堆排序(HEAP SORT)" href="/computer/2013/04/06/heapsort" target="_blank">
        堆排序(HEAP SORT)
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-04-06
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/computer/2013/04/06/heapsort" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-110"><br /></h2>

<h2>
    <a id="3种交换值的方法" href="/computer/2013/04/02/swap" target="_blank">
        3种交换值的方法
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-04-02
</p>
<p class="description">交换值是比较常用的步骤，也比较简单，这里总结了3个方法。</p>

<p class="read-all">
    <a href="/computer/2013/04/02/swap" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-111"><br /></h2>

<h2>
    <a id="简单选择排序(SELECT SORT)" href="/computer/2013/04/02/selectsort" target="_blank">
        简单选择排序(SELECT SORT)
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-04-02
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/computer/2013/04/02/selectsort" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-112"><br /></h2>

<h2>
    <a id="构造函数成员初始化列表" href="/computer/2013/03/31/constructor" target="_blank">
        构造函数成员初始化列表
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-31
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/computer/2013/03/31/constructor" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-113"><br /></h2>

<h2>
    <a id="蒙特卡洛求圆周率" href="/kaleidoscope/2013/03/27/pi" target="_blank">
        蒙特卡洛求圆周率
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-27
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/kaleidoscope/2013/03/27/pi" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-114"><br /></h2>

<h2>
    <a id="STATIC数据成员" href="/computer/2013/03/21/static" target="_blank">
        STATIC数据成员
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-21
</p>
<p class="description">如果要在程序任意点统计已创建特定类型的数量，就需要一个全局变量。但是全局变量会破坏封装。static数据成员可以解决这个问题。</p>

<p class="read-all">
    <a href="/computer/2013/03/21/static" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-115"><br /></h2>

<h2>
    <a id="字符数组、字符指针、字符串" href="/computer/2013/03/20/string" target="_blank">
        字符数组、字符指针、字符串
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-20
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/computer/2013/03/20/string" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-116"><br /></h2>

<h2>
    <a id="SIZEOF操作符" href="/computer/2013/03/20/sizeof" target="_blank">
        SIZEOF操作符
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-20
</p>
<p class="description"> sizeof是C/C++中的单目操作符，在编译期间计算，测量这个变量占用的内存大小，它并不是函数。</p>

<p class="read-all">
    <a href="/computer/2013/03/20/sizeof" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-117"><br /></h2>

<h2>
    <a id="内存分配方式" href="/computer/2013/03/20/memory" target="_blank">
        内存分配方式
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-20
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/computer/2013/03/20/memory" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-118"><br /></h2>

<h2>
    <a id="CONST关键字" href="/computer/2013/03/14/const" target="_blank">
        CONST关键字
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-14
</p>
<p class="description">const关键字可以用来限定变量不可以被修改，基于这个特性，程序员可以用这个关键字来提高程序的健壮性，这也是程序员面试常见考点。</p>

<p class="read-all">
    <a href="/computer/2013/03/14/const" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-119"><br /></h2>

<h2>
    <a id="位运算符(Bitwise Operators)" href="/computer/2013/03/11/bitwise_operators" target="_blank">
        位运算符(Bitwise Operators)
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-11
</p>
<p class="description">位运算符的一些介绍。</p>

<p class="read-all">
    <a href="/computer/2013/03/11/bitwise_operators" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-120"><br /></h2>

<h2>
    <a id="又一年" href="/notebook/2013/03/05/yet_another_year" target="_blank">
        又一年
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-03-05
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/notebook/2013/03/05/yet_another_year" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/notebook/2013/03/05/yet_another_year" target="_blank">
        <img src="http://cyeam.qiniudn.com/north_drift.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-121"><br /></h2>

<h2>
    <a id="你永远都不知道自己有多傻逼" href="/notebook/2013/01/23/u_will_never_know_how_stupid_u_r" target="_blank">
        你永远都不知道自己有多傻逼
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-01-23
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/notebook/2013/01/23/u_will_never_know_how_stupid_u_r" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/notebook/2013/01/23/u_will_never_know_how_stupid_u_r" target="_blank">
        <img src="http://cyeam.qiniudn.com/north_drift.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-122"><br /></h2>

<h2>
    <a id="Connect the dots" href="/notebook/2013/01/23/connect_the_dots" target="_blank">
        Connect the dots
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-01-23
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/notebook/2013/01/23/connect_the_dots" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/notebook/2013/01/23/connect_the_dots" target="_blank">
        <img src="http://cyeam.qiniudn.com/north_drift.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-123"><br /></h2>

<h2>
    <a id="2012读书总结" href="/notebook/2013/01/13/2012_reading" target="_blank">
        2012读书总结
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2013-01-13
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/notebook/2013/01/13/2012_reading" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/notebook/2013/01/13/2012_reading" target="_blank">
        <img src="http://cyeam.qiniudn.com/%E4%B9%A1%E5%85%B3%E4%BD%95%E5%A4%84.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-124"><br /></h2>

<h2>
    <a id="请求网页所需协议" href="/network/2012/11/02/Protocol_Of_HTTP" target="_blank">
        请求网页所需协议
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2012-11-02
</p>
<p class="description">网页一次请求用的协议和访问流程的总结。</p>

<p class="read-all">
    <a href="/network/2012/11/02/Protocol_Of_HTTP" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/network/2012/11/02/Protocol_Of_HTTP" target="_blank">
        <img src="http://cyeam.qiniudn.com/http.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-125"><br /></h2>

<h2>
    <a id="沉淀心情" href="/notebook/2012/05/23/precipitation_mood" target="_blank">
        沉淀心情
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2012-05-23
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/notebook/2012/05/23/precipitation_mood" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/notebook/2012/05/23/precipitation_mood" target="_blank">
        <img src="http://cyeam.qiniudn.com/bit_logo.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-126"><br /></h2>

<h2>
    <a id="很多事情感觉都来不及计划" href="/notebook/2012/03/26/feeling_a_lot_of_things_are_too_late_to_plan" target="_blank">
        很多事情感觉都来不及计划
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2012-03-26
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/notebook/2012/03/26/feeling_a_lot_of_things_are_too_late_to_plan" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/notebook/2012/03/26/feeling_a_lot_of_things_are_too_late_to_plan" target="_blank">
        <img src="http://cyeam.qiniudn.com/postgraduate_exam.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-127"><br /></h2>

<h2>
    <a id="在main函数中创建2个线程，一个运行冒泡排序算法，另一个运行选择排序算法，分别对两个一样的数组进行" href="/computer/2011/01/11/cpp_thread" target="_blank">
        在main函数中创建2个线程，一个运行冒泡排序算法，另一个运行选择排序算法，分别对两个一样的数组进行
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2011-01-11
</p>
<p class="description">操作系统作业。</p>

<p class="read-all">
    <a href="/computer/2011/01/11/cpp_thread" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-128"><br /></h2>

<h2>
    <a id="用CreateProcess创建线程，指定运行QQ.EXE" href="/computer/2011/01/11/cpp_process" target="_blank">
        用CreateProcess创建线程，指定运行QQ.EXE
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2011-01-11
</p>
<p class="description">操作系统作业。</p>

<p class="read-all">
    <a href="/computer/2011/01/11/cpp_process" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-129"><br /></h2>

<h2>
    <a id="DX基础知识" href="/directx/2011/01/09/dx" target="_blank">
        DX基础知识
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2011-01-09
</p>
<p class="description">DirectX学习笔记。</p>

<p class="read-all">
    <a href="/directx/2011/01/09/dx" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-130"><br /></h2>

<h2>
    <a id="动画与游戏设计基础" href="/game/2010/12/29/game" target="_blank">
        动画与游戏设计基础
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2010-12-29
</p>
<p class="description">游戏开发学习笔记。</p>

<p class="read-all">
    <a href="/game/2010/12/29/game" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-131"><br /></h2>

<h2>
    <a id="3Q大战" href="/ctalk/2010/11/03/360" target="_blank">
        3Q大战
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2010-11-03
</p>
<p class="description"></p>

<p class="read-all">
    <a href="/ctalk/2010/11/03/360" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<p class="figure center">
    <a href="/ctalk/2010/11/03/360" target="_blank">
        <img src="http://cyeam.qiniudn.com/3q_fight.jpg" alt="IMG-THUMBNAIL" />
    </a>
</p>

<h2 id="br-132"><br /></h2>

<h2>
    <a id="我的阿凡达" href="/kaleidoscope/2010/05/31/avatar" target="_blank">
        我的阿凡达
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2010-05-31
</p>
<p class="description">当前读数字媒体专业学习的P图和flash作业。。。雷死人了</p>

<p class="read-all">
    <a href="/kaleidoscope/2010/05/31/avatar" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-133"><br /></h2>

<h2>
    <a id="又要走了，可是不怎么想走啊" href="/notebook/2009/02/16/do_not_wanna_leave" target="_blank">
        又要走了，可是不怎么想走啊
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2009-02-16
</p>
<p class="description">当年还想着拿奖学金的，这辈子没戏了。。。</p>

<p class="read-all">
    <a href="/notebook/2009/02/16/do_not_wanna_leave" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-134"><br /></h2>

<h2>
    <a id="又过了一年" href="/notebook/2008/12/31/another_year" target="_blank">
        又过了一年
    </a>
</h2>
<p class="date">
    <span class="icon-calendar">
    </span>
    2008-12-31
</p>
<p class="description">2008，就这么过去了，我还要加油啊。努力吧，Superman.</p>

<p class="read-all">
    <a href="/notebook/2008/12/31/another_year" target="_blank">
        <span class="icon-resize-full">
        </span>
        阅读全文
    </a>
</p>

<h2 id="br-135"><br /></h2>

  </div>
</div>


      </div>

<!--
      <script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
      <ins class="adsbygoogle"
           style="display:inline-block;width:728px;height:90px"
           data-ad-client="ca-pub-1651120361108148"
           data-ad-slot="4548335295"></ins>
      <script>
      (adsbygoogle = window.adsbygoogle || []).push({});
      </script>
-->  

<script type="text/javascript">
    /*234*60 创建于 2015-02-11*/
    var cpro_id = "u1949560";
</script>
<script src="http://cpro.baidustatic.com/cpro/ui/c.js" type="text/javascript"></script>

      <hr>
      <footer>
    <script type="text/javascript">
      (function(w,d,t,u,n,s,e){w['SwiftypeObject']=n;w[n]=w[n]||function(){
      (w[n].q=w[n].q||[]).push(arguments);};s=d.createElement(t);
      e=d.getElementsByTagName(t)[0];s.async=1;s.src=u;e.parentNode.insertBefore(s,e);
      })(window,document,'script','//s.swiftypecdn.com/install/v1/st.js','_st');

      _st('install','wU4JEZ3TeKgM-qap8TxG');
    </script>
        <p>&copy; 2015 Bryce
          with help from <a href="http://jekyllbootstrap.com" target="_blank" title="The Definitive Jekyll Blogging Framework">Jekyll Bootstrap</a>
          and <a href="http://twitter.github.com/bootstrap/" target="_blank">Twitter Bootstrap</a>
        </p>
    <a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="Creative Commons License" style="border-width:0" src="http://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br /><span xmlns:dct="http://purl.org/dc/terms/" href="http://purl.org/dc/dcmitype/Dataset" property="dct:title" rel="dct:type">Cyeam</span> by <a xmlns:cc="http://creativecommons.org/ns#" href="www.cyeam.com" property="cc:attributionName" rel="cc:attributionURL">Bryce</a> is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.

        


  <script type="text/javascript">
  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-50349737-1']);
  _gaq.push(['_trackPageview']);

  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();
</script>



        <script type="text/javascript">var cnzz_protocol = (("https:" == document.location.protocol) ? " https://" : " http://");document.write(unescape("%3Cspan id='cnzz_stat_icon_1000474528'%3E%3C/span%3E%3Cscript src='" + cnzz_protocol + "s19.cnzz.com/z_stat.php%3Fid%3D1000474528%26show%3Dpic1' type='text/javascript'%3E%3C/script%3E"));</script>
      </footer>

    </div>

  </body>
</html>
`

var raw = `
<!DOCTYPE html>
<html lang="en" xmlns:wb="http://open.weibo.com/wb">
  <head>
    <meta charset="utf-8">
    <title>垂直搜索爬虫——maodou</title>
    <meta name="description" content="用Golang实现的搜索引擎爬虫maodou。">
    <meta name="author" content="Bryce">

    <meta name="google-translate-customization" content="a4136e955b3e09f2-45a74b56dc13e741-gf616ffda6e6360e0-11"></meta>

    <!-- Enable responsive viewport -->
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <!-- Le HTML5 shim, for IE6-8 support of HTML elements -->
    <!--[if lt IE 9]>
      <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->

    <!-- Le styles -->
    <link href="/assets/themes/twitter/bootstrap/css/bootstrap.2.2.2.min.css" rel="stylesheet">
    <link href="/assets/themes/twitter/css/style.css?body=1" rel="stylesheet" type="text/css" media="all">

    <!-- Le fav and touch icons -->
  <!-- Update these with your own images -->
    <link rel="shortcut icon" href="/assets/images/c32.ico">
  <!--  <link rel="apple-touch-icon" href="images/apple-touch-icon.png">
    <link rel="apple-touch-icon" sizes="72x72" href="images/apple-touch-icon-72x72.png">
    <link rel="apple-touch-icon" sizes="114x114" href="images/apple-touch-icon-114x114.png">
  -->

    <!-- atom & rss feed -->
    <link href="/atom.xml" type="application/atom+xml" rel="alternate" title="Sitewide ATOM Feed">
    <link href="/rss.xml" type="application/rss+xml" rel="alternate" title="Sitewide RSS Feed">

    <!-- jQuery -->
     <!--<script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>-->
    <script src="http://libs.baidu.com/jquery/1.7.2/jquery.min.js"></script>
    <script>!window.jQuery && document.write('<script src="/assets/js/jquery-1.7.2.min.js"><\/script>');</script>
  <!--<script type="text/javascript" src="http://app.stacktack.com/jquery.stacktack.min.js"></script>-->
  <script type="text/javascript" src="/assets/js/cyeam.js"></script>
  <script type="text/javascript">
    $().ready(function(){
      //wallpaper();
      // $(document).stacktack();
    })
  </script>
  </head>

  <body>
    <div class="navbar">
      <div class="navbar-inner">
        <div class="container-narrow">
          <div class="wrapperLogo" id="wrapperLogo" style="float: left; width: 80px;">
            <a href="/"><img src="/assets/images/logo.jpg" style="width:50px"></a>
          </div>
          <ul class="nav">
               <li><a href="/"><span class="icon-home"></span> Home</a></li>
            <li><a href="/categories.html"><span class="icon-th-list"></span> Categories</a></li>
            <li><a href="/archive.html"><span class="icon-ok"></span> Archive</a></li>
            <li><a href="/tags.html"><span class="icon-tags"></span> Tags</a></li>
            <li><a href="/pages.html"><span class="icon-th"></span> Pages</a></li>
            <!--<li><a href="/cyeam.html"><span class="icon-bookmark"></span> Cyeam</a></li>-->
            <!--
           
           


 
   
     
          
           <li><a href="/archive.html">Archive</a></li>
          
     
   
 
   
     
   
 
   
     
          
           <li><a href="/categories.html">Categories</a></li>
          
     
   
 
   
     
   
 
   
     
          
           <li><a href="/pages.html">Pages</a></li>
          
     
   
 
   
     
   
 
   
     
   
 
   
 
   
     
          
           <li><a href="/tags.html">Tags</a></li>
          
     
   
 


-->
          </ul>
        </div>
      </div>
    </div>

    <div id="container-narrow" class="container-narrow opacity">
     
      <div class="content" id="content">
       
<div class="page-header c6">
  <h1>垂直搜索爬虫——maodou </h1>
</div>

<div class="row-fluid post-full">
  <div class="span12">
    <div class="date">
      <span class="icon-calendar">
      </span>
      <span>05 April 2015</span>
    </div>
    
     <blockquote>用Golang实现的搜索引擎爬虫maodou。</blockquote>
    <div class="content">
      <p>去年开始学习搜索引擎，目前主要研究的是爬虫这块。主要是因为爬虫相对简单一些，索引这块不是很了解，只能借助于Solr实现。</p>

<p>之前文档<a href="http://blog.cyeam.com/se/2014/12/26/search_engine/">《如何开发搜索引擎——爬虫（一）》</a>介绍了一个比较传统的抓取方式。就是把互联网当作是一张有向图，通过遍历这张图来抓取网站。后来也实现了抓取的爬虫，但一直没有啥实际用途，因为这么抓只能做成Google那种传统搜索引擎。</p>

<p>后来发现了大神binux的垂直搜索爬虫<a href="https://github.com/binux/pyspider">pyspider</a>，用Python写的一个垂直抓取框架。试用了一翻，确实很好用，还有他的<a href="http://demo.pyspider.org/">demo</a>。这是他抓取豆瓣害羞组的<a href="https://f.binux.me/haixiuzu.html">效果页</a>。看到这个才发现，做这个挺有意思的。</p>

<p>Python基本不会，索性就用Golang写了，看了看他的抓取效果，想着自己基于也有的知识也是能够实现的，就用Golang写了一个抓取框架<a href="https://github.com/mnhkahn/maodou">maodou</a>（maodou是媳妇的名字）。抓取害羞组的代码如下，我也是将结果存在了<a href="http://duoshuo.com/">多说</a>。</p>

<pre><code>package main

import (
     "github.com/PuerkitoBio/goquery"
     "github.com/mnhkahn/maodou"
     "github.com/mnhkahn/maodou/cygo"
     "github.com/mnhkahn/maodou/dao"
     "github.com/mnhkahn/maodou/models"
     "strings"
)

type Haixiu struct {
     maodou.MaoDou
}

func (this *Haixiu) Start() {
     this.Index(this.Cawl("http://www.douban.com/group/haixiuzu/discussion"))
}

func (this *Haixiu) Index(resp *maodou.Response) {
     resp.Doc(#content &gt; div &gt; div.article &gt; div:nth-child(2) &gt; table &gt; tbody &gt; tr &gt; td.title &gt; a).Each(func(i int, s *goquery.Selection) {
          href, has := s.Attr("href")
          if has {
               this.Detail(this.Cawl(href))
          }
     })
}

func (this *Haixiu) Detail(resp *maodou.Response) {
     res := new(models.Result)
     res.Id = strings.Split(resp.Url, "/")[5]
     res.Title = resp.Doc("#content &gt; h1").Text()
     res.Author = resp.Doc("#content &gt; div &gt; div.article &gt; div.topic-content.clearfix &gt; div.topic-doc &gt; h3 &gt; span.from &gt; a").Text()
     res.Figure, _ = resp.Doc("#link-report &gt; div.topic-content &gt; div.topic-figure.cc &gt; img").Attr("src")
     res.Link = resp.Url
     res.Source = "www.douban.com/group/haixiuzu/discussion"
     res.ParseDate = cygo.Now()
     res.Description = "haixiuzu"
     this.Result(res)
}

func (this *Haixiu) Result(result *models.Result) {
     if result.Figure != "" {
          Dao, err := dao.NewDao("duoshuo", {"short_name":"cyeam","secret":"df66f048bd56cba5bf219b51766dec0d"})
          if err != nil {
               panic(err)
          }
          Dao.AddResult(result)
     }
}

func main() {
     maodou.Register(new(Haixiu), 30)
}
</code></pre>

<p>介绍一下抓取害羞租时候遇到的问题。抓取的难度就在于模拟浏览。因为小组里的东西肯定是不让随便抓走的，所以就要防止机器抓取。</p>

<ol>
  <li>第一个是请求头里面的<code>UserAgent</code>，本来想写自己的<code>Cyeambot</code>，发现不行，人家会拦掉不认识的设备；</li>
  <li>第二个是请求头里面的<code>Referer</code>。这个代表了当前请求来源页面。如果是空，就认为是直接在浏览器打开的，或者是从书签打开的。。如果请求的来源一直是空，就有可能是爬虫。但是具体问题具体分析，比如害羞组的首页<a href="http://www.douban.com/group/haixiuzu/discussion">http://www.douban.com/group/haixiuzu/discussion</a>，常用用户肯定会把这个地址保存在书签里面，所以前面的防爬虫机制尽量不起作用，而里面的帖子，<a href="http://www.douban.com/group/topic/73646091/">http://www.douban.com/group/topic/73646091/</a>，帖子地址是不可预估的，理论上只能通过前面讨论组的地址进入，那么如果打开具体帖子的<code>Referer</code>也还是空，这样就可以认为是爬虫。我抓取的时候就有遇到这样的问题，有时候会有<code>302</code>的情况出现，就是访问帖子，有一定的概率会跳转到登陆页，让你登陆来防范爬虫。这样应对也不难，就是抓取帖子的时候把<code>Referer</code>加上害羞组的地址即可；</li>
  <li>关于访问次数限制，害怕被封杀IP，我都是没30分钟抓取一次，每访问一次就等待5秒钟，尽量模拟得像人一些。这样没半个小时会访问21次，量不算大。程序运行了快一个月了，IP没有被拦截的迹象。之前抓取CSDN，没有休息5秒钟这一步，爬了一周就不能爬了。。。</li>
</ol>

<p>最后再说一些没用到的但是很牛逼的技术：</p>

<ol>
  <li>IP代理抓取。如果你有21台代理机器，那么上面的抓取可以同时进行。这样能大大降低抓取时间。这一般是企业级爬虫要用到的，对抓取速度有要求，像我这种半小时抓21页的就不需要用这个。如果有代理服务器，抓取豆瓣所有帖子也可以很快。当然，这样也会有问题，就是你抓取这么多东西就会有被封杀IP的风险了，所以大规模抓取还需要定时更换代理IP。其实这招可以被认为是<strong><em>黑科技</em></strong>，就是去搞中了病毒的电脑（俗称肉鸡），用这个进行代理，一般网上都有卖的；</li>
  <li>关于封杀IP。你只要别太恶意去抓，一般也是不会去封你IP的。现在都还是IPv4，IP地址是有限的，而设备越来越多，封杀一个IP有可能导致整个小区都没办法访问该网站了，所以也不用特别担心；</li>
  <li>如何更像人去抓取。有一个模拟器，叫做selenium。我发现Golang也有相应的客户端<a href="https://github.com/tebeka/selenium">tebeka - selenium</a>。用这个，你模拟登陆，自动记录你的<code>Cookie</code>并附带到抓取请求头里面，这样有了登陆信息，就更难进行封杀了。</li>
</ol>

<p>最后就是展示部分，<a href="https://www.cyeam.com/haixiuzu">点这里访问</a>。关于前端图片瀑布流的东西我一窍不通，完全抄袭了binux大神写的前端代码，请大神原谅。。。此爬虫并没有去抓取豆瓣小组里面的图片，只是记录了地址，所以在展示的时候还是需要访问豆瓣的图片地址，这就属于盗链图片。豆瓣肯定会对此进行处理。</p>

<p>你把豆瓣的图片放在你自己的网页上面展示，这就属于盗链了。浏览器在加载完HTML之后会去加载图片，这时图片请求会自动追加<code>Referer</code>，也就是你自己网页的地址，豆瓣的图片服务器发现请求来源是其他站点的，就会取消响应。</p>

<ol>
  <li>
    <p>盗链图片有两种方式，第一种是在网页里面加<code>iframe</code>，这个标签能新建一个嵌套的HTML，这样浏览器在请求图片的时候不会自动追加<code>Referer</code>，图片就能打开了；</p>
  </li>
  <li>
    <p>还有一种方式，就是用HTTPS发送请求。你把自己站点做成HTTPS的请求，这样在请求图片的时候，浏览器会把<code>Referer</code>去掉，保护用户隐私。搭建HTTPS服务器需要SSL证书，一般是一年$10。我去申请一个免费的，人家就是不同意。无奈之下，只能自己给自己颁发证书。坏处就是，用户第一次登陆我的站点的时候，浏览器会提示你不安全。铁道部不是也这样么。。。盗链图片的详细说明请参考最下面的两篇引用文章。</p>
  </li>
</ol>

<hr />

<p>最近GitHub被DDos了，我的博客访问量也有点降低了。本来量就不多。。。是不是需要把博客牵走了。</p>

<hr />

<p>最近有个同事离职了，跟我一起入职的，有点惋惜。可能互联网就是这样。</p>

<hr />

<h6 id="section"><em>参考文献</em></h6>
<ol>
  <li><a href="http://bindog.github.io/blog/2014/11/18/http-referer-security-and-anti-anti-hotlink/">《HTTP Referer安全与反反盗链》</a></li>
  <li><a href="http://www.zhujun.org/web/nginx-selfsign-ssl-cert/">《Nginx下绑定自己签发的免费SSL数字安全证书》 - 朱俊</a></li>
</ol>


    </div>
    
     <p>–<abbr title="End of file">EOF</abbr>–</p>

 
    <ul class="tag_box inline">
      <li><i class="icon-folder-open"></i></li>
     
     


 
    
         <li><a href="/categories.html#se-ref">
              se <span>2</span>
         </a></li>
   
 


    </ul>
   

 
    <ul class="tag_box inline">
      <li><i class="icon-tags"></i></li>
     
     


 
    
         <li><a href="/tags.html#Search Engine-ref">Search Engine <span>2</span></a></li>
   
 



    </ul>
  

     <div id="aboutme">
        <h3 id="about_me">关于作者</h3>
        <p>Bryce，80后程序员，从没想过改变世界，只是想快快乐乐的码东西。这是我的个人博客，在这里写点小技术，说点小心请，谈点小想法。观点比较偏激，还请大家海涵。</p>
        <p>最近想做一些微信服务器开发，需要进行认证。条件就差500粉丝。如果大家觉得我的文章还可以的话，劳驾关注一下我的个人公共微信，作为对在下的奖励。</p>
        <img src="http://cyeam.qiniudn.com/qrcode_for_gh_aa23c6563b3d_258.jpg" alt="Alt text">
      </div>
      <div class="bdsharebuttonbox"><a href="#" class="bds_tsina" data-cmd="tsina" title="Weibo"></a><a href="#" class="bds_weixin" data-cmd="weixin" title="WeChat"></a><a href="#" class="bds_linkedin" data-cmd="linkedin" title="Linkedin"></a><a href="#" class="bds_twi" data-cmd="twi" title="Twitter"></a><a href="#" class="bds_fbook" data-cmd="fbook" title="Facebook"></a></div>
<script>window._bd_share_config={"common":{"bdSnsKey":{},"bdText":"","bdMini":"1","bdMiniList":false,"bdPic":"http://cyeam.qiniudn.com/qrcode_for_gh_aa23c6563b3d_258.jpg","bdStyle":"2","bdSize":"24"},"share":{}};with(document)0[(getElementsByTagName('head')[0]||body).appendChild(createElement('script')).src='http://bdimg.share.baidu.com/static/api/js/share.js?v=89860593.js?cdnversion='+~(-new Date()/36e5)];</script>

    <hr>
    <div class="pagination">
      <ul>
     
        <li class="prev"><a href="/network/2015/03/16/fing" title="探测局域网里面的设备">&larr; Previous</a></li>
     
        <li><a href="/archive.html">Archive</a></li>
     
        <li class="next disabled"><a>Next &rarr;</a>
     
      </ul>
    </div>
    <hr>
   


  <div id="disqus_thread"></div>
<script type="text/javascript">
   
    var disqus_shortname = 'cyeam'; // required: replace example with your forum shortname
   
    /* * * DON'T EDIT BELOW THIS LINE * * */
    (function() {
        var dsq = document.createElement('script'); dsq.type = 'text/javascript'; dsq.async = true;
        dsq.src = 'http://' + disqus_shortname + '.disqus.com/embed.js';
        (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(dsq);
    })();
</script>
<noscript>Please enable JavaScript to view the <a href="http://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
<a href="http://disqus.com" class="dsq-brlink">blog comments powered by <span class="logo-disqus">Disqus</span></a>




  </div>
</div>


      </div>

<!--
      <script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
      <ins class="adsbygoogle"
           style="display:inline-block;width:728px;height:90px"
           data-ad-client="ca-pub-1651120361108148"
           data-ad-slot="4548335295"></ins>
      <script>
      (adsbygoogle = window.adsbygoogle || []).push({});
      </script>
--> 

<script type="text/javascript">
    /*234*60 创建于 2015-02-11*/
    var cpro_id = "u1949560";
</script>
<script src="http://cpro.baidustatic.com/cpro/ui/c.js" type="text/javascript"></script>

      <hr>
      <footer>
          <script type="text/javascript">
            (function(w,d,t,u,n,s,e){w['SwiftypeObject']=n;w[n]=w[n]||function(){
            (w[n].q=w[n].q||[]).push(arguments);};s=d.createElement(t);
            e=d.getElementsByTagName(t)[0];s.async=1;s.src=u;e.parentNode.insertBefore(s,e);
            })(window,document,'script','//s.swiftypecdn.com/install/v1/st.js','_st');

            _st('install','wU4JEZ3TeKgM-qap8TxG');
          </script>
        <p>&copy; 2015 Bryce
          with help from <a href="http://jekyllbootstrap.com" target="_blank" title="The Definitive Jekyll Blogging Framework">Jekyll Bootstrap</a>
          and <a href="http://twitter.github.com/bootstrap/" target="_blank">Twitter Bootstrap</a>
        </p>
          <a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="Creative Commons License" style="border-width:0" src="http://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br /><span xmlns:dct="http://purl.org/dc/terms/" href="http://purl.org/dc/dcmitype/Dataset" property="dct:title" rel="dct:type">Cyeam</span> by <a xmlns:cc="http://creativecommons.org/ns#" href="www.cyeam.com" property="cc:attributionName" rel="cc:attributionURL">Bryce</a> is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.

       


  <script type="text/javascript">
  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-50349737-1']);
  _gaq.push(['_trackPageview']);

  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();
</script>



        <script type="text/javascript">var cnzz_protocol = (("https:" == document.location.protocol) ? " https://" : " http://");document.write(unescape("%3Cspan id='cnzz_stat_icon_1000474528'%3E%3C/span%3E%3Cscript src='" + cnzz_protocol + "s19.cnzz.com/z_stat.php%3Fid%3D1000474528%26show%3Dpic1' type='text/javascript'%3E%3C/script%3E"));</script>
      </footer>

    </div>

  </body>
</html>
`
