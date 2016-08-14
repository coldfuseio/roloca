package main // import "github.com/nalexeric/roloca"

import (
	"log"
	"net/http"
	"fmt"
	"github.com/yosssi/gohtml"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	h := `
	<!doctype html>
    <html class="no-js" lang="">
        <head>
            <meta charset="utf-8">
            <meta http-equiv="x-ua-compatible" content="ie=edge">
            <title>Romanian counties and cities - open API</title>
            <meta name="description" content="An open API for Romanian addresses which lists all counties in Romania, all cities in a Romania and all cities in a romanian county, developed at coldfuse.io">
            <meta name="keywords" content="open api, romania, romanian, cities, counties">
            <meta name="author" content="coldfuse.io">
            <meta name="viewport" content="width=device-width, initial-scale=1">
    
            <!-- Place favicon.ico in the root directory -->
            <link rel="shortcut icon" href="http://coldfuse.io/favicon.ico" />
    
            <style>
                /*! normalize.css v3.0.3 | MIT License | github.com/necolas/normalize.css */a,html{color:#888}fieldset,hr,img,legend{border:0}fieldset,hr,legend,td,th{padding:0}pre,textarea{overflow:auto}.logoType,.u-tac{text-align:center}body{margin:0;background:#fafafa}article,aside,details,figcaption,figure,footer,header,hgroup,main,menu,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block;vertical-align:baseline}audio:not([controls]){display:none;height:0}[hidden],template{display:none}.logo,hr{display:block}a{background-color:transparent;text-decoration:underline;transition:.2s ease-in-out}a:active,a:hover{outline:0}abbr[title]{border-bottom:1px dotted}b,optgroup,strong{font-weight:700}dfn{font-style:italic}h1{font-size:2em;margin:.67em 0}mark{background:#ff0;color:#000}small{font-size:80%}sub,sup{font-size:75%;line-height:0;position:relative;vertical-align:baseline}.u-fpc>.u-fpc-content,audio,canvas,iframe,img,svg,video{vertical-align:middle}sup{top:-.5em}sub{bottom:-.25em}svg:not(:root){overflow:hidden}figure{margin:1em 40px}hr{box-sizing:content-box;height:1px;border-top:1px solid #ccc;margin:1em 0}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}button,input,optgroup,select,textarea{color:inherit;font:inherit;margin:0}html,pre{font-family:Nitti,monospace}button{overflow:visible}button,select{text-transform:none}button,html input[type=button],input[type=reset],input[type=submit]{-webkit-appearance:button;cursor:pointer}button[disabled],html input[disabled]{cursor:default}button::-moz-focus-inner,input::-moz-focus-inner{border:0;padding:0}input{line-height:normal}input[type=checkbox],input[type=radio]{box-sizing:border-box;padding:0}input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button{height:auto}input[type=search]{-webkit-appearance:textfield;box-sizing:content-box}input[type=search]::-webkit-search-cancel-button,input[type=search]::-webkit-search-decoration{-webkit-appearance:none}textarea{resize:vertical}table{border-collapse:collapse;border-spacing:0}/*! HTML5 Boilerplate v5.3.0 | MIT License | https://html5boilerplate.com/ */@font-face{font-family:Nitti;src:url(http://coldfuse.io/css/nitti/NittiWM2-Medium.eot?#iefix) format('embedded-opentype'),url(http://coldfuse.io/css/nitti/NittiWM2-Medium.otf) format('opentype'),url(http://coldfuse.io/css/nitti/NittiWM2-Medium.woff) format('woff'),url(http://coldfuse.io/css/nitti/NittiWM2-Medium.ttf) format('truetype'),url(http://coldfuse.io/css/nitti/NittiWM2-Medium.svg#NittiWM2-Medium) format('svg');font-weight:500;font-style:normal}@font-face{font-family:Nitti;src:url(http://coldfuse.io/css/nitti/NittiWM2-Light.eot?#iefix) format('embedded-opentype'),url(http://coldfuse.io/css/nitti/NittiWM2-Light.otf) format('opentype'),url(http://coldfuse.io/css/nitti/NittiWM2-Light.woff) format('woff'),url(http://coldfuse.io/css/nitti/NittiWM2-Light.ttf) format('truetype'),url(http://coldfuse.io/css/nitti/NittiWM2-Light.svg#NittiWM2-Light) format('svg');font-weight:400;font-style:normal}*,:after,:before{box-sizing:border-box}html{-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;font-size:17px;line-height:1.4}::-moz-selection{background:#b3d4fc;text-shadow:none}::selection{background:#b3d4fc;text-shadow:none}fieldset{margin:0}a:hover{color:#5CC4F2}.browserupgrade{margin:.2em 0;background:#ccc;color:#000;padding:.2em 0}.logo{width:54px;margin-top:100px;margin-bottom:20px}.logoType{font-size:28px;letter-spacing:1px;margin-bottom:80px}pre{background-color:#eee;border:1px solid #ddd;padding:.5rem 1rem;margin-top:0}.u-lnu{text-decoration:none}.code-linkPath{font-weight:500}.footer{padding:50px 0}.u-fpc{display:table;width:100%;min-height:100vh}.u-fpc>.u-fpc-content{display:table-cell}.hidden{display:none!important}.visuallyhidden{border:0;clip:rect(0 0 0 0);height:1px;margin:-1px;overflow:hidden;padding:0;position:absolute;width:1px}.visuallyhidden.focusable:active,.visuallyhidden.focusable:focus{clip:auto;height:auto;margin:0;overflow:visible;position:static;width:auto}.invisible{visibility:hidden}.clearfix:after,.clearfix:before{content:" ";display:table}.clearfix:after{clear:both}.center{margin-left:auto;margin-right:auto}.container--text{width:100%;max-width:600px;padding-left:20px;padding-right:20px}.u-taj{text-align:justify}.u-fs12{font-size:12px}.h-p{font-size:1rem;font-weight:400}.u-mb20{margin-bottom:2rem}.u-mb40{margin-bottom:4rem}
            </style>
        </head>
        <body>
            <!--[if lt IE 8]>
                <p class="browserupgrade">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
            <![endif]-->
            <div class="u-fpc">
                <div class="u-fpc-content">
    
                    <img src="http://coldfuse.io/coldfuse-logo.png" alt="coldfuse logo" class="logo center">
                    <div class="logoType">
                        roloca
                    </div>
    
                    <!-- Add your site or application content here -->
                    <div class="container--text center">
    
    
                        <h1 class="u-tac h-p u-mb40">An open API for Romanian counties and cities</h1>
    
                        <div class="u-mb20">
                            <h2 class="h-p">All counties in Romania</h2>
                            <pre><code><a href="/judete" target="_blank" class="u-lnu">roloca.coldfuse.io/judete</a></code></pre>
                        </div>
    
                        <div class="u-mb20">
                            <h2 class="h-p">All cities in Romania</h2>
                            <pre><code><a href="/orase" target="_blank" class="u-lnu">roloca.coldfuse.io/orase</a></code></pre>
                        </div>
    
                        <div class="u-mb40">
                            <h2 class="h-p">All cities in a county</h2>
                            <pre><code><a href="/orase/GL" target="_blank" class="u-lnu">roloca.coldfuse.io/orase/GL</a></code></pre>
                        </div>
    
                        <p class="u-tac u-mb40">Drop us a line at <a href="mailto:hello@coldfuse.io">hello@coldfuse.io</a> for any suggestions/issues</p>
    
    
                        <p class="u-tac">
                            developed at <a href='http://coldfuse.io'>coldfuse.io</a><br>using data from <a href='https://github.com/romania/localitati'>github.com/romania/localitati</a>
                        </p>
    
    
                    </div>
    
                    <footer class="footer">
                        <p class="u-fs12 u-tac">© 2016 coldfuse<br></p>
                    </footer>
    
                </div>
            </div>
    
    
            <script>
                (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
                            (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                        m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
                })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');
    
                ga('create', 'UA-67119922-8', 'auto');
                ga('send', 'pageview');
    
            </script>
        </body>
    </html>

	`
	fmt.Fprintf(w, gohtml.Format(h))
}
