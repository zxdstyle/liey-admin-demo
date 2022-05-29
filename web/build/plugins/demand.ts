import { createStyleImportPlugin, VxeTableResolve } from 'vite-plugin-style-import';

export default createStyleImportPlugin({
  resolves: [VxeTableResolve()]
});
