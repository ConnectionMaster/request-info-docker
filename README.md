# request-info-docker
Simple docker web app to output Request path, parameters and headers. It's useful if you want to see how application sees requests from behind proxy/load balancer. 

<b>Run example:</b>

docker run -d -p 8080:8080 --name show-headers keeperlink/request-info-docker:latest

<hr/>
<b>Sample output:</b>

<i>Request Host</i>:  test.mysite.com </br>
<i>Request URL</i>:  /test/path?param1=value1&param2=value2 </br>
<i>Context path</i>:  test/path </br>
<i>Client IP (RemoteAddr)</i>:  172.17.0.1:33870 </br>
<i>Request TLS</i>:  <nil> </br>
</br><b>Request Headers:</b></br>
<ul>
<li><i> Accept </i>: [text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8] </li>
<li><i> Accept-Encoding </i>: [gzip, deflate, sdch] </li>
<li><i> Accept-Language </i>: [en-GB,en;q=0.8,en-US;q=0.6,ru;q=0.4] </li>
<li><i> Cookie </i>: [JSESSIONID=0FADD492D8D1D0E46F9454B75E5720BE; _ym_isad=1] </li>
<li><i> Server-Name </i>: [test] </li>
<li><i> Upgrade-Insecure-Requests </i>: [1] </li>
<li><i> User-Agent </i>: [Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36] </li>
<li><i> X-Forwarded-For </i>: [111.222.333.444] </li>
<li><i> X-Forwarded-Proto </i>: [http] </li>
</ul>
</br><b>Form parameters:</b></br>
<ul>
<li><i> param1 </i>: [value1] </li>
<li><i> param2 </i>: [value2] </li>
</ul>
