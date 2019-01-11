$(function () {

    // 表格初始化
    $('#table').DataTable({
        ajax: {
            url: '/ajax/GetComponents',
            type: 'POST',
            dataSrc: ''
        },
        columns: [
            { "data": "CompNo" },
            { "data": "CompName", className: 'compname' },
            { "data": "CompCmd" },
            { "data": "ModifyTime" },
            { "data": null, className: 'operation', width: '10%' }
        ],
        columnDefs: [
            {
                targets: -1,
                render: function (data, type, row) {
                    // return data + ' (' + row[3] + ')';
                    var id = '"' + row.id + '"';
                    var html = "<a href='#' class='edit'>编辑</a><span> </span><a href='#' class='del'>删除</a><span> </span><a href='#' class='run'>执行</a>";

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
    $('#table tbody').on('click', 'a.edit', function () {
        var data = $('#table').DataTable().row($(this).parents('tr')).data();
        console.log(data.CompId, data.CompName);
        var editpage = window.open("/html/ComponentDetail?CompId="+data.CompId);
    });

    // 删除
    $('#table tbody').on('click', 'a.del', function () {
        var data = $('#table').DataTable().row($(this).parents('tr')).data();
        console.log(data.CompId, data.CompName);
        /*
            Ajax
        */
        var params = {};
        params['from'] = 'Component';
        params['data'] = {};
        params['data']['CompId'] = data.CompId;

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/DelComponentByID',
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
                $('#table').DataTable().ajax.reload();
            },
        });
        /*
             Ajax end
        */
    });

    // 增加
    $('#add').click(function () {
        var editpage = window.open("/html/ComponentAdd");
    });

}); 
