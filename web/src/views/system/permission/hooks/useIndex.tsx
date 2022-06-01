import { NTag } from 'naive-ui';
import { iconifyRender } from '@/utils';
import { TableActionOption, useTable, TableAction, TableActionHandler } from '@/components/basic/table';
import ApiPermission from '@/service/api/scaffold/permission';

export default function useIndex(actionHandler: TableActionHandler<Api.Permission>) {
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

  const [registerTable, { reload }] = useTable<Api.Permission>({
    api: ApiPermission.TreeData,
    columns: [
      { type: 'selection' },
      { key: 'id', title: 'ID', sorter: { multiple: 2 } },
      { key: 'name', title: '权限名称' },
      {
        key: 'slug',
        title: '标识',
        render(row: Api.Permission) {
          return <NTag type="primary">{row.slug}</NTag>;
        }
      },
      {
        key: 'rules',
        title: '权限规则',
        render(row: Api.Permission) {
          return (
            <div>
              {row.rules?.map(rule => (
                <div class="mb-2">
                  {rule.http_methods.map(method => (
                    <NTag size="small" type="primary">
                      {method}
                    </NTag>
                  ))}
                  <NTag size="small">{rule.http_path}</NTag>
                </div>
              ))}
            </div>
          );
        }
      },
      { key: 'created_at', title: '创建时间' },
      { key: 'sort_num', title: '排序值' },
      {
        key: 'action',
        title: '操作',
        render(row) {
          return <TableAction actions={actions} onSelect={(key: string) => handleAction(key, row)} />;
        }
      }
    ],
    pagination: false
  });

  async function handleAction(key: string, row: Api.Permission) {
    actionHandler(key, row);
    await reload();
  }

  return { registerTable, reload };
}
