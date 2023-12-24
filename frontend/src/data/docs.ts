import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import documentsApi from "../middlewares/documents";

export enum EDocumentType {
    passport = "passport",
    snils = "snils"
}

export enum EDocumentState {
    planned =    "accepted",
    inProgress = "in-progress",
    completed =  "completed"
}

export interface IDocument {
    id: number;
    name: string;
    format: string;
    type: EDocumentType;
    status: EDocumentState;
}

const initialState: Array<IDocument> = [];

const docsSlice = createSlice({
    name: "docs-slice",
    initialState,
    reducers: {
        append: (state, action: PayloadAction<IDocument>) => {
            state.push(action.payload);
        }
    },

    extraReducers(builder) {
        builder.addMatcher(
            documentsApi.endpoints.getAll.matchFulfilled,
            (state, action: PayloadAction<Array<IDocument>>) => {
                state = action.payload;
                return state;
            }
        )
    },
});

export default docsSlice;
