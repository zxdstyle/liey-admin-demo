import { NTag } from 'naive-ui';
import { iconifyRender } from '@/utils';
import { TableActionOption, useTable, TableAction, TableActionHandler } from '@/components/basic/table';
import ApiRole from '@/service/api/scaffold/role';

export default function useIndex(actionHandler: TableActionHandler<Api.Role>) {
  const actions: TableActionOption = [
    { key: 'edit', label: '编辑', icon: iconifyRender('ep:edit', '', 18) },
    { key: 'detail', label: '详情', icon: iconifyRender('icon-park-outline:view-grid-detail', '', 18) },
    { type: 'divider' },
    {
      key: 'delete',
      label: '删除',
      icon: iconifyRender('ci:trash-full', 'red', 18),
      props: { class: 'important-text-red' }
    }
  ];

  const [registerTable, { reload }] = useTable<Api.Role>({
    api: ApiRole.Index,
    columns: [
      { type: 'selection' },
      { key: 'id', title: 'ID', sorter: { multiple: 2 } },
      { key: 'name', title: '角色名称' },
      {
        key: 'slug',
        title: '标识',
        render(row: Api.Role) {
          return <NTag type="primary">{row.slug}</NTag>;
        }
      },
      { key: 'created_at', title: '创建时间' },
      {
        key: 'action',
        title: '操作',
        render(row) {
          return <TableAction actions={actions} onSelect={(key: string) => handleAction(key, row)} />;
        }
      }
    ]
  });

  async function handleAction(key: string, row: Api.Role) {
    actionHandler(key, row);
    await reload();
  }

  return { registerTable, reload };
}
