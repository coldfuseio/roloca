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
    <html class="no-js" lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="x-ua-compatible" content="ie=edge">
        <title>Roloca: an open API for romanian counties and cities</title>
        <meta name="description" content="An open API for romanian addresses which lists all counties in Romania, all cities in Romania and all cities in a romanian county, developed at coldfuse.io">
        <meta name="keywords" content="open api, romania, romanian, cities, counties">
        <meta name="author" content="coldfuse.io">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="alternate" hreflang="ro" href="http://roloca.coldfuse.io/ro">

        <!-- Place favicon.ico in the root directory -->
        <link rel="shortcut icon" href="http://coldfuse.io/favicon.ico" />

        <link rel="stylesheet" href="http://s3.eu-central-1.amazonaws.com/coldfuse.io/css/normalize.css">
        <link rel="stylesheet" href="http://s3.eu-central-1.amazonaws.com/coldfuse.io/css/roloca/main.css">

    </head>
    <body>
    <!--[if lt IE 8]>
    <p class="browserupgrade">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
    <![endif]-->

            <header class="header">
                <img src="http://coldfuse.io/logos/regular/mono-inverted/coldfuseLogo.svg" alt="coldfuse logo" class="logo center">
                <div class="logoType">
                    roloca
                </div>
            </header>

            <div class="container--text center">


                <h1 class="u-tac h-p u-mb40">An open API for Romanian counties and cities</h1>

                <div class="u-mb20">
                    <h2 class="h-p">All counties in Romania</h2>
                    <a href="/judete" target="_blank" class="code-link">roloca.coldfuse.io/judete</a>
                </div>

                <div class="u-mb20">
                    <h2 class="h-p">All cities in Romania</h2>
                    <a href="/orase" target="_blank" class="code-link">roloca.coldfuse.io/orase</a>
                </div>

                <div class="u-mb40">
                    <h2 class="h-p">All cities in a county</h2>
                    <a href="/orase/GL" target="_blank" class="code-link">roloca.coldfuse.io/orase/GL</a>
                </div>

                <p class="u-tac u-mb40">Drop us a line at <a href="mailto:hello@coldfuse.io">hello@coldfuse.io</a> for any issues / suggestions</p>


                <p class="u-tac">
                    developed at <a href='http://coldfuse.io'>coldfuse.io</a><br>using data from <a href='https://github.com/romania/localitati'>github.com/romania/localitati</a>
                </p>


            </div>

            <footer class="footer">
                <p class="u-fs12 u-tac">© 2016 coldfuse<br></p>
            </footer>



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

func HomeRoHandler(w http.ResponseWriter, r *http.Request) {
	h := `
	<!doctype html>
        <html class="no-js" lang="ro">
            <head>
                <meta charset="utf-8">
                <meta http-equiv="x-ua-compatible" content="ie=edge">
                <title>Roloca: un API public pentru localităţile şi judeţele din România</title>
                <meta name="description" content="API public ce conţine lista tuturor judeţelor din România şi lista tuturor oraşelor din România, organizate pe judeţe; dezvoltat de coldfuse.io">
                <meta name="keywords" content="open api, românia, romanesti, orase, judete">
                <meta name="author" content="coldfuse.io">
                <meta name="viewport" content="width=device-width, initial-scale=1">
                <link rel="alternate" hreflang="en" href="http://roloca.coldfuse.io">

                <!-- Place favicon.ico in the root directory -->
                <link rel="shortcut icon" href="http://coldfuse.io/favicon.ico" />

                <link rel="stylesheet" href="http://s3.eu-central-1.amazonaws.com/coldfuse.io/css/normalize.css">
                <link rel="stylesheet" href="http://s3.eu-central-1.amazonaws.com/coldfuse.io/css/roloca/main.css">

            </head>
            <body>
                <!--[if lt IE 8]>
                    <p class="browserupgrade">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
                <![endif]-->

                <header class="header">
                    <img src="http://coldfuse.io/logos/regular/mono-inverted/coldfuseLogo.svg" alt="coldfuse logo" class="logo center">
                    <div class="logoType">
                        roloca
                    </div>
                </header>

                <div class="container--text center">

                    <h1 class="u-tac h-p u-mb40">API public pentru localităţile şi judeţele din România</h1>

                    <div class="u-mb20">
                        <h2 class="h-p">Toate judeţele din România</h2>
                        <a href="/judete" target="_blank" class="code-link">roloca.coldfuse.io/judete</a>
                    </div>

                    <div class="u-mb20">
                        <h2 class="h-p">Toate oraşele din România</h2>
                        <a href="/orase" target="_blank" class="code-link">roloca.coldfuse.io/orase</a>
                    </div>

                    <div class="u-mb40">
                        <h2 class="h-p">Toate oraşele dintr-un anumit judeţ</h2>
                        <a href="/orase/GL" target="_blank" class="code-link">roloca.coldfuse.io/orase/GL</a>
                    </div>

                    <p class="u-tac u-mb40">Trimite-ne un email la <a href="mailto:hello@coldfuse.io">hello@coldfuse.io</a> pentru probleme / sugestii</p>


                    <p class="u-tac">
                        dezvoltat de <a href='http://coldfuse.io'>coldfuse.io</a><br> folosind date de la <a href='https://github.com/romania/localitati'>github.com/romania/localitati</a>
                    </p>


                </div>

                <footer class="footer">
                    <p class="u-fs12 u-tac">© 2016 coldfuse<br></p>
                </footer>



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
