import { FetchBaseQueryError } from "@reduxjs/toolkit/query";
import DocumentItem from "../components/DocumentItem";
import DocumentUpload from "../components/DocumentUpload";
import { useAppSelector } from "../hooks";
import { useGetAllQuery } from "../middlewares/documents";
import styles from "./internal-index.module.css";

export default (): React.JSX.Element => {
    const { isLoading, isError, error } = useGetAllQuery();
    const documents = useAppSelector(state => state.documents);

    return (
        <>
            <h1>Hello!</h1>
            {
                isLoading ? <>Data loading...</> : (
                    isError ? <>Error: {(error as FetchBaseQueryError).status}</> : (
                        <div className={styles.grid}>
                            <DocumentUpload/>
                            {documents.map(
                                (value, index) =>
                                <DocumentItem name={value.name} id={index} type={value.format} status={value.status}/>
                            )}
                        </div>
                    )
                )
            }
        </>
    );
};
