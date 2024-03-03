<h3>Instruction for use</h3>
Edit the <code>config</code> package
<p><code>d, err := gorm.Open("mysql", "username:password@tcp(IP-Address:3306)/u_shortner?charset=utf8&parseTime=True&loc=Local")</code></p>
Change the DB username, password and ip address respectively

<h3>Endpoints</h3>
The root path <code>/</code> with the <code>GET</code> method retrives all created url from the database
<p>To create a new url, use the <code>/create</code> route with the <code>POST</code></p>
