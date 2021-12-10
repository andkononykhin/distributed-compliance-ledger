export interface ModelModel {
    /** @format int32 */
    vid?: number;
    /** @format int32 */
    pid?: number;
    /** @format int32 */
    deviceTypeId?: number;
    productName?: string;
    productLabel?: string;
    partNumber?: string;
    /** @format int32 */
    commissioningCustomFlow?: number;
    commissioningCustomFlowUrl?: string;
    /** @format uint64 */
    commissioningModeInitialStepsHint?: string;
    commissioningModeInitialStepsInstruction?: string;
    /** @format uint64 */
    commissioningModeSecondaryStepsHint?: string;
    commissioningModeSecondaryStepsInstruction?: string;
    userManualUrl?: string;
    supportUrl?: string;
    productUrl?: string;
    creator?: string;
}
export interface ModelModelVersion {
    /** @format int32 */
    vid?: number;
    /** @format int32 */
    pid?: number;
    /** @format uint64 */
    softwareVersion?: string;
    softwareVersionString?: string;
    /** @format int32 */
    cdVersionNumber?: number;
    firmwareDigests?: string;
    softwareVersionValid?: boolean;
    otaUrl?: string;
    /** @format uint64 */
    otaFileSize?: string;
    otaChecksum?: string;
    /** @format int32 */
    otaChecksumType?: number;
    /** @format uint64 */
    minApplicableSoftwareVersion?: string;
    /** @format uint64 */
    maxApplicableSoftwareVersion?: string;
    releaseNotesUrl?: string;
    creator?: string;
}
export interface ModelModelVersions {
    /** @format int32 */
    vid?: number;
    /** @format int32 */
    pid?: number;
    softwareVersions?: string[];
    creator?: string;
}
export declare type ModelMsgCreateModelResponse = object;
export declare type ModelMsgCreateModelVersionResponse = object;
export declare type ModelMsgCreateModelVersionsResponse = object;
export declare type ModelMsgDeleteModelResponse = object;
export declare type ModelMsgDeleteModelVersionResponse = object;
export declare type ModelMsgDeleteModelVersionsResponse = object;
export declare type ModelMsgUpdateModelResponse = object;
export declare type ModelMsgUpdateModelVersionResponse = object;
export declare type ModelMsgUpdateModelVersionsResponse = object;
export interface ModelProduct {
    /** @format int32 */
    pid?: number;
    name?: string;
    partNumber?: string;
}
export interface ModelQueryAllModelResponse {
    model?: ModelModel[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface ModelQueryAllModelVersionResponse {
    modelVersion?: ModelModelVersion[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface ModelQueryAllModelVersionsResponse {
    modelVersions?: ModelModelVersions[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface ModelQueryAllVendorProductsResponse {
    vendorProducts?: ModelVendorProducts[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface ModelQueryGetModelResponse {
    model?: ModelModel;
}
export interface ModelQueryGetModelVersionResponse {
    modelVersion?: ModelModelVersion;
}
export interface ModelQueryGetModelVersionsResponse {
    modelVersions?: ModelModelVersions;
}
export interface ModelQueryGetVendorProductsResponse {
    vendorProducts?: ModelVendorProducts;
}
export interface ModelVendorProducts {
    /** @format int32 */
    vid?: number;
    products?: ModelProduct[];
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
    /**
     * key is a value returned in PageResponse.next_key to begin
     * querying the next page most efficiently. Only one of offset or key
     * should be set.
     * @format byte
     */
    key?: string;
    /**
     * offset is a numeric offset that can be used when key is unavailable.
     * It is less efficient than using key. Only one of offset or key should
     * be set.
     * @format uint64
     */
    offset?: string;
    /**
     * limit is the total number of results to be returned in the result page.
     * If left empty it will default to a value to be set by each app.
     * @format uint64
     */
    limit?: string;
    /**
     * count_total is set to true  to indicate that the result set should include
     * a count of the total number of items available for pagination in UIs.
     * count_total is only respected when offset is used. It is ignored when key
     * is set.
     */
    countTotal?: boolean;
    /**
     * reverse is set to true if results are to be returned in the descending order.
     *
     * Since: cosmos-sdk 0.43
     */
    reverse?: boolean;
}
/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
    /** @format byte */
    nextKey?: string;
    /** @format uint64 */
    total?: string;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title model/genesis.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryModelAll
     * @summary Queries a list of Model items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/model/model
     */
    queryModelAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ModelQueryAllModelResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryModel
     * @summary Queries a Model by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/model/model/{vid}/{pid}
     */
    queryModel: (vid: number, pid: number, params?: RequestParams) => Promise<HttpResponse<ModelQueryGetModelResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryModelVersionAll
     * @summary Queries a list of ModelVersion items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/model/model_version
     */
    queryModelVersionAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ModelQueryAllModelVersionResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryModelVersion
     * @summary Queries a ModelVersion by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/model/model_version/{vid}/{pid}/{softwareVersion}
     */
    queryModelVersion: (vid: number, pid: number, softwareVersion: string, params?: RequestParams) => Promise<HttpResponse<ModelQueryGetModelVersionResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryModelVersionsAll
     * @summary Queries a list of ModelVersions items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/model/model_versions
     */
    queryModelVersionsAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ModelQueryAllModelVersionsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryModelVersions
     * @summary Queries a ModelVersions by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/model/model_versions/{vid}/{pid}
     */
    queryModelVersions: (vid: number, pid: number, params?: RequestParams) => Promise<HttpResponse<ModelQueryGetModelVersionsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryVendorProductsAll
     * @summary Queries a list of VendorProducts items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/model/vendor_products
     */
    queryVendorProductsAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ModelQueryAllVendorProductsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryVendorProducts
     * @summary Queries a VendorProducts by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/model/vendor_products/{vid}
     */
    queryVendorProducts: (vid: number, params?: RequestParams) => Promise<HttpResponse<ModelQueryGetVendorProductsResponse, RpcStatus>>;
}
export {};
