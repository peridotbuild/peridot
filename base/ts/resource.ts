export interface StandardResource {
  name?: string;

  [key: string]: any;
}

export interface ResourceListResponse {
  nextPageToken?: string;
}
