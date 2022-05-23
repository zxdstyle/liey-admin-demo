import { reactive } from 'vue';
import { DataTableSortState } from 'naive-ui';
import { isArray } from 'lodash-es';

export default function useSorter(reload: (opt?: Recordable) => void) {
  const sortParams = reactive<Recordable>({});

  const resolveSort = (option: DataTableSortState) => {
    const key = `order.${option.columnKey}`;
    switch (option.order) {
      case 'ascend':
        sortParams[key] = 'ASC';
        break;
      case 'descend':
        sortParams[key] = 'DESC';
        break;
      case false:
        delete sortParams[key];
        break;
      default:
    }
  };

  const handleSorterChange = (options: DataTableSortState | DataTableSortState[] | null) => {
    if (!options) return;

    if (isArray(options)) {
      options.forEach(option => resolveSort(option));
    } else {
      resolveSort(options);
    }

    reload(sortParams);
  };

  return { sortParams, handleSorterChange };
}
