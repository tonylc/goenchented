package enchented

const page7HTML = headerHTML + `
<body id="event_details">
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
      <li><a href="/pages/2">The Wedding</a></li>
      <li><a href="/pages/5">Wedding Party</a></li>
      <li class="current"><strong><a href="/pages/7">Event Details</a></strong></li>
      <li><a href="/pages/6">Guest &amp; Travel Information</a></li>
      <li><a href="/pages/9">Things To Do</a></li>
      <li><a href="/pages/10">Registry</a></li>
      <li><a href="/rsvps/new?wedding_id=1">Rsvp</a></li>
    </ul>
  </nav>

  <div id="page">
    <div>


    <!-- WIDGET GROUP -->
<section class="events tabbed" style="text-align: center">

 <h1>Wedding Events</h1>

 <ul>
   <li><a href="#wedding_and_reception">wedding &amp; reception</a></li>
   <li><a href="#after_party">after party</a></li>
 </ul>


<!-- EVENTS WIDGET -->
<div id="wedding_and_reception" class="events">
  <div class="event vevent">
    <h3 class="summary">wedding &amp; reception</h3>
    <time class="dtstart" datetime=2011-08-27>Saturday, August 27, 2011 at <span class="value"> 6:30pm</span></time>
    <h4 class="location">Thurmans Mansion at the Salt Lick BBQ</h4>
    <p class="description">The outdoor ceremony will be followed by cocktail hour and a bbq dinner reception.</p>
  </div>

  <ol class="places">
    <li>
      <div class="adr">
        <h3>Thurmans Mansion at the Salt Lick BBQ</h3>
        <span class="street-address">18300 Farm to Market Road 1826</span><br />
        <span class="locality">Driftwood</span>,&nbsp;
        <span class="region">TX</span><br />
        <a class="directions" href="http://maps.google.com/maps?f=d&amp;source=s_d&amp;saddr=700+San+Jacinto+at+8th+Street,+Austin,+TX&amp;daddr=18300+Farm+to+Market+Road+1826,+Driftwood,+TX&amp;hl=en&amp;geocode=FXjezQEd_5ks-ilFVkg_p7VEhjGd-ppcWLIwIg%3BFYTEywEdkm4o-im7wzFj5URbhjFy0mcmUBc5uA&amp;mra=ls&amp;sll=30.2656" target=_blank>Get Directions</a>
      </div>
      <div class="geo">
        <span class="latitude">30.131695</span>
        <span class="longitude">-98.017504</span>
      </div>
    </li>
  </ol>

  <menu>
    <li>
      <a href="/rsvps/new?wedding_id=1">RSVP for wedding &amp; reception</a>

    </li>
  </menu>

</div>
<!-- / EVENTS WIDGET -->


<!-- EVENTS WIDGET -->
<div id="after_party" class="events">
  <div class="event vevent">
    <h3 class="summary">after party</h3>
    <time class="dtstart" datetime=2011-08-27>Saturday, August 27, 2011 at <span class="value">11:30pm</span></time>
    <h4 class="location">Colorado St &amp; W 4th St</h4>
    <p class="description">Immediately following the reception, we will be heading back to downtown Austin to continue the celebration on 4th Street.  Hope to see yall there!</p>
  </div>

  <ol class="places">
    <li>
      <div class="adr">
        <h3>Colorado St & W 4th St</h3>
        <span class="street-address">Colorado St &amp; W 4th St</span><br />
        <span class="locality">Austin</span>,&nbsp;
        <span class="region">TX</span><br />
        <a class="directions" href="http://maps.google.com/maps?q=4th+and+colorado,+austin,+TX&amp;hl=en&amp;sll=37.0625,-95.677068&amp;sspn=45.197878,93.076172&amp;z=16&amp;iwloc=A" target=_blank>Get Directions</a>
      </div>
      <div class="geo">
        <span class="latitude">30.266569</span>
        <span class="longitude">-97.744811</span>
      </div>
    </li>
  </ol>


</div>
<!-- / EVENTS WIDGET -->


</section>
<!-- / WIDGET GROUP -->



    </div>
  </div>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.6/jquery.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/jqueryui/1.8.13/jquery-ui.min.js"></script>
<script src="/javascripts/jquery/plugins/jquery-pjax.js?1325274968" type="text/javascript"></script>
<script src="http://maps.google.com/maps/api/js?sensor=false" type="text/javascript"></script>
<script src="/javascripts/maps.js?1325274968" type="text/javascript"></script>
<script src="/javascripts/groups.js?1325274968" type="text/javascript"></script>
<script src="/javascripts/rsvp.js?1325274968" type="text/javascript"></script>

    <script src="/javascripts/navigation.js?1325274968" type="text/javascript"></script>

</body>
</html>
` + footerHTML
