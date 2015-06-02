<!DOCTYPE html>
<html lang="zh-CN">
<head>
<title>《职业卫生安全评价，从入门到精通》网上报名</title>
<meta charset="utf-8">
<meta name="robots" content="all" />
<meta name="author" content="3203317@qq.com,新" />
<meta name="generator" content="Foreworld" />
<meta http-equiv="content-type" content="text/html;charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="shortcut icon" href="/static/img/favicon.ico">
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
<link rel="stylesheet" type="text/css"
	href="/static/js/bootstrap-datetimepicker/2.2.0/css/bootstrap-datetimepicker.min.css"
	media="screen">
<style>
body {
	margin-top: 10px;
}

.container {
	width: 1000px;
}

.olx-form-required {
	color: red;
}
</style>
</head>
<body>

	<div class="container">
		<div class="row">
			<div class="col-md-12">
				<div class="panel panel-default">
					<div class="panel-heading">
						<b>《职业卫生安全评价，从入门到精通》网上报名</b>
						<b>{{.msg}}</b>
					</div>
					<div class="panel-body">
						<form id="addFrm" class="form-horizontal"  method="post"
							action="addUser">
							<div class="form-group">
								<label for="addFrm_xm" class="col-sm-2 control-label"><span
									class="olx-form-required">*</span>姓名</label>
								<div class="col-sm-4">
								 {{ .xsrfdata }}
								{{.token}}
									<input type="text" class="form-control" id="addFrm_xm"
										name="xm" placeholder="姓名">
								</div>
								<label for="addFrm_xb" class="col-sm-2 control-label">性别</label>
								<div class="col-sm-4">
									<select class="form-control" id="addFrm_xb" name="xb">
										<option value="男">男</option>
										<option value="女">女</option>
									</select>
								</div>
							</div>
							<div class="form-group">
								<label for="addFrm_sfzh" class="col-sm-2 control-label"><span
									class="olx-form-required">*</span>身份证号</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_sfzh"
										name='sfzh' placeholder="身份证号">
								</div>
								<label for="addFrm_zzmm" class="col-sm-2 control-label">政治面貌</label>
								<div class="col-sm-4">
									<select class="form-control" id="addFrm_zzmm" name="zzmm">
										<option value="中共党员">中共党员</option>
										<option value="中共预备党员">中共预备党员</option>
										<option value="共青团员">共青团员</option>
										<option value="民革党员">民革党员</option>
										<option value="民盟盟员">民盟盟员</option>
										<option value="民建会员">民建会员</option>
										<option value="民进会员">民进会员</option>
										<option value="农工党党员">农工党党员</option>
										<option value="致公党党员">致公党党员</option>
										<option value="九三学社社员">九三学社社员</option>
										<option value="台盟盟员">台盟盟员</option>
										<option value="无党派人士">无党派人士</option>
										<option value="群众（现称普通公民）">群众（现称普通公民）</option>
										<option value="港澳同胞">港澳同胞</option>
									</select>
								</div>
							</div>
							<div class="form-group">
								<label for="addFrm_mz" class="col-sm-2 control-label">民族</label>
								<div class="col-sm-4">
									<select class="form-control" id="addFrm_mz" name="mz">
										<option value="汉族">汉族</option>
										<option value="蒙古族">蒙古族</option>
										<option value="彝族">彝族</option>
										<option value="侗族">侗族</option>
										<option value="哈萨克族">哈萨克族</option>
										<option value="畲族">畲族</option>
										<option value="纳西族">纳西族</option>
										<option value="仫佬族">仫佬族</option>
										<option value="仡佬族">仡佬族</option>
										<option value="怒族">怒族</option>
										<option value="保安族">保安族</option>
										<option value="鄂伦春族">鄂伦春族</option>
										<option value="回族">回族</option>
										<option value="壮族">壮族</option>
										<option value="瑶族">瑶族</option>
										<option value="傣族">傣族</option>
										<option value="高山族">高山族</option>
										<option value="景颇族">景颇族</option>
										<option value="羌族">羌族</option>
										<option value="锡伯族">锡伯族</option>
										<option value="乌孜别克族">乌孜别克族</option>
										<option value="裕固族">裕固族</option>
										<option value="赫哲族">赫哲族</option>
										<option value="藏族">藏族</option>
										<option value="布依族">布依族</option>
										<option value="白族">白族</option>
										<option value="黎族">黎族</option>
										<option value="拉祜族">拉祜族</option>
										<option value="柯尔克孜族">柯尔克孜族</option>
										<option value="布朗族">布朗族</option>
										<option value="阿昌族">阿昌族</option>
										<option value="俄罗斯族">俄罗斯族</option>
										<option value="京族">京族</option>
										<option value="门巴族">门巴族</option>
										<option value="维吾尔族">维吾尔族</option>
										<option value="朝鲜族">朝鲜族</option>
										<option value="土家族">土家族</option>
										<option value="傈僳族">傈僳族</option>
										<option value="水族">水族</option>
										<option value="土族">土族</option>
										<option value="撒拉族">撒拉族</option>
										<option value="普米族">普米族</option>
										<option value="鄂温克族">鄂温克族</option>
										<option value="塔塔尔族">塔塔尔族</option>
										<option value="珞巴族">珞巴族</option>
										<option value="苗族">苗族</option>
										<option value="满族">满族</option>
										<option value="哈尼族">哈尼族</option>
										<option value="佤族">佤族</option>
										<option value="东乡族">东乡族</option>
										<option value="达斡尔族">达斡尔族</option>
										<option value="毛南族">毛南族</option>
										<option value="塔吉克族">塔吉克族</option>
										<option value="德昂族">德昂族</option>
										<option value="独龙族">独龙族</option>
										<option value="基诺族">基诺族</option>
									</select>
								</div>
								<label for="addFrm_jg" class="col-sm-2 control-label">籍贯</label>
								<div class="col-sm-4">
									<select class="form-control" id="addFrm_jg" name="jg">
										<option value="111">北京</option>
										<option value="112">天津</option>
										<option value="113">河北</option>
										<option value="114">山西</option>
										<option value="115">内蒙古</option>
										<option value="121">辽宁</option>
										<option value="122">吉林</option>
										<option value="123">黑龙江</option>
										<option value="131">上海</option>
										<option value="132">江苏</option>
										<option value="133">浙江</option>
										<option value="134">安徽</option>
										<option value="135">福建</option>
										<option value="136">江西</option>
										<option value="137">山东</option>
										<option value="141" selected>河南</option>
										<option value="142">湖北</option>
										<option value="143">湖南</option>
										<option value="144">广东</option>
										<option value="145">广西</option>
										<option value="146">海南</option>
										<option value="150">重庆</option>
										<option value="151">四川</option>
										<option value="152">贵州</option>
										<option value="153">云南</option>
										<option value="154">西藏</option>
										<option value="161">陕西</option>
										<option value="162">甘肃</option>
										<option value="163">青海</option>
										<option value="164">宁夏</option>
										<option value="165">新疆</option>
										<option value="171">台湾</option>
										<option value="172">香港</option>
										<option value="173">澳门</option>
									</select>
								</div>
							</div>
							<div class="form-group">
								<label for="addFrm_byyx" class="col-sm-2 control-label"><span
									class="olx-form-required">*</span>毕业院校</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_byyx"
										name="byyx" placeholder="毕业院校">
								</div>
								<label for="addFrm_shadow_bysj" class="col-sm-2 control-label">毕业时间</label>
								<div class="col-sm-4">
									<div class="input-group date form_date col-md-12" data-date=""
										data-date-format="yyyy-mm-dd" data-link-field="addFrm_bysj"
										data-link-format="yyyy-mm-dd">
										<input class="form-control" id='addFrm_shadow_bysj'
											type="text" value="" readonly placeholder="毕业时间"> <span
											class="input-group-addon"><span
											class="glyphicon glyphicon-remove"></span>
										</span> <span class="input-group-addon"><span
											class="glyphicon glyphicon-calendar"></span>
										</span>
									</div>
									<input type="hidden" data-olx-type="olx.form.input"
										class="form-control" id="addFrm_bysj" name="bysj"
										placeholder="毕业时间">
								</div>
							</div>
							<div class="form-group">
								<label for="addFrm_xl" class="col-sm-2 control-label">学历</label>
								<div class="col-sm-4">
									<select class="form-control" id="addFrm_xl" name='xl'>
										<option value="5">大专</option>
										<option value="4" selected>本科</option>
										<option value="3">硕士</option>
										<option value="1">博士</option>
										<option value="10">MBA</option>
										<option value="11">EMBA</option>
										<option value="12">中专</option>
										<option value="13">中技</option>
										<option value="7">高中</option>
										<option value="9">初中</option>
										<option value="8">其他</option>
									</select>
								</div>
								<label for="addFrm_zy" class="col-sm-2 control-label">专业</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_zy"
										name="zy" placeholder="专业">
								</div>
							</div>
							<div class="form-group">
								<label for="addFrm_gzdw" class="col-sm-2 control-label">工作单位</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_gzdw"
										name="gzdw" placeholder="工作单位">
								</div>
								<label for="addFrm_szbm" class="col-sm-2 control-label">所在部门</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_szbm"
										name="szbm" placeholder="所在部门">
								</div>
							</div>
							<div class="form-group">
								<label for="addFrm_cszy" class="col-sm-2 control-label">从事专业</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_cszy"
										name="cszy" placeholder="从事专业">
								</div>
								<label for="addFrm_zw" class="col-sm-2 control-label">职位</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_zw"
										name="zw" placeholder="职位">
								</div>
							</div>
							<div class="form-group">
								<label for="addFrm_zc" class="col-sm-2 control-label">职称</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_zc"
										name="zc" placeholder="职称">
								</div>
								<label for="addFrm_lxdh" class="col-sm-2 control-label">联系电话</label>
								<div class="col-sm-4">
									<input type="text" class="form-control" id="addFrm_lxdh"
										name="lxdh" placeholder="联系电话">
								</div>
							</div>
							<div class="form-group">
								<label for="addFrm_lxdz" class="col-sm-2 control-label">联系地址</label>
								<div class="col-sm-10">
									<input type="text" class="form-control" id="addFrm_lxdz"
										name="lxdz" placeholder="联系地址">
								</div>
							</div>

							<div class="form-group">
								<div class="col-sm-offset-2 col-sm-10">
									<button type="button" class="btn btn-primary" id="btn_submit">提交</button>
									<button type="reset" class="btn btn-default">重置</button>
								</div>
							</div>
						</form>
					</div>
				</div>
			</div>
		</div>
	</div>

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
	<script type="text/javascript"
		src="/static/js/bootstrap-datetimepicker/2.2.0/js/bootstrap-datetimepicker.js"
		charset="UTF-8"></script>
	<script type="text/javascript"
		src="/static/js/bootstrap-datetimepicker/2.2.0/js/locales/bootstrap-datetimepicker.zh-CN.js"
		charset="UTF-8"></script>
	<script type="text/javascript" src="/static/js/addFrm.js"></script>
	<script type="text/javascript">
			$('.form_date').datetimepicker({
				language: 'zh-CN',
				weekStart: 1,
				todayBtn: 1,
				autoclose: 1,
				todayHighlight: 1,
				startView: 2,
				minView: 2,
				forceParse: 0
			});
			function findUser() {
				$('#addFrm').submit();
				/*var name = $('#addFrm_xm').val();
				var ids = $('#addFrm_sfzh').val();
	            $.ajax({
	        		url : 'QueryUser',
	        		type : "post",
	        		dataType : 'json',
	        		data:{'name':name,'ids':ids},
	        		contentType : "application/x-www-form-urlencoded",
	        		success : function(data) {
	        			if("null"!=data.id){
	        				alert("已经存在此人信息,请核实个人信息");
	        			}else{
		        			$('#addFrm').submit();
		        		}
	        		}
	        	});*/
			}
			$('#btn_submit').click(function(){
				var s = $('#addFrm').serializeObjectForm();
				var v = valiAddFrm(s);
				if(v){
					return $('#addFrm_'+ v[1]).olxFormInput('validate', v);
				}
				findUser();
			});
		</script>
</body>
</html>