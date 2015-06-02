function valiAddFrm(data){
	if(!data) return '参数异常';
	if(!data.xm || !/^[\u4E00-\u9FA5]{2,10}$/.test(data.xm)) return ['仅支持2-10位中文。', 'xm'];
	if(!data.sfzh || !/^[\w]{15,18}$/.test(data.sfzh)) return ['请输入正确的身份证格式。', 'sfzh'];
	if(!data.byyx || !data.byyx.length) return ['必填项。', 'byyx'];
}