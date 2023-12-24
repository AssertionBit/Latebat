import { useEffect, useState } from "react";
import styles from "./DocumentUpload.module.css";


export default (): React.JSX.Element => {
    const [clicked, setClicked] = useState<boolean>(false);

    // const handleUpload = (e: FormDataEvent) => {
    //     const formReq = new FormData();

    //     const el = document.getElementById("file-upload").files[0];
    //     formReq.append("file", el.)
    // };

    useEffect(() => {
        const dialog = document.querySelector("dialog");
        if(clicked) {
            dialog?.showModal();
        } else {
            dialog?.close();
        }
    }, [clicked]);

    return (
        <>
            <div className={styles.wrap} onClick={() => setClicked(!clicked)}>
                <p>Click here to upload document</p>
            </div>

            {clicked && <dialog className={styles.dialog}>
                <form method="dialog">
                    <form method="post"
                          action="http://localhost:8080/api/v1/docs" 
                          encType="multipart/form-data">
                        <input type="file" name="file" id="file-upload" accept=".png,.jpeg,.jpg,.pdf" multiple/>
                        <input type="submit" value="Submit" />
                    </form>
                    <button 
                        className={styles.formCloseBtn}
                        onClick={() => setClicked(!clicked)}>
                        Close
                    </button>
                </form>
            </dialog> }
        </>
    );
};
