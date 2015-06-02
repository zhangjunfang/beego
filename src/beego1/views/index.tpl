<!DOCTYPE html>
<html lang="zh-CN">
<head>
<title>《职业卫生安全评价，从入门到精通》网上报名</title>
<meta charset="utf-8">
<meta name="robots" content="all" />
<meta name="author" content="ocean" />
<meta name="generator" content="Foreworld" />
<meta http-equiv="content-type" content="text/html;charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="shortcut icon" href="/static/img/favicon.ico">
<link rel="icon" href="/favicon.ico" type="image/x-icon" />
<link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
<link rel="Bookmark" href="/static/img/favicon.ico">
<!--[if lte IE 9]>
		<link rel="stylesheet" type="text/css" href="/static/js/bsie/0.0.0/css/bootstrap-ie6.css">
		<link rel="stylesheet" type="text/css" href="/static/js/bsie/0.0.0/css/ie.css">
		<link rel="stylesheet" type="text/css" href="/static/js/verybsie/1.0.0/css/bootstrap-verybsie.css">
		<![endif]-->

<!--[if lte IE 7]>
		<link rel="stylesheet" type="text/css" href="/static/js/verybsie/1.0.0/css/bootstrap-verybsie.css">
		<![endif]-->

<link rel="stylesheet" type="text/css"
	href="/static/js/bootstrap/3.0.3/dist/css/bootstrap.css?v3.0.3"
	media="screen">
<link rel="stylesheet" type="text/css"
	href="/static/js/bootstrap/3.0.3/dist/css/bootstrap-theme.min.css?v3.0.3"
	media="screen">
<style>
body {
	background-color: #F5F5F5;
	padding-bottom: 40px;
	padding-top: 40px;
}

#addFrm {
	background-color: #FFFFFF;
	border: 1px solid #E5E5E5;
	border-radius: 5px;
	box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
	margin: 0 auto 20px;
	max-width: 350px;
	padding: 19px 29px 29px;
}

#alert_info {
	max-width: 350px;
	margin: 0 auto 20px;
}
</style>
</head>
<body>
	<section class="container">
	    <center>{{.msg}}</center>
		<form id="addFrm" method="post"  action="login2">
			<h3>《职业卫生安全评价，从入门到精通》</h3>
			<div class="form-group">
			 {{ .xsrfdata }}
			{{.token}}
				<input id="addFrm_xm" name="xm" type="text" class="form-control"
					placeholder="姓名" value="{{.name}}">
			</div>
			<div class="form-group">
				<input id="addFrm_sfzh" name="sfzh" type="text" class="form-control"
					placeholder="身份证号" value="{{.idCard}}">
			</div>
			<div class="form-group">
			   {{create_captcha}}
	           <input name="captcha" type="text">
			</div>
			<button id="btn_submit" type="submit" class="btn btn-primary">登录</button>
			<a  class="btn btn-default" href="signup.jsp">报名</a>
		</form>
	</section>
	<script type="text/javascript"
		src="/static/js/jquery/2.0.2/jquery.min.js"></script>
	<script type="text/javascript"
		src="/static/js/bootstrap/3.0.3/dist/js/bootstrap.js"></script>
	<script type="text/javascript"
		src="/static/js/jquery/ext/ext.form.js"></script>
	<script type="text/javascript"
		src="/static/js/underscore/1.5.1/underscore.min.js"></script>
	<script type="text/javascript"
		src="/static/js/olxjs/1.0.0/js/olx.min.js?v12"></script>
	<script type="text/javascript" src="/static/js/step0.js"></script>
</body>
</html>



