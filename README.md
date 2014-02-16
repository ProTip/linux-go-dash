linux-go-dash
=============
Forked from linux-dash to provide a golang based back end.  The primary goals for this fork are:

<ul>
 <li>Produce a lightweight Go backend
    <ul>
      <li>No language platform requirements</li>
      <li>Compiled for speed</li>
      <li>Massage into a general purpose linux info REST service</li>
    </ul>
 </li>
 <li>Produce Go linux information package
   <ul>
     <li>Retrieve information from running linux system</li>
     <li>Avoids forking at all costs: collect info from proc, sys, and kernel sockets</li>
   </ul>
 </li>
</ul>


linux-dash
==========

A low-overhead monitoring web dashboard for a linux machine. Simply drop-in the app and go!

<a href="http://afaq.dreamhosters.com/linux-dash/"> View Demo </a>



<img width="700px" src="http://afaq.dreamhosters.com/linux-dash.PNG">

<h2>Installation</h2>
<ol>
  <li>Download the zip/repo/package</li>
  <li>Place it in /var/www/ (for Apache)</li>
  <li>Secure access to the page via .htaccess or method of your choice</li>
</ol>  

<b>Please note: If you would like to limit access to the webpage, please add .htaccess or other security measure.</b>

<h2>Support</h2>

<em>The information listed here is currently limited and will expand shortly.</em>

<ul>
 <li>OS
    <ul>
      <li>Debian 6  </li>
      <li>Debian 7 </li>
      <li>Ubuntu 11.04 +</li>
    </ul>
 </li>
 
 <li>Apache 2</li>
 <li>PHP 5</li>
 <li>Browsers
  <ul>
          <li>Modern Browsers (Not IE)</li>
        
  </ul>
 </li>
</ul>
