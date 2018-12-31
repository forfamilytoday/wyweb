<ul class="breadcrumb">
  <li>
    <i class="ace-icon fa fa-home home-icon"></i>
    <a href="/">主页</a>
  </li>
</ul>
<div class="nav-search" id="nav-search" style="top:3px;">

</div>

<script type="text/javascript">
    // 切换Index库
    function onchangeindexlist(){
        var args = new Object();
        args["input-indexlist"] = $("#input-indexlist").val()

        //alert(JSON.stringify(args))
        $.ajax({
            type:'POST',
            url:"/ajax/setsource",
            async:true,
            data:args,
            success:function(){
                window.location.reload();//刷新当前页面                
            },
            beforeSend:function(){},
            complete:function(){}
        });

    }
</script>