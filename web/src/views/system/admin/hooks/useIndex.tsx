import { iconifyRender } from '@/utils';
import { TableActionOption, useTable, TableAction, TableActionHandler } from '@/components/basic/table';
import ApiAdmin from '@/service/api/scaffold/admin';
import { ApiSwitch } from '@/components/basic/form';

export default function useIndex(actionHandler: TableActionHandler<Api.Admin>) {
  const actions: TableActionOption = [
    { key: 'edit', label: '编辑', icon: iconifyRender('ep:edit', '', 18) },
    { type: 'divider' },
    {
      key: 'delete',
      label: '删除',
      icon: iconifyRender('ci:trash-full', 'red', 18),
      props: { class: 'important-text-red' }
    }
  ];

  const [registerTable, { reload }] = useTable<Api.Admin>({
    api: ApiAdmin.Index,
    columns: [
      { type: 'selection' },
      { key: 'id', title: 'ID', sorter: { multiple: 2 } },
      { key: 'username', title: '用户名' },
      { key: 'email', title: '邮箱' },
      { key: 'avatar', title: '头像' },
      {
        key: 'is_active',
        title: '是否启用',
        render(row) {
          return (
            <ApiSwitch
              v-model:value={row.is_active}
              api={(params: any) => ApiAdmin.Update(row.id, params)}
              field="is_active"
              checkedValue={1}
              uncheckedValue={0}
            ></ApiSwitch>
          );
        }
      },
      {
        key: 'action',
        title: '操作',
        render(row) {
          return <TableAction actions={actions} onSelect={(key: string) => handleAction(key, row)} />;
        }
      }
    ],
    immediate: true
  });

  async function handleAction(key: string, row: Api.Admin) {
    actionHandler(key, row);
    await reload();
  }

  return { registerTable, reload };
}
