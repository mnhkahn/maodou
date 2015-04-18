package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestContent(t *testing.T) {
	reader := strings.NewReader(raw)
	fmt.Println(Content(reader))
}

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
