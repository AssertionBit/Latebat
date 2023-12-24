import DocumentItem from "../components/DocumentItem";
import { EDocumentState } from "../data/docs";
import { useAppSelector } from "../hooks";
import styles from "./download.module.css";

export default (): React.JSX.Element => {
    const documents = useAppSelector(state => state.documents);

    return (
        <>
            <button className={styles.btn}>Download all</button>
            {
                documents.map((value, index) =>
                    value.status === EDocumentState.completed ? 
                        <DocumentItem 
                            id={index}
                            name={value.name} 
                            type={value.type} 
                            status={value.status}
                            key={index}/> : 
                        <></>
                )
            }
        </>
    );
};
