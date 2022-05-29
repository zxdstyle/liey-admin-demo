// Code generated .* DO NOT EDIT
declare namespace AuthRoute {
  /** 路由的key */
  type RouteKey =
    // 固定的路由
    | 'root' // 根路由
    | 'login'
    | 'not-found'
    | 'no-permission'
    | 'service-error'
    | 'not-found-page' // 捕获无效path的路由
    // 自定义路由
    | 'dashboard'
    | 'dashboard_analysis'
    | 'dashboard_workbench'
    | 'system'
    | 'system_menu'
    | 'system_admin'
    | 'document'
    | 'document_vue'
    | 'document_vue-new'
    | 'document_vite'
    | 'document_naive'
    | 'document_project'
    | 'component'
    | 'component_form'
    | 'component_form_cropper'
    | 'component_button'
    | 'component_card'
    | 'plugin'
    | 'plugin_map'
    | 'plugin_video'
    | 'plugin_editor'
    | 'plugin_editor_quill'
    | 'plugin_editor_markdown'
    | 'plugin_copy'
    | 'plugin_icon'
    | 'plugin_print'
    | 'plugin_swiper'
    | 'exception'
    | 'exception_403'
    | 'exception_404'
    | 'exception_500'
    | 'multi-menu'
    | 'multi-menu_first'
    | 'multi-menu_first_second'
    | 'multi-menu_first_second-new'
    | 'multi-menu_first_second-new_third';
}
