/** 统一失败和成功的请求结果的数据类型 */
export async function handleServiceResult(error: Service.RequestError | null, data: any) {
  if (error) {
    const fail: Service.FailedResult = {
      error,
      data
    };
    return Promise.reject(fail);
  }
}

export async function handleServiceSuccess<T = any>(data: any, meta: ApiResponseMeta) {
  const success: Service.SuccessResult<T> = {
    error: null,
    meta,
    data
  };
  return success;
}
