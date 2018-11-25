$(function () {

    // 表格初始化
    $('#table').DataTable({
        ajax: {
            url: '/test/ajax/datatables',
            type: 'POST',
            dataSrc: ''
        },
        columns: [
            { "data": "WfiName" },
            { "data": "WfiDesc" },
            { "data": "WfiStatus" },
            { "data": "ModifyTime" },
            { "data": null }
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

    // 点击事件
    $('#table tbody').on('click', 'tr', function () {
        var table = $('#table').DataTable();
        var data = table.row(this).data();
        // alert( 'You clicked on '+data[0]+'\'s row' );
        console.log(data);
    });

    // 修改
    $('#table tbody').on('click', 'a.edit', function () {
        var data = $('#table').DataTable().row($(this).parents('tr')).data();
        // alert("查看修改：" + data[1] + "," + data[2]);
        console.log(data);
    });

    // 删除
    $('a.delete').click(function () {
        var data = $('#table').DataTable().row($(this).parents('tr')).data();
        alert("删除：" + data[1] + "," + data[2]);
    });

}); 