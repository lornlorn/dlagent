$(function () {

    window.onunload = function () {
        window.opener.location.reload();
        // window.close();
    }

    // $('#WfiStatus').val("{{ .WfiStatus }}");

    // 表格初始化
    $('#wfd').DataTable({
        ajax: {
            url: '/ajax/getworkflowdtl?WfiId=' + GetQuery('WfiId'),
            type: 'POST',
            dataSrc: ''
        },
        columns: [
            { "data": "WfdSeq", width: '5%' },
            { "data": "WfdName" },
            { "data": "WfdStatus" },
            { "data": "WfdShell" },
            { "data": "WfdCmd" },
            { "data": "ModifyTime" },
            { "data": null, width: '10%' }
        ],
        columnDefs: [
            {
                targets: -1,
                render: function (data, type, row) {
                    // return data + ' (' + row[3] + ')';
                    var id = '"' + row.id + '"';
                    var html = "<a href='#' class='wfdEdit'>编辑</a><span> </span><a href='#' class='wfdDelete'>删除</a>";

                    return html;
                }
            }
        ]
    });

    // 修改
    $('#wfd tbody').on('click', 'a.wfdEdit', function () {
        var data = $('#wfd').DataTable().row($(this).parents('tr')).data();
        console.log(data.WfdId, data.WfdName);
        var editpage = window.open("/html/workflowparam?WfdId=" + data.WfdId);
    });

    // 删除
    $('#wfd tbody').on('click', 'a.wfdDelete', function () {
        var data = $('#wfd').DataTable().row($(this).parents('tr')).data();
        console.log(data.WfdId, data.WfdName);
        /*
            Ajax
        */
        var params = {};
        params['from'] = 'workflow_dtl';
        params['data'] = {};
        params['data']['WfdId'] = data.WfdId;

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/wfddelete',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                console.log("请求成功");
            },
            error: function (result) {
                console.log("请求失败");
            },
            complete: function () {
                console.log("Ajax finish");
                $('#wfd').DataTable().ajax.reload();
            },
        });
        /*
             Ajax end
        */
    });

    // 增加
    $('#add').click(function () {
        var editpage = window.open("/html/workflow_dtl_add?WfiId=" + GetQuery('WfiId'));
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
