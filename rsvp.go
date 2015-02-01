package enchented

const rsvpHTML = headerHTML + `
<body id="rsvp">
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
      <li><a href="/pages/7">Event Details</a></li>
      <li><a href="/pages/6">Guest &amp; Travel Information</a></li>
      <li><a href="/pages/9">Things To Do</a></li>
      <li><a href="/pages/10">Registry</a></li>
      <li class="current"><strong><a href="/rsvps/new?wedding_id=1">Rsvp</a></strong></li>
    </ul>
  </nav>

  <div id="page">
    <div>

<!-- RSVP WIDGET -->
                <div class="rsvp">

                    <h1>RSVP</h1>

                    <div class="event vevent">

                        <h3 class="summary">Caroline &amp; Tony's Wedding </h3>
                        <time class="dtstart" datetime="2011-08-27"><a href="#" style="cursor: default;">August 27, 2011</a></time>

                    </div>

<form accept-charset="UTF-8" action="/rsvps" class="new_rsvp" id="new_rsvp" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="+LgU8l+aQtzuJuIsAVx8fOpyDWscXMjbUD06/NAvL1I=" /></div>                       <div class="form">


      <input id="rsvp_wedding_id" name="rsvp[wedding_id]" type="hidden" value="1" />

                            <p>Kindly reply for each guest by 6/27/11. Thanks!</p>

                            <fieldset class="rsvp">
                                <legend>Will you be able to attend?</legend>

        <input id="rsvp_can_attend_true" name="rsvp[can_attend]" type="radio" value="true" required="required" /><label for="#">Yes</label>
        <input checked="checked" id="rsvp_can_attend_false" name="rsvp[can_attend]" type="radio" value="false" required="required" /><label for="#">No</label>
                            </fieldset>

                            <label for="rsvp_name">Your Name</label>
          <input id="rsvp_name" name="rsvp[name]" type="text" required="required" />

                            <section id="details">

                                <h1>Details</h1>


        <fieldset>

  <legend>Food Preference:</legend>



</fieldset>

        <label class="optional" for="rsvp_questions_2">
  Have you booked accommodations for your stay?<br>If so, please let us know where:</label>
  <input id=rsvp_questions_2 type="text" name="rsvp_questions[2]"/>

<small>Optional</small>

        <fieldset>

  <legend>Transportation is provided to/from the wedding at the Omni Austin Downtown Hotel. Please indicate if youd like a ride:</legend>



</fieldset>

        <p>For more information, visit the Guest & Travel tab.</p>


                             <label class="optional" for="rsvp_additional_notes">Additional Notes <small>Optional</small></label>
            <textarea cols="40" id="rsvp_additional_notes" name="rsvp[additional_notes]" placeholder="Please let us know if you have any special requests" rows="4"></textarea>

                            </section>

                        </div>

                    <menu class="actions">
                        <li><input type="submit" value="Submit" /></li>
                    </menu>
</form>                </div>
                <!-- / RSVP WIDGET -->



    </div>
  </div>


<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.6/jquery.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/jqueryui/1.8.13/jquery-ui.min.js"></script>
<script src="/javascripts/jquery/plugins/jquery-pjax.js?1325274968" type="text/javascript"></script>
<script src="/javascripts/rsvp.js?1325274968" type="text/javascript"></script>
<script>

  var required_fields;

  $(document).ready(function() {
  required_fields = new Array()
  $('#rsvp_can_attend_true').change(function(){toggleCanAttend();});
  $('#rsvp_can_attend_false').change(function(){toggleCanAttend();});
  toggleCanAttend();
  });

  function toggleCanAttend(){
 if($('#rsvp_can_attend_true').is(':checked')){

 for (var i = 0; i < required_fields.length; i++){
 $('#' +required_fields[i]).attr('required', true);
  }

  required_files = new Array();

  }else{

 $('#details :input').each(function (index, field) {
  if($(this).attr('required') === 'required'){
  required_fields.push($(this).attr('id'));
  $(this).removeAttr('required');
  }
  });

  }
}

</script>
    <script src="/javascripts/navigation.js?1325274968" type="text/javascript"></script>
</body>
</html>
`
