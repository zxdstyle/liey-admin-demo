import ApiAdmin from '@/service/api/scaffold/admin';

export default function useDestroy() {
  const destroy = async (id: number) => {
    await ApiAdmin.Destroy(id);
  };

  return { destroy };
}
