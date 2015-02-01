package enchented

const page1HTML = headerHTML +
`
<body id="welcome">
  <header>
    <h1>Caroline <em>&amp;</em> Tony</h1>
    <time datetime="2011-08-27">August 27, 2011</time>
    <em>Austin, Texas</em>
  </header>

  <nav>
    <h1>Navigation</h1>
    <ul>
      <li class="current"><strong><a href="/pages/1">Welcome</a></strong></li>
      <li><a href="/pages/3">About Us</a></li>
      <li><a href="/pages/2">The Wedding</a></li>
      <li><a href="/pages/5">Wedding Party</a></li>
      <li><a href="/pages/7">Event Details</a></li>
      <li><a href="/pages/6">Guest &amp; Travel Information</a></li>
      <li><a href="/pages/9">Things To Do</a></li>
      <li><a href="/pages/10">Registry</a></li>
      <li><a href="/rsvps/new?wedding_id=1">Rsvp</a></li>
    </ul>
  </nav>

  <div id="page">
    <div>


   <!-- TEXT WIDGET -->
<div class="text freeform" style="text-align: center">
<h1>Howdy!</h1>
<p>Welcome to our wedding website! You can find all the details about our upcoming wedding here so please come back often for new information. We are so excited for you to be a part of our big day and cant wait to see yall there!</p><p><strong>Love,<br><span class=bride>Caroline</span> &amp; Tony</strong></p>
</div>
<!-- / TEXT WIDGET -->


    </div>
  </div>
` + footerHTML
