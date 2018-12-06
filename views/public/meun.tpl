<script type="text/javascript">try{ace.settings.check('sidebar' , 'fixed')}catch(e){}</script>
<ul class="nav nav-list">
<!-- <li {{if eq $.GActive "Additionals"}}class="active"{{end}}>
		<a href="/"> <i class="menu-icon fa fa-home"></i>
			<span class="menu-text">首页</span>
		</a>
	</li> -->
	{{range $k,$v:=$.Menus}}
	<li id="menu{{$v.Id}}">
		<a href="#" class="dropdown-toggle"> <i class="menu-icon fa fa-list"></i>
			<span class="menu-text">{{$v.Title}}</span>
		</a>
		<ul class="submenu">
			{{range $k2,$v2:=$v.Children}}
			<li {{if eq $.GActive $v2.Url}}
					{{if eq $.GActive "VertApis" "VerttoApis"}}
						{{if eq $.GAction $v2.Actionname}}class="active"{{end}}
					{{else}}
						class="active"
					{{end}}
				{{end}}>
				<a href="/{{$v2.Additional}}/{{$v2.Url}}/{{$v2.Actionname}}"> <i class="menu-icon fa fa-caret-right"></i>
					{{$v2.Title}}
				</a> <b class="arrow"></b>
				{{if eq $.GActive $v2.Url}}
				<script type="text/javascript">
					// jQuery(function($) {
					// 	$("#menu{{$v.Id}}").addClass("open")
					// });
					document.getElementById("menu{{$v.Id}}").setAttribute("class", "open")
				</script>
				{{end}}
			</li>
			{{end}}
		</ul>
	</li>
	{{end}}
</ul>