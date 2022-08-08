/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * OpenBooks Search Service
 * This service provides advanced search capabilities to OpenBooks as well as a management interface.
 * OpenAPI spec version: 1.0
 */
import axios from "axios";
import type { AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";
import { useQuery, useMutation } from "@tanstack/react-query";
import type {
  UseQueryOptions,
  UseMutationOptions,
  QueryFunction,
  MutationFunction,
  UseQueryResult,
  QueryKey,
} from "@tanstack/react-query";
import type {
  ModelSchedule,
  ControllersHTTPError,
  ModelScheduleBody,
} from ".././models";

/**
 * @summary Get all schedules
 */
export const getApiSchedule = (
  options?: AxiosRequestConfig
): Promise<AxiosResponse<ModelSchedule[]>> => {
  return axios.get(`/api/schedule`, options);
};

export const getGetApiScheduleQueryKey = () => [`/api/schedule`];

export type GetApiScheduleQueryResult = NonNullable<
  Awaited<ReturnType<typeof getApiSchedule>>
>;
export type GetApiScheduleQueryError = AxiosError<ControllersHTTPError>;

export const useGetApiSchedule = <
  TData = Awaited<ReturnType<typeof getApiSchedule>>,
  TError = AxiosError<ControllersHTTPError>
>(options?: {
  query?: UseQueryOptions<
    Awaited<ReturnType<typeof getApiSchedule>>,
    TError,
    TData
  >;
  axios?: AxiosRequestConfig;
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetApiScheduleQueryKey();

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getApiSchedule>>> = ({
    signal,
  }) => getApiSchedule({ signal, ...axiosOptions });

  const query = useQuery<
    Awaited<ReturnType<typeof getApiSchedule>>,
    TError,
    TData
  >(queryKey, queryFn, queryOptions) as UseQueryResult<TData, TError> & {
    queryKey: QueryKey;
  };

  query.queryKey = queryKey;

  return query;
};

/**
 * @summary Create a schedule
 */
export const postApiSchedule = (
  modelScheduleBody: ModelScheduleBody,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<ModelSchedule>> => {
  return axios.post(`/api/schedule`, modelScheduleBody, options);
};

export type PostApiScheduleMutationResult = NonNullable<
  Awaited<ReturnType<typeof postApiSchedule>>
>;
export type PostApiScheduleMutationBody = ModelScheduleBody;
export type PostApiScheduleMutationError = AxiosError<ControllersHTTPError>;

export const usePostApiSchedule = <
  TError = AxiosError<ControllersHTTPError>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof postApiSchedule>>,
    TError,
    { data: ModelScheduleBody },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof postApiSchedule>>,
    { data: ModelScheduleBody }
  > = (props) => {
    const { data } = props ?? {};

    return postApiSchedule(data, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof postApiSchedule>>,
    TError,
    { data: ModelScheduleBody },
    TContext
  >(mutationFn, mutationOptions);
};
/**
 * @summary Update a schedule
 */
export const putApiSchedule = (
  modelScheduleBody: ModelScheduleBody,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<ModelSchedule>> => {
  return axios.put(`/api/schedule`, modelScheduleBody, options);
};

export type PutApiScheduleMutationResult = NonNullable<
  Awaited<ReturnType<typeof putApiSchedule>>
>;
export type PutApiScheduleMutationBody = ModelScheduleBody;
export type PutApiScheduleMutationError = AxiosError<ControllersHTTPError>;

export const usePutApiSchedule = <
  TError = AxiosError<ControllersHTTPError>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof putApiSchedule>>,
    TError,
    { data: ModelScheduleBody },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof putApiSchedule>>,
    { data: ModelScheduleBody }
  > = (props) => {
    const { data } = props ?? {};

    return putApiSchedule(data, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof putApiSchedule>>,
    TError,
    { data: ModelScheduleBody },
    TContext
  >(mutationFn, mutationOptions);
};
/**
 * @summary Delete a schedule
 */
export const deleteApiScheduleScheduleId = (
  scheduleId: number,
  options?: AxiosRequestConfig
): Promise<AxiosResponse<void>> => {
  return axios.delete(`/api/schedule/${scheduleId}`, options);
};

export type DeleteApiScheduleScheduleIdMutationResult = NonNullable<
  Awaited<ReturnType<typeof deleteApiScheduleScheduleId>>
>;

export type DeleteApiScheduleScheduleIdMutationError =
  AxiosError<ControllersHTTPError>;

export const useDeleteApiScheduleScheduleId = <
  TError = AxiosError<ControllersHTTPError>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteApiScheduleScheduleId>>,
    TError,
    { scheduleId: number },
    TContext
  >;
  axios?: AxiosRequestConfig;
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {};

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof deleteApiScheduleScheduleId>>,
    { scheduleId: number }
  > = (props) => {
    const { scheduleId } = props ?? {};

    return deleteApiScheduleScheduleId(scheduleId, axiosOptions);
  };

  return useMutation<
    Awaited<ReturnType<typeof deleteApiScheduleScheduleId>>,
    TError,
    { scheduleId: number },
    TContext
  >(mutationFn, mutationOptions);
};
