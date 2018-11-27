$(function () {

       // 表格初始化
       $('#wfd').DataTable({
        ajax: {
            url: '/ajax/getworkflowdtl?WfiId='+GetRequest()['WfiId'],
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
                    var html = "<a href='#' class='wfdEdit'>编辑</a><span> </span><a href='#' class='wfdDelete'>删除</a>";

                    return html;
                }
            }
        ]
    });
});

function GetRequest() {
    var url = location.search; //获取url中"?"符后的字串 
    var theRequest = new Object();
    if (url.indexOf("?") != -1) {
        var str = url.substr(1);
        strs = str.split("&");
        for (var i = 0; i < strs.length; i++) {
            theRequest[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
        }
    }
    return theRequest;
} 