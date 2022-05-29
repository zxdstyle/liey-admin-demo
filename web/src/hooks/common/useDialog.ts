import { useDialog } from 'naive-ui';

export default function useBasicDialog() {
  const dialog = useDialog();
  function warning(text: string, handler: () => void) {
    dialog.warning({
      title: '警告',
      content: text,
      positiveText: '确认',
      negativeText: '取消',
      onPositiveClick: handler
    });
  }

  return { dialog, warning };
}
