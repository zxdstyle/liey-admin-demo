import { h } from 'vue';
import { Icon } from '@iconify/vue';

/**
 * 动态渲染iconify
 * @param icon - 图标名称
 * @param color - 图标颜色
 * @param size - 图标大小
 */
export function iconifyRender(icon: string, color?: string, size?: number) {
  const style: { color?: string; fontSize?: string } = {};
  if (color && color.length > 0) {
    style.color = color;
  }
  if (size) {
    style.fontSize = `${size}px`;
  }
  return () => h(Icon, { icon, style });
}
