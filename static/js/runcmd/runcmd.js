$(function () {

    // 表格初始化
    $('#wfi').DataTable({
        ajax: {
            url: '/ajax/getworkflows',
            type: 'POST',
            dataSrc: ''
        },
        columns: [
            { "data": "WfiName", className: 'wfiname' },
            { "data": "WfiDesc" },
            { "data": "WfiStatus" },
            { "data": "ModifyTime" },
            { "data": null, className: 'operation', width: '10%' }
        ],
        columnDefs: [
            {
                targets: -1,
                render: function (data, type, row) {
                    // return data + ' (' + row[3] + ')';
                    var id = '"' + row.id + '"';
                    var html = "<a href='#' class='wfiShow'>查看</a><span> </span><a href='#' class='wfiRun'>执行</a>";

                    return html;
                }
            }
        ]
    });

    /*
    // 点击事件
    $('#table tbody').on('click', 'tr', function () {
        var table = $('#table').DataTable();
        var data = table.row(this).data();
        console.log(data);
    });
    */

    // 修改
    $('#wfi tbody').on('click', 'a.wfiEdit', function () {
        var data = $('#wfi').DataTable().row($(this).parents('tr')).data();
        console.log(data.WfiId, data.WfiName);
        var editpage = window.open("/html/workflowdetail?WfiId="+data.WfiId);
    });

    // 删除
    $('#wfi tbody').on('click', 'a.wfiDelete', function () {
        var data = $('#wfi').DataTable().row($(this).parents('tr')).data();
        console.log(data.WfiId, data.WfiName);
        /*
            Ajax
        */
        var params = {};
        params['from'] = 'workflow_inf';
        params['data'] = {};
        params['data']['WfiId'] = data.WfiId;

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/wfidelete',
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
                $('#wfi').DataTable().ajax.reload();
            },
        });
        /*
             Ajax end
        */
    });

    // 增加
    $('#add').click(function () {
        var editpage = window.open("/html/workflow_inf_add");
    });

}); 
