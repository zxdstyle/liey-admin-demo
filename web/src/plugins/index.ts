import { App } from 'vue';
import setupAssets from './assets';
import setupVexTable from './vex-table';

function setupPlugins(app: App) {
  setupVexTable(app);
}

export { setupAssets, setupPlugins };
