/**
 * 作者: 黄鑫
 * Email: 3203317@qq.com 
 ***/
$.fn.serializeObjectForm = function(){	
	var __a = this.serializeArray();
	var __b = _.pluck(__a, "name");
	var __c = _.pluck(__a, "value");
	var __d = _.object(__b, __c);
	return __d;
};