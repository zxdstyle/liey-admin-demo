import type { PaginationProps as NPaginationProps } from 'naive-ui';

export interface PaginationRenderProps {
  page: number;
  type: 'page' | 'prev' | 'next';
  originalElement: any;
}

export interface pageSizeOption {
  value: number;
  label: string;
}

export interface PaginationProps extends NPaginationProps {}
