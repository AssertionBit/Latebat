import { configureStore } from "@reduxjs/toolkit";
import docsSlice from "./data/docs";
import documentsApi from "./middlewares/documents";

const store = configureStore({
    reducer: {
        documents: docsSlice.reducer,
        [documentsApi.reducerPath]: documentsApi.reducer,


    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(documentsApi.middleware)
});

export default store;