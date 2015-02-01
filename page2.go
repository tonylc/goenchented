package enchented

const page2HTML = headerHTML + `
<body id="the_wedding">
  <header>
    <h1>Caroline <em>&amp;</em> Tony</h1>
    <time datetime="2011-08-27">August 27, 2011</time>
    <em>Austin, Texas</em>
  </header>

  <nav>
    <h1>Navigation</h1>
    <ul>
      <li><a href="/pages/1">Welcome</a></li>
      <li><a href="/pages/3">About Us</a></li>
      <li class="current"><strong><a href="/pages/2">The Wedding</a></strong></li>
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

<p>together with their family and friends</p><p><strong style=\"color: #da8b91; font-size: 24px;\">Caroline Ryu</strong><br />and<br /><strong style=\"font-size: 24px;\">Tony Chen</strong></p><p>invite you to join the celebration<br />of their love and new life together</p><p><strong>saturday, august 27, 2011</strong><br />at six-thirty in the evening</p><p><strong>thurmans mansion</strong> at the <strong>salt lick bbq</strong><br /><strong>driftwood, texas</strong></p><p>dinner and dancing to follow<br /><strong style=\"color: #da8b91\">please rsvp on this site by june 27</strong></p>
</div>
<!-- / TEXT WIDGET -->


    </div>
  </div>
` + footerHTML
