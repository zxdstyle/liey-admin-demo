import { reactive } from 'vue';
import { DataTableColumn, DataTableFilterState } from 'naive-ui';
import { isArray, isEmpty } from 'lodash-es';

export default function useFilter(reload: (opt?: Recordable) => void) {
  let filterParams = reactive<Recordable>({});

  const handleFilterChange = (filters: DataTableFilterState, initiatorColumn: DataTableColumn) => {
    filterParams = {};
    console.log(filters);
    Object.keys(filters).forEach(key => {
      let val = filters[key];
      if (!val || isEmpty(val)) return;

      let operate = 'eq';
      if (isArray(val)) {
        operate = 'in';
        val = val.join(',');
      }
      const field = `where.${key}.${operate}`;

      filterParams[field] = val;
    });

    reload(filterParams);
  };

  return { handleFilterChange };
}
