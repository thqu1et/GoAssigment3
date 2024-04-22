﻿<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>Login - Student Information System</title>
<link rel="stylesheet" href="css/loginAuth.css" />
<link rel="icon" href="images/icons/favicon.ico" type="image/png">
<script>
	function OnL() {
		var el = document.getElementById("username");
		if(typeof el == "undefined") return;
		if(el.value != "") {
			var elPsw = document.getElementById("password");
			elPsw.focus();
		}
		else el.focus();
	}
</script>
</head>
<body onload="OnL();">
<div style="position:absolute;width:100%;height:100%;left:0;top:0;overflow:hidden;min-height:700px;padding:0;margin:0;">
<div style="position:relative; text-align:right; top:5pt">
<a href="?lng=kz"><img height="15" src="images/kz.png" title="Қазақ тілі" border="0"></a>&nbsp;
<a href="?lng=en"><img height="15" src="images/gb.png" title="English" border="0"></a>&nbsp;
 </div>
<div style="position:relative; top:0pt; text-align:right; height:6pt; width:100%; margin-top:10px;
 background-color:#16296E"; " > </div>
<div id=top_back style="top:-15pt">
	<div style="text-align:center;"><img src="images/login-logo-ps.png" /></div>
</div>

<div style="position:relative; top:0pt; text-align:right; height:6pt; width:100%; background-color:#16296E"; " > </div>
<br /><br /><br /><br />
<div style="border:2px solid #16296E;padding:6px;border-radius:5px;" class="wac">
<div class="loginDiv" align="left" style="width:auto;">
    <div style="margin-bottom:12px; text-align:center">
	
 
        <span style="color:#999; line-height:22px;font-size:16px;">
            <i>
                Student Information System            </i>
        </span>
    </div>
    <br style="clear:both" />
    
    <form method="post" action="loginAuth.php" autocomplete="off">
        <div>
            <label style="display: block; margin:6px 0;color:#005490;"><b>Student number</b></label>
            <input id="username" name="username" value="" type="text" class="inp" autocomplete="off"/>
        </div>
        <div style="margin-top:10px">
            <label style="display: block; margin:6px 0;color:#005490;"><b>Password</b></label>
            <input id="password" name="password" type="password" class="inp" autocomplete="off"/>
        </div>
                <br />
        <input type="hidden" name="modstring" value="">
        <input type="submit" class="q-button" name="LogIn" value=" Log in ">
        <div align="right"><a class="loginLink" href="/passwForgot">Can't access my account</a></div>
    </form>
</div>
</div>
<div style="padding:10px 2px" align="right" class="wac"><img style="" src="images/gmail.jpg"/>&nbsp;<a title="Login to SduMail" target="_blank" class="loginLink" href="http://mail.sdu.edu.kz/">Login to SduMail</a></div>
<div style="position:absolute; left:0; bottom:0; text-align:right; height:19pt; width:100%; margin-top:10px; background-color:#16296E"; >
	<div style="padding:4px 20px; border-top:1px solid #FFE0C1; color:#999; color:#FFFFFF; font-size:12pt; font:Arial, Helvetica, sans-serif; text-align:center" >SDU © 2024
    </div>
</div>
</div>
</body>
</html>
