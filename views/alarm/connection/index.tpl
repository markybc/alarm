<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <title>layout 管理系统大布局 - Layui</title>
    <link rel="stylesheet" href="//unpkg.com/layui@2.6.5/dist/css/layui.css">
</head>
<body>
<div class="layui-layout layui-layout-admin">
    <div class="layui-header">
        <div class="layui-logo">Alarm 告警系统</div>
        <!-- 头部区域（可配合layui 已有的水平导航） -->
        <ul class="layui-nav layui-layout-left">
            <li class="layui-nav-item"><a href="">nav 1</a></li>
            <li class="layui-nav-item"><a href="">nav 2</a></li>
            <li class="layui-nav-item"><a href="">nav 3</a></li>
            <li class="layui-nav-item">
                <a href="javascript:;">nav groups</a>
                <dl class="layui-nav-child">
                    <dd><a href="">menu 11</a></dd>
                    <dd><a href="">menu 22</a></dd>
                    <dd><a href="">menu 33</a></dd>
                </dl>
            </li>
        </ul>
        <ul class="layui-nav layui-layout-right">
            <li class="layui-nav-item">
                <a href="javascript:;">
                    <img src="//tva1.sinaimg.cn/crop.0.0.118.118.180/5db11ff4gw1e77d3nqrv8j203b03cweg.jpg" class="layui-nav-img">
                    tester
                </a>
                <dl class="layui-nav-child">
                    <dd><a href="">set 1</a></dd>
                    <dd><a href="">set 2</a></dd>
                </dl>
            </li>
            <li class="layui-nav-item"><a href="">Sign out</a></li>
        </ul>
    </div>

    <div class="layui-side layui-bg-black">
        <div class="layui-side-scroll">
            <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
            <ul class="layui-nav layui-nav-tree"  lay-filter="test">
                <li class="layui-nav-item layui-nav-itemed">
                    <a class="" href="javascript:;">消息配置</a>
                    <dl class="layui-nav-child">
                        <dd class="layui-this"><a href="/alarm-api/alarm/connection/index">数据库列表</a></dd>
                        <dd><a href="/alarm-api/alarm/topic/index">消息主题列表</a></dd>
                        <dd><a href="/alarm-api/alarm/monitor/index">业务SQL监控</a></dd>
                        <dd><a href="/alarm-api/alarm/listener/index">接收者列表</a></dd>
                    </dl>
                </li>
            </ul>
        </div>
    </div>

    <div class="layui-body">
        <!-- 内容主体区域 -->
        <div style="padding: 15px;">
            <div class="layui-btn-container" align="right">
                <button type="button" class="layui-btn" onclick="alert('功能正在升级中!')">添加</button>
            </div>
            <table class="layui-hide" id="test"></table>
        </div>
    </div>

    <div class="layui-footer">
        <!-- 底部固定区域 -->
        底部固定区域
    </div>
</div>

<script src="//unpkg.com/layui@2.6.5/dist/layui.js"></script>
<script>
    //JavaScript代码区域
    layui.use('table', function(){
        var table = layui.table;

        table.render({
            elem: '#test'
            ,url:'/alarm-api/alarm/list/connection'
            ,cellMinWidth: 80 //全局定义常规单元格的最小宽度，layui 2.2.1 新增
            ,cols: [[
                {field:'Id', width:80, title: 'ID'}
                ,{field:'AliasName', title: '数据库别名'}
                ,{field:'Name', title: '数据库'}

            ]]
        });

    });

</script>
</body>
</html>