$(function () {

    // 表格初始化
    $('#table').DataTable({
        ajax: {
            url: '/test/ajax/datatables',
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
                    var html = "<a href='#' class='edit'>编辑</a><span> </span><a href='#' class='delete'>删除</a>";

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
        console.log(data.WfiId, data.WfiName);
        var editpage = window.open("/test/detail?WfiId="+data.WfiId);
        // var timer = window.setInterval("IfWindowClosed("+editpage+")", 500);
    });

    // 删除
    $('#table tbody').on('click', 'a.delete', function () {
        var data = $('#table').DataTable().row($(this).parents('tr')).data();
        console.log(data.WfiId, data.WfiName);
        /*
            Ajax
        */
        var params = {};
        params['from'] = 'datatables';
        params['data'] = {};
        params['data']['WfiId'] = data.WfiId;

        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/test/ajax/delete',
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
        // $('#table').DataTable().ajax.reload();
    });

}); 

//判断子窗口是否关闭，关闭刷新页面
function IfWindowClosed(page) {
       //判断B页面打开事件
       if (page.closed == true)          
       {
          //执行A页面的相关方法操作
          XXX();
          //关闭监听器
          window.clearInterval(timer);
       }
}

//A页面的相关方法
function XXX(){
    alert("XXX");
}
