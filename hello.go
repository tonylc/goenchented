package enchented

import (
    "fmt"
    "html/template"
    "net/http"
)

func init() {
    http.HandleFunc("/", page1)
    http.HandleFunc("/pages/1", page1)
    http.HandleFunc("/pages/2", page2)
    http.HandleFunc("/pages/3", page3)
    http.HandleFunc("/pages/5", page5)
    http.HandleFunc("/pages/6", page6)
    http.HandleFunc("/pages/7", page7)
    http.HandleFunc("/pages/9", page9)
    http.HandleFunc("/pages/10", page10)
    http.HandleFunc("/rsvps/new", rsvp)
    http.HandleFunc("/javascripts/jquery/plugins/jquery-pjax.js", jquerypjax)
    http.HandleFunc("/javascripts/rsvp.js", rsvpjs)
    http.HandleFunc("/javascripts/navigation.js", navigationjs)
    http.HandleFunc("/javascripts/maps.js", mapsjs)
    http.HandleFunc("/javascripts/groups.js", groupsjs)
}

func page1(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, page1HTML)
}

func page2(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, page2HTML)
}

func page3(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, page3HTML)
}

func page5(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, page5HTML)
}

func page6(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, page6HTML)
}

func page7(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, page7HTML)
}

func page9(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, page9HTML)
}

func page10(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, page10HTML)
}

func rsvp(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, rsvpHTML)
}

func jquerypjax(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, jquerypjaxHTML)
}

func rsvpjs(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, rsvpjsHTML)
}

func navigationjs(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, navigationjsHTML)
}

func groupsjs(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, groupsjsHTML)
}

func mapsjs(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, mapsjsHTML)
}

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

const signTemplateHTML = `
<html>
  <body>
    <p>You wrote:</p>
    <pre>{{.}}</pre>
  </body>
</html>
`

const headerHTML = `
<!DOCTYPE html>
<html>
<head>
  <title>Caroline &amp; Tony &infin; August 27, 2011 &infin; Austin, Texas</title>

    <!--
    /*
     * MyFonts Webfont Build ID 804514, 2011-05-11T23:39:13-0400
     *
     * The fonts listed in this notice are subject to the End User License
     * Agreement(s) entered into by the website owner. All other parties are
     * explicitly restricted from using the Licensed Webfonts(s).
     *
     * You may obtain a valid license at the URLs below.
     *
     * Webfont: Bodoni Egyptian Pro Bold
     * URL: http://new.myfonts.com/fonts/shinn/bodoni-egyptian-pro/bold/
     * Foundry: ShinnType
     * Copyright: Copyright (c) 2010 by Nick Shinn. Published by Shinn Type Foundry. All rights reserved.
     * License: http://www.myfonts.com/viewlicense?1056
     * Licensed pageviews: 10,000/month
     * CSS font-family: BodoniEgyptianPro-Bold
     * CSS font-weight: normal
     *
     * Webfont: Bodoni Egyptian Pro Regular
     * URL: http://new.myfonts.com/fonts/shinn/bodoni-egyptian-pro/regular/
     * Foundry: ShinnType
     * Copyright: Copyright (c) 2010 by Nick Shinn. Published by Shinn Type Foundry. All rights reserved.
     * License: http://www.myfonts.com/viewlicense?1056
     * Licensed pageviews: 10,000/month
     * CSS font-family: BodoniEgyptianPro-Regular
     * CSS font-weight: normal
     *
     * (c) 2011 Bitstream Inc
    */
    -->

    <!--[if lte IE 8]>
  <script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
  <![endif]-->
  <style>` + mastercssHTML + `</style>

<!--
  <link href="http://enchented2.herokuapp.com/stylesheets/master.css?1325274968" media="screen" rel="stylesheet" type="text/css" />
<link href="http://enchented2.herokuapp.com/stylesheets/layout.css?1325274968" media="screen" rel="stylesheet" type="text/css" />
<link href="http://enchented2.herokuapp.com/stylesheets/global.css?1325274968" media="screen" rel="stylesheet" type="text/css" />
<link href="http://enchented2.herokuapp.com/stylesheets/groups.css?1325274968" media="screen" rel="stylesheet" type="text/css" />
<link href="http://enchented2.herokuapp.com/stylesheets/patterns.css?1325274968" media="screen" rel="stylesheet" type="text/css" />
<link href="http://enchented2.herokuapp.com/stylesheets/widgets.css?1325274968" media="screen" rel="stylesheet" type="text/css" /> -->



  <style>
    .bride { color: #da8b91; }
  </style>

  <script type="text/javascript">var _gaq = _gaq || [];_gaq.push(["_setAccount", "UA-2311903-14"]);_gaq.push(["_trackPageview"]);(function() {var ga = document.createElement("script"); ga.type = "text/javascript"; ga.async = true;ga.src = ("https:" == document.location.protocol ? "https://ssl" : "http://www") + ".google-analytics.com/ga.js";var s = document.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ga, s);})();</script>
</head>
`

const footerHTML = `
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.6/jquery.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/jqueryui/1.8.13/jquery-ui.min.js"></script>
<script src="/javascripts/jquery/plugins/jquery-pjax.js?1325274968" type="text/javascript"></script>
<script src="/javascripts/navigation.js?1325274968" type="text/javascript"></script>
</body>
</html>
`
