/*!
 * Olx v0.0.1 (http://www.foreworld.net)
 * Copyright 2013 Foreworld.Net, Inc.
 * Licensed under http://www.apache.org/licenses/LICENSE-2.0
 */

if (typeof jQuery === "undefined") { throw new Error("Olx requires jQuery") }



/**
 * Util
 */
;(function($,undefined){ "use strict";
	// console.log("Hello, Util!");

	/**
	 * Cookie获取或设置
	 *
	 * @name 名称
	 * @return 值
	*/
	$.fn.olxCookie = function($name){
		var m = document.cookie.match(new RegExp("(^| )"+ $name +"=([^;]*)(;|$)"));
		return !m ? "" : unescape(m[2]);
	};

	/**
	 * 查询Url参数值
	 *
	 * @key 参数名称
	 * @return 值
	*/
	$.fn.olxQueryString = function($key){
		var uri = window.location.search;
		var re = new RegExp(""+ $key +"\=([^\&\?]*)", "ig");
		return ((uri.match(re)) ? (uri.match(re)[0].substr($key.length + 1)) : "");
	};

	/**
	 * Ajax
	 *
	 * @url
	 * @obj
	 * @cb
	 * @return this
	*/
	$.fn.olxAjax = function(url, obj, cb){
		$.ajax({
			url: $.fn.olxUtilRandomUrl(url),
			type: "POST",
			dataType: "json",
			data: {
				data: JSON.stringify(obj)
			}
		}).done(function (data) {
			cb(data);
		});
	};

	/**
	 * 地址加参，防止缓存
	 * index.html?ts=1389339795 
	 * index.html?userid=123&ts=1389339795 
	 *
	 * @method
	 * @params url
	 * @return 字符串
	*/
	$.fn.olxUtilRandomUrl = function(url){
		var href = [];
		href[0] = url;
		href[1] = (url+"").indexOf("?") == -1 ? "?" : "&";
		href[2] = "ts=";
		href[3] = Date.parse(new Date()) / 1000;
		return href.join("");
	};

	/**
	 * 与JQ Load相似，但更方便
	 *
	 * @url
	 * @params
	 * @callback
	 * @return this
	*/
	$.fn.olxUtilLoad = function(url, params, callback){
		var that = this;

		$.ajax({
			url: $.fn.olxUtilRandomUrl(url),
			type: "POST",
			dataType: "html",
			data: params
		}).done(function(responseText) {
			var parent = that.parent();
			var length = parent.children().length - 1;
			var prev = that.prev();
			var next = that.next();

			that.remove();

			if(length){
				if(prev.length){
					prev.after(responseText);
				}else{
					next.before(responseText);
				}
			}else{
				parent.append(responseText);
			}
		}).complete(callback);

		return this;
	};

	$.fn.olxUtilLoad2 = function(url, params, callback){
		var that = this;

		$.ajax({
			url: $.fn.olxUtilRandomUrl(url),
			type: "GET",
			dataType: "json"
		}).done(function(responseText) {
			var parent = that.parent();
			var length = parent.children().length - 1;
			var prev = that.prev();
			var next = that.next();

			that.remove();

			if(responseText.success){
				if(length){
					if(prev.length){
						prev.after(responseText.data);
					}else{
						next.before(responseText.data);
					}
				}else{
					parent.append(responseText.data);
				}
			}
		}).complete(callback);

		return this;
	};

})(jQuery);	

/**
 * Grid
 */
;(function($,undefined){ "use strict";
	// console.log("Hello, Grid!");

	var Grid = function (element, options){
		this._element = $(element);
		this.options = options;
	};

	// TODO 全选
	Grid.prototype.selectAll = function(){
		return "huangxin";
	};

	/**
	 * 加载数据
	 *
	 * @method
	 * @params 数组 [params, callback]
	 * @return 
	*/
	Grid.prototype.loadData = function(params){
		var that = this._element,
			url = that.data("url"),
			id = that.attr("id");
		
		$(that).olxUtilLoad2(url, params[0], function(){
			initGrid($("#"+id));
			var fn = params[1];
			fn.apply(this, arguments);
		});
	};

	/**
	 * 获取选中的Checkboxs的值
	 *
	 * @method
	 * @param col 列号,从1开始
	 * @return 数组
	*/
	Grid.prototype.getCheckedRowsValue = function(col){
		var vals = [];
		this.getCheckedRows(col).each(function(){
			vals.push($(this).val());
		});
		return vals;
	};

	/**
	 * 获取选中的Checkboxs的对象
	 *
	 * @method
	 * @param col 列号,从1开始
	 * @return 数组
	*/
	Grid.prototype.getCheckedRowsCheckbox = function(col){
		var vals = [];
		this.getCheckedRows(col).each(function(){
			vals.push($(this));
		});
		return vals;
	};

	/**
	 * 获取选中的Checkboxs的总数
	 *
	 * @method
	 * @param col 列号,从1开始
	 * @return 数字
	 */
	Grid.prototype.getCheckedRowsCount = function(col){
		return this.getCheckedRows(col).length;
	};

	/**
	 * 获取选中行
	 *
	 * @method
	 * @param col 列号,从1开始
	 * @return
	 */
	Grid.prototype.getCheckedRows = function(col){
		return this._element.find("tbody tr td:nth-child("+col+") input:checkbox:checked");
	};

	/**
	 * 获取行数
	 *
	 * @method
	 * @return 数字
	 */
	Grid.prototype.getRowsNum = function(){
		return this._element.find("tbody tr").length;
	};

	/**
	 * 获取列数
	 *
	 * @method
	 * @return 数字
	 */
	Grid.prototype.getColumnCount = function(){
		return this._element.find("thead tr th").length;
	};

	var old = $.fn.olxGrid;

	$.fn.olxGrid = function(option, _relatedTarget){
		var $this = $(this);
		var data = $this.data("olx.grid");
		var options = typeof option == "object" && option;

		if(!data) $this.data("olx.grid", (data = new Grid(this, options)));
		return data[option](_relatedTarget);
	};

	$.fn.olxGrid.Constructor = Grid;

	$.fn.olxGrid.noConflict = function () {
		$.fn.olxGrid = old;
		return this;
	};

	function initGrid(grid){
		/* 获取表头列中全部的checkbox组件 */
		var checkboxs = grid.find("thead tr th input[type='checkbox']");

		for(var i=0;i<checkboxs.length;i++){
			/* 绑定Checkbox事件，实现全选/取消 */
			(function (checkbox){
				/* 当前Checkbox的父节点的索引 */
				var index = checkbox.parent().index();
				index++;

				/* 当前列所有的Checkbox */
				var td_checkboxs = grid.find("tbody tr td:nth-child("+index+") input[type='checkbox']");

				/* 主Checkbox绑定事件 */
				$(checkbox).bind("change",function(){
					var checked = $(this).is(":checked");
					/* 选中/全选 */
					for(var i=0;i<td_checkboxs.length;i++){
						$(td_checkboxs[i]).prop("checked", checked);
					}
				});

				/* tbody中的Checkbox */
				for(var i=0,j=td_checkboxs.length;i<j;i++){
					$(td_checkboxs[i]).bind("change", function(){
						var count,
							checked = $(this).is(":checked");
						if(checked){
							/* 选中的Checkbox的总数 */
							count = 0;
							for(var m=0;m<j;m++){
								if(!$(td_checkboxs[m]).prop("checked")){
									break;
								}
								count++;
							}
							checkbox.prop("checked", j === count);
						}else{
							checkbox.prop("checked", checked);
						}
					});
				}
			})($(checkboxs[i]));
		}
	}

	function initGrids(){
		var tables = $("table[data-olx-type='olx.grid']");
		var i,j;
		for(i=0,j=tables.length;i<j;i++){
			initGrid($(tables[i]));
		}
	}

	$(document).ready(function(){
		initGrids();
	});

})(jQuery);

/**
 * Pagination
 */
;(function($,undefined){ "use strict";
	// console.log("Hello, Pagination!");

	var Pagination = function (element, options){
		this._element = $(element);
		this.options = options;
	};

	/**
	 * 页面跳转
	 *
	 * @method
	 * @params 数组 [params, callback]
	 * @return 
	*/
	Pagination.prototype.turnPage = function(params){
		var _element = this._element;
		/* 获取目标的tableId */
		var target = _element.data("target");

		var newParams = [];
		newParams.push(params[0]);
		newParams.push(function(){
			initPagination(_element);
			var fn = params[1];
			fn.apply(this, arguments);
		})
		/* 开始跳转 */
		$(target).olxGrid("loadData", newParams);
	};

	Pagination.prototype.init = function(){
		console.log(arguments);
	};

	var old = $.fn.olxPagination;

	$.fn.olxPagination = function(option, _relatedTarget){
		var $this = $(this);
		var data = $this.data("olx.pagination");
		var options = typeof option == "object" && option;

		if(!data) $this.data("olx.pagination", (data = new Pagination(this, options)));
		return data[option](_relatedTarget);
	};

	$.fn.olxPagination.Constructor = Pagination;

	$.fn.olxPagination.noConflict = function () {
		$.fn.olxPagination = old;
		return this;
	};

	function initPagination(pagination){
		var target =  pagination.data("target"),
			i,j,
			grid = $(target),

			count = grid.data("count"),			//总数
			pagesize = grid.data("pagesize"),	//每页大小
			current = grid.data("current"),		//当前页

			html = [],
			pageCount = Math.ceil(count/pagesize),
			pagination_id = pagination.attr("id");


		html.push('<li><span>第<span class="text-success">'+ current +'</span>页 共<span class="text-success">'+ count +'</span>条</span></li>');

		if(1 == current){
			html.push('<li class="disabled"><a href="javascript:void(0);">首页</a></li>');
		}else{
			html.push('<li><a href="javascript:void(0);" onclick="'+ pagination_id +'(1);"))>首页</a></li>');
		}

		if(1 == current){
			html.push('<li class="disabled"><a href="javascript:void(0);">上页</a></li>');
		}else{
			html.push('<li><a href="javascript:void(0);" onclick="'+ pagination_id +'('+ (current-1) +');"))>上页</a></li>');
		}

		if(pageCount == current){
			html.push('<li class="disabled"><a href="javascript:void(0);">下页</a></li>');
		}else{
			html.push('<li><a href="javascript:void(0);" onclick="'+ pagination_id +'('+ (current+1) +');"))>下页</a></li>');
		}

		if(pageCount == current){
			html.push('<li class="disabled"><a href="javascript:void(0);">尾页</a></li>');
		}else{
			html.push('<li><a href="javascript:void(0);" onclick="'+ pagination_id +'('+ pageCount +');"))>尾页</a></li>');
		}

		/* 数字分页 */
		// html.push('<li class="disabled"><a href="javascript:void(0);">&laquo;</a></li>');

		// if(10 >= pageCount){
		// 	for(i=1,j=(pageCount+1);i<j;i++){
		// 		if(i == current){
		// 			html.push('<li class="active"><a href="javascript:void(0);">'+ i +'<span class="sr-only">(current)</span></a></li>');
		// 		}else{
		// 			html.push('<li><a href="javascript:void(0);" onclick="'+ pagination_id +'('+ i +')">'+ i +'</a></li>');
		// 		}
		// 	}			
		// }else{
		// 	for(i=1,j=8;i<j;i++){
		// 		if(i == current){
		// 			html.push('<li class="active"><a href="javascript:void(0);">'+ i +'<span class="sr-only">(current)</span></a></li>');
		// 		}else{
		// 			html.push('<li><a href="javascript:void(0);" onclick="'+ pagination_id +'('+ i +')">'+ i +'</a></li>');
		// 		}
		// 	}

		// 	html.push('<li><a>...</a></li>');

		// 	for(i=(pageCount-1),j=(pageCount+1);i<j;i++){
		// 		if(i == current){
		// 			html.push('<li class="active"><a href="javascript:void(0);">'+ i +'<span class="sr-only">(current)</span></a></li>');
		// 		}else{
		// 			html.push('<li><a href="javascript:void(0);" onclick="'+ pagination_id +'('+ i +')">'+ i +'</a></li>');
		// 		}
		// 	}
		// }

		// html.push('<li class="disabled"><a href="javascript:void(0);">&raquo;</a></li>');
		/* 数字分页 */

		pagination.html(html.join(""));
	}

	function initPaginations(){
		var paginations = $("ul[data-olx-type='olx.pagination']"),
			i,j;
		for(i=0,j=paginations.length;i<j;i++){
			initPagination($(paginations[i]));
		}
	}

	$(document).ready(function(){
		initPaginations();
	})
})(jQuery);

/**
 * Form
 */
;(function($,undefined){ "use strict";
	// console.log("Hello, Form!");

	var Form = function (element, options){
		this._element = $(element);
		this.options = options;
	};

	/**
	 * 表单初始化
	 *
	 * @form
	 * @return 
	*/
	Form.prototype.init = function(){
		var that = this,
			form = this._element;

		form.on("reset", function(){
			return that.reset();
		});
	};

	/**
	 * 表单验证
	 *
	 * @params
	 * @return 
	*/
	Form.prototype.validate = function(){
		var form = this._element,
			i,j,
			widget,
			widgets = form.find("[data-olx-type^='olx.form.']"),
			validateResult;
		
		for(i=0,j=widgets.length;i<j;i++){
			widget = $(widgets[i]);

			switch(widget.data("olx-type")){
				case "olx.form.input":
					validateResult = widget.olxFormInput("validate");
					break;
				default:
					break;
			}
			if(!validateResult) break;
		}

		return validateResult;
	};

	/**
	 * 表单序列化
	 *
	 * @return 
	*/
	Form.prototype.serializeObjectForm = function(){
		var __a = this._element.serializeArray();
		var __b = _.pluck(__a, "name");
		var __c = _.pluck(__a, "value");
		var __d = _.object(__b, __c);
		return __d;
	};

	/**
	 * 表单提交
	 *
	 * @return 
	*/
	Form.prototype.submit = function(params){
		var that = this;

		var frmObj = that.serializeObjectForm();
		/* 表单验证 */
		var vali = params[0],
			fail = params[1],
			valiResu = vali(frmObj),
			frm = that._element[0];

		// if(valiResu){
		// 	if(fail) return fail(valiResu);
		// 	widget = $('#'+ frm.id +'_'+ valiResu[1]);
		// 	if(widget) widget.olxFormInput('validate', valiResu);
		// 	return;
		// }

		if(valiResu) return fail(valiResu);

		$(that).olxAjax(that._element.data('url'), frmObj, params[2]);
		return false;
	};

	/**
	 * 表单重置
	 *
	 * @params
	 * @return 
	*/
	Form.prototype.reset = function(){
		console.log("reset")
		return true;
	};

	var old = $.fn.olxForm;

	$.fn.olxForm = function(option, _relatedTarget){
		var $this = $(this);
		var data = $this.data("olx.form");
		var options = typeof option == "object" && option;

		if(!data) $this.data("olx.form", (data = new Form(this, options)));
		return data[option](_relatedTarget);
	};

	$.fn.olxForm.Constructor = Form;

	$.fn.olxForm.noConflict = function () {
		$.fn.olxForm = old;
		return this;
	};

	function initForms(){
		var forms = $("form[data-olx-type='olx.form']"),
			i,j;
		for(i=0,j=forms.length;i<j;i++){
			$(forms[i]).olxForm("init");
		}
	}

	$(document).ready(function(){
		initForms();
	})
})(jQuery);


/**
 * FormInput
 */
;(function($,undefined){ "use strict";
	// console.log("Hello, FormInput!");

	var Input = function(element, options){
		this._element = $(element);
		this.options = options;
	};

	/**
	 * Input数据验证
	 *
	 * @params
	 * @return 
	*/
	Input.prototype.validate = function(valiResu){
		var $this = this._element;
		/* 启用tooltip */
		$this.tooltip("enable");
		$this.attr("data-original-title", valiResu[0]);
		$this.tooltip("toggle");
		$this.focus();
		return false;
	};

	var old = $.fn.olxFormInput;

	$.fn.olxFormInput = function(option, _relatedTarget){
		var $this = $(this);
		var data = $this.data("olx.form.input");
		var options = typeof option == "object" && option;

		if(!data) $this.data("olx.form.input", (data = new Input(this, options)));
		return data[option](_relatedTarget);
	};

	$.fn.olxFormInput.Constructor = Input;

	$.fn.olxFormInput.noConflict = function () {
		$.fn.olxFormInput = old;
		return this;
	};

})(jQuery);

/**
 * BackTop
 */
;(function($,undefined){ "use strict";
	// console.log("Hello, BackTop!");

	function initBackTop(backTop){
		var element=document.documentElement,
			body=document.body;
		window.onscroll=set;
		backTop.onclick=function(){
			backTop.style.display="none";
			window.onscroll=null;
			this.timer=setInterval(function(){
				element.scrollTop-=Math.ceil((element.scrollTop+body.scrollTop)*0.1);
				body.scrollTop-=Math.ceil((element.scrollTop+body.scrollTop)*0.1);
				if((element.scrollTop+body.scrollTop)==0) clearInterval(backTop.timer,window.onscroll=set);
			},10);
		};
		function set(){backTop.style.display=(element.scrollTop+body.scrollTop>100)?"block":"none"}
	}

	function initBackTops(){
		var backTop = $("div[data-olx-type='olx.backTop']"),
			i,j;
		for(i=0,j=backTop.length;i<j;i++){
			initBackTop(backTop[i]);
		}
	}

	$(document).ready(function(){
		initBackTops();
	})
})(jQuery);