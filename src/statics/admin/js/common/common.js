// JavaScript Document
showMsg = function(title, msg) {
	$("#msg_content").html(msg);
	return $("#msg_box_box").dialog({
				title : title,
				draggable : false,
				modal : false,
				resizable : false
			});
}

showError = function(title, msg) {
	$("#err_content").html(msg);
	return $("#error_msg_box").dialog({
				title : title,
				draggable : false,
				modal : false,
				resizable : false
			});
}

function showConfirm(msg, callback) {
	if (confirm(msg)) {
		callback.call();
	} else {
		return false;
	}
}
logout = function (url) {
	if (top) parent.window.location.href = url;
	location.href = url;
}
AjaxLoader = function() {
	var _self = this;
	_self.show = function() {
		$('#ajax_loader').dialog({
					title : '处理中...',
					draggable : false,
					modal : true,
					resizable : false,
					close : function() {
					}
				});
		$('.ui-dialog-titlebar-close').hide();
	}
	_self.hide = function() {
		$('#ajax_loader').dialog('close');
		$('.ui-dialog-titlebar-close').show();
	}

}
var editor;

var EDITOR_ITEMS = [
                    'source', 'preview', '|', 'plainpaste', 'wordpaste', '|', 
                    'justifyleft', 'justifycenter', 'justifyright',
        			'justifyfull', 'insertorderedlist', 'insertunorderedlist', 'indent', 'outdent', 'subscript',
        			'superscript', 'clearhtml', 'quickformat','|', 'formatblock', 'fontname', 'fontsize', '|', 
        			'forecolor', 'hilitecolor', 'bold','italic', 'underline', 'strikethrough', 'lineheight', 
        			'removeformat', '|', 'image','table', 'hr', 'link', 'unlink'];
var ajaxLoader = new AjaxLoader();
function ajaxForm(formName, callback, presubmit) {
	if (undefined == callback)	callback = function() {};
	if (undefined == presubmit)	presubmit = function() {};
	if ($('#' + formName)) {
		$('#' + formName).submit(function() {
					if(editor) editor.sync();
					presubmit.call();
					var options = {
						dataType : 'json',
						success : function(data) {
							ajaxLoader.hide();
							callback(data);
						}
					};
					ajaxLoader.show();
					$(this).ajaxSubmit(options);
					return false;
				});
	}
}

// ajax默认回调函数
function ajaxCall(ret) {
	if (ret == '')
		return false;
	ret = ('object' == typeof(ret)) ? ret : eval('(' + ret + ')');
	if (ret.success) {
		showMsg('', ret.msg);
	} else {
		showError('', ret.msg);
	}
}

// ajax跳转
function ajaxRedirect(ret, url) {
	if (ret == '')
		return false;
	if (ret) {
		if (ret.success) {
			showMsg('', ret.msg);
			setTimeout(function() {
						location.href = url;
					}, 500);
		} else {
			showError('', ret.msg);
		}
	}
}

// 删除单个信息
function deleteOne(url, msg, e) {
	if (msg == '')
		msg = '确认删除该条信息？';
	showConfirm(msg, function() {
				$.ajax({
							url : url,
							type : 'POST',
							dataType : 'json',
							data : 'token='+token,
							success : function(ret) {
								if (ret.success) {
									showMsg('', ret.msg);
									setTimeout(function() {
												location.reload();
											}, 500);
								} else {
									showError('', ret.msg);
								}
							}
						});
			}, e);

}

$(document).ready(function() {
	$.datepicker.regional['zh-CN'] = {
		clearText : '清除',
		clearStatus : '清除已选日期',
		closeText : '关闭',
		closeStatus : '不改变当前选择',
		prevText : '&lt;上月',
		prevStatus : '显示上月',
		nextText : '下月&gt;',
		nextStatus : '显示下月',
		currentText : '今天',
		currentStatus : '显示本月',
		monthNames : ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月',
				'十月', '十一月', '十二月'],
		monthNamesShort : ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月',
				'九月', '十月', '十一月', '十二月'],
		monthStatus : '选择月份',
		yearStatus : '选择年份',
		weekHeader : '周',
		weekStatus : '年内周次',
		dayNames : ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'],
		dayNamesShort : ['周日', '周一', '周二', '周三', '周四', '周五', '周六'],
		dayNamesMin : ['日', '一', '二', '三', '四', '五', '六'],
		dayStatus : '设置 DD 为一周起始',
		dateStatus : '选择 m月 d日, DD',
		dateFormat : 'yy-mm-dd',
		firstDay : 1,
		initStatus : '请选择日期',
		isRTL : false,
		changeMonth : true,
		changeYear : true
	};
	$.datepicker.setDefaults($.datepicker.regional['zh-CN']);
	$('#stime').datepicker({
				buttonImage : 'img/calendar.gif',
				buttonImageOnly : true,
				showOn : 'both'
			});
	$('#etime').datepicker({
				buttonImage : 'img/calendar.gif',
				buttonImageOnly : true,
				showOn : 'both'
			});
});

function showAttach(id, imgsrc, value) {
	$('#'+id).children('img').attr('src', imgsrc);
	$('#'+id).children('input').attr('value', value);
}

function showFile(data, path) {
	var data = data;
	var str = '';
	str += '<input type="hidden" name="file" value="'+data.file+'">';
	str += '缩略图：<img src="'+path+'/'+data.icon+'" />&nbsp;';
	str += '<input type="hidden" name="icon" value="'+data.icon+'">';
	str += '预览图gif：<img src="'+path+'/'+data.img_gif+'" />&nbsp;';
	str += '<input type="hidden" name="img_gif" value="'+data.img_gif+'">';
	str += '预览图png：<img src="'+path+'/'+data.img_png+'" />&nbsp;';
	str += '<input type="hidden" name="img_png" value="'+data.img_png+'">';
	str += '<input type="hidden" name="file_size" value="'+data.file_size+'">';
	$("#File").html(str);
}



