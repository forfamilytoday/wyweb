
    <!--[if !IE]> -->
    <script type="text/javascript">
      window.jQuery || document.write("<script src='/static/assets/js/jquery.js'>"+"<"+"/script>");
    </script>

    <!-- <![endif]-->

    <!--[if IE]>
    <script type="text/javascript">
     window.jQuery || document.write("<script src='/static/assets/js/jquery1x.js'>"+"<"+"/script>");
    </script>
    <![endif]-->
    <script type="text/javascript">
      if('ontouchstart' in document.documentElement) document.write("<script src='/static/assets/js/jquery.mobile.custom.js'>"+"<"+"/script>");
    </script>
    <script src="/static/assets/js/bootstrap.js"></script>

    <!-- page specific plugin scripts -->
    <script src="/static/assets/js/jquery.dataTables.js"></script>
    <script src="/static/assets/js/jquery.dataTables.bootstrap.js"></script>
    <script src="/static/assets/js/jquery-ui.js"></script>
    <script src="/static/assets/js/jquery.ui.touch-punch.js"></script>

    <!--[if lte IE 8]>
      <script src="/static/assets/js/excanvas.js"></script>
    <![endif]-->
    <script src="/static/assets/js/jquery-ui.custom.js"></script>
    <script src="/static/assets/js/jquery.ui.touch-punch.js"></script>
    <script src="/static/assets/js/jquery.easypiechart.js"></script>
    <script src="/static/assets/js/jquery.sparkline.js"></script>
    <script src="/static/assets/js/flot/jquery.flot.js"></script>
    <script src="/static/assets/js/flot/jquery.flot.pie.js"></script>
    <script src="/static/assets/js/flot/jquery.flot.resize.js"></script>
    <script src="/static/assets/js/chosen.jquery.js"></script>
    <script src="/static/assets/js/fuelux/fuelux.spinner.js"></script>
    <script src="/static/assets/js/date-time/bootstrap-datepicker.js"></script>
    <script src="/static/assets/js/date-time/bootstrap-timepicker.js"></script>
    <script src="/static/assets/js/date-time/moment.js"></script>
    <script src="/static/assets/js/date-time/daterangepicker.js"></script>
    <script src="/static/assets/js/date-time/bootstrap-datetimepicker.js"></script>
    <script src="/static/assets/js/bootstrap-colorpicker.js"></script>
    <script src="/static/assets/js/jquery.knob.js"></script>
    <script src="/static/assets/js/jquery.autosize.js"></script>
    <script src="/static/assets/js/jquery.inputlimiter.1.3.1.js"></script>
    <script src="/static/assets/js/jquery.maskedinput.js"></script>
    <script src="/static/assets/js/bootstrap-tag.js"></script>
    <!-- ace scripts -->

    <script src="/static/assets/js/ace/elements.scroller.js"></script>
    <script src="/static/assets/js/ace/elements.colorpicker.js"></script>
    <script src="/static/assets/js/ace/elements.fileinput.js"></script>
    <script src="/static/assets/js/ace/elements.typeahead.js"></script>
    <script src="/static/assets/js/ace/elements.wysiwyg.js"></script>
    <script src="/static/assets/js/ace/elements.spinner.js"></script>
    <script src="/static/assets/js/ace/elements.treeview.js"></script>
    <script src="/static/assets/js/ace/elements.wizard.js"></script>
    <script src="/static/assets/js/ace/elements.aside.js"></script>
    <script src="/static/assets/js/ace/ace.js"></script>
    <script src="/static/assets/js/ace/ace.ajax-content.js"></script>
    <script src="/static/assets/js/ace/ace.touch-drag.js"></script>
    <script src="/static/assets/js/ace/ace.sidebar.js"></script>
    <script src="/static/assets/js/ace/ace.sidebar-scroll-1.js"></script>
    <script src="/static/assets/js/ace/ace.submenu-hover.js"></script>
    <script src="/static/assets/js/ace/ace.widget-box.js"></script>
    <script src="/static/assets/js/ace/ace.settings.js"></script>
    <script src="/static/assets/js/ace/ace.settings-rtl.js"></script>
    <script src="/static/assets/js/ace/ace.settings-skin.js"></script>
    <script src="/static/assets/js/ace/ace.widget-on-reload.js"></script>
    <script src="/static/assets/js/ace/ace.searchbox-autocomplete.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/wangEditor/dist/css/wangEditor.min.css">
    <script type="text/javascript" src="/static/wangEditor/dist/js/wangEditor.min.js"></script>
    <!-- 数据验证 -->
    <script src="/static/assets/js/jquery.validate.js"></script>
    <script src="/static/assets/js/validate.message.js"></script>

    <!-- 弹出层 -->
    <script src="/static/assets/layer/layer.min.js"></script>

    <script src="/static/js/web.comm.js"></script>
	<script src="/static/js/web.bindevent.js"></script>
    <script src="/static/js/datatable.http.js"></script>


    <link rel="stylesheet" href="/static/assets/flatpickr/flatpickr.min.css">
    <script src="/static/assets/flatpickr/flatpickr.min.js"></script>
    <!-- inline scripts related to this page -->
    <script type="text/javascript">
      jQuery(function($) {            

            //时间选择器
            $('input[name=input-date-range-picker]').daterangepicker({
              'applyClass' : 'btn-sm btn-success',
              'cancelClass' : 'btn-sm btn-default',
              locale: {
                applyLabel: '确定',
                cancelLabel: '取消',
              }
            })
            .prev().on(ace.click_event, function(){
              $(this).next().focus();
            });            
       })



      $.widget("ui.dialog", $.extend({}, $.ui.dialog.prototype, {
          _title: function(title) {
              var $title = this.options.title || '&nbsp;'
              if( ("title_html" in this.options) && this.options.title_html == true )
                  title.html($title);
              else title.text($title);
          }
      }));
    </script>