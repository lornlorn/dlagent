$(function () {

    window.onunload = function () {
        window.opener.location.reload();
        // window.close();
    }

    // 增加
    $('#add').click(function () {
        var editpage = window.open("/html/workflow_param_add?WfdId=" + GetQuery('WfdId'));
    });

    // 保存
    $('#update').click(function () {
        var params = {};
        params['from'] = 'workflow_param';
        params['data'] = {};
        params['data']['WfdId'] = GetQuery('WfdId');
        params['data']['WfdSeq'] = $('#WfdSeq').val();
        params['data']['WfdName'] = $('#WfdName').val();
        params['data']['WfdStatus'] = $('#WfdStatus').val();
        params['data']['WfdShell'] = $('#WfdShell').val();
        params['data']['WfdCmd'] = $('#WfdCmd').val();

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/wfdupdate',
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
        params['from'] = 'workflow_param';
        params['data'] = {};
        params['data']['WfdId'] = GetQuery('WfdId');

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/wfpdelete',
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
        params['from'] = 'workflow_param';
        params['data'] = {};
        params['data']['paramlist'] = [];

        $('#params .param').each(function () {
            var param = {};
            var wfpid = $(this).children("div.field").children("label#WfpSeq").attr("data-WfpId");
            var wfpname = $(this).children("div.field").children("div.control").children("input#WfpName").val();
            var wfpdefault = $(this).children("div.field").children("div.control").children("input#WfpDefault").val();
            param['WfpId'] = wfpid;
            param['WfpName'] = wfpname;
            param['WfpDefault'] = wfpdefault;
            // console.log(param);
            params['data']['paramlist'].push(param);
            // console.log($(this));
        });

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/wfpupdate',
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
