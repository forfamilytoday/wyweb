<ul class="breadcrumb">
  <li>
    <i class="ace-icon fa fa-home home-icon"></i>
    <a href="/">主页</a>
  </li>
</ul>
<div class="nav-search" id="nav-search" style="top:3px;">
	
    <div style="float:right; width:150px; margin:0 5px 0 0;">
        <select class="form-control" id="input-indexlist" name="input-indexlist" style="font-size:12px;" onchange="onchangeindexlist()">
            {{range $k, $v := $.GIndexList}}
               <option {{if eq $.GIndexSel $v.agentId }} selected="selected"{{end}} value="{{$v.agentId}}">{{$v.agentName}}</option>
            {{end}}            
        </select>
        <input type="hidden" name="ajaxDataAction" id="ajaxDataAction" value="{{$.GAjaxDataAction}}" />
        {{range $k, $v := $.GUrlParams}}
            <input type="hidden" name="{{$k}}" id="{{$k}}" value="{{$v}}" filter="true" /> 
        {{end}}
    </div>

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