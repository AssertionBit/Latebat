import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { IDocument } from "../data/docs";

const documentsApi = createApi({
    baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:8080/api/v1" }),
    reducerPath: "documentsApi",
    endpoints: (build) => ({
        getAll: build.query<Array<IDocument>, void>({
            query: () => ({
                url: "/docs"  
            })
        })
    }),
});

export const {
    useGetAllQuery,
} = documentsApi;
export default documentsApi;
