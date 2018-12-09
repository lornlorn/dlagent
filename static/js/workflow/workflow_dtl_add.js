$(function () {

    window.onunload = function () {
        window.opener.location.reload();
        // window.close();
    }

    $('#submit').click(function () {
        var params = {};
        params['from'] = 'workflow_dtl_add';
        params['data'] = {};
        params['data']['WfiId'] = GetQuery('WfiId');
        params['data']['WfdSeq'] = $('#WfdSeq').val();
        params['data']['WfdName'] = $('#WfdName').val();
        params['data']['WfdStatus'] = $('#WfdStatus').val();
        params['data']['WfdShell'] = $('#WfdShell').val();
        params['data']['WfdCmd'] = $('#WfdCmd').val();

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/wfdadd',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                console.log("请求成功");
                alert('成功');
                window.close();
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