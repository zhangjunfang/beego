<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>图片上传</title>
</head>

<body>
<form action="upload" method="post" enctype="multipart/form-data" name="upload" id="upload">
  <label>
{{.xsrfdata}}
{{.token}}
  图片上传：<input type="file" name="imgs" />
  </label>
<label>
 性别： <input type="text" name="sex" />
  </label>
  <label>
  <input type="submit" name="Submit" value="提交" />
  </label>
</form>
</body>
</html>
