$(function () {

    window.onunload = function () {
        window.opener.location.reload();
        // window.close();
    }

    // 增加
    $('#add').click(function () {
        var editpage = window.open("/html/ParameterAdd?CompId=" + GetQuery('CompId'));
    });

    // 保存
    $('#update').click(function () {
        var params = {};
        params['from'] = 'ComponentDetail';
        params['data'] = {};
        params['data']['CompId'] = GetQuery('CompId');
        params['data']['CompName'] = $('#CompName').val();
        params['data']['CompCmd'] = $('#CompCmd').val();

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/UpdateComponent',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                console.log("请求成功");
                alert("成功");
                window.location.reload();
            },
            error: function (result) {
                console.log("请求失败");
            },
            complete: function () {
                console.log("Ajax finish");
            },
        });
    });

    // 删除最后一个参数
    $('#del').click(function () {
        var params = {};
        params['from'] = 'ComponentDetail';
        params['data'] = {};
        params['data']['CompId'] = GetQuery('CompId');

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/DelLastParameterByCompID',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                console.log("请求成功");
                alert("成功");
                window.location.reload();
            },
            error: function (result) {
                console.log("请求失败");
            },
            complete: function () {
                console.log("Ajax finish");
            },
        });
    });

    // 更新参数信息
    $('#save').click(function () {
        var params = {};
        params['from'] = 'ComponentDetail';
        params['data'] = {};
        params['data']['paramlist'] = [];

        $('#params .param').each(function () {
            var param = {};
            var paramid = $(this).children("div.field").children("label#ParamSeq").attr("data-ParamId");
            var paramname = $(this).children("div.field").children("div.control").children("input#ParamName").val();
            var paramdefault = $(this).children("div.field").children("div.control").children("input#ParamDefault").val();
            param['ParamId'] = paramid;
            param['CompId'] = GetQuery('CompId');
            param['ParamName'] = paramname;
            param['ParamDefault'] = paramdefault;
            // console.log(param);
            params['data']['paramlist'].push(param);
            // console.log($(this));
        });

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/UpdateParameters',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                console.log("请求成功");
                alert("成功");
                window.location.reload();
            },
            error: function (result) {
                console.log("请求失败");
            },
            complete: function () {
                console.log("Ajax finish");
            },
        });
    });

});

function GetQuery(param) {
    var url = location.search; //获取url中"?"符后的字串 
    var query = new Object();
    if (url.indexOf("?") != -1) {
        var str = url.substr(1);
        strs = str.split("&");
        for (var i = 0; i < strs.length; i++) {
            query[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
        }
    }
    return query[param];
}
