package enchented

const mastercssHTML =
`
/* LOVE FOR DESIGN
/* Master Stylesheet
----------------------------------------------------------------------------- */


    @import url('layout.css');
    @import url('global.css');
    @import url('groups.css');
    @import url('patterns.css');
    @import url('widgets.css');

    @font-face {

        font-family: 'Bodoni Egyptian';
        src: url('http://enchented2.herokuapp.com/fonts/eot/style_196043.eot');
        src: url('http://enchented2.herokuapp.com/fonts/eot/style_196043.eot?#iefix') format('embedded-opentype'),url('http://enchented2.herokuapp.com/fonts/woff/style_196043.woff') format('woff'),url('http://enchented2.herokuapp.com/fonts/ttf/style_196043.ttf') format('truetype'),url('http://enchented2.herokuapp.com/fonts/svg/style_196043.svg#BodoniEgyptianPro-Regular') format('svg');

    }

    @font-face {

        font-family: 'Bodoni Egyptian';
        font-weight: bold;
        src: url('http://enchented2.herokuapp.com/fonts/eot/style_196039.eot');
        src: url('http://enchented2.herokuapp.com/fonts/eot/style_196039.eot?#iefix') format('embedded-opentype'),url('http://enchented2.herokuapp.com/fonts/woff/style_196039.woff') format('woff'),url('http://enchented2.herokuapp.com/fonts/ttf/style_196039.ttf') format('truetype'),url('http://enchented2.herokuapp.com/fonts/svg/style_196039.svg#BodoniEgyptianPro-Bold') format('svg');

    }


/* Reset
----------------------------------------------------------------------------- */


    body {

        font-family: "Bodoni Egyptian", serif;

    }

    nav h1 {

        display: none;

    }

    form, h3, h4 {

        margin: 0;
        padding: 0;

    }

    p {

        margin: 0 0 20px 0;

    }

    #about_us h3 { font-size: 12pt; }


/* LOVE FOR DESIGN
/* Layout
----------------------------------------------------------------------------- */


    body {

        min-width: 980px;
        margin: 0 0 20px 0;
        padding: 0;
        background: #c7af82 url(https://lh4.googleusercontent.com/-iS0WdB83870/VNgShlMRRII/AAAAAAABQWE/aDfF9ugbPWw/w800-h544-no/body.jpg);

    }


/* Page */


    div#page {

        width: 500px;
        margin: 0 auto;
                padding: 5px;
        background: url(https://lh3.googleusercontent.com/-BGaTHyVlvBI/VNgSimUtYhI/AAAAAAABQVs/2O-a7GnKebQ/w610-h10-no/border-top.png), url(https://lh6.googleusercontent.com/-FSQSqpQsYQY/VNgShqWEgFI/AAAAAAABQV4/iicuWdXE9DA/w610-h10-no/border-bottom.png), url(https://lh4.googleusercontent.com/-FbMVLIcXMh0/VNgSh8MUE-I/AAAAAAABQV0/Ypc8yWqmIvY/w10-h884-no/border-left.png), url(https://lh4.googleusercontent.com/-vp_X4Jpx1gk/VNgSiVGC78I/AAAAAAABQVw/o2tsG-0KEvI/w10-h884-no/border-right.png);
        background-repeat: repeat-x, repeat-x, repeat-y, repeat-y;
        background-position: top left, bottom left, top left, top right;
                display: block;
    }

    div#page:after, div#page > div > *:after {

        content: ".";
        display: block;
        height: 0;
        clear: both;
        visibility: hidden;

    }

    div#page > div {

        width: 500px;
                float: left;
        background: url(https://lh3.googleusercontent.com/-nBfVWhL07Io/VNgShYt_umI/AAAAAAABQV8/1yvOpG96G9o/w744-h866-no/background-page.png) repeat top left;

    }


/* LOVE FOR DESIGN
/* Global Elements
----------------------------------------------------------------------------- */


    header {

        display: block;
        width: 500px;
        margin: 50px auto 30px auto;
        text-align: center;

    }

    header > * {

        margin: 0 auto;

    }

    header h1, header h1 a {

        margin-bottom: 5px;
        color: #625a52;
        font-weight: normal;
        font-size: 40px;
        line-height: 40px;
        text-decoration: none;

    }

    header h1 em {

        color: #fdd5d5;
        font-style: normal;

    }

    header time, header time + em {

        color: #fff;
        font-style: normal;
        font-size: 15px;
        line-height: 15px;
        text-transform: lowercase;

    }

    header time:after {

        content: " \2014"

    }


/* Navigation
----------------------------------------------------------------------------- */


    nav {

        display: block;
        position: fixed;
        top: 145px;
        right: 0;
                z-index: 1;

    }

    nav ul, nav li {

        margin: 0;
        padding: 0;
        list-style-type: none;

    }

    nav li {

        margin-bottom: 10px;
        padding-left: 5px;
        background: url(https://lh4.googleusercontent.com/-Q_z-cOrixC0/VNgSjgY4Q4I/AAAAAAABQVc/StjIRNsXw0M/w10-h180-no/navigation.png) no-repeat 0 0;

    }

    nav li.current {

        background-image: url(https://lh6.googleusercontent.com/-Yy87I0Y7XNQ/VNgSjA0sR9I/AAAAAAABQVg/0l4af6aVbLA/w10-h180-no/navigation-current.png);

    }

    nav li + li {

        background-position: 0 -30px;

    }

    nav li + li + li {

        background-position: 0 -60px;

    }


/* Text */


    nav a {

        display: block;
        padding: 2px 20px 0 7px;
        background-color: #625a52;
        color: #c7af82;
        font-weight: bold;
        text-decoration: none;
        text-transform: lowercase;
        font-size: 14px;
        line-height: 28px;

    }

    nav a:hover {

        color: #da8b91;

    }

    nav li strong a {

        background-color: #fff;
        color: #da8b91;

    }

/* LOVE FOR DESIGN
/* Widget Groups Stylesheet
----------------------------------------------------------------------------- */


    section.accordion > h1, section.tabbed > h1 {

        text-align: center;

    }


/* Accordion Format */


    section.accordion > h2 {

        display: block;
        width: 100%;
        clear: left;
        float: left;
        margin: 0 0 10px 0;
        border-radius: 5px;
        outline: none;
        background-color: #c7af82;
        opacity: 0.5;
        line-height: 40px;

    }

    section.accordion > h2.ui-state-active {

        margin-bottom: 20px;
        border-bottom-left-radius: 0;
        border-bottom-right-radius: 0;
        opacity: 1;
        background-color: #da8b91;

    }

    section.accordion > h2 a {

        display: block;
        color: #f7f7f7;
        text-align: center;
        text-decoration: none;
        text-shadow: #8a4a4e 1px 1px 1px;

    }

    section.accordion > h2.ui-state-active a {

        color: #fff;
        text-shadow: #625a52 1px 1px 1px;

    }

    section.accordion > h2 + * {

        display: block;
        width: 100%;
        clear: left;
        float: left;
        margin-bottom: 20px;
        background: none;

    }


/* Tabbed Format */


    .ui-tabs-hide {

        display: none;

    }

    section.tabbed > ul {

        display: block;
        width: 100%;
        float: left;
        margin: 0;
        border-bottom: 2px solid #625a52;
        padding: 0;

    }

    section.tabbed > ul li {

        display: block;
        float: left;
        margin: 0 10px 0 0;
        padding: 0;
        cursor: pointer;

    }

    section.tabbed > div {

        clear: left;
        float: left;
        width: 100%;
        text-align: left;

    }

    section.tabbed > ul a {

        display: block;
        outline: none;
        border-top-left-radius: 5px;
        border-top-right-radius: 5px;
        padding: 2px 7px 0 7px;
        background-color: #c8c8c4;
        color: #fff;
        font-size: 12px;
        font-weight: bold;
        line-height: 20px;
        text-decoration: none;
        cursor: pointer;

        -moz-user-select: none;
        -webkit-user-select: none;

    }

    section.tabbed > ul li.ui-tabs-selected a {

       background-color: #625a52;
       color: #c7af82;

    }

    section.tabbed > div h4, section.tabbed > div h4 a {

        color: #625a52;
        margin: 0;
        padding: 0;
        font-size: 20px;
        line-height: 30px;
        text-decoration: none;

    }

/* LOVE FOR DESIGN
/* Patterns Stylesheet
----------------------------------------------------------------------------- */


    .events p, .form p, .freeform p {

        width: 90%;
        clear: left;
        margin: 0 auto 20px auto;
        padding: 0;
        font-size: 14px;
        line-height: 20px;
        letter-spacing: 0.1em;

    }


/* Calls to Action
----------------------------------------------------------------------------- */


    menu {

        display: block;
        float: left;
        clear: left;
        width: 90%;
        margin: 0 5%;
        padding: 0;

    }

    menu li {

        margin: 0;
        padding: 0;
        list-style-type: none;

    }

    menu li + li {

        margin-top: 20px;

    }

    menu li a, menu li input {

        display: block;
        width: 99%;
        margin: 0 0 1px 0;
        border: 0;
        border-radius: 10px;
                -moz-border-radius: 10px;
                -webkit-border-radius: 10px;
        padding-top: 5px;
        box-shadow: 0px 0px 5px #a4a4a4;
                -moz-box-shadow: 0 0 5px #a4a4a4;
                -webkit-box-shadow: 0 0 5px #a4a4a4;
        background-color: #da8b91;
        color: #fff;
        text-shadow: 1px 1px 1px #8a4a4e;
        font-weight: bold;
        font-family: 'Bodoni Egyptian';
        font-size: 20px;
        line-height: 30px;
        text-align: center;
        text-decoration: none;

    }

    menu li input:hover {

        cursor: pointer;

    }

    menu a:hover {

        color: #fff;

    }


/* Events
----------------------------------------------------------------------------- */


    .event {

        margin-top: 20px;

    }

    .event h3, .event time {

        margin: 0;
        padding: 0;
        font-size: 24px;
        line-height: 30px;

    }

    .event p {

        width: 100%;

    }

    .event time {

        display: block;
        margin: 5px 0 20px 0;
        padding: 0 0 0 22px;
        background: url(https://lh5.googleusercontent.com/-XaxHkQQccmE/VNgSi9HksVI/AAAAAAABQVo/r1ysOh2PtBg/w34-h32-no/date.png) no-repeat 0 0 ;
        color: #da8b91;
        font-weight: bold;
        font-size: 14px;
        line-height: 20px;

    }

    .event time a {

        color: #da8b91;
        text-decoration: none;

    }


/* Forms
----------------------------------------------------------------------------- */


    form, .form menu {

        display: block;
        float: left;
        width: 100%;

    }

    .form p {

        margin: 0 0 20px 0;

    }


/* Controls */


    input, select, textarea {

        float: left;
        margin: 0 5px 0 0;

    }

    input[type=email], input[type=text], input[type=url], select {

        clear: left;
        margin-bottom: 20px;

    }

    input[type=radio] {

        margin-top: 3px;

    }

    textarea {

        clear: left;
        resize: none;
        margin: 5px 0 20px 0;
        border: 0;
        outline: none;
        padding: 10px;
        border-radius: 5px;
                -moz-border-radius: 5px;
                -webkit-border-radius: 5px;
        box-shadow: #a4a4a4 0 0 5px;
                -moz-box-shadow: #a4a4a4 0 0 5px;
                -webkit-box-shadow: #a4a4a4 0 0 5px;

    }

    textarea:focus {

        box-shadow: #646464 0 0 5px;

    }


/* Groups */


    fieldset {

        clear: left;
        float: left;
        margin: 0 0 20px 0;
        border: 1px dashed #aaa;
        border-radius: 3px;
                -moz-border-radius: 3px;
                -webkit-boder-radius: 3px;
        padding: 10px;

    }


/* Labels */


    label, legend {

        margin: 0;
        padding: 0;
        font-weight: bold;
        font-style: normal;
        font-size: 14px;
        line-height: 20px;
        white-space: normal;
        cursor: pointer;

        -moz-user-select: none;
        -webkit-user-select: none;

    }

    label {

        clear: left;
        float: left;
        margin: 0 10px 5px 0;

    }

    label.optional {

        color: #aca8a3;

    }

    label.optional.focused {

        color: #625a52;

    }

    fieldset label {

        clear: none;
        margin: 0 10px 0 0;
        color: #625a52;
        font-weight: normal;

    }

    .form small {

        display: block;
        float: left;
        color: #aca8a3;
        font-size: 12px;
        line-height: 20px;

    }

    .form label small {

        display: inline;
        padding-left: 5px;
        font-weight: normal;
        float: none;

    }


/* Validation */


    label.error, legend.error {

        float: left;
        color: #da8b91;

    }

    .form aside.error, .form strong.error {

        margin: 0 0 20px 0;
        border-radius: 5px;
                -moz-border-radius: 5px;
                -webkit-border-radius: 5px;
        padding: 10px;
        color: #fdd5d5;
        background-color: #625a52;
        font-size: 12px;

    }

    .form aside.error h3 {

        margin: 0 0 5px 0;
        font-size: 20px;

    }

    .form aside.error p {

        margin: 0;
        color: #d8d8d8;
        font-size: 12px;

    }

    .form strong.error {

        float: left;
        margin: 0;
        padding: 0 5px;
        line-height: 20px;

    }


/* Freeform Content
----------------------------------------------------------------------------- */


    .freeform {

        color: #625a52;

    }

    .freeform a {

        color: #625a52;

    }

    .freeform a:hover {

        color: #da8b91;

    }

    .freeform h3 {

        margin: 20px 0;
        font-size: 24px;
        line-height: 24px;

    }

    .freeform h4 {

        margin: 0 0 20px 0;
        font-size: 18px;
        line-height: 20px;

    }

    .freeform hr {

        display: block;
        width: 90%;
        height: 7px;
        margin: 20px auto;
        border: 0;
        padding: 0;
        background: url(https://lh6.googleusercontent.com/-QDZvHTzr9SM/VNgSj3tG-QI/AAAAAAABQVk/8UzvB4lbzmE/w732-h14-no/thread.png) repeat-x bottom left;
        color: transparent;
    clear: both;

    }

    .freeform figure {

        float: left;
        margin: 0 20px 20px 20px;

    }

        #wedding_party .freeform figure { width: 175px; }

    .freeform img {

        margin-bottom: 20px;
        border: 5px solid #625a52;
        border-radius: 5px;
                -moz-border-radius: 5px;
                -webkit-border-radius: 5px;
        box-shadow: #625a52 0 0 5px;
                -moz-box-shadow: #625a52 0 0 5px;
                -webkit-box-shadow: #625a52 0 0 5px;

    }

    .freeform p {
        width: 100%;
    }

    .freeform p strong {

        letter-spacing: 0;

    }

    .freeform ul, .freeform ol {

        clear: left;
        margin: 20px 0;

    }


/* Places
----------------------------------------------------------------------------- */


    .places {

        display: block;
        width: 100%;
        margin: 0 0 20px 0;
        border-bottom-left-radius: 5px;
        border-bottom-right-radius: 5px;
        padding: 10px 0 0 0;
        background: #625a52;
        box-shadow: #625a52 0 0 5px;
                -moz-box-shadow: #625a52 0 0 5px;
                -webkit-box-shadow: #625a52 0 0 5px;
        color: #fff;
        font-size: 14px;
        line-height: 20px;

    }

    .places, .places li {

        display: block;
        float: left;
        list-style-type: none;

    }

    .places li {

        width: 96%;
        margin: 0;
        padding: 0 2% 10px 2%;
        color: #fff;

    }

    .places li .adr {

        width: 100%;
        float: left;
        border-bottom: 1px dashed #95856a;
        padding-bottom: 10px;

    }

    .places li:last-child .adr, .places li:only-child .adr {

        border-bottom: 0;
        padding-bottom: 0;

    }

    .places .adr h3 {

        color: #c7af82;
        font-size: 18px;
        line-height: 28px;
        font-weight: bold;

    }

    .places li .adr h3 em {

        color: #da8b91;

    }

    .places .adr a {

        color: #da8b91;
        font-size: 12px;
        font-weight: bold;
        text-decoration: none;

    }

    .places .geo {

        display: none;

    }

    .map {

        display: block;
        visibility: hidden;
        width: 100%;
        height: 300px;
        clear: left;
        float: left;
        background-color: #c7af82;
        box-shadow: #625a52 0 0 5px;
                -moz-box-shadow: #625a52 0 0 5px;
                -webkit-box-shadow: #625a52 0 0 5px;

    }

/* LOVE FOR DESIGN
/* Widgets Stylesheet
----------------------------------------------------------------------------- */


    div#page > div > * > h1, div#page > div > form > * > h1 {

        margin: 0 0 20px 0;
        color: #da8b91;
        font-weight: normal;
        font-size: 30px;
        line-height: 40px;

    }

    div#page > div > * {

        float: left;
        width: 460px;
        margin: 20px;
        padding: 25px 5px 25px 5px;
        background: url(https://lh6.googleusercontent.com/-QDZvHTzr9SM/VNgSj3tG-QI/AAAAAAABQVk/8UzvB4lbzmE/w732-h14-no/thread.png), url(https://lh6.googleusercontent.com/-QDZvHTzr9SM/VNgSj3tG-QI/AAAAAAABQVk/8UzvB4lbzmE/w732-h14-no/thread.png), url(http://enchented2.herokuapp.com/images/skin/widget.png);
        background-repeat: repeat-x, repeat-x, repeat;
        background-position: top center, bottom center, top left;
        color: #625a52;

    }

    div#page > div > section + section {

        margin-top: 0;
        padding-top: 0;
        background: url(https://lh6.googleusercontent.com/-QDZvHTzr9SM/VNgSj3tG-QI/AAAAAAABQVk/8UzvB4lbzmE/w732-h14-no/thread.png);
        background-repeat: repeat-x, repeat;
        background-position: bottom center, top left;

    }


/* Events Widget
----------------------------------------------------------------------------- */


    section.events .event h4.location {

        display: none;

    }


/* RSVP Widget
----------------------------------------------------------------------------- */


    .rsvp section, .rsvp section h1 {

        clear: left;
        display: none;

    }
`